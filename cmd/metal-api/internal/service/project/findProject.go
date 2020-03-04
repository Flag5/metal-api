package project

import (
	"context"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	mdmv1 "github.com/metal-stack/masterdata-api/api/v1"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	v1 "github.com/metal-stack/metal-api/cmd/metal-api/internal/service/v1"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/utils"
	"github.com/metal-stack/metal-lib/httperrors"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r projectResource) addFindProjectRoute(ws *restful.WebService, tags []string) {
	ws.Route(ws.GET("/{id}").
		To(helper.Viewer(r.findProject)).
		Operation("findProject").
		Doc("get project by id").
		Param(ws.PathParameter("id", "identifier of the project").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.ProjectResponse{}).
		Returns(http.StatusOK, "OK", v1.ProjectResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))
}

func (r projectResource) findProject(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")

	p, err := r.mdc.Project().Get(context.Background(), &mdmv1.ProjectGetRequest{Id: id})
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	err = response.WriteHeaderAndEntity(http.StatusOK, p.Project)
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}
