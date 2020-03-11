package size

import (
	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	v1 "github.com/metal-stack/metal-api/pkg/proto/v1"
	"github.com/metal-stack/metal-api/pkg/util"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r *sizeResource) listSizes(request *restful.Request, response *restful.Response) {
	ss, err := r.ds.ListSizes()
	if helper.CheckError(request, response, util.CurrentFuncName(), err) {
		return
	}

	var result []*v1.SizeResponse
	for i := range ss {
		result = append(result, service.NewSizeResponse(&ss[i]))
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, result)
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}
