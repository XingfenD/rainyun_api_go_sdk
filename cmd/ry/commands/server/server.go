package server

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Manage cloud servers",
	}

	serverListCmd := &cobra.Command{
		Use:   "list",
		Short: "List servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRcsList(&rcs.GetRcsListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			servers := make([]model.Server, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				servers[i] = toServer(r)
			}
			return (*out).Print(servers)
		},
	}

	serverGetCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get server details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid server id: %s", args[0])
			}
			resp, err := (*rySDK).GetRcsDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(toServerDetail(resp.Data))
		},
	}

	serverStartCmd := &cobra.Command{
		Use:   "start <id>",
		Short: "Start a server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid server id: %s", args[0])
			}
			if _, err := (*rySDK).StartRcs(id); err != nil {
				return err
			}
			fmt.Printf("Server %s started\n", args[0])
			return nil
		},
	}

	serverStopCmd := &cobra.Command{
		Use:   "stop <id>",
		Short: "Stop a server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid server id: %s", args[0])
			}
			if _, err := (*rySDK).StopRcs(id); err != nil {
				return err
			}
			fmt.Printf("Server %s stopped\n", args[0])
			return nil
		},
	}

	serverRebootCmd := &cobra.Command{
		Use:   "reboot <id>",
		Short: "Reboot a server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid server id: %s", args[0])
			}
			if _, err := (*rySDK).RebootRcs(id); err != nil {
				return err
			}
			fmt.Printf("Server %s rebooted\n", args[0])
			return nil
		},
	}

	serverReinstallCmd := &cobra.Command{
		Use:   "reinstall <id>",
		Short: "Reinstall server OS",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid server id: %s", args[0])
			}
			osID, _ := cmd.Flags().GetInt("os-id")
			if osID == 0 {
				return fmt.Errorf("--os-id flag is required (e.g. --os-id 123)")
			}
			if _, err := (*rySDK).ReinstallRcs(id, &rcs.ReinstallRcsRequest{OsID: osID}); err != nil {
				return err
			}
			fmt.Printf("Server %s reinstalling with os-id %d\n", args[0], osID)
			return nil
		},
	}

	serverResetPasswordCmd := &cobra.Command{
		Use:   "reset-password <id>",
		Short: "Reset server password",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid server id: %s", args[0])
			}
			if _, err := (*rySDK).ResetRcsPassword(id, ""); err != nil {
				return err
			}
			fmt.Printf("Password reset for server %s\n", args[0])
			return nil
		},
	}

	serverVNCCmd := &cobra.Command{
		Use:   "vnc <id>",
		Short: "Get VNC connection URL",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid server id: %s", args[0])
			}
			v, err := (*rySDK).GetRcsVnc(id, "novnc")
			if err != nil {
				return err
			}
			url, uerr := common.GetVncConnectURL(v)
			if uerr != nil || url == "" {
				fmt.Println(v.Data.VNCProxyURL)
			} else {
				fmt.Println(url)
			}
			return nil
		},
	}

	serverReinstallCmd.Flags().Int("os-id", 0, "OS template ID (required)")
	serverReinstallCmd.MarkFlagRequired("os-id")

	serverCmd.AddCommand(serverListCmd)
	serverCmd.AddCommand(serverGetCmd)
	serverCmd.AddCommand(serverStartCmd)
	serverCmd.AddCommand(serverStopCmd)
	serverCmd.AddCommand(serverRebootCmd)
	serverCmd.AddCommand(serverReinstallCmd)
	serverCmd.AddCommand(serverResetPasswordCmd)
	serverCmd.AddCommand(serverVNCCmd)

	addLifecycleCommands(serverCmd, rySDK, out)
	addCreateCommand(serverCmd, rySDK, out)
	addRenewCommands(serverCmd, rySDK, out)
	addUpgradeCommand(serverCmd, rySDK, out)
	addEDiskCommands(serverCmd, rySDK, out)
	addMonitorCommand(serverCmd, rySDK, out)
	addBackupCommands(serverCmd, rySDK, out)
	addEIPCommands(serverCmd, rySDK, out)

	return serverCmd
}

func toServer(r rcs.RcsRecord) model.Server {
	return model.Server{
		ID:       strconv.Itoa(r.ID),
		Name:     r.HostName,
		Status:   r.Status,
		IP:       r.MainIPv4,
		CPU:      r.CPU,
		Memory:   r.Memory,
		Disk:     r.Disk,
		OS:       r.OsName,
		Region:   r.Zone,
		ExpireAt: time.Unix(int64(r.ExpDate), 0),
	}
}

