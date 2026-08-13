package rcdn

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

func stubService(t *testing.T, wantMethod, wantPath string, body string) *RcdnService {
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
	return NewRcdnService(c)
}

func TestGetRcdnInstanceListPath(t *testing.T) {
	svc := stubService(t, "GET", "/product/rcdn/instance", `{"code":200,"data":{"TotalRecords":0,"Records":[]}}`)
	req := &GetRcdnInstanceListRequest{Options: common.StandQueryParameters{Page: 1, PerPage: 50}}
	resp, err := svc.GetRcdnInstanceList(req)
	if err != nil {
		t.Fatalf("GetRcdnInstanceList() error = %v", err)
	}
	if resp.Code != 200 || resp.Data.TotalRecords != 0 {
		t.Errorf("resp = %+v, want code 200", resp)
	}
}

func TestGetRcdnInstanceDetailDeserialization(t *testing.T) {
	raw := `{"code":200,"data":{"id":7,"planID":1,"status":"running","tag":"cdn-a",
		"createDate":1700000000,"expDate":1710000000,"autoRenew":true,"usage_traffic":12345,
		"node":{"uuid":"n1","region":"mainland_china","chineseName":"大陆"},
		"plan":{"id":1,"chinese":"入门套餐","traffic_in_gb":100}}}`
	svc := stubService(t, "GET", "/product/rcdn/instance/7", raw)
	resp, err := svc.GetRcdnInstanceDetail(7)
	if err != nil {
		t.Fatalf("GetRcdnInstanceDetail() error = %v", err)
	}
	if resp.Data.ID != 7 || !resp.Data.AutoRenew || resp.Data.UsageTraffic != 12345 {
		t.Errorf("Data = %+v", resp.Data)
	}
	if resp.Data.Node.ChineseName != "大陆" || resp.Data.Plan.Chinese != "入门套餐" {
		t.Errorf("nested unmarshal wrong: %+v", resp.Data)
	}
}

func TestToggleRcdnDomainWafPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rcdn/domain/9/toggle_waf", `{"code":200,"data":"ok"}`)
	if _, err := svc.ToggleRcdnDomainWaf(9); err != nil {
		t.Fatalf("ToggleRcdnDomainWaf() error = %v", err)
	}
}

func TestRefreshRcdnCachePath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rcdn/instance/3/domain/9/cache_refresh", `{"code":200,"data":"ok"}`)
	req := &RefreshRcdnCacheRequest{Type: "dir", Urls: []string{"https://a.com/x"}}
	if _, err := svc.RefreshRcdnCache(3, 9, req); err != nil {
		t.Fatalf("RefreshRcdnCache() error = %v", err)
	}
}

func TestGetRcdnPlanListDeserialization(t *testing.T) {
	svc := stubService(t, "GET", "/product/rcdn/plans", `{"code":200,"data":[{"id":1,"price":9.9,"traffic_in_gb":100,"domain_limit":3}]}`)
	resp, err := svc.GetRcdnPlanList()
	if err != nil {
		t.Fatalf("GetRcdnPlanList() error = %v", err)
	}
	if len(resp.Data) != 1 || resp.Data[0].Price != 9.9 || resp.Data[0].DomainLimit != 3 {
		t.Errorf("Data = %+v", resp.Data)
	}
}
