package rcdn

import (
	"testing"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcdn"
)

func TestToRcdnInstance(t *testing.T) {
	rec := rcdn.RcdnInstance{
		ID: 3, Status: "running", Tag: "cdn-1", PlanID: 1,
		CreateDate: 1700000000, ExpDate: 1710000000,
		UsageTraffic: 12345, AutoRenew: true,
		Node: rcdn.RcdnNode{Region: "mainland_china"},
		Plan: rcdn.RcdnPlan{Chinese: "入门套餐"},
	}
	m := toRcdnInstance(rec)
	if m.ID != "3" || m.Status != "running" || m.Plan != "入门套餐" {
		t.Errorf("m = %+v", m)
	}
	if m.Region != "mainland_china" || m.ExpireAt != time.Unix(1710000000, 0) {
		t.Errorf("m = %+v", m)
	}
}

func TestToRcdnDomain(t *testing.T) {
	d := rcdn.RcdnDomain{ID: 9, Domain: "a.example.com", CNAME: "x.rainyun.com", Region: "mainland_china"}
	m := toRcdnDomain(d)
	if m.ID != "9" || m.Domain != "a.example.com" || m.CNAME != "x.rainyun.com" {
		t.Errorf("m = %+v", m)
	}
}
