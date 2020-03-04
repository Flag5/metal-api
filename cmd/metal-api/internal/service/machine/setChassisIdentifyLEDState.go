package machine

import (
	"fmt"
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

func (r machineResource) addSetChassisIdentifyLEDStateRoute(ws *restful.WebService, tags []string) {
	ws.Route(ws.POST("/{id}/chassis-identify-led-state").
		To(helper.Editor(r.setChassisIdentifyLEDState)).
		Operation("setChassisIdentifyLEDState").
		Doc("set the state of a chassis identify LED").
		Param(ws.PathParameter("id", "identifier of the machine").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(v1.ChassisIdentifyLEDState{}).
		Writes(v1.MachineResponse{}).
		Returns(http.StatusOK, "OK", v1.MachineResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))
}

func (r machineResource) setChassisIdentifyLEDState(request *restful.Request, response *restful.Response) {
	var requestPayload v1.ChassisIdentifyLEDState
	err := request.ReadEntity(&requestPayload)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	ledState, err := metal.LEDStateFrom(requestPayload.Value)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	if ledState == metal.LEDStateOff && requestPayload.Description == "" {
		// we want a cause why this chassis identify LED is off
		if helper.CheckError(request, response, utils.CurrentFuncName(), fmt.Errorf("you must supply a description")) {
			return
		}
	}

	id := request.PathParameter("id")
	oldMachine, err := r.DS.FindMachineByID(id)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	newMachine := *oldMachine

	newMachine.LEDState = metal.ChassisIdentifyLEDState{
		Value:       ledState,
		Description: requestPayload.Description,
	}

	err = r.DS.UpdateMachine(oldMachine, &newMachine)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	err = response.WriteHeaderAndEntity(http.StatusOK, helper.MakeMachineResponse(&newMachine, r.DS, utils.Logger(request).Sugar()))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}
