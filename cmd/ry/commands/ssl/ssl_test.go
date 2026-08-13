package ssl

import (
	"strconv"
	"testing"
	"time"

	apisssl "github.com/XingfenD/rainyun_api_go_sdk/apis/ssl"

	"github.com/spf13/cobra"
)

func TestToSslCertRecord(t *testing.T) {
	r := apisssl.SslCertificateRecord{ID: 30136, Domain: "a.example.com", Issuer: "TrustAsia", StartDate: 1784084400, ExpDate: 1791860399}
	m := toSslCertRecord(r)
	if m.ID != 30136 || m.Start != "2026-07-15" || m.Expire != "2026-10-13" {
		t.Errorf("m = %+v", m)
	}
}

func TestToSslOrder(t *testing.T) {
	o := apisssl.SslOrder{
		ID: 5951, Status: "issued", Price: 0, CertExpireAt: 1791860399,
		CsrInfo: apisssl.SslOrderCsrInfo{CommonName: "a.example.com", DNSNames: []string{"b.example.com"}},
		Product: apisssl.SslProduct{Name: "TrustAsia免费SSL证书"},
	}
	m := toSslOrder(o)
	if m.Domain != "a.example.com, b.example.com" || m.Expire != "2026-10-13" {
		t.Errorf("m = %+v", m)
	}
	d := toSslOrderDetail(o)
	if d.Product != "TrustAsia免费SSL证书" {
		t.Errorf("d = %+v", d)
	}
	wantRemain := strconv.Itoa(int(time.Until(time.Unix(1791860399, 0)).Hours() / 24))
	if d.Remain != wantRemain {
		t.Errorf("Remain = %s, want %s", d.Remain, wantRemain)
	}
}

func TestToSslProduct(t *testing.T) {
	p := apisssl.SslProduct{ID: 13, Name: "TrustAsia域名型SSL证书", Type: "dv", Brand: "trustasia",
		PriceMap: map[string]int{"12": 559}, OriginalPriceMap: map[string]int{"12": 699}}
	m := toSslProduct(p)
	if m.ID != 13 || m.Price12 != 559 || m.Original != 699 {
		t.Errorf("m = %+v", m)
	}
}

func TestOrderFromFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addOrderFlags(cmd)
	cmd.Flags().Set("domains", "a.example.com")
	cmd.Flags().Set("duration", "12")
	cmd.Flags().Set("product-id", "3")
	cmd.Flags().Set("price", "9.9")
	cmd.Flags().Set("coupon", "7")

	req, err := orderFromFlags(cmd)
	if err != nil {
		t.Fatalf("orderFromFlags() error = %v", err)
	}
	if req.Domains != "a.example.com" || req.Duration != 12 || req.ProductID != 3 ||
		req.Price != 9.9 || req.WithCouponID != 7 {
		t.Errorf("req = %+v", req)
	}
}

func TestOrderFromFlagsMissingRequired(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addOrderFlags(cmd)
	cmd.Flags().Set("domains", "a.example.com")

	if _, err := orderFromFlags(cmd); err == nil {
		t.Fatal("orderFromFlags() error = nil, want missing flags error")
	}
}
