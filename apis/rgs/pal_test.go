package rgs

import "testing"

func TestGetPalConfigPath(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/3/mcsm/pal/config", `{"code":200,"data":null}`)
	if _, err := svc.GetPalConfig(3); err != nil {
		t.Fatalf("GetPalConfig() error = %v", err)
	}
}

func TestSetPalConfigPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/3/mcsm/pal/config", `{"code":200,"data":"ok"}`)
	if _, err := svc.SetPalConfig(3, map[string]any{"max_players": 32}); err != nil {
		t.Fatalf("SetPalConfig() error = %v", err)
	}
}

func TestInitPalPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/3/mcsm/pal/init", `{"code":200,"data":"ok"}`)
	if _, err := svc.InitPal(3); err != nil {
		t.Fatalf("InitPal() error = %v", err)
	}
}

func TestGetPalLangPath(t *testing.T) {
	svc := stubService(t, "GET", "/product/rgs/3/mcsm/pal/lang", `{"code":200,"data":null}`)
	if _, err := svc.GetPalLang(3); err != nil {
		t.Fatalf("GetPalLang() error = %v", err)
	}
}

func TestPalRconPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/3/mcsm/pal/rcon", `{"code":200,"data":"ok"}`)
	if _, err := svc.PalRcon(3, "broadcast hi"); err != nil {
		t.Fatalf("PalRcon() error = %v", err)
	}
}

func TestRestartPalPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/3/mcsm/pal/restart", `{"code":200,"data":"ok"}`)
	if _, err := svc.RestartPal(3); err != nil {
		t.Fatalf("RestartPal() error = %v", err)
	}
}

func TestStopPalPath(t *testing.T) {
	svc := stubService(t, "POST", "/product/rgs/3/mcsm/pal/stop", `{"code":200,"data":"ok"}`)
	if _, err := svc.StopPal(3); err != nil {
		t.Fatalf("StopPal() error = %v", err)
	}
}
