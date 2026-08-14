package rgs

import "testing"

func TestGetMcsmUserListPath(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/mcsm/panel_user/", `{"code":200,"data":[]}`)
	resp, err := svc.GetMcsmUserList()
	if err != nil {
		t.Fatalf("GetMcsmUserList() error = %v", err)
	}
	if resp.Data == nil {
		t.Errorf("Data should be empty slice, got nil")
	}
}

func TestMcsmSftpInitPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/5/mcsm/sftp_init", `{"code":200,"data":"ok"}`)
	if _, err := svc.McsmSftpInit(5); err != nil {
		t.Fatalf("McsmSftpInit() error = %v", err)
	}
}

func TestStartMcsmInstancePath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/5/mcsm/start", `{"code":200,"data":"ok"}`)
	if _, err := svc.StartMcsmInstance(5); err != nil {
		t.Fatalf("StartMcsmInstance() error = %v", err)
	}
}

func TestGetMcsmStatusPath(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/5/mcsm/status", `{"code":200,"data":null}`)
	if _, err := svc.GetMcsmStatus(5); err != nil {
		t.Fatalf("GetMcsmStatus() error = %v", err)
	}
}
