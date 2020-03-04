package service

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	mdmv1 "github.com/metal-stack/masterdata-api/api/v1"
	mdm "github.com/metal-stack/masterdata-api/pkg/client"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/datastore"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/ipam"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	v1 "github.com/metal-stack/metal-api/cmd/metal-api/internal/service/v1"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/utils"
	"github.com/metal-stack/metal-lib/httperrors"
	"github.com/metal-stack/metal-lib/zapup"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (r machineResource) addAllocateMachineRoute(ws *restful.WebService, tags []string) {
	ws.Route(ws.POST("/allocate").
		To(editor(r.allocateMachine)).
		Operation("allocateMachine").
		Doc("allocate a machine").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(v1.MachineAllocateRequest{}).
		Returns(http.StatusOK, "OK", v1.MachineResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))
}

func (r machineResource) allocateMachine(request *restful.Request, response *restful.Response) {
	var requestPayload v1.MachineAllocateRequest
	err := request.ReadEntity(&requestPayload)
	if checkError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	var uuid string
	if requestPayload.UUID != nil {
		uuid = *requestPayload.UUID
	}
	var name string
	if requestPayload.Name != nil {
		name = *requestPayload.Name
	}
	var description string
	if requestPayload.Description != nil {
		description = *requestPayload.Description
	}
	hostname := "metal"
	if requestPayload.Hostname != nil {
		hostname = *requestPayload.Hostname
	}
	var userdata string
	if requestPayload.UserData != nil {
		userdata = *requestPayload.UserData
	}

	image, err := r.ds.FindImage(requestPayload.ImageID)
	if checkError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	if !image.HasFeature(metal.ImageFeatureMachine) {
		if checkError(request, response, utils.CurrentFuncName(), fmt.Errorf("given image is not usable for a machine, features: %s", image.ImageFeatureString())) {
			return
		}
	}

	spec := machineAllocationSpec{
		UUID:        uuid,
		Name:        name,
		Description: description,
		Hostname:    hostname,
		ProjectID:   requestPayload.ProjectID,
		PartitionID: requestPayload.PartitionID,
		SizeID:      requestPayload.SizeID,
		Image:       image,
		SSHPubKeys:  requestPayload.SSHPubKeys,
		UserData:    userdata,
		Tags:        requestPayload.Tags,
		Networks:    requestPayload.Networks,
		IPs:         requestPayload.IPs,
		HA:          false,
		IsFirewall:  false,
	}

	m, err := allocateMachine(r.ds, r.ipamer, &spec, r.mdc)
	if checkError(request, response, utils.CurrentFuncName(), err) {
		// TODO: Trigger network garbage collection
		utils.Logger(request).Sugar().Errorf("machine allocation went wrong, triggered network garbage collection", "error", err)
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, makeMachineResponse(m, r.ds, utils.Logger(request).Sugar()))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}

func allocateMachine(ds *datastore.RethinkStore, ipamer ipam.IPAMer, allocationSpec *machineAllocationSpec, mdc mdm.Client) (*metal.Machine, error) {
	err := validateAllocationSpec(allocationSpec)
	if err != nil {
		return nil, err
	}
	projectID := allocationSpec.ProjectID
	p, err := mdc.Project().Get(context.Background(), &mdmv1.ProjectGetRequest{Id: projectID})
	if err != nil {
		return nil, err
	}

	// Check if more machine would be allocated than project quota permits
	if p.GetProject() != nil && p.GetProject().GetQuotas() != nil && p.GetProject().GetQuotas().GetMachine() != nil {
		mq := p.GetProject().GetQuotas().GetMachine()
		maxMachines := mq.GetQuota().GetValue()
		var actualMachines metal.Machines
		err := ds.SearchMachines(&datastore.MachineSearchQuery{AllocationProject: &projectID}, &actualMachines)
		if err != nil {
			return nil, err
		}
		machineCount := int32(-1)
		imageMap, err := ds.ListImages()
		if err != nil {
			return nil, err
		}
		for _, m := range actualMachines {
			if m.IsFirewall(imageMap.ByID()) {
				continue
			}
			machineCount++
		}
		if machineCount >= maxMachines {
			return nil, fmt.Errorf("project quota for machines reached max:%d", maxMachines)
		}
	}

	machineCandidate, err := findMachineCandidate(ds, allocationSpec)
	if err != nil {
		return nil, err
	}
	// as some fields in the allocation spec are optional, they will now be clearly defined by the machine candidate
	allocationSpec.UUID = machineCandidate.ID
	allocationSpec.PartitionID = machineCandidate.PartitionID
	allocationSpec.SizeID = machineCandidate.SizeID

	networks, err := makeNetworks(ds, ipamer, allocationSpec)
	if err != nil {
		return nil, err
	}

	alloc := &metal.MachineAllocation{
		Created:         time.Now(),
		Name:            allocationSpec.Name,
		Description:     allocationSpec.Description,
		Hostname:        allocationSpec.Hostname,
		Project:         projectID,
		ImageID:         allocationSpec.Image.ID,
		UserData:        allocationSpec.UserData,
		SSHPubKeys:      allocationSpec.SSHPubKeys,
		MachineNetworks: getMachineNetworks(networks),
	}

	// refetch the machine to catch possible updates after dealing with the network...
	machine, err := ds.FindMachineByID(machineCandidate.ID)
	if err != nil {
		return nil, err
	}
	if machine.Allocation != nil {
		return nil, fmt.Errorf("machine %q already allocated", machine.ID)
	}

	old := *machine
	machine.Allocation = alloc
	machine.Tags = makeMachineTags(machine, networks, allocationSpec.Tags)

	err = ds.UpdateMachine(&old, machine)
	if err != nil {
		return nil, fmt.Errorf("error when allocating machine %q, %v", machine.ID, err)
	}

	err = ds.UpdateWaitingMachine(machine)
	if err != nil {
		updateErr := ds.UpdateMachine(machine, &old) // try rollback allocation
		if updateErr != nil {
			return nil, fmt.Errorf("during update rollback due to an error (%v), another error occurred: %v", err, updateErr)
		}
		return nil, fmt.Errorf("cannot allocate machine in DB: %v", err)
	}

	return machine, nil
}

