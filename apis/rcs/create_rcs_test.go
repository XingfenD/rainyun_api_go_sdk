package rcs

import (
	"testing"

	"github.com/bytedance/sonic"
)

func TestCreateRcsRequest_SecurityGroupIDs(t *testing.T) {
	req := &CreateRcsRequest{
		PlanID:           int(1),
		SecurityGroupIDs: []int{10, 20},
		Duration:         int(1),
	}
	b, err := sonic.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]any
	if err := sonic.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	sg, ok := m["security_group_ids"].([]any)
	if !ok {
		t.Fatalf("security_group_ids missing or wrong type: %v", m["security_group_ids"])
	}
	if len(sg) != 2 || sg[0].(float64) != 10 || sg[1].(float64) != 20 {
		t.Fatalf("security_group_ids value mismatch: %v", sg)
	}
}

func TestRcsManagesElasticCloudDisks_MoveStorage(t *testing.T) {
	req := &RcsManagesElasticCloudDisksRequest{
		Actions: []struct {
			Type   string `json:"type"`
			Action any    `json:"action"`
		}{
			{
				Type: "move_storage",
				Action: RcsManagesElasticCloudDisksMoveStorage{
					EdiskID:      int(5),
					ToInstanceID: int(7),
				},
			},
		},
	}
	b, err := sonic.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m struct {
		Actions []struct {
			Type   string `json:"type"`
			Action struct {
				EdiskID      int `json:"edisk_id"`
				ToInstanceID int `json:"to_instance_id"`
			} `json:"action"`
		} `json:"actions"`
	}
	if err := sonic.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.Actions[0].Type != "move_storage" {
		t.Fatalf("type mismatch: %v", m.Actions[0].Type)
	}
	if m.Actions[0].Action.EdiskID != 5 || m.Actions[0].Action.ToInstanceID != 7 {
		t.Fatalf("move_storage action mismatch: %+v", m.Actions[0].Action)
	}
}
