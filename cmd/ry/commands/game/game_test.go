package game

import (
	"testing"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
)

func TestToGameServer(t *testing.T) {
	r := rgs.RgsRecord{
		ID: 1, HostName: "mc-1", Status: "RUNNING", MainIPv4: "1.2.3.4",
		CPU: 4, Memory: 8, BaseDisk: 30, DataDisk: 20, OsName: "Ubuntu 22.04",
		Zone: "cn-sq1", ExpDate: 1700000000,
	}
	g := toGameServer(r)
	if g.ID != "1" || g.Name != "mc-1" || g.Status != "RUNNING" || g.IP != "1.2.3.4" {
		t.Errorf("g = %+v", g)
	}
	if g.CPU != 4 || g.Memory != 8 || g.BaseDisk != 30 || g.DataDisk != 20 {
		t.Errorf("g = %+v", g)
	}
	if !g.ExpireAt.Equal(time.Unix(1700000000, 0)) {
		t.Errorf("ExpireAt = %v", g.ExpireAt)
	}
}

func TestToGameServerDetail(t *testing.T) {
	d := rgs.RgsDetailData{
		Data: rgs.RgsRecord{
			ID: 2, HostName: "mc-2", Status: "STOPPED", MainIPv4: "1.2.3.5",
			IntIPv4: "10.0.0.5", OsName: "Ubuntu 22.04", Zone: "cn-sq1", Tag: "t",
			CPU: 2, Memory: 4, BaseDisk: 30, DataDisk: 0, NetMode: "nat",
			NatPublicIP: "1.2.3.5", NatPublicDomain: "x.rainyun.com",
			DailyMode: true, CPULimitMode: false, McsmUserName: "u1",
			AutoRenew: true, CreateDate: 1690000000, ExpDate: 1700000000,
		},
		NatList:         []rgs.RgsNatItem{{ID: 9, PortIn: 25565, PortOut: 25565, PortType: "tcp", Tag: "mc"}},
		RenewPointPrice: rgs.RgsRenewPointPrice{Num7: 100, Num31: 400},
	}
	gd := toGameServerDetail(d)
	if gd.ID != "2" || gd.NetMode != "nat" || gd.McsmUser != "u1" || !gd.DailyMode || !gd.AutoRenew {
		t.Errorf("gd = %+v", gd)
	}
	if gd.Renew7d != 100 || gd.Renew31d != 400 {
		t.Errorf("renew points = %d/%d", gd.Renew7d, gd.Renew31d)
	}
	if len(gd.NatList) != 1 || gd.NatList[0].PortIn != 25565 {
		t.Errorf("NatList = %+v", gd.NatList)
	}
}

func TestToGamePlan(t *testing.T) {
	p := rgs.RgsPlan{ID: 1, Region: "cn-sq1", Subtype: "kvm", PlanName: "large",
		Machine: "5900X", ChargeType: "elastic_dynamic", Chinese: "Ryzen 5900X 动态计费",
		IsSelling: true, CPUPrice: 10, MemoryPrice: 10, NetOutPrice: 7}
	m := toGamePlan(p)
	if m.ID != 1 || m.Plan != "large" || m.Chinese != "Ryzen 5900X 动态计费" || m.NetOutPrice != 7 {
		t.Errorf("m = %+v", m)
	}
}

func TestToGameEggServer(t *testing.T) {
	s := rgs.RgsEggServer{Server: "ArclightFabric", EggName: "mc_fabric", Desc: "d", Order: 30}
	m := toGameEggServer(s)
	if m.Server != "ArclightFabric" || m.EggName != "mc_fabric" || m.Order != 30 {
		t.Errorf("m = %+v", m)
	}
}