func validateAllocationSpec(allocationSpec *machineAllocationSpec) error {
	if allocationSpec.ProjectID == "" {
		return fmt.Errorf("project id must be specified")
	}

	if allocationSpec.UUID == "" && allocationSpec.PartitionID == "" {
		return fmt.Errorf("when no machine id is given, a partition id must be specified")
	}

	if allocationSpec.UUID == "" && allocationSpec.SizeID == "" {
		return fmt.Errorf("when no machine id is given, a size id must be specified")
	}

	for _, ip := range allocationSpec.IPs {
		if net.ParseIP(ip) == nil {
			return fmt.Errorf("%q is not a valid IP address", ip)
		}
	}

	for _, pubKey := range allocationSpec.SSHPubKeys {
		_, _, _, _, err := ssh.ParseAuthorizedKey([]byte(pubKey))
		if err != nil {
			return fmt.Errorf("invalid public SSH key: %s", pubKey)
		}
	}

	// A firewall must have either IP or Network with auto IP acquire specified.
	if allocationSpec.IsFirewall {
		if len(allocationSpec.IPs) == 0 && allocationSpec.autoNetworkN() == 0 {
			return fmt.Errorf("when no ip is given at least one auto acquire network must be specified")
		}
	}

	if noautoNetN := allocationSpec.noautoNetworkN(); noautoNetN > len(allocationSpec.IPs) {
		return fmt.Errorf("missing ip(s) for network(s) without automatic ip allocation")
	}

	return nil
}

func findMachineCandidate(ds *datastore.RethinkStore, allocationSpec *machineAllocationSpec) (*metal.Machine, error) {
	var err error
	var machine *metal.Machine
	if allocationSpec.UUID == "" {
		// requesting allocation of an arbitrary machine in partition with given size
		machine, err = findAvailableMachine(ds, allocationSpec.PartitionID, allocationSpec.SizeID)
		if err != nil {
			return nil, err
		}
	} else {
		// requesting allocation of a specific, existing machine
		machine, err = ds.FindMachineByID(allocationSpec.UUID)
		if err != nil {
			return nil, fmt.Errorf("machine cannot be found: %v", err)
		}

		if machine.Allocation != nil {
			return nil, fmt.Errorf("machine is already allocated")
		}
		if allocationSpec.PartitionID != "" && machine.PartitionID != allocationSpec.PartitionID {
			return nil, fmt.Errorf("machine %q is not in the requested partition: %s", machine.ID, allocationSpec.PartitionID)
		}

		if allocationSpec.SizeID != "" && machine.SizeID != allocationSpec.SizeID {
			return nil, fmt.Errorf("machine %q does not have the requested size: %s", machine.ID, allocationSpec.SizeID)
		}
	}
	return machine, err
}

