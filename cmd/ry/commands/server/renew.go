package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addRenewCommands(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	renewPriceCmd := &cobra.Command{
		Use:   "renew-price <id>",
		Short: "Get server renewal price",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			duration, _ := cmd.Flags().GetInt("duration")
			coupon, _ := cmd.Flags().GetInt("coupon")
			resp, err := (*rySDK).GetRcsRenewPrice(&rcs.GetRcsRenewPriceRequest{
				ProductID:    id,
				Duration:     duration,
				WithCouponID: coupon,
			})
			if err != nil {
				return err
			}
			return (*out).Print(toRenewPrice(resp.Data))
		},
	}
	renewPriceCmd.Flags().Int("duration", 1, "Renewal duration in months")
	renewPriceCmd.Flags().Int("coupon", 0, "Coupon ID")
	renewPriceCmd.MarkFlagRequired("duration")

	renewCmd := &cobra.Command{
		Use:   "renew <id>",
		Short: "Renew a server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			duration, _ := cmd.Flags().GetInt("duration")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if _, err := (*rySDK).RenewRcs(id, rcs.RenewRcsRequest{
				Duration:     duration,
				WithCouponID: coupon,
			}); err != nil {
				return err
			}
			fmt.Printf("Server %d renewed for %d month(s)\n", id, duration)
			return nil
		},
	}
	renewCmd.Flags().Int("duration", 1, "Renewal duration in months")
	renewCmd.Flags().Int("coupon", 0, "Coupon ID")
	renewCmd.MarkFlagRequired("duration")

	autoRenewCmd := &cobra.Command{
		Use:   "auto-renew <id>",
		Short: "Enable or disable server auto renewal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			enable, _ := cmd.Flags().GetBool("enable")
			disable, _ := cmd.Flags().GetBool("disable")
			if enable == disable {
				return fmt.Errorf("specify exactly one of --enable or --disable")
			}
			if _, err := (*rySDK).EnableRcsAutoRenew(id, rcs.EnableRcsAutoRenewRequest{
				AutoRenewOption: enable,
			}); err != nil {
				return err
			}
			fmt.Printf("Auto renew %s for server %d\n", map[bool]string{true: "enabled", false: "disabled"}[enable], id)
			return nil
		},
	}
	autoRenewCmd.Flags().Bool("enable", false, "Enable auto renewal")
	autoRenewCmd.Flags().Bool("disable", false, "Disable auto renewal")

	serverCmd.AddCommand(renewPriceCmd, renewCmd, autoRenewCmd)
}

func toRenewPrice(p rcs.RcsRenewPrice) model.RenewPrice {
	return model.RenewPrice{
		Price:    p.Price,
		Renew:    p.Detail.PerScene.Renew,
		RenewEIP: p.Detail.PerScene.RenewEip,
		Coupon:   p.Detail.CouponValue,
	}
}
