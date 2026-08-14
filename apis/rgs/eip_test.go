package rgs

import "testing"

func TestSetRgsEipDescriptionPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/8/eip/description", `{"code":200,"data":"ok"}`)
	if _, err := svc.SetRgsEipDescription(8, "1.2.3.4", "web"); err != nil {
		t.Fatalf("SetRgsEipDescription() error = %v", err)
	}
}
