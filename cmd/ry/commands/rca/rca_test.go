package rca

import (
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rca"
)

func TestToRcaProject(t *testing.T) {
	rec := rca.RcaProject{
		ID: 5, Name: "proj-a", Status: "running",
		Region:         &rca.RcaRegion{ID: 3, ChineseName: "宁波"},
		ResourceLimits: rca.RcaResourceLimits{MaxCPU: 1000, MaxMemory: 4096},
	}
	m := toRcaProject(rec)
	if m.ID != "5" || m.Name != "proj-a" || m.Region != "宁波" {
		t.Errorf("m = %+v", m)
	}
}
