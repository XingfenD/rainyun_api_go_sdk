package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addRenewCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	renewCmd := &cobra.Command{
		Use:   "renew <id>",
		Short: "Renew a game server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			duration, _ := cmd.Flags().GetInt("duration")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if _, err := (*rySDK).RenewRgs(id, rcs.RenewRcsRequest{
				Duration:     duration,
				WithCouponID: coupon,
			}); err != nil {
				return err
			}
			fmt.Printf("Game server %d renewed for %d month(s)\n", id, duration)
			return nil
		},
	}
	renewCmd.Flags().Int("duration", 1, "Renewal duration in months")
	renewCmd.Flags().Int("coupon", 0, "Coupon ID")
	renewCmd.MarkFlagRequired("duration")

	autoRenewCmd := &cobra.Command{
		Use:   "auto-renew <id>",
		Short: "Enable or disable game server auto renewal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			enable, _ := cmd.Flags().GetBool("enable")
			disable, _ := cmd.Flags().GetBool("disable")
			if enable == disable {
				return fmt.Errorf("specify exactly one of --enable or --disable")
			}
			if _, err := (*rySDK).EnableRgsAutoRenew(id, rcs.EnableRcsAutoRenewRequest{
				AutoRenewOption: enable,
			}); err != nil {
				return err
			}
			state := "disabled"
			if enable {
				state = "enabled"
			}
			fmt.Printf("Auto renew %s for game server %d\n", state, id)
			return nil
		},
	}
	autoRenewCmd.Flags().Bool("enable", false, "Enable auto renewal")
	autoRenewCmd.Flags().Bool("disable", false, "Disable auto renewal")

	scaleCmd := &cobra.Command{
		Use:   "scale <id>",
		Short: "Scale game server configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			req := &rgs.ScaleRgsRequest{
				DestPlan: mustFlagInt(cmd, "plan-id"),
				DestConfig: rgs.RgsConfig{
					CPU:      mustFlagInt(cmd, "cpu"),
					Memory:   mustFlagInt(cmd, "memory"),
					NetOut:   mustFlagInt(cmd, "net-out"),
					BaseDisk: mustFlagInt(cmd, "base-disk"),
					DataDisk: mustFlagInt(cmd, "data-disk"),
				},
			}
			if _, err := (*rySDK).ScaleRgs(id, req); err != nil {
				return err
			}
			fmt.Printf("Game server %d scaling\n", id)
			return nil
		},
	}
	scaleCmd.Flags().Int("plan-id", 0, "Destination plan ID (0 keeps current)")
	scaleCmd.Flags().Int("cpu", 0, "CPU cores")
	scaleCmd.Flags().Int("memory", 0, "Memory GB")
	scaleCmd.Flags().Int("net-out", 0, "Out bandwidth Mbps")
	scaleCmd.Flags().Int("base-disk", 0, "Base disk GB")
	scaleCmd.Flags().Int("data-disk", 0, "Data disk GB")

	renewPriceCmd := &cobra.Command{
		Use:   "renew-price <id>",
		Short: "Get game server renewal price",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			duration, _ := cmd.Flags().GetInt("duration")
			coupon, _ := cmd.Flags().GetInt("coupon")
			resp, err := (*rySDK).GetRgsRenewPrice(&rgs.GetRgsRenewPriceRequest{
				ProductID:    id,
				Duration:     duration,
				WithCouponID: coupon,
			})
			if err != nil {
				return err
			}
			return (*out).Print(toGameRenewPrice(resp.Data))
		},
	}
	renewPriceCmd.Flags().Int("duration", 1, "Renewal duration in months")
	renewPriceCmd.Flags().Int("coupon", 0, "Coupon ID")
	renewPriceCmd.MarkFlagRequired("duration")

	upgradePriceCmd := &cobra.Command{
		Use:   "upgrade-price <id>",
		Short: "Get game server upgrade price",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			duration, _ := cmd.Flags().GetInt("duration")
			coupon, _ := cmd.Flags().GetInt("coupon")
			resp, err := (*rySDK).GetRgsUpgradePrice(&rgs.GetRgsUpgradePriceRequest{
				ProductID:    id,
				Duration:     duration,
				WithCouponID: coupon,
				Config: rgs.RgsConfig{
					CPU:      mustFlagInt(cmd, "cpu"),
					Memory:   mustFlagInt(cmd, "memory"),
					NetOut:   mustFlagInt(cmd, "net-out"),
					BaseDisk: mustFlagInt(cmd, "base-disk"),
					DataDisk: mustFlagInt(cmd, "data-disk"),
				},
			})
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}
	upgradePriceCmd.Flags().Int("duration", 1, "Duration in months")
	upgradePriceCmd.Flags().Int("coupon", 0, "Coupon ID")
	upgradePriceCmd.Flags().Int("cpu", 0, "Target CPU cores")
	upgradePriceCmd.Flags().Int("memory", 0, "Target memory GB")
	upgradePriceCmd.Flags().Int("net-out", 0, "Target out bandwidth Mbps")
	upgradePriceCmd.Flags().Int("base-disk", 0, "Target base disk GB")
	upgradePriceCmd.Flags().Int("data-disk", 0, "Target data disk GB")
	upgradePriceCmd.MarkFlagRequired("duration")

	gameCmd.AddCommand(renewCmd)
	gameCmd.AddCommand(autoRenewCmd)
	gameCmd.AddCommand(scaleCmd)
	gameCmd.AddCommand(renewPriceCmd)
	gameCmd.AddCommand(upgradePriceCmd)
}

func toGameRenewPrice(p rcs.RcsRenewPrice) model.RenewPrice {
	return model.RenewPrice{
		Price:    p.Price,
		Renew:    p.Detail.PerScene.Renew,
		RenewEIP: p.Detail.PerScene.RenewEip,
		Coupon:   p.Detail.CouponValue,
	}
}
