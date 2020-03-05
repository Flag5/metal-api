package size

import (
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	v1 "github.com/metal-stack/metal-api/cmd/metal-api/internal/service/v1"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/utils"
	"github.com/metal-stack/metal-lib/httperrors"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r sizeResource) addUpdateSizeRoute(ws *restful.WebService, tags []string) {
	ws.Route(ws.POST("/").
		To(helper.Admin(r.updateSize)).
		Operation("updateSize").
		Doc("updates a size. if the size was changed since this one was read, a conflict is returned").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(v1.SizeUpdateRequest{}).
		Returns(http.StatusOK, "OK", v1.SizeResponse{}).
		Returns(http.StatusConflict, "Conflict", httperrors.HTTPErrorResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))
}

func (r sizeResource) updateSize(request *restful.Request, response *restful.Response) {
	var requestPayload v1.SizeUpdateRequest
	err := request.ReadEntity(&requestPayload)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	oldSize, err := r.DS.FindSize(requestPayload.ID)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	newSize := *oldSize

	if requestPayload.Name != nil {
		newSize.Name = *requestPayload.Name
	}
	if requestPayload.Description != nil {
		newSize.Description = *requestPayload.Description
	}
	var constraints []metal.Constraint
	if requestPayload.SizeConstraints != nil {
		sizeConstraints := *requestPayload.SizeConstraints
		for i := range sizeConstraints {
			constraint := metal.Constraint{
				Type: sizeConstraints[i].Type,
				Min:  sizeConstraints[i].Min,
				Max:  sizeConstraints[i].Max,
			}
			constraints = append(constraints, constraint)
		}
		newSize.Constraints = constraints
	}

	err = r.DS.UpdateSize(oldSize, &newSize)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, v1.NewSizeResponse(&newSize))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}