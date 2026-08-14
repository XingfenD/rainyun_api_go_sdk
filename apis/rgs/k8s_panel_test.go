package rgs

import "testing"

func TestSetK8SPanelDatabasePath(t *testing.T) {
	svc := stubService(t, "PATCH", "/product/rgs/1/k8s-panel/database", `{"code":200,"data":"ok"}`)
	if _, err := svc.SetK8SPanelDatabase(1, &SetK8SPanelDatabaseRequest{IsEnabled: true, Version: "8.0"}); err != nil {
		t.Fatalf("SetK8SPanelDatabase() error = %v", err)
	}
}

func TestSetK8SPanelStartCommandPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/1/k8s-panel/set-start-command", `{"code":200,"data":"ok"}`)
	if _, err := svc.SetK8SPanelStartCommand(1, &SetK8SPanelStartCommandRequest{Command: "java -jar server.jar"}); err != nil {
		t.Fatalf("SetK8SPanelStartCommand() error = %v", err)
	}
}

func TestSetK8SPanelSFTPPath(t *testing.T) {
	svc := stubService(t, "PATCH", "/product/rgs/1/k8s-panel/sftp", `{"code":200,"data":"ok"}`)
	if _, err := svc.SetK8SPanelSFTP(1, &SetK8SPanelSFTPRequest{Username: "u", Password: "p"}); err != nil {
		t.Fatalf("SetK8SPanelSFTP() error = %v", err)
	}
}
