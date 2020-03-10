package ip

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful"
	mdmv1 "github.com/metal-stack/masterdata-api/api/v1"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	"github.com/metal-stack/metal-api/pkg/helper"
	v1 "github.com/metal-stack/metal-api/pkg/proto/v1"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r *ipResource) allocateIP(request *restful.Request, response *restful.Response) {
	r.allocateSpecificIP(request, response)
}

func (r *ipResource) allocateSpecificIP(request *restful.Request, response *restful.Response) {
	specificIP := request.PathParameter("ip")
	var requestPayload v1.IPAllocateRequest
	err := request.ReadEntity(&requestPayload)
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
		return
	}

	if requestPayload.NetworkID == "" {
		if helper.CheckError(request, response, helper.CurrentFuncName(), fmt.Errorf("networkid should not be empty")) {
			return
		}
	}
	if requestPayload.ProjectID == "" {
		if helper.CheckError(request, response, helper.CurrentFuncName(), fmt.Errorf("projectid should not be empty")) {
			return
		}
	}

	var name string
	if requestPayload.Name != nil {
		name = *requestPayload.Name
	}
	var description string
	if requestPayload.Description != nil {
		description = *requestPayload.Description
	}

	nw, err := r.ds.FindNetworkByID(requestPayload.NetworkID)
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
		return
	}

	p, err := r.mdc.Project().Get(context.Background(), &mdmv1.ProjectGetRequest{Id: requestPayload.ProjectID})
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
		return
	}

	tags := requestPayload.Tags
	if requestPayload.MachineID != nil {
		tags = append(tags, metal.IpTag(metal.TagIPMachineID, *requestPayload.MachineID))
	}

	tags, err = helper.ProcessTags(tags)
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
		return
	}

	// TODO: Following operations should span a database transaction if possible

	ipAddress, ipParentCidr, err := helper.AllocateIP(nw, specificIP, r.ipamer)
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
		return
	}
	helper.Logger(request).Sugar().Debugw("found an ip to allocate", "ip", ipAddress, "network", nw.ID)

	ipType := metal.Ephemeral
	if requestPayload.Type == metal.Static {
		ipType = metal.Static
	}

	ip := &metal.IP{
		IPAddress:        ipAddress,
		ParentPrefixCidr: ipParentCidr,
		Name:             name,
		Description:      description,
		NetworkID:        nw.ID,
		ProjectID:        p.GetProject().GetMeta().GetId(),
		Type:             ipType,
		Tags:             tags,
	}

	err = r.ds.CreateIP(ip)
	if helper.CheckError(request, response, helper.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusCreated, service.NewIPResponse(ip))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}
