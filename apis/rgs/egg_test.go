package rgs

import "testing"

func TestGetRgsEggServerListDeserialization(t *testing.T) {
	raw := `{"code":200,"data":[{"server":"ArclightFabric",
		"desc":"高性能Mod服务端，兼容Bukkit插件",
		"official_url":"https://github.com/IzzelAliz/Arclight/releases",
		"icon_url":"https://rainyun-public.cn-nb1.rains3.com/assets/rgs-server-icon/Arclight.png",
		"order":30,"egg_name":"mc_fabric"}]}`
	svc := stubService(t, "GET", "/product/rgs/egg_server", raw)
	resp, err := svc.GetRgsEggServerList()
	if err != nil {
		t.Fatalf("GetRgsEggServerList() error = %v", err)
	}
	if len(resp.Data) != 1 || resp.Data[0].Server != "ArclightFabric" || resp.Data[0].EggName != "mc_fabric" || resp.Data[0].Order != 30 {
		t.Errorf("Data = %+v", resp.Data)
	}
}
