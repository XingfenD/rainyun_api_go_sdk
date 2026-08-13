package ssl

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

func stubService(t *testing.T, wantMethod, wantPath string, body string) *SslService {
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
	return NewSslService(c)
}

func TestGetSSLOrderListQuery(t *testing.T) {
	raw := `{"code":200,"data":{"totalRecords":1,"records":[{"id":5951,"uid":1175873,
		"csrInfo":{"commonName":"yoresee.cc","dnsNames":[],"keyAlgo":"RSA","keyLen":4096,"signHash":"SHA256","country":"CN"},
		"product":{"id":29,"name":"TrustAsia免费SSL证书","verifyMethods":["dns","file"],"algorithms":["RSA","ECC"],
		"priceMap":{"3":0},"originalPriceMap":{"3":0},"isOnSale":true,"provider":"certcloud",
		"providerProductId":"trustasia_d5_ssl_dv_free","providerCaa":"digicert.com","multiDomain":true,
		"wildcard":false,"type":"dv","brand":"trustasia"},
		"duration":3,"price":0,"status":"issued",
		"validationRecords":[{"ID":42233,"domain":"yoresee.cc","host":"_dnsauth","value":"x","type":"TXT","purpose":"issue"}],
		"description":null,"rewardToBeCollect":0,"unsubscribeAble":false,"unsubscribeTime":0,
		"unsubscribeReason":null,"notUnsubscribeAbleReason":null,
		"certIssuedAt":1784088638,"certStartAt":1784084400,"certExpireAt":1791860399,
		"createdAt":1784088639,"updatedAt":1784088725}]}}`
	svc := stubService(t, "GET", "/product/sslcenter/order", raw)
	req := &GetSSLOrderListRequest{Options: common.StandQueryParameters{Page: 1, PerPage: 50}}
	resp, err := svc.GetSSLOrderList(req)
	if err != nil {
		t.Fatalf("GetSSLOrderList() error = %v", err)
	}
	if resp.Data.TotalRecords != 1 || len(resp.Data.Records) != 1 {
		t.Fatalf("Data = %+v", resp.Data)
	}
	o := resp.Data.Records[0]
	if o.ID != 5951 || o.Status != "issued" || o.CsrInfo.CommonName != "yoresee.cc" ||
		o.Product.Brand != "trustasia" || len(o.ValidationRecords) != 1 ||
		o.CertExpireAt != 1791860399 {
		t.Errorf("order = %+v", o)
	}
}

func TestApplyFreeSSLCertificatePath(t *testing.T) {
	svc := stubService(t, "POST", "/product/sslcenter/cert/order", `{"code":200,"data":"ok"}`)
	req := &ApplyFreeSSLCertRequest{Domains: "a.example.com", VerifyMethod: "auto"}
	if _, err := svc.ApplyFreeSSLCertificate(req); err != nil {
		t.Fatalf("ApplyFreeSSLCertificate() error = %v", err)
	}
}

func TestVerifyFreeSSLCertificatePath(t *testing.T) {
	svc := stubService(t, "POST", "/product/sslcenter/cert/order_verify", `{"code":200,"data":"ok"}`)
	if _, err := svc.VerifyFreeSSLCertificate(7); err != nil {
		t.Fatalf("VerifyFreeSSLCertificate() error = %v", err)
	}
}

func TestGetSSLCertApplyListQuery(t *testing.T) {
	svc := stubService(t, "GET", "/product/sslcenter/cert/orders", `{"code":200,"data":{"TotalRecords":0,"Records":[]}}`)
	req := &GetSSLOrderListRequest{Options: common.StandQueryParameters{Page: 1, PerPage: 50}}
	if _, err := svc.GetSSLCertApplyList(req); err != nil {
		t.Fatalf("GetSSLCertApplyList() error = %v", err)
	}
}

func TestCreateSSLOrderPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/sslcenter/order", `{"code":200,"data":"ok"}`)
	req := &CreateSSLOrderRequest{Domains: "a.example.com", Duration: 12, Price: 9.9, ProductID: 1}
	if _, err := svc.CreateSSLOrder(req); err != nil {
		t.Fatalf("CreateSSLOrder() error = %v", err)
	}
}

