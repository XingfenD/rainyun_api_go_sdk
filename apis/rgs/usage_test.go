package rgs

import (
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
)

func TestGetRgsUsageListQuery(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/usage", `{"code":200,"data":{"TotalRecords":0,"Records":null}}`)
	req := &GetRgsUsageListRequest{Options: common.StandQueryParameters{Page: 1, PerPage: 50}}
	resp, err := svc.GetRgsUsageList(req)
	if err != nil {
		t.Fatalf("GetRgsUsageList() error = %v", err)
	}
	if resp.Data.TotalRecords != 0 {
		t.Errorf("TotalRecords = %d", resp.Data.TotalRecords)
	}
}

func TestGetRgsUsagePath(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/9/usage", `{"code":200,"data":null}`)
	if _, err := svc.GetRgsUsage(9); err != nil {
		t.Fatalf("GetRgsUsage() error = %v", err)
	}
}