func findAvailableMachine(ds *datastore.RethinkStore, partitionID, sizeID string) (*metal.Machine, error) {
	size, err := ds.FindSize(sizeID)
	if err != nil {
		return nil, fmt.Errorf("size cannot be found: %v", err)
	}
	partition, err := ds.FindPartition(partitionID)
	if err != nil {
		return nil, fmt.Errorf("partition cannot be found: %v", err)
	}
	machine, err := ds.FindAvailableMachine(partition.ID, size.ID)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

func makeNetworks(ds *datastore.RethinkStore, ipamer ipam.IPAMer, allocationSpec *machineAllocationSpec) (allocationNetworkMap, error) {
	networks, err := gatherNetworks(ds, ipamer, allocationSpec)
	if err != nil {
		return nil, err
	}

	for _, n := range networks {
		machineNetwork, err := makeMachineNetwork(ds, ipamer, allocationSpec, n)
		if err != nil {
			return nil, err
		}
		n.machineNetwork = machineNetwork
	}

	// the metal-networker expects to have the same unique ASN on all networks of this machine
	asn, err := makeASN(networks)
	if err != nil {
		return nil, err
	}
	for _, n := range networks {
		n.machineNetwork.ASN = asn
	}

	return networks, nil
}

func gatherNetworks(ds *datastore.RethinkStore, ipamer ipam.IPAMer, allocationSpec *machineAllocationSpec) (allocationNetworkMap, error) {
	partition, err := ds.FindPartition(allocationSpec.PartitionID)
	if err != nil {
		return nil, fmt.Errorf("partition cannot be found: %v", err)
	}

	var privateSuperNetworks metal.Networks
	boolTrue := true
	err = ds.SearchNetworks(&datastore.NetworkSearchQuery{PrivateSuper: &boolTrue}, &privateSuperNetworks)
	if err != nil {
		return nil, errors.Wrap(err, "partition has no private super network")
	}

	specNetworks, err := gatherNetworksFromSpec(ds, allocationSpec, partition, privateSuperNetworks)
	if err != nil {
		return nil, err
	}

	var underlayNetwork *allocationNetwork
	if allocationSpec.IsFirewall {
		underlayNetwork, err = gatherUnderlayNetwork(ds, allocationSpec, partition)
		if err != nil {
			return nil, err
		}
	}

	// assemble result
	result := specNetworks
	if underlayNetwork != nil {
		result[underlayNetwork.network.ID] = underlayNetwork
	}

	return result, nil
}

func gatherNetworksFromSpec(ds *datastore.RethinkStore, allocationSpec *machineAllocationSpec, partition *metal.Partition, privateSuperNetworks metal.Networks) (allocationNetworkMap, error) {
	var partitionPrivateSuperNetwork *metal.Network
	for _, privateSuperNetwork := range privateSuperNetworks {
		if partition.ID == privateSuperNetwork.PartitionID {
			partitionPrivateSuperNetwork = &privateSuperNetwork
			break
		}
	}
	if partitionPrivateSuperNetwork == nil {
		return nil, fmt.Errorf("partition %s does not have a private super network", partition.ID)
	}

	// what do we have to prevent:
	// - user wants to place his machine in a network that does not belong to the project in which the machine is being placed
	// - user wants a machine with a private network that is not in the partition of the machine
	// - user wants to define multiple private networks for his machine
	// - user must define one private network
	// - user specifies administrative networks, i.e. underlay or privatesuper networks
	// - user's private network is specified with noauto, which would make the machine have no ip address

	specNetworks := make(map[string]*allocationNetwork)
	var privateNetwork *allocationNetwork

	for _, networkSpec := range allocationSpec.Networks {
		auto := true
		if networkSpec.AutoAcquireIP != nil {
			auto = *networkSpec.AutoAcquireIP
		}

		network, err := ds.FindNetworkByID(networkSpec.NetworkID)
		if err != nil {
			return nil, err
		}

		if network.Underlay {
			return nil, fmt.Errorf("underlay networks are not allowed to be set explicitly: %s", network.ID)
		}
		if network.PrivateSuper {
			return nil, fmt.Errorf("private super networks are not allowed to be set explicitly: %s", network.ID)
		}

		n := &allocationNetwork{
			network:   network,
			auto:      auto,
			ips:       []metal.IP{},
			isPrivate: false,
		}

		for _, privateSuperNetwork := range privateSuperNetworks {
			if network.ParentNetworkID == privateSuperNetwork.ID {
				// this is the user given private network
				if privateNetwork != nil {
					return nil, fmt.Errorf("multiple private networks provided, which is not allowed")
				}
				if network.PartitionID != partitionPrivateSuperNetwork.PartitionID {
					return nil, fmt.Errorf("the private network must be in the partition where the machine is going to be placed")
				}
				n.isPrivate = true
				privateNetwork = n
				break
			}
		}

		specNetworks[network.ID] = n
	}

	if len(specNetworks) != len(allocationSpec.Networks) {
		return nil, fmt.Errorf("given network ids are not unique")
	}

	if privateNetwork == nil {
		return nil, fmt.Errorf("no private network given")
	}

	if privateNetwork.network.ProjectID != allocationSpec.ProjectID {
		return nil, fmt.Errorf("the given private network does not belong to the project, which is not allowed")
	}

	for _, ipString := range allocationSpec.IPs {
		ip, err := ds.FindIPByID(ipString)
		if err != nil {
			return nil, err
		}
		if ip.ProjectID != allocationSpec.ProjectID {
			return nil, fmt.Errorf("given ip %q with project id %q does not belong to the project of this allocation: %s", ip.IPAddress, ip.ProjectID, allocationSpec.ProjectID)
		}
		network, ok := specNetworks[ip.NetworkID]
		if !ok {
			return nil, fmt.Errorf("given ip %q is not in any of the given networks, which is required", ip.IPAddress)
		}
		s := ip.GetScope()
		if s != metal.ScopeMachine && s != metal.ScopeProject {
			return nil, fmt.Errorf("given ip %q is not available for direct attachment to machine because it is already in use", ip.IPAddress)
		}

		network.auto = false
		network.ips = append(network.ips, *ip)
	}

	if !privateNetwork.auto && len(privateNetwork.ips) == 0 {
		return nil, fmt.Errorf("the private network has no auto ip acquisition, but no suitable IPs were provided, which would lead into a machine having no ip address")
	}

	return specNetworks, nil
}

func gatherUnderlayNetwork(ds *datastore.RethinkStore, allocationSpec *machineAllocationSpec, partition *metal.Partition) (*allocationNetwork, error) {
	boolTrue := true
	var underlays metal.Networks
	err := ds.SearchNetworks(&datastore.NetworkSearchQuery{PartitionID: &partition.ID, Underlay: &boolTrue}, &underlays)
	if err != nil {
		return nil, err
	}
	if len(underlays) == 0 {
		return nil, fmt.Errorf("no underlay found in the given partition: %v", err)
	}
	if len(underlays) > 1 {
		return nil, fmt.Errorf("more than one underlay network in partition %s in the database, which should not be the case", partition.ID)
	}
	underlay := &underlays[0]

	return &allocationNetwork{
		network:   underlay,
		auto:      true,
		isPrivate: false,
	}, nil
}

func makeMachineNetwork(ds *datastore.RethinkStore, ipamer ipam.IPAMer, allocationSpec *machineAllocationSpec, n *allocationNetwork) (*metal.MachineNetwork, error) {
	if n.auto {
		ipAddress, ipParentCidr, err := allocateIP(n.network, "", ipamer)
		if err != nil {
			return nil, fmt.Errorf("unable to allocate an ip in network: %s %#v", n.network.ID, err)
		}
		ip := &metal.IP{
			IPAddress:        ipAddress,
			ParentPrefixCidr: ipParentCidr,
			Name:             allocationSpec.Name,
			Description:      "autoassigned",
			NetworkID:        n.network.ID,
			Type:             metal.Ephemeral,
			ProjectID:        allocationSpec.ProjectID,
		}
		ip.AddMachineId(allocationSpec.UUID)
		err = ds.CreateIP(ip)
		if err != nil {
			return nil, err
		}
		n.ips = append(n.ips, *ip)
	}

	var ipAddresses []string
	for _, ip := range n.ips {
		new := ip
		new.AddMachineId(allocationSpec.UUID)
		err := ds.UpdateIP(&ip, &new)
		if err != nil {
			return nil, err
		}
		ipAddresses = append(ipAddresses, ip.IPAddress)
	}

	machineNetwork := metal.MachineNetwork{
		NetworkID:           n.network.ID,
		Prefixes:            n.network.Prefixes.String(),
		IPs:                 ipAddresses,
		DestinationPrefixes: n.network.DestinationPrefixes.String(),
		Private:             n.isPrivate,
		Underlay:            n.network.Underlay,
		Nat:                 n.network.Nat,
		Vrf:                 n.network.Vrf,
	}

	return &machineNetwork, nil
}

// makeASN we can use the IP of the private network (which always have to be present and unique)
// for generating a unique ASN.
func makeASN(networks allocationNetworkMap) (int64, error) {
	privateNetwork, err := getPrivateNetwork(networks)
	if err != nil {
		return 0, err
	}

	if len(privateNetwork.ips) == 0 {
		return 0, fmt.Errorf("private network has no IPs, which would result in a machine without an IP")
	}

	asn, err := privateNetwork.ips[0].ASN()
	if err != nil {
		return 0, err
	}

	return asn, nil
}

// makeMachineTags constructs the tags of the machine.
// following tags are added in the following precedence (from lowest to highest in case of duplication):
// - external network labels (concatenated, from all machine networks that this machine belongs to)
// - private network labels (concatenated)
// - user given tags (from allocation spec)
// - system tags (immutable information from the metal-api that are useful for the end user, e.g. machine rack and chassis)
func makeMachineTags(m *metal.Machine, networks allocationNetworkMap, userTags []string) []string {
	labels := make(map[string]string)

	for _, n := range networks {
		if !n.isPrivate {
			for k, v := range n.network.Labels {
				labels[k] = v
			}
		}
	}

	privateNetwork, _ := getPrivateNetwork(networks)
	if privateNetwork != nil {
		for k, v := range privateNetwork.network.Labels {
			labels[k] = v
		}
	}

	// as user labels are given as an array, we need to figure out if label-like tags were provided.
	// otherwise the user could provide confusing information like:
	// - machine.metal-pod.io/chassis=123
	// - machine.metal-pod.io/chassis=789
	userLabels := make(map[string]string)
	var actualUserTags []string
	for _, tag := range userTags {
		if strings.Contains(tag, "=") {
			parts := strings.SplitN(tag, "=", 2)
			userLabels[parts[0]] = parts[1]
		} else {
			actualUserTags = append(actualUserTags, tag)
		}
	}
	for k, v := range userLabels {
		labels[k] = v
	}

	for k, v := range makeMachineSystemLabels(m) {
		labels[k] = v
	}

	tags := actualUserTags
	for k, v := range labels {
		tags = append(tags, fmt.Sprintf("%s=%s", k, v))
	}

	return uniqueTags(tags)
}

func makeMachineSystemLabels(m *metal.Machine) map[string]string {
	labels := make(map[string]string)
	for _, n := range m.Allocation.MachineNetworks {
		if n.Private {
			if n.ASN != 0 {
				labels[fmt.Sprintf("%s/%s", metal.MachineLabelPrefix, "network.primary.asn")] = strconv.FormatInt(n.ASN, 10)
				break
			}
		}
	}
	if m.RackID != "" {
		labels[fmt.Sprintf("%s/%s", metal.MachineLabelPrefix, "rack")] = m.RackID
	}
	if m.IPMI.Fru.ChassisPartSerial != "" {
		labels[fmt.Sprintf("%s/%s", metal.MachineLabelPrefix, "chassis")] = m.IPMI.Fru.ChassisPartSerial
	}
	return labels
}

// uniqueTags the last added tags will be kept!
func uniqueTags(tags []string) []string {
	tagSet := make(map[string]bool)
	for _, t := range tags {
		tagSet[t] = true
	}
	var uniqueTags []string
	for k := range tagSet {
		uniqueTags = append(uniqueTags, k)
	}
	return uniqueTags
}

func (s machineAllocationSpec) noautoNetworkN() int {
	result := 0
	for _, net := range s.Networks {
		if net.AutoAcquireIP != nil && !*net.AutoAcquireIP {
			result++
		}
	}
	return result
}

func (s machineAllocationSpec) autoNetworkN() int {
	return len(s.Networks) - s.noautoNetworkN()
}

// allocationNetwork is intermediate struct to create machine networks from regular networks during machine allocation
type allocationNetwork struct {
	network        *metal.Network
	machineNetwork *metal.MachineNetwork
	ips            []metal.IP
	auto           bool
	isPrivate      bool
}

// allocationNetworkMap is a map of allocationNetworks with the network id as the key
type allocationNetworkMap map[string]*allocationNetwork

// getPrivateNetwork extracts the private network from an allocationNetworkMap
func getPrivateNetwork(networks allocationNetworkMap) (*allocationNetwork, error) {
	var privateNetwork *allocationNetwork
	for _, n := range networks {
		if n.isPrivate {
			privateNetwork = n
			break
		}
	}
	if privateNetwork == nil {
		return nil, fmt.Errorf("no private network contained")
	}
	return privateNetwork, nil
}

// getMachineNetworks extracts the machines networks from an allocationNetworkMap
func getMachineNetworks(networks allocationNetworkMap) []*metal.MachineNetwork {
	var machineNetworks []*metal.MachineNetwork
	for _, n := range networks {
		machineNetworks = append(machineNetworks, n.machineNetwork)
	}
	return machineNetworks
}