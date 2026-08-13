package server

import (
	"testing"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
)

func TestToRenewPrice(t *testing.T) {
	p := rcs.RcsRenewPrice{
		Price: 100.5,
		Detail: rcs.RcsRenewPriceDetail{
			CouponValue: 10,
			PerScene:    rcs.RcsRenewPerScene{Renew: 80.0, RenewEip: 20.5},
		},
	}

	rp := toRenewPrice(p)

	if rp.Price != 100.5 {
		t.Errorf("Price = %v, want 100.5", rp.Price)
	}
	if rp.Renew != 80.0 {
		t.Errorf("Renew = %v, want 80.0", rp.Renew)
	}
	if rp.RenewEIP != 20.5 {
		t.Errorf("RenewEIP = %v, want 20.5", rp.RenewEIP)
	}
	if rp.Coupon != 10 {
		t.Errorf("Coupon = %d, want 10", rp.Coupon)
	}
}

func TestToMonitorSamples(t *testing.T) {
	d := rcs.RcsMonitorData{
		Columns: []string{"time", "cpu", "mem"},
		Values: [][]float64{
			{1700000000, 12.5, 50.0},
			{1700000060, 20.0, 55.5},
		},
	}

	samples := toMonitorSamples(d)

	if len(samples) != 2 {
		t.Fatalf("len(samples) = %d, want 2", len(samples))
	}
	wantTime := time.Unix(1700000000, 0).Format("2006-01-02 15:04:05")
	if samples[0].Time != wantTime {
		t.Errorf("Time = %q, want %q", samples[0].Time, wantTime)
	}
	if samples[0].Metrics != "cpu=12.50 mem=50.00" {
		t.Errorf("Metrics = %q, want %q", samples[0].Metrics, "cpu=12.50 mem=50.00")
	}
}

func TestFormatMonitorTime(t *testing.T) {
	want := time.Unix(1700000000, 0).Format("2006-01-02 15:04:05")

	if got := formatMonitorTime(1700000000); got != want {
		t.Errorf("seconds: got %q, want %q", got, want)
	}
	if got := formatMonitorTime(1700000000000); got != want {
		t.Errorf("milliseconds: got %q, want %q", got, want)
	}
}

func TestToServerEIPs(t *testing.T) {
	eips := toServerEIPs([]rcs.EIPItem{
		{ID: 11, IP: "1.2.3.4", Region: "襄阳", Type: "ipv4", Gateway: "1.2.3.1", Description: "web"},
	})

	if len(eips) != 1 {
		t.Fatalf("len(eips) = %d, want 1", len(eips))
	}
	if eips[0].ID != 11 || eips[0].IP != "1.2.3.4" || eips[0].Type != "ipv4" || eips[0].Gateway != "1.2.3.1" {
		t.Errorf("eips[0] = %+v", eips[0])
	}
}

func TestToFirewallRules(t *testing.T) {
	records := []rcs.RcsFirewallRule{
		{ID: 1, Protocol: "tcp", SourceAddress: "1.2.3.0/24", DestPort: "80", Action: "accept", IsEnable: true, Description: "web"},
		{ID: 2, Protocol: "udp", SourceAddress: "", DestPort: "53", Action: "drop", IsEnable: false},
	}

	rules := toFirewallRules(records)

	if len(rules) != 2 {
		t.Fatalf("len(rules) = %d, want 2", len(rules))
	}
	if rules[0].ID != "1" || rules[0].Protocol != "tcp" || rules[0].Source != "1.2.3.0/24" ||
		rules[0].DestPort != "80" || rules[0].Action != "accept" || rules[0].Enabled != true || rules[0].Desc != "web" {
		t.Errorf("rule[0] = %+v", rules[0])
	}
	if rules[1].Enabled != false {
		t.Errorf("rule[1].Enabled = %v, want false", rules[1].Enabled)
	}
}
