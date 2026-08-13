package server

import (
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
)

func TestToServerDetail(t *testing.T) {
	d := rcs.RcsDetail{
		Data: rcs.RcsRecord{
			ID:          370018,
			HostName:    "instance-d0fxYXSB",
			Status:      "running",
			MainIPv4:    "171.80.11.70",
			IntIPv4:     "10.0.0.2",
			OsName:      "btpanel10ub",
			Zone:        "襄阳网络1区",
			Tag:         "prod",
			CPU:         2,
			Memory:      4096,
			Disk:        30,
			NetMode:     "bridge",
			NetIn:       100,
			NetOut:      100,
			NatPublicIP: "1.2.3.4",
			UsageData: rcs.UsageData{
				CPU:       12.5,
				UsedMem:   1024,
				FreeMem:   3072,
				MaxMem:    4096,
				NetIn:     3.5,
				NetOut:    4.5,
				SmartTemp: 40,
			},
			TrafficBytes:         1024,
			TrafficBytesToday:    512,
			TrafficBytesDayLimit: 10240,
			TrafficOnLimit:       0,
			TrafficResetDate:     1700000000,
			AutoRenew:            true,
			CreateDate:           1600000000,
			ExpDate:              1800000000,
			Plan: rcs.Plan{
				PlanName:      "通用型-2C4G",
				ChargeType:    "monthly",
				TrafficBaseGb: 1024,
			},
		},
		RenewPointPrice: rcs.RenewPointPrice{Num7: 100, Num31: 400},
		EDiskList: []rcs.EDiskItem{
			{ID: 101, Slot: 1, DiskType: "ssd", Size: 30},
			{ID: 102, Slot: 2, DiskType: "hdd", Size: 50},
		},
		EIPList: []rcs.EIPItem{
			{IP: "5.6.7.8", Region: "襄阳", Gateway: "5.6.7.1"},
		},
		RBSList: []rcs.RBSItem{
			{ID: 5, Label: "auto", FileName: "b1.tar.gz", PackSize: 100, Status: "success", CreateTime: 1600000000},
		},
		UpgradeablePlans: []rcs.UpgradeablePlan{
			{PlanName: "通用型-4C8G", CPU: 4, Memory: 8192, Price: 60},
		},
	}

	sd := toServerDetail(d)

	if sd.ID != "370018" {
		t.Errorf("ID = %q, want 370018", sd.ID)
	}
	if sd.IntranetIP != "10.0.0.2" {
		t.Errorf("IntranetIP = %q, want 10.0.0.2", sd.IntranetIP)
	}
	if sd.CPUUsage != 12.5 {
		t.Errorf("CPUUsage = %v, want 12.5", sd.CPUUsage)
	}
	if sd.Renew7d != 100 || sd.Renew31d != 400 {
		t.Errorf("Renew = %d/%d, want 100/400", sd.Renew7d, sd.Renew31d)
	}
	if sd.TrafficUsed != 1024 || sd.TrafficToday != 512 || sd.TrafficLimit != 10240 {
		t.Errorf("Traffic = %d/%d/%d, want 1024/512/10240", sd.TrafficUsed, sd.TrafficToday, sd.TrafficLimit)
	}
	if sd.AutoRenew != true {
		t.Error("AutoRenew = false, want true")
	}
	if sd.PlanName != "通用型-2C4G" {
		t.Errorf("PlanName = %q", sd.PlanName)
	}

	if want := "#101 slot1 ssd 30GB, #102 slot2 hdd 50GB"; sd.EDiskSummary != want {
		t.Errorf("EDiskSummary = %q, want %q", sd.EDiskSummary, want)
	}
	if want := "5.6.7.8"; sd.EIPSummary != want {
		t.Errorf("EIPSummary = %q, want %q", sd.EIPSummary, want)
	}
	if want := "auto(#5, success)"; sd.BackupSummary != want {
		t.Errorf("BackupSummary = %q, want %q", sd.BackupSummary, want)
	}
	if want := "通用型-4C8G cpu4/mem8192"; sd.UpgradeSummary != want {
		t.Errorf("UpgradeSummary = %q, want %q", sd.UpgradeSummary, want)
	}

	if len(sd.EDiskList) != 2 || len(sd.EIPList) != 1 || len(sd.BackupList) != 1 || len(sd.UpgradeablePlans) != 1 {
		t.Errorf("nested lists = %d/%d/%d/%d, want 2/1/1/1",
			len(sd.EDiskList), len(sd.EIPList), len(sd.BackupList), len(sd.UpgradeablePlans))
	}
}

func TestToServerDetailEmptyNested(t *testing.T) {
	sd := toServerDetail(rcs.RcsDetail{Data: rcs.RcsRecord{ID: 1}})
	if sd.EDiskSummary != "" || sd.EIPSummary != "" || sd.BackupSummary != "" || sd.UpgradeSummary != "" {
		t.Errorf("empty summaries = %q/%q/%q/%q, want all empty",
			sd.EDiskSummary, sd.EIPSummary, sd.BackupSummary, sd.UpgradeSummary)
	}
}
