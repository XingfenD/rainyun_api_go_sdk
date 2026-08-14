package game

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addCreateCommand(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a game server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := &rgs.CreateRgsRequest{
				Duration:     mustFlagInt(cmd, "duration"),
				PayMode:      mustFlagString(cmd, "pay-mode"),
				Subtype:      mustFlagString(cmd, "subtype"),
				PlanID:       mustFlagInt(cmd, "plan-id"),
				EggTypeID:    mustFlagInt(cmd, "egg-type-id"),
				OsID:         mustFlagInt(cmd, "os-id"),
				WithEipNum:   mustFlagInt(cmd, "eip-num"),
				WithEipFlags: mustFlagString(cmd, "eip-flags"),
				WithEipType:  mustFlagString(cmd, "eip-type"),
				Zone:         mustFlagString(cmd, "zone"),
				PanelUser:    mustFlagString(cmd, "panel-user"),
				Config: rgs.RgsConfig{
					CPU:      mustFlagInt(cmd, "cpu"),
					Memory:   mustFlagInt(cmd, "memory"),
					NetOut:   mustFlagInt(cmd, "net-out"),
					BaseDisk: mustFlagInt(cmd, "base-disk"),
					DataDisk: mustFlagInt(cmd, "data-disk"),
				},
			}
			resp, err := (*rySDK).CreateRgs(req)
			if err != nil {
				return err
			}
			return (*out).Print(toGameServer(resp.Data.RgsRecord))
		},
	}
	createCmd.Flags().Int("duration", 1, "Duration in months (1/3/6/12, -1 for pay-as-you-go)")
	createCmd.Flags().String("pay-mode", "month", "Pay mode (day/month)")
	createCmd.Flags().String("subtype", "kvm", "Subtype (kvm/ptero/mcsm/k8s_panel)")
	createCmd.Flags().Int("plan-id", 0, "Plan ID")
	createCmd.Flags().Int("egg-type-id", 0, "Game type ID")
	createCmd.Flags().Int("os-id", 0, "OS template ID")
	createCmd.Flags().Int("cpu", 0, "CPU cores (config mode)")
	createCmd.Flags().Int("memory", 0, "Memory GB (config mode)")
	createCmd.Flags().Int("net-out", 0, "Out bandwidth Mbps (config mode)")
	createCmd.Flags().Int("base-disk", 0, "Base disk GB (config mode)")
	createCmd.Flags().Int("data-disk", 0, "Data disk GB (config mode)")
	createCmd.Flags().Int("eip-num", 0, "Extra IP count")
	createCmd.Flags().String("eip-flags", "", "Extra IP flags (us_ddosip/nb_ddosip)")
	createCmd.Flags().String("eip-type", "", "Extra IP type (IPv4/IPv6)")
	createCmd.Flags().String("zone", "", "Intranet zone")
	createCmd.Flags().String("panel-user", "", "Panel user")

	gameCmd.AddCommand(createCmd)
}
