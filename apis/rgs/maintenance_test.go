package rgs

import "testing"

func TestGetRgsMaintenance(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/456/maintenance", `{"code":200,"data":null}`)
	resp, err := svc.GetRgsMaintenance(456)
	if err != nil {
		t.Fatalf("GetRgsMaintenance() error = %v", err)
	}
	if resp.Code != 200 {
		t.Errorf("code = %d, want 200", resp.Code)
	}
	if resp.Data != nil {
		t.Errorf("data = %v, want nil", resp.Data)
	}
}
