package partition

import (
	restful "github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	v1 "github.com/metal-stack/metal-api/pkg/proto/v1"
	"github.com/metal-stack/metal-lib/httperrors"
	"net/http"
)

func (r *partitionResource) webService() *restful.WebService {
	return service.Build(&service.WebService{
		Version: service.V1,
		Path:    "partition",
		Routes: []*service.Route{
			{
				Method:  http.MethodGet,
				SubPath: "/",
				Doc:     "get all partitions",
				Writes:  []v1.PartitionResponse{},
				Handler: r.listPartitions,
			},
			{
				Method:        http.MethodGet,
				SubPath:       "/{id}",
				PathParameter: service.NewPathParameter("id", "identifier of the partition"),
				Doc:           "get partition by id",
				Writes:        v1.PartitionResponse{},
				Handler:       r.findPartition,
			},
			{
				Method:  http.MethodPut,
				SubPath: "/",
				Doc:     "creates a partition. If the given ID already exists a conflict is returned",
				Access:  metal.AdminAccess,
				Reads:   v1.PartitionCreateRequest{},
				Writes:  v1.PartitionResponse{},
				Returns: []*service.Return{
					service.NewReturn(http.StatusCreated, "Created", v1.PartitionResponse{}),
					service.NewReturn(http.StatusConflict, "Conflict", httperrors.HTTPErrorResponse{}),
				},
				Handler: r.createPartition,
			},
			{
				Method:  http.MethodPost,
				SubPath: "/",
				Doc:     "updates a partition. If the partition was changed since this one was read, a conflict is returned",
				Access:  metal.AdminAccess,
				Reads:   v1.PartitionUpdateRequest{},
				Writes:  v1.PartitionResponse{},
				Returns: []*service.Return{
					service.NewReturn(http.StatusOK, "OK", v1.PartitionResponse{}),
					service.NewReturn(http.StatusConflict, "Conflict", httperrors.HTTPErrorResponse{}),
				},
				Handler: r.updatePartition,
			},
			{
				Method:        http.MethodDelete,
				SubPath:       "/{id}",
				PathParameter: service.NewPathParameter("id", "identifier of the partition"),
				Doc:           "deletes a partition and returns the deleted entity",
				Access:        metal.AdminAccess,
				Writes:        v1.PartitionResponse{},
				Handler:       r.deletePartition,
			},
			{
				Method:  http.MethodGet,
				SubPath: "/capacity",
				Doc:     "get partition capacities",
				Writes:  []v1.PartitionCapacity{},
				Handler: r.listPartitionCapacities,
			},
		},
	})
}
