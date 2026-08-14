package rgs

import "testing"

func TestScaleRgsPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/10/scale", `{"code":200,"data":null}`)
	req := &ScaleRgsRequest{
		DestConfig: RgsConfig{CPU: 4, Memory: 8, NetOut: 10, BaseDisk: 30},
		DestPlan:   1,
	}
	if _, err := svc.ScaleRgs(10, req); err != nil {
		t.Fatalf("ScaleRgs() error = %v", err)
	}
}
