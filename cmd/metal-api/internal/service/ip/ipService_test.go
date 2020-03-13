package ip

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	v1 "github.com/metal-stack/metal-api/pkg/proto/v1"
	"github.com/metal-stack/metal-api/pkg/util"
	"github.com/metal-stack/metal-lib/pkg/tag"
	"net/http"
	"net/http/httptest"
	"testing"

	mdmv1 "github.com/metal-stack/masterdata-api/api/v1"
	mdmock "github.com/metal-stack/masterdata-api/api/v1/mocks"
	mdm "github.com/metal-stack/masterdata-api/pkg/client"

	"github.com/metal-stack/metal-api/cmd/metal-api/internal/datastore"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/ipam"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/testdata"

	"github.com/metal-stack/metal-lib/httperrors"

	"github.com/google/go-cmp/cmp"
	goipam "github.com/metal-stack/go-ipam"
	"github.com/stretchr/testify/require"

	restful "github.com/emicklei/go-restful"
)

func TestGetIPs(t *testing.T) {
	ds, mock := datastore.InitMockDB()
	testdata.InitMockDBData(mock)

	ipService := NewIPService(ds, ipam.New(goipam.New()), nil)
	container := restful.NewContainer().Add(ipService)
	req := httptest.NewRequest("GET", "/v1/ip", nil)
	container = service.InjectViewer(container, req)
	w := httptest.NewRecorder()
	container.ServeHTTP(w, req)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode, w.Body.String())
	var result []v1.IPResponse
	err := json.NewDecoder(resp.Body).Decode(&result)

	require.Nil(t, err)
	require.Len(t, result, 3)
	require.Equal(t, testdata.IP1.IPAddress, result[0].Identifiable.IPAddress)
	require.Equal(t, testdata.IP1.Name, result[0].IP.Common.Name.GetValue())
	require.Equal(t, testdata.IP2.IPAddress, result[1].Identifiable.IPAddress)
	require.Equal(t, testdata.IP2.Name, result[1].IP.Common.Name.GetValue())
	require.Equal(t, testdata.IP3.IPAddress, result[2].Identifiable.IPAddress)
	require.Equal(t, testdata.IP3.Name, result[2].IP.Common.Name.GetValue())
}

func TestGetIP(t *testing.T) {
	ds, mock := datastore.InitMockDB()
	testdata.InitMockDBData(mock)

	ipService := NewIPService(ds, ipam.New(goipam.New()), nil)
	container := restful.NewContainer().Add(ipService)
	req := httptest.NewRequest("GET", "/v1/ip/1.2.3.4", nil)
	container = service.InjectViewer(container, req)
	w := httptest.NewRecorder()
	container.ServeHTTP(w, req)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode, w.Body.String())
	var result v1.IPResponse
	err := json.NewDecoder(resp.Body).Decode(&result)

	require.Nil(t, err)
	require.Equal(t, testdata.IP1.IPAddress, result.Identifiable.IPAddress)
	require.Equal(t, testdata.IP1.Name, result.IP.Common.Name.GetValue())
}

func TestGetIPNotFound(t *testing.T) {
	ds, mock := datastore.InitMockDB()
	testdata.InitMockDBData(mock)

	ipService := NewIPService(ds, ipam.New(goipam.New()), nil)
	container := restful.NewContainer().Add(ipService)
	req := httptest.NewRequest("GET", "/v1/ip/9.9.9.9", nil)
	container = service.InjectViewer(container, req)
	w := httptest.NewRecorder()
	container.ServeHTTP(w, req)

	resp := w.Result()
	require.Equal(t, http.StatusNotFound, resp.StatusCode, w.Body.String())
	var result httperrors.HTTPErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&result)

	require.Nil(t, err)
	require.Contains(t, result.Message, "9.9.9.9")
	require.Equal(t, 404, result.StatusCode)
}

func TestDeleteIP(t *testing.T) {
	ds, mock := datastore.InitMockDB()
	ipamer, err := testdata.InitMockIpamData(mock, true)
	require.Nil(t, err)
	testdata.InitMockDBData(mock)

	ipService := NewIPService(ds, ipamer, nil)
	container := restful.NewContainer().Add(ipService)

	tests := []struct {
		name         string
		ip           string
		wantedStatus int
	}{
		{
			name:         "free an ip",
			ip:           testdata.IPAMIP.IPAddress,
			wantedStatus: http.StatusOK,
		},
		{
			name:         "free an machine-ip should fail",
			ip:           testdata.IP3.IPAddress,
			wantedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:         "free an cluster-ip should fail",
			ip:           testdata.IP2.IPAddress,
			wantedStatus: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/v1/ip/free/"+testdata.IPAMIP.IPAddress, nil)
			container = service.InjectEditor(container, req)
			req.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			container.ServeHTTP(w, req)

			resp := w.Result()
			require.Equal(t, tt.wantedStatus, resp.StatusCode, w.Body.String())
			var result v1.IPResponse
			err = json.NewDecoder(resp.Body).Decode(&result)

			require.Nil(t, err)
		})
	}
}

