package rgs

import "testing"

func TestSetRgsFirewallModePath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/2/firewall/mode", `{"code":200,"data":"ok"}`)
	if _, err := svc.SetRgsFirewallMode(2, "black"); err != nil {
		t.Fatalf("SetRgsFirewallMode() error = %v", err)
	}
}

func TestGetRgsFirewallSyncTimePath(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/2/firewall/sync_time", `{"code":200,"data":null}`)
	if _, err := svc.GetRgsFirewallSyncTime(2); err != nil {
		t.Fatalf("GetRgsFirewallSyncTime() error = %v", err)
	}
}
