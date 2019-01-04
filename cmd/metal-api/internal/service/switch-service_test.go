package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.f-i-ts.de/cloud-native/metal/metal-api/cmd/metal-api/internal/datastore"
	"git.f-i-ts.de/cloud-native/metal/metal-api/cmd/metal-api/internal/metal"
	"git.f-i-ts.de/cloud-native/metal/metal-api/cmd/metal-api/internal/netbox"
	"git.f-i-ts.de/cloud-native/metal/metal-api/netbox-api/client/switches"
	nbswitch "git.f-i-ts.de/cloud-native/metal/metal-api/netbox-api/client/switches"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/require"

	restful "github.com/emicklei/go-restful"
)

func TestCreateSwitch(t *testing.T) {
	ds, mock := datastore.InitMockDB()
	metal.InitMockDBData(mock)

	nb := netbox.New()
	called := false
	nb.DoRegisterSwitch = func(params *switches.NetboxAPIProxyAPISwitchRegisterParams, authInfo runtime.ClientAuthInfoWriter) (*nbswitch.NetboxAPIProxyAPISwitchRegisterOK, error) {
		called = true
		return &nbswitch.NetboxAPIProxyAPISwitchRegisterOK{}, nil
	}
	switchservice := NewSwitch(testlogger, ds, nb)
	container := restful.NewContainer().Add(switchservice)

	js, _ := json.Marshal(metal.RegisterSwitch{
		ID:     "switch999",
		SiteID: "1",
		RackID: "1",
	})
	body := bytes.NewBuffer(js)
	req := httptest.NewRequest("POST", "/v1/switch/register", body)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	container.ServeHTTP(w, req)

	resp := w.Result()
	require.Equal(t, http.StatusCreated, resp.StatusCode, w.Body.String())
	var result metal.Switch
	err := json.NewDecoder(resp.Body).Decode(&result)
	require.Nil(t, err)
	require.True(t, called)
	require.Equal(t, "switch999", result.ID)
	require.Equal(t, "switch999", result.Name)
	require.Equal(t, "1", result.RackID)
	require.Equal(t, "1", result.SiteID)
	require.Len(t, result.Connections, 0)
}

func TestUpdateSwitch(t *testing.T) {
	ds, mock := datastore.InitMockDB()
	metal.InitMockDBData(mock)

	nb := netbox.New()
	called := false
	nb.DoRegisterSwitch = func(params *switches.NetboxAPIProxyAPISwitchRegisterParams, authInfo runtime.ClientAuthInfoWriter) (*nbswitch.NetboxAPIProxyAPISwitchRegisterOK, error) {
		called = true
		return &nbswitch.NetboxAPIProxyAPISwitchRegisterOK{}, nil
	}

	switchservice := NewSwitch(testlogger, ds, nb)
	container := restful.NewContainer().Add(switchservice)

	js, _ := json.Marshal(metal.RegisterSwitch{
		ID:     metal.Switch1.ID,
		SiteID: metal.Switch1.SiteID,
		RackID: metal.Switch1.RackID,
	})
	body := bytes.NewBuffer(js)
	req := httptest.NewRequest("POST", "/v1/switch/register", body)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	container.ServeHTTP(w, req)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode, w.Body.String())
	var result metal.Switch
	err := json.NewDecoder(resp.Body).Decode(&result)
	require.Nil(t, err)
	require.True(t, called)
	require.Equal(t, metal.Switch1.ID, result.ID)
	require.Equal(t, metal.Switch1.ID, result.Name)
	require.Equal(t, metal.Switch1.RackID, result.RackID)
	require.Equal(t, metal.Switch1.SiteID, result.SiteID)
	require.Len(t, result.Connections, 1)
	con := result.Connections[0]
	require.Equal(t, metal.Switch1.DeviceConnections["d1"][0].Nic.MacAddress, con.Nic.MacAddress)
}