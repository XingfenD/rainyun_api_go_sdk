package rgs

import (
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
)

func TestListRgsMpQuery(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs-mp/", `{"code":200,"data":null}`)
	req := &ListRgsMpRequest{Options: common.StandQueryParameters{Page: 1, PerPage: 10}}
	if _, err := svc.ListRgsMp(req); err != nil {
		t.Fatalf("ListRgsMp() error = %v", err)
	}
}

func TestCreateRgsMpBody(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs-mp/", `{"code":200,"data":null}`)
	req := &CreateRgsMpRequest{
		Duration:  1,
		EggTypeID: 2,
		DstInfo:   &RgsMpDstInfo{ClusterName: "c1"},
	}
	if _, err := svc.CreateRgsMp(req); err != nil {
		t.Fatalf("CreateRgsMp() error = %v", err)
	}
}

func TestRenewRgsMpPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs-mp/77/renew/", `{"code":200,"data":null}`)
	if _, err := svc.RenewRgsMp(77, &RenewRgsMpRequest{Duration: 1}); err != nil {
		t.Fatalf("RenewRgsMp() error = %v", err)
	}
}
