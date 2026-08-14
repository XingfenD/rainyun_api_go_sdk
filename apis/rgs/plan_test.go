package rgs

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

func stubService(t *testing.T, wantMethod, wantPath string, body string) *RgsService {
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
	return NewRgsService(c)
}

func TestGetRgsPlanListDeserialization(t *testing.T) {
	raw := `{"code":200,"data":[{"id":1,"region":"cn-sq1","line":"","subtype":"kvm",
		"plan_name":"large","machine":"5900X","charge_type":"elastic_dynamic",
		"chinese":"Ryzen 5900X 动态计费","is_free":false,"point_renew_price":null,
		"is_selling":true,"stock_discount":0,"eip_stock_discount":0.8,
		"ip_prices":{"":150},"ip_selling":null,"cpu_price":10,"memory_price":10,
		"net_in_price":0,"net_out_price":7,"base_disk_price":1,"data_disk_price":1,
		"config":[{"cpu":0,"memory":2,"net_in":100,"cpu_max":6,"cpu_min":1,"net_out":5,
			"base_disk":30,"data_disk":20,"base_price":30,"memory_max":4,"memory_min":2,
			"net_in_max":200,"net_in_min":100,"net_out_max":50,"net_out_min":5,
			"base_disk_max":100,"base_disk_min":30,"data_disk_max":200,"data_disk_min":20}]}]}`
	svc := stubService(t, "GET", "/product/rgs/plans", raw)
	resp, err := svc.GetRgsPlanList()
	if err != nil {
		t.Fatalf("GetRgsPlanList() error = %v", err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("len(Data) = %d, want 1", len(resp.Data))
	}
	p := resp.Data[0]
	if p.PlanName != "large" || p.Chinese != "Ryzen 5900X 动态计费" {
		t.Errorf("plan = %+v", p)
	}
	if p.EipStockDiscount != 0.8 {
		t.Errorf("EipStockDiscount = %v, want 0.8", p.EipStockDiscount)
	}
	if p.IPPrices[""] != 150 {
		t.Errorf("IPPrices = %+v", p.IPPrices)
	}
	if len(p.Config) != 1 || p.Config[0].CPUMax != 6 || p.Config[0].DataDiskMin != 20 {
		t.Errorf("Config = %+v", p.Config)
	}
}

func TestGetRgsDiscountPercentQuery(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/discount-percent", `{"code":200,"data":null}`)
	resp, err := svc.GetRgsDiscountPercent(&GetRgsDiscountPercentRequest{})
	if err != nil {
		t.Fatalf("GetRgsDiscountPercent() error = %v", err)
	}
	if resp.Code != 200 {
		t.Errorf("Code = %d", resp.Code)
	}
}
