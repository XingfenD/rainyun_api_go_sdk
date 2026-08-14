package rgs

import "testing"

func TestSendRgsFaiTaskPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/8/fai-send", `{"code":200,"data":"ok"}`)
	req := &SendRgsFaiTaskRequest{AppVars: []RgsAppVar{{AppID: 1}}}
	if _, err := svc.SendRgsFaiTask(8, req); err != nil {
		t.Fatalf("SendRgsFaiTask() error = %v", err)
	}
}
