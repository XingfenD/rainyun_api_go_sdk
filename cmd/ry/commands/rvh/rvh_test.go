package rvh

import (
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rvh"
)

func TestToRvhPlan(t *testing.T) {
	p := rvh.RvhPlan{ID: 1, PlanName: "small", Chinese: "宿迁EP二代 入门版", Price: 6, Disk: 500}
	m := toRvhPlan(p)
	if m.ID != 1 || m.Plan != "small" || m.Chinese != "宿迁EP二代 入门版" || m.Disk != 500 {
		t.Errorf("m = %+v", m)
	}
}
