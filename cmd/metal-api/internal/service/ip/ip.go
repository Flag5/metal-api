package ip

import (
	"fmt"
	"github.com/emicklei/go-restful"
	v12 "github.com/metal-stack/masterdata-api/api/v1"
	mdm "github.com/metal-stack/masterdata-api/pkg/client"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/datastore"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/ipam"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/pkg/proto/v1"
	"github.com/metal-stack/metal-api/pkg/util"
	"strings"
)

type ipResource struct {
	ds     *datastore.RethinkStore
	ipamer ipam.IPAMer
	mdc    mdm.Client
}

// NewIPService returns a webservice for ip specific endpoints.
func NewIPService(ds *datastore.RethinkStore, ipamer ipam.IPAMer, mdc mdm.Client) *restful.WebService {
	r := ipResource{
		ds:     ds,
		ipamer: ipamer,
		mdc:    mdc,
	}
	return r.webService()
}

func AllocateIP(parent *metal.Network, specificIP string, ipamer ipam.IPAMer) (string, string, error) {
	var errors []error
	var err error
	var ipAddress string
	var parentPrefixCidr string
	for _, prefix := range parent.Prefixes {
		if specificIP == "" {
			ipAddress, err = ipamer.AllocateIP(prefix)
		} else {
			ipAddress, err = ipamer.AllocateSpecificIP(prefix, specificIP)
		}
		if err != nil {
			errors = append(errors, err)
			continue
		}
		if ipAddress != "" {
			parentPrefixCidr = prefix.String()
			break
		}
	}
	if ipAddress == "" {
		if len(errors) > 0 {
			return "", "", fmt.Errorf("cannot allocate free ip in ipam: %v", errors)
		}
		return "", "", fmt.Errorf("cannot allocate free ip in ipam")
	}

	return ipAddress, parentPrefixCidr, nil
}

func NewIPResponse(ip *metal.IP) *v1.IPResponse {
	return &v1.IPResponse{
		IP: ToIP(ip),
		Identifiable: &v1.IPIdentifiable{
			IPAddress: ip.IPAddress,
		},
	}
}

func ToIP(ip *metal.IP) *v1.IP {
	return &v1.IP{
		Common: &v1.Common{
			Meta: &v12.Meta{
				Id:          ip.GetID(),
				Apiversion:  "v1",
				Version:     1,
				CreatedTime: util.TimestampProto(ip.Created),
				UpdatedTime: util.TimestampProto(ip.Changed),
			},
			Name:        util.StringProto(ip.Name),
			Description: util.StringProto(ip.Description),
		},
		NetworkID: ip.NetworkID,
		ProjectID: ip.ProjectID,
		Type:      toIPType(ip.Type),
		Tags:      util.StringSliceProto(ip.Tags...),
	}
}

func toIPType(ipType metal.IPType) v1.IP_Type {
	if strings.EqualFold(string(ipType), "ephemeral") {
		return v1.IP_EPHEMERAL
	}
	return v1.IP_STATIC
}