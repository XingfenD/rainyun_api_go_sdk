package game

import (
	"fmt"
	"strconv"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	gameCmd := &cobra.Command{
		Use:   "game",
		Short: "Manage game servers",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List game servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRgsList(&rgs.GetRgsListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			games := make([]model.GameServer, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				games[i] = toGameServer(r)
			}
			return (*out).Print(games)
		},
	}

	getCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get game server details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRgsDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(toGameServerDetail(resp.Data))
		},
	}

	startCmd := &cobra.Command{
		Use:   "start <id>",
		Short: "Start a game server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).StartRgs(id); err != nil {
				return err
			}
			fmt.Printf("Game server %s started\n", args[0])
			return nil
		},
	}

	stopCmd := &cobra.Command{
		Use:   "stop <id>",
		Short: "Stop a game server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).StopRgs(id); err != nil {
				return err
			}
			fmt.Printf("Game server %s stopped\n", args[0])
			return nil
		},
	}

	rebootCmd := &cobra.Command{
		Use:   "reboot <id>",
		Short: "Reboot a game server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).RebootRgs(id); err != nil {
				return err
			}
			fmt.Printf("Game server %s rebooted\n", args[0])
			return nil
		},
	}

	freeCmd := &cobra.Command{
		Use:   "free <id>",
		Short: "Release a game server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).FreeRgs(id); err != nil {
				return err
			}
			fmt.Printf("Game server %s released\n", args[0])
			return nil
		},
	}

	resetPasswordCmd := &cobra.Command{
		Use:   "reset-password <id>",
		Short: "Reset game server password",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).ResetRgsPassword(id, ""); err != nil {
				return err
			}
			fmt.Printf("Password reset for game server %s\n", args[0])
			return nil
		},
	}

	vncCmd := &cobra.Command{
		Use:   "vnc <id>",
		Short: "Get VNC connection URL",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			v, err := (*rySDK).GetRgsVnc(id, "novnc")
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

	setTagCmd := &cobra.Command{
		Use:   "set-tag <id> <tag>",
		Short: "Set game server tag",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).SetRgsTag(id, args[1]); err != nil {
				return err
			}
			fmt.Printf("Tag set for game server %s\n", args[0])
			return nil
		},
	}

	reinstallCmd := &cobra.Command{
		Use:   "reinstall <id>",
		Short: "Reinstall game server OS",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			osID, _ := cmd.Flags().GetInt("os-id")
			if osID == 0 {
				return fmt.Errorf("--os-id flag is required (e.g. --os-id 123)")
			}
			if _, err := (*rySDK).ReinstallRgs(id, &rcs.ReinstallRcsRequest{OsID: osID}); err != nil {
				return err
			}
			fmt.Printf("Game server %s reinstalling with os-id %d\n", args[0], osID)
			return nil
		},
	}
	reinstallCmd.Flags().Int("os-id", 0, "OS template ID (required)")
	reinstallCmd.MarkFlagRequired("os-id")

	faiSendCmd := &cobra.Command{
		Use:   "fai-send <id>",
		Short: "Send a fast app install task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			appID, _ := cmd.Flags().GetInt("app-id")
			if appID == 0 {
				return fmt.Errorf("--app-id flag is required")
			}
			if _, err := (*rySDK).SendRgsFaiTask(id, &rgs.SendRgsFaiTaskRequest{
				AppVars: []rgs.RgsAppVar{{AppID: appID}},
			}); err != nil {
				return err
			}
			fmt.Printf("Fast app install task sent for game server %s\n", args[0])
			return nil
		},
	}
	faiSendCmd.Flags().Int("app-id", 0, "App ID to install (required)")

	gameCmd.AddCommand(listCmd)
	gameCmd.AddCommand(getCmd)
	gameCmd.AddCommand(startCmd)
	gameCmd.AddCommand(stopCmd)
	gameCmd.AddCommand(rebootCmd)
	gameCmd.AddCommand(freeCmd)
	gameCmd.AddCommand(resetPasswordCmd)
	gameCmd.AddCommand(vncCmd)
	gameCmd.AddCommand(setTagCmd)
	gameCmd.AddCommand(reinstallCmd)
	gameCmd.AddCommand(faiSendCmd)

	addCreateCommand(gameCmd, rySDK, out)
	addRenewCommands(gameCmd, rySDK, out)
	addBackupCommands(gameCmd, rySDK, out)

	return gameCmd
}

func toGameServer(r rgs.RgsRecord) model.GameServer {
	return model.GameServer{
		ID:       strconv.Itoa(r.ID),
		Name:     r.HostName,
		Status:   r.Status,
		IP:       r.MainIPv4,
		CPU:      r.CPU,
		Memory:   r.Memory,
		BaseDisk: r.BaseDisk,
		DataDisk: r.DataDisk,
		OS:       r.OsName,
		Region:   r.Zone,
		ExpireAt: time.Unix(int64(r.ExpDate), 0),
	}
}

func toGameServerDetail(d rgs.RgsDetailData) model.GameServerDetail {
	r := d.Data
	gd := model.GameServerDetail{
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
		BaseDisk:        r.BaseDisk,
		DataDisk:        r.DataDisk,
		NetMode:         r.NetMode,
		NatPublicIP:     r.NatPublicIP,
		NatPublicDomain: r.NatPublicDomain,
		DailyMode:       r.DailyMode,
		CPULimitMode:    r.CPULimitMode,
		McsmUser:        r.McsmUserName,
		AutoRenew:       r.AutoRenew,
		CreatedAt:       time.Unix(int64(r.CreateDate), 0),
		ExpireAt:        time.Unix(int64(r.ExpDate), 0),
		Renew7d:         d.RenewPointPrice.Num7,
		Renew31d:        d.RenewPointPrice.Num31,
	}
	gd.NatList = make([]model.GameNatMapping, 0, len(d.NatList))
	for _, n := range d.NatList {
		gd.NatList = append(gd.NatList, model.GameNatMapping{
			ID: n.ID, PortIn: n.PortIn, PortOut: n.PortOut, PortType: n.PortType, Tag: n.Tag,
		})
	}
	return gd
}

func mustFlagInt(cmd *cobra.Command, name string) int {
	v, _ := cmd.Flags().GetInt(name)
	return v
}

func mustFlagString(cmd *cobra.Command, name string) string {
	v, _ := cmd.Flags().GetString(name)
	return v
}
