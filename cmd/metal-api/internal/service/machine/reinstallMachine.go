package machine

import (
	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	v1 "github.com/metal-stack/metal-api/pkg/proto/v1"
	"github.com/metal-stack/metal-api/pkg/util"
	"github.com/metal-stack/metal-lib/httperrors"
	"github.com/metal-stack/metal-lib/zapup"
)

func (r *machineResource) reinstallMachine(request *restful.Request, response *restful.Response) {
	log := util.Logger(request).Sugar()
	var requestPayload v1.MachineReinstallRequest
	err := request.ReadEntity(&requestPayload)
	if service.CheckError(request, response, util.CurrentFuncName(), err) {
		return
	}

	err = r.reinstallOrDeleteMachine(request, response, &requestPayload.ImageID)
	if err != nil {
		service.SendError(log.Desugar(), response, util.CurrentFuncName(), httperrors.InternalServerError(err))
	}
}

func (r *machineResource) abortReinstall(machineID string) {
	log := zapup.MustRootLogger().Sugar()

	m, err := r.ds.FindMachineByID(machineID)
	if err != nil {
		log.Errorw("unable to find machine", "machineID", machineID, "error", err)
		return
	}

	if m.Allocation != nil && m.Allocation.Reinstall {
		old := *m
		m.Allocation.Reinstall = false
		err = r.ds.UpdateMachine(&old, m)
		if err != nil {
			log.Errorw("unable to find machine", "machineID", machineID, "error", err)
		}
	}

	err = PublishMachineCmd(log, m, r, metal.MachineAbortReinstall)
	if err != nil {
		log.Errorw("unable to publish ’Abort Reinstall' command", "machineID", machineID, "error", err)
	}
}