func toServerDetail(d rcs.RcsDetail) model.ServerDetail {
	r := d.Data
	sd := model.ServerDetail{
		ID:              strconv.Itoa(r.ID),
		Name:            r.HostName,
		Status:          r.Status,
		IP:              r.MainIPv4,
		IntranetIP:      r.IntIPv4,
		OS:              r.OsName,
		Region:          r.Zone,
		Tag:             r.Tag,
		CPU:             r.CPU,
		Memory:          r.Memory,
		Disk:            r.Disk,
		NetMode:         r.NetMode,
		BandwidthIn:     r.NetIn,
		BandwidthOut:    r.NetOut,
		NatPublicIP:     r.NatPublicIP,
		NatPublicDomain: r.NatPublicDomain,
		CPUUsage:        r.UsageData.CPU,
		UsedMem:         r.UsageData.UsedMem,
		FreeMem:         r.UsageData.FreeMem,
		MaxMem:          r.UsageData.MaxMem,
		NetInNow:        r.UsageData.NetIn,
		NetOutNow:       r.UsageData.NetOut,
		DiskTemp:        r.UsageData.SmartTemp,
		TrafficUsed:     r.TrafficBytes,
		TrafficToday:    r.TrafficBytesToday,
		TrafficLimit:    r.TrafficBytesDayLimit,
		TrafficOnLimit:  r.TrafficOnLimit,
		TrafficResetAt:  unixPtr(r.TrafficResetDate),
		AutoRenew:       r.AutoRenew,
		CreatedAt:       time.Unix(int64(r.CreateDate), 0),
		ExpireAt:        time.Unix(int64(r.ExpDate), 0),
		PlanName:        r.Plan.PlanName,
		ChargeType:      r.Plan.ChargeType,
		BaseTraffic:     r.Plan.TrafficBaseGb,
		Renew7d:         d.RenewPointPrice.Num7,
		Renew31d:        d.RenewPointPrice.Num31,
	}

	sd.EDiskList = make([]model.ServerEDisk, 0, len(d.EDiskList))
	for _, e := range d.EDiskList {
		sd.EDiskList = append(sd.EDiskList, model.ServerEDisk{
			Slot: e.Slot, Type: e.DiskType, Size: e.Size, Backup: e.Backup,
		})
	}
	sd.EIPList = make([]model.ServerEIP, 0, len(d.EIPList))
	for _, e := range d.EIPList {
		sd.EIPList = append(sd.EIPList, model.ServerEIP{
			IP: e.IP, Region: e.Region, Gateway: e.Gateway, Description: e.Description,
		})
	}
	sd.BackupList = make([]model.ServerBackup, 0, len(d.RBSList))
	for _, b := range d.RBSList {
		sd.BackupList = append(sd.BackupList, model.ServerBackup{
			Label:      b.Label,
			FileName:   b.FileName,
			SizeBytes:  b.PackSize,
			Status:     b.Status,
			CreatedAt:  time.Unix(int64(b.CreateTime), 0),
			FinishedAt: time.Unix(int64(b.FinishTime), 0),
		})
	}
	sd.UpgradeablePlans = make([]model.ServerPlan, 0, len(d.UpgradeablePlans))
	for _, p := range d.UpgradeablePlans {
		sd.UpgradeablePlans = append(sd.UpgradeablePlans, model.ServerPlan{
			Name: p.PlanName, CPU: p.CPU, Memory: p.Memory, Price: p.Price,
		})
	}

	sd.EDiskSummary = summarizeEDisks(sd.EDiskList)
	sd.EIPSummary = summarizeEIPs(sd.EIPList)
	sd.BackupSummary = summarizeBackups(sd.BackupList)
	sd.UpgradeSummary = summarizePlans(sd.UpgradeablePlans)

	return sd
}

func unixPtr(sec int) *time.Time {
	if sec == 0 {
		return nil
	}
	t := time.Unix(int64(sec), 0)
	return &t
}

func summarizeEDisks(items []model.ServerEDisk) string {
	parts := make([]string, 0, len(items))
	for _, e := range items {
		parts = append(parts, fmt.Sprintf("slot%d %s %dGB", e.Slot, e.Type, e.Size))
	}
	return strings.Join(parts, ", ")
}

func summarizeEIPs(items []model.ServerEIP) string {
	parts := make([]string, 0, len(items))
	for _, e := range items {
		parts = append(parts, e.IP)
	}
	return strings.Join(parts, ", ")
}

func summarizeBackups(items []model.ServerBackup) string {
	parts := make([]string, 0, len(items))
	for _, b := range items {
		parts = append(parts, fmt.Sprintf("%s(%s)", b.Label, b.Status))
	}
	return strings.Join(parts, ", ")
}

func summarizePlans(items []model.ServerPlan) string {
	parts := make([]string, 0, len(items))
	for _, p := range items {
		parts = append(parts, fmt.Sprintf("%s cpu%d/mem%d", p.Name, p.CPU, p.Memory))
	}
	return strings.Join(parts, ", ")
}
