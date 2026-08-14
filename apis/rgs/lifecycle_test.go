package rgs

import "testing"

func TestSwitchRgsPanelUserPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/4/switch-user", `{"code":200,"data":"ok"}`)
	req := &SwitchRgsPanelUserRequest{Name: "u1", Password: "p1"}
	if _, err := svc.SwitchRgsPanelUser(4, req); err != nil {
		t.Fatalf("SwitchRgsPanelUser() error = %v", err)
	}
}