func TestAllocateIP(t *testing.T) {
	ds, mock := datastore.InitMockDB()
	ipamer, err := testdata.InitMockIpamData(mock, false)
	require.Nil(t, err)
	testdata.InitMockDBData(mock)

	psc := mdmock.ProjectServiceClient{}
	psc.On("Get", context.Background(), &mdmv1.ProjectGetRequest{Id: "123"}).Return(&mdmv1.ProjectResponse{
		Project: &mdmv1.Project{
			Meta: &mdmv1.Meta{Id: "project-1"},
		},
	}, nil,
	)
	tsc := mdmock.TenantServiceClient{}

	mdc := mdm.NewMock(&psc, &tsc)

	ipService := NewIPService(ds, ipamer, mdc)
	container := restful.NewContainer().Add(ipService)

	tests := []struct {
		name            string
		allocateRequest v1.IPAllocateRequest
		wantedStatus    int
		wantedType      v1.IP_Type
		wantedIP        string
	}{
		{
			name: "allocate an ephemeral ip",
			allocateRequest: v1.IPAllocateRequest{
				IP: &v1.IP{
					Common: &v1.Common{
						Name:        util.StringProto(""),
						Description: util.StringProto(""),
					},
					ProjectID: "123",
					NetworkID: testdata.NwIPAM.ID,
					Type:      v1.IP_EPHEMERAL,
				},
			},
			wantedStatus: http.StatusCreated,
			wantedType:   v1.IP_EPHEMERAL,
			wantedIP:     "10.0.0.1",
		},
		{
			name: "allocate a static ip",
			allocateRequest: v1.IPAllocateRequest{
				IP: &v1.IP{
					Common: &v1.Common{
						Name:        util.StringProto(""),
						Description: util.StringProto(""),
					},
					ProjectID: "123",
					NetworkID: testdata.NwIPAM.ID,
					Type:      v1.IP_STATIC,
				},
			},
			wantedStatus: http.StatusCreated,
			wantedType:   v1.IP_STATIC,
			wantedIP:     "10.0.0.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.allocateRequest.IP.Common.Name = util.StringProto(tt.name)
			js, _ := json.Marshal(tt.allocateRequest)
			body := bytes.NewBuffer(js)
			req := httptest.NewRequest("POST", "/v1/ip/allocate", body)
			container = service.InjectEditor(container, req)
			req.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			container.ServeHTTP(w, req)
			resp := w.Result()

			require.Equal(t, tt.wantedStatus, resp.StatusCode, w.Body.String())
			var result v1.IPResponse
			err = json.NewDecoder(resp.Body).Decode(&result)

			require.Nil(t, err)
			require.Equal(t, tt.wantedType, result.IP.Type)
			require.Equal(t, tt.wantedIP, result.Identifiable.IPAddress)
			require.Equal(t, tt.name, result.IP.Common.Name.GetValue())
		})
	}
}

func TestUpdateIP(t *testing.T) {
	ds, mock := datastore.InitMockDB()
	testdata.InitMockDBData(mock)

	ipService := NewIPService(ds, ipam.New(goipam.New()), nil)
	container := restful.NewContainer().Add(ipService)
	machineIDTag1 := tag.MachineID + "=" + "1"
	tests := []struct {
		name                 string
		updateRequest        v1.IPUpdateRequest
		wantedStatus         int
		wantedIPIdentifiable *v1.IPIdentifiable
	}{
		{
			name: "update ip name",
			updateRequest: v1.IPUpdateRequest{
				Common: &v1.Common{
					Name:        util.StringProto(testdata.IP2.Name),
					Description: util.StringProto(testdata.IP2.Description),
				},
				Identifiable: &v1.IPIdentifiable{
					IPAddress: testdata.IP1.IPAddress,
				},
			},
			wantedStatus: http.StatusOK,
		},
		{
			name: "moving from ephemeral to static",
			updateRequest: v1.IPUpdateRequest{
				Common: &v1.Common{
					Name:        util.StringProto(""),
					Description: util.StringProto(""),
				},
				Identifiable: &v1.IPIdentifiable{
					IPAddress: testdata.IP1.IPAddress,
				},
				Type: v1.IP_STATIC,
			},
			wantedStatus: http.StatusOK,
		},
		{
			name: "moving from static to ephemeral must not be allowed",
			updateRequest: v1.IPUpdateRequest{
				Common: &v1.Common{
					Name:        util.StringProto(""),
					Description: util.StringProto(""),
				},
				Identifiable: &v1.IPIdentifiable{
					IPAddress: testdata.IP2.IPAddress,
				},
				Type: v1.IP_EPHEMERAL,
			},
			wantedStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "internal tag machine is allowed",
			updateRequest: v1.IPUpdateRequest{
				Common: &v1.Common{
					Name:        util.StringProto(""),
					Description: util.StringProto(""),
				},
				Identifiable: &v1.IPIdentifiable{
					IPAddress: testdata.IP3.IPAddress,
				},
				Type: v1.IP_STATIC,
				Tags: util.StringSliceProto(machineIDTag1),
			},
			wantedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js, _ := json.Marshal(tt.updateRequest)
			body := bytes.NewBuffer(js)
			req := httptest.NewRequest("POST", "/v1/ip", body)
			container = service.InjectEditor(container, req)
			req.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			container.ServeHTTP(w, req)

			resp := w.Result()
			require.Equal(t, tt.wantedStatus, resp.StatusCode, w.Body.String())
			var result v1.IPResponse
			err := json.NewDecoder(resp.Body).Decode(&result)

			require.Nil(t, err)
			if tt.wantedIPIdentifiable != nil {
				require.Equal(t, *tt.wantedIPIdentifiable, result.Identifiable)
			}
		})
	}
}

func TestProcessTags(t *testing.T) {
	tests := []struct {
		name    string
		tags    []string
		wanted  []string
		wantErr bool
	}{
		{
			name:   "distinct and sorted",
			tags:   []string{"2", "1", "2"},
			wanted: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := helper.ProcessTags(tt.tags)
			if tt.wantErr && err == nil {
				t.Fatalf("expected error")
			}
			if !cmp.Equal(got, tt.wanted) {
				t.Errorf("%v", cmp.Diff(got, tt.wanted))
			}
		})
	}
}
