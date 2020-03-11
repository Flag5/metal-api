package machine

import (
	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/datastore"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	v1 "github.com/metal-stack/metal-api/pkg/proto/v1"
	"github.com/metal-stack/metal-api/pkg/util"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r *machineResource) findIPMIMachine(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")

	m, err := r.ds.FindMachineByID(id)
	if helper.CheckError(request, response, util.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, makeMachineIPMIResponse(m, r.ds, util.Logger(request).Sugar()))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}

func makeMachineIPMIResponse(m *metal.Machine, ds *datastore.RethinkStore, logger *zap.SugaredLogger) *v1.MachineIPMIResponse {
	s, p, i, ec := helper.FindMachineReferencedEntities(m, ds, logger)
	return service.NewMachineIPMIResponse(m, s, p, i, ec)
}
