package sw

import (
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	v1 "github.com/metal-stack/metal-api/cmd/metal-api/internal/service/v1"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/utils"
	"github.com/metal-stack/metal-lib/httperrors"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r switchResource) addDeleteSwitchRoute(ws *restful.WebService, tags []string) {
	ws.Route(ws.DELETE("/{id}").
		To(helper.Editor(r.deleteSwitch)).
		Operation("deleteSwitch").
		Doc("deletes an switch and returns the deleted entity").
		Param(ws.PathParameter("id", "identifier of the switch").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.SwitchResponse{}).
		Returns(http.StatusOK, "OK", v1.SwitchResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))
}

func (r switchResource) deleteSwitch(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")

	s, err := r.DS.FindSwitch(id)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	err = r.DS.DeleteSwitch(s)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, helper.MakeSwitchResponse(s, r.DS, utils.Logger(request).Sugar()))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}