package network

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/datastore"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	"github.com/metal-stack/metal-api/pkg/util"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r *networkResource) deleteNetwork(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")

	nw, err := r.ds.FindNetworkByID(id)
	if helper.CheckError(request, response, util.CurrentFuncName(), err) {
		return
	}

	var children metal.Networks
	err = r.ds.SearchNetworks(&datastore.NetworkSearchQuery{ParentNetworkID: &nw.ID}, &children)
	if helper.CheckError(request, response, util.CurrentFuncName(), err) {
		return
	}

	if len(children) != 0 {
		if helper.CheckError(request, response, util.CurrentFuncName(), fmt.Errorf("network cannot be deleted because there are children of this network")) {
			return
		}
	}

	allIPs, err := r.ds.ListIPs()
	if helper.CheckError(request, response, util.CurrentFuncName(), err) {
		return
	}

	err = helper.CheckAnyIPOfPrefixesInUse(allIPs, nw.Prefixes)
	if err != nil {
		if helper.CheckError(request, response, util.CurrentFuncName(), fmt.Errorf("unable to delete Network: %v", err)) {
			return
		}
	}

	for _, p := range nw.Prefixes {
		err := r.ipamer.DeletePrefix(p)
		if helper.CheckError(request, response, util.CurrentFuncName(), err) {
			return
		}
	}

	if nw.Vrf != 0 {
		err = r.ds.ReleaseUniqueInteger(nw.Vrf)
		if err != nil {
			if helper.CheckError(request, response, util.CurrentFuncName(), fmt.Errorf("could not release vrf: %v", err)) {
				return
			}
		}
	}

	err = r.ds.DeleteNetwork(nw)
	if helper.CheckError(request, response, util.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, service.NewNetworkResponse(nw, &metal.NetworkUsage{}))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}
