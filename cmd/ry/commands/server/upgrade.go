package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addUpgradeCommand(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	upgradeCmd := &cobra.Command{
		Use:   "upgrade <id>",
		Short: "Upgrade server plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			plan, _ := cmd.Flags().GetInt("plan")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if _, err := (*rySDK).UpgradeRcs(id, plan, coupon); err != nil {
				return err
			}
			fmt.Printf("Server %d upgrading to plan %d\n", id, plan)
			return nil
		},
	}
	upgradeCmd.Flags().Int("plan", 0, "Destination plan ID (required)")
	upgradeCmd.Flags().Int("coupon", 0, "Coupon ID")
	upgradeCmd.MarkFlagRequired("plan")

	serverCmd.AddCommand(upgradeCmd)
}
