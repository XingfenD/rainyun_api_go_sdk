package rgs

import "testing"

func TestRgsToBridgePath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/6/to-bridge", `{"code":200,"data":"ok"}`)
	if _, err := svc.RgsToBridge(6); err != nil {
		t.Fatalf("RgsToBridge() error = %v", err)
	}
}

func TestRgsBridgeSetIntIPPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/6/bridge_setintip", `{"code":200,"data":"ok"}`)
	if _, err := svc.RgsBridgeSetIntIP(6, &RgsBridgeSetIntIPRequest{IP: "10.0.0.2"}); err != nil {
		t.Fatalf("RgsBridgeSetIntIP() error = %v", err)
	}
}

func TestCreateRgsVnetPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/6/vnet", `{"code":200,"data":"ok"}`)
	if _, err := svc.CreateRgsVnet(6, &RgsVnetRequest{Name: "lan1"}); err != nil {
		t.Fatalf("CreateRgsVnet() error = %v", err)
	}
}

func TestRenameRgsVnetPath(t *testing.T) {
	svc := stubService(t, "PATCH", "/product/rgs/6/vnet", `{"code":200,"data":"ok"}`)
	if _, err := svc.RenameRgsVnet(6, &RgsVnetRequest{NewName: "lan2"}); err != nil {
		t.Fatalf("RenameRgsVnet() error = %v", err)
	}
}
