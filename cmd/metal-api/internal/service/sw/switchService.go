package sw

import (
	restful "github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	v1 "github.com/metal-stack/metal-api/cmd/metal-api/internal/service/v1"
	"net/http"
)

func (r switchResource) webService() *restful.WebService {
	return service.Build(service.WebResource{
		Version: service.V1,
		Path:    "switch",
		Routes: []service.Route{
			{
				Method:  http.MethodGet,
				SubPath: "/",
				Doc:     "get all switches",
				Writes:  []v1.SwitchResponse{},
				Handler: r.listSwitches,
			},
			{
				Method:        http.MethodGet,
				SubPath:       "/{id}",
				PathParameter: service.NewPathParameter("id", "identifier of the switch"),
				Doc:           "get switch by id",
				Writes:        v1.SwitchResponse{},
				Handler:       r.findSwitch,
			},
			{
				Method:  http.MethodPost,
				SubPath: "/register",
				Doc:     "register a switch",
				Access:  metal.EditAccess,
				Reads:   v1.SwitchRegisterRequest{},
				Writes:  []v1.SwitchResponse{},
				Returns: []*service.Return{
					service.NewReturn(http.StatusOK, "OK", v1.SwitchResponse{}),
					service.NewReturn(http.StatusCreated, "Created", v1.SwitchResponse{}),
				},
				Handler: r.registerSwitch,
			},
			{
				Method:        http.MethodDelete,
				SubPath:       "/{id}",
				PathParameter: service.NewPathParameter("id", "identifier of the switch"),
				Doc:           "deletes an switch and returns the deleted entity",
				Access:        metal.EditAccess,
				Writes:        v1.SwitchResponse{},
				Handler:       r.deleteSwitch,
			},
		},
	})
}
