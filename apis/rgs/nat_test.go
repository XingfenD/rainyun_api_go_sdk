package rgs

import "testing"

func TestDeleteRgsNatPortMappingQuery(t *testing.T) {
	svc := stubService(t, "DELETE", "/product/rgs/8/nat", `{"code":200,"data":"ok"}`)
	if _, err := svc.DeleteRgsNatPortMapping(8, &DeleteRgsNatPortMappingRequest{NatID: 3}); err != nil {
		t.Fatalf("DeleteRgsNatPortMapping() error = %v", err)
	}
}
