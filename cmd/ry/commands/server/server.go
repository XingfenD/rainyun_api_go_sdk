package server

import (
	"fmt"
	"strconv"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer, raw *bool) *cobra.Command {
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
			if *raw {
				return (*out).Print(resp)
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
			s := toServer(resp.Data.Data)
			if *raw {
				return (*out).Print(resp)
			}
			return (*out).Print(s)
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
