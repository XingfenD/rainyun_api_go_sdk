package rgs

import "testing"

func TestGetPteroUserListPath(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/ptero/panel_user/", `{"code":200,"data":[]}`)
	if _, err := svc.GetPteroUserList(); err != nil {
		t.Fatalf("GetPteroUserList() error = %v", err)
	}
}

func TestCreatePteroUserPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/ptero/panel_user", `{"code":200,"data":"ok"}`)
	if _, err := svc.CreatePteroUser("u1", "p1"); err != nil {
		t.Fatalf("CreatePteroUser() error = %v", err)
	}
}

func TestEditPteroUserPath(t *testing.T) {
	svc := stubService(t, "PATCH", "/product/rgs/ptero/panel_user", `{"code":200,"data":"ok"}`)
	if _, err := svc.EditPteroUser("u1", "p2"); err != nil {
		t.Fatalf("EditPteroUser() error = %v", err)
	}
}

func TestDeletePteroUserPath(t *testing.T) {
	svc := stubService(t, "DELETE", "/product/rgs/ptero/panel_user/u1", `{"code":200,"data":"ok"}`)
	if _, err := svc.DeletePteroUser("u1"); err != nil {
		t.Fatalf("DeletePteroUser() error = %v", err)
	}
}
