package project

import (
	restful "github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	v12 "github.com/metal-stack/metal-api/cmd/metal-api/internal/service/proto/v1"
	"net/http"
)

func (r *projectResource) webService() *restful.WebService {
	return service.Build(&service.WebService{
		Version: service.V1,
		Path:    "project",
		Routes: []*service.Route{
			{
				Method:  http.MethodGet,
				SubPath: "/",
				Doc:     "get all projects",
				Access:  metal.ViewAccess,
				Writes:  []v12.ProjectResponse{},
				Handler: r.listProjects,
			},
			{
				Method:        http.MethodGet,
				SubPath:       "/{id}",
				PathParameter: service.NewPathParameter("id", "identifier of the project"),
				Doc:           "get project by id",
				Access:        metal.ViewAccess,
				Writes:        v12.ProjectResponse{},
				Handler:       r.findProject,
			},
			{
				Method:  http.MethodPost,
				SubPath: "/find",
				Doc:     "get all projects that match given properties",
				Access:  metal.ViewAccess,
				Reads:   v12.ProjectFindRequest{},
				Writes:  []v12.ProjectResponse{},
				Handler: r.findProjects,
			},
		},
	})
}
