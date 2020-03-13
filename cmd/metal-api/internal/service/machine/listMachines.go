package machine

import (
	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	"github.com/metal-stack/metal-api/pkg/util"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r *machineResource) listMachines(request *restful.Request, response *restful.Response) {
	ms, err := r.ds.ListMachines()
	if service.CheckError(request, response, util.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, helper.MakeMachineResponseList(ms, r.ds, util.Logger(request).Sugar()))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}
