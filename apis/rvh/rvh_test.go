package rvh

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func stubService(t *testing.T, wantMethod, wantPath string, body string) *RvhService {
	t.Helper()
	c := apis.NewRyClient("test-key")
	c.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != wantMethod {
			t.Errorf("method = %s, want %s", req.Method, wantMethod)
		}
		if req.URL.Path != wantPath {
			t.Errorf("path = %s, want %s", req.URL.Path, wantPath)
		}
		return jsonResponse(200, body), nil
	}))
	return NewRvhService(c)
}

func TestGetRvhListQuery(t *testing.T) {
	svc := stubService(t, "GET", "/product/rvh/", `{"code":200,"data":{"TotalRecords":0,"Records":[]}}`)
	req := &GetRvhListRequest{Options: common.StandQueryParameters{Page: 1, PerPage: 50}}
	if _, err := svc.GetRvhList(req); err != nil {
		t.Fatalf("GetRvhList() error = %v", err)
	}
}

func TestUnbindRvhDomainQuery(t *testing.T) {
	c := apis.NewRyClient("test-key")
	var gotQuery string
	c.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		gotQuery = req.URL.RawQuery
		return jsonResponse(200, `{"code":200,"data":"ok"}`), nil
	}))
	svc := NewRvhService(c)
	if _, err := svc.UnbindRvhDomain(5, "a.example.com"); err != nil {
		t.Fatalf("UnbindRvhDomain() error = %v", err)
	}
	if gotQuery != "domain=a.example.com" {
		t.Errorf("query = %q, want domain=a.example.com", gotQuery)
	}
}

func TestGetRvhPlanListDeserialization(t *testing.T) {
	raw := `{"code":200,"data":[{"id":1,"region":"cn-sq1","subtype":"ep","plan_name":"small",
		"charge_type":"package","chinese":"宿迁EP二代 入门版","is_selling":true,"price":6,
		"tools":2,"disk":500,"epdb":200,"bt_cpu":0,"ep_bandwidth":30,"bt_ram":0,
		"bt_net_in":0,"bt_net_out":0,"is_support_backup":true,"auto_restock":0,
		"available_stock":0,"ip_prices":{"":150},"point_renew_price":null}]}`
	svc := stubService(t, "GET", "/product/rvh/plans", raw)
	resp, err := svc.GetRvhPlanList()
	if err != nil {
		t.Fatalf("GetRvhPlanList() error = %v", err)
	}
	if len(resp.Data) != 1 || resp.Data[0].Chinese != "宿迁EP二代 入门版" || resp.Data[0].Disk != 500 {
		t.Errorf("Data = %+v", resp.Data)
	}
	if resp.Data[0].IPPrices[""] != 150 {
		t.Errorf("IPPrices = %+v", resp.Data[0].IPPrices)
	}
}

func TestCreateRvhBackupPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rvh/3/backup/", `{"code":200,"data":"ok"}`)
	if _, err := svc.CreateRvhBackup(3, CreateRvhBackupRequest{Label: "nightly"}); err != nil {
		t.Fatalf("CreateRvhBackup() error = %v", err)
	}
}

func TestSetRvhBackupSettingPath(t *testing.T) {
	svc := stubService(t, "PATCH", "/product/rvh/3/backup/setting", `{"code":200,"data":"ok"}`)
	req := SetRvhBackupSettingRequest{AutoBackupHour: 3, AutoBackupMinute: 30, KeepLast: 7}
	if _, err := svc.SetRvhBackupSetting(3, req); err != nil {
		t.Fatalf("SetRvhBackupSetting() error = %v", err)
	}
}