func TestGetSSLOrderDetailPath(t *testing.T) {
	raw := `{"code":200,"data":{"id":5951,"status":"issued","duration":3,"price":0,
		"certStartAt":1784084400,"certExpireAt":1791860399}}`
	svc := stubService(t, "GET", "/product/sslcenter/order/3", raw)
	resp, err := svc.GetSSLOrderDetail(3)
	if err != nil {
		t.Fatalf("GetSSLOrderDetail() error = %v", err)
	}
	if resp.Data.ID != 5951 || resp.Data.Status != "issued" {
		t.Errorf("Data = %+v", resp.Data)
	}
}

func TestAssignSSLOrderPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/sslcenter/order/3/assign", `{"code":200,"data":"ok"}`)
	if _, err := svc.AssignSSLOrder(3); err != nil {
		t.Fatalf("AssignSSLOrder() error = %v", err)
	}
}

func TestGetSSLOrderCertPath(t *testing.T) {
	raw := `{"code":200,"data":{"Cert":"CERT","Key":"KEY","DomainName":"a.example.com","Issuer":"TrustAsia",
		"StartDate":1784084400,"ExpDate":1791860399,"RemainDays":89}}`
	svc := stubService(t, "GET", "/product/sslcenter/order/3/cert", raw)
	resp, err := svc.GetSSLOrderCert(3)
	if err != nil {
		t.Fatalf("GetSSLOrderCert() error = %v", err)
	}
	if resp.Data.Cert != "CERT" || resp.Data.Key != "KEY" || resp.Data.RemainDays != 89 {
		t.Errorf("Data = %+v", resp.Data)
	}
}

func TestUpdateSSLOrderDescriptionPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/sslcenter/order/3/description", `{"code":200,"data":"ok"}`)
	if _, err := svc.UpdateSSLOrderDescription(3, "主站证书"); err != nil {
		t.Fatalf("UpdateSSLOrderDescription() error = %v", err)
	}
}

func TestRevokeSSLOrderPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/sslcenter/order/3/revoke", `{"code":200,"data":"ok"}`)
	if _, err := svc.RevokeSSLOrder(3, "不再需要", ""); err != nil {
		t.Fatalf("RevokeSSLOrder() error = %v", err)
	}
}

func TestVerifySSLOrderPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/sslcenter/order/3/verify", `{"code":200,"data":"ok"}`)
	if _, err := svc.VerifySSLOrder(3, true); err != nil {
		t.Fatalf("VerifySSLOrder() error = %v", err)
	}
}

func TestGetSSLOrderPricePath(t *testing.T) {
	svc := stubService(t, "POST", "/product/sslcenter/price", `{"code":200,"data":{"price":559,"reward":100.62,"rewardPoints":201240}}`)
	req := &CreateSSLOrderRequest{Domains: "a.example.com", Duration: 12, ProductID: 13, Price: 559}
	resp, err := svc.GetSSLOrderPrice(req)
	if err != nil {
		t.Fatalf("GetSSLOrderPrice() error = %v", err)
	}
	if resp.Data.Price != 559 || resp.Data.Reward != 100.62 || resp.Data.RewardPoints != 201240 {
		t.Errorf("Data = %+v", resp.Data)
	}
}

func TestGetSSLProductListPath(t *testing.T) {
	raw := `{"code":200,"data":[{"id":13,"name":"TrustAsia域名型SSL证书","verifyMethods":["dns","file"],
		"algorithms":["RSA","ECC"],"priceMap":{"12":559},"originalPriceMap":{"12":699},
		"isOnSale":true,"provider":"certcloud","providerProductId":"trustasia_d5_ssl_dv",
		"providerCaa":"digicert.com","multiDomain":false,"wildcard":false,"type":"dv","brand":"trustasia"}]}`
	svc := stubService(t, "GET", "/product/sslcenter/product", raw)
	resp, err := svc.GetSSLProductList()
	if err != nil {
		t.Fatalf("GetSSLProductList() error = %v", err)
	}
	if len(resp.Data) != 1 || resp.Data[0].ID != 13 || resp.Data[0].PriceMap["12"] != 559 {
		t.Errorf("Data = %+v", resp.Data)
	}
}
