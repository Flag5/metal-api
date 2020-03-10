package firewall

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	"github.com/metal-stack/metal-api/pkg/helper"
	v1 "github.com/metal-stack/metal-api/pkg/proto"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r *firewallResource) allocateFirewall(request *restful.Request, response *restful.Response) {
	var requestPayload v1.FirewallCreateRequest
	err := request.ReadEntity(&requestPayload)
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
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
	if requestPayload.Networks != nil && len(requestPayload.Networks) <= 0 {
		if helper.CheckError(request, response, helper.CurrentFuncName(), fmt.Errorf("network ids cannot be empty")) {
			return
		}
	}
	ha := false
	if requestPayload.HA != nil {
		ha = *requestPayload.HA
	}
	if ha {
		if helper.CheckError(request, response, helper.CurrentFuncName(), fmt.Errorf("highly-available firewall not supported for the time being")) {
			return
		}
	}

	image, err := r.ds.FindImage(requestPayload.ImageID)
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
		return
	}

	if !image.HasFeature(metal.ImageFeatureFirewall) {
		if helper.CheckError(request, response, helper.CurrentFuncName(), fmt.Errorf("given image is not usable for a firewall, features: %s", image.ImageFeatureString())) {
			return
		}
	}

	spec := helper.MachineAllocationSpec{
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
		HA:          ha,
		IsFirewall:  true,
	}

	m, err := helper.AllocateMachine(r.ds, r.ipamer, &spec, r.mdc)
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, helper.MakeMachineResponse(m, r.ds, helper.Logger(request).Sugar()))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}
