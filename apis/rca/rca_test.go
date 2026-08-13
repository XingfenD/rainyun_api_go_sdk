package rca

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis"
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

func stubService(t *testing.T, wantMethod, wantPath string, body string) *RcaService {
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
	return NewRcaService(c)
}

func TestListRcaProjectsDeserialization(t *testing.T) {
	raw := `{"code":200,"data":{"TotalRecords":1,"Records":[{"ID":5,"UID":1,"Status":"running",
		"name":"proj-a","region_id":3,"charge_type":"elastic",
		"region":{"id":3,"name":"cn-nb1","chinese_name":"宁波"},
		"resource_limits":{"max_cpu":1000,"max_memory":4096},
		"usage_data":{"cpu":12,"memory":30,"app_count":1,"website_count":0},
		"ExpDate":1710000000,"AutoRenew":true}]}}`
	svc := stubService(t, "GET", "/product/rca/project/", raw)
	resp, err := svc.ListRcaProjects(`{"page":1,"perPage":50}`)
	if err != nil {
		t.Fatalf("ListRcaProjects() error = %v", err)
	}
	rec := resp.Data.Records[0]
	if rec.ID != 5 || rec.Name != "proj-a" || rec.UsageData.AppCount != 1 {
		t.Errorf("record = %+v", rec)
	}
	if rec.Region.ChineseName != "宁波" || !rec.AutoRenew {
		t.Errorf("nested unmarshal wrong: %+v", rec)
	}
}

func TestGetRcaAppListPath(t *testing.T) {
	// project_id 以查询参数传递
	c := apis.NewRyClient("test-key")
	var gotQuery string
	c.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		gotQuery = req.URL.RawQuery
		return jsonResponse(200, `{"code":200,"data":null}`), nil
	}))
	svc := NewRcaService(c)
	if _, err := svc.GetRcaAppList(7); err != nil {
		t.Fatalf("GetRcaAppList() error = %v", err)
	}
	if gotQuery != "project_id=7" {
		t.Errorf("query = %q, want project_id=7", gotQuery)
	}
}

func TestCreateRcaWebsitePath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rca/website/", `{"code":200,"data":"ok"}`)
	req := &CreateRcaWebsiteRequest{Name: "w1", Type: "reverse_proxy", ProjectID: 3,
		Domains: []string{"a.example.com"}, ReverseProxyURL: "https://x.com"}
	if _, err := svc.CreateRcaWebsite(req); err != nil {
		t.Fatalf("CreateRcaWebsite() error = %v", err)
	}
}

func TestStartRcaAppPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rca/app/9/start", `{"code":200,"data":"ok"}`)
	if _, err := svc.StartRcaApp(9); err != nil {
		t.Fatalf("StartRcaApp() error = %v", err)
	}
}
