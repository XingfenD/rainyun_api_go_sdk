package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addCreateCommand(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a server",
		RunE: func(cmd *cobra.Command, args []string) error {
			planID, _ := cmd.Flags().GetInt("plan-id")
			duration, _ := cmd.Flags().GetInt("duration")
			osID, _ := cmd.Flags().GetInt("os-id")
			addDisk, _ := cmd.Flags().GetInt("add-disk-size")
			eipNum, _ := cmd.Flags().GetInt("with-eip-num")
			eipType, _ := cmd.Flags().GetString("with-eip-type")
			eipFlags, _ := cmd.Flags().GetString("with-eip-flags")
			coupon, _ := cmd.Flags().GetInt("with-coupon-id")
			try, _ := cmd.Flags().GetBool("try")
			zone, _ := cmd.Flags().GetString("zone")

			resp, err := (*rySDK).CreateRcs(&rcs.CreateRcsRequest{
				PlanID:       planID,
				Duration:     duration,
				OsID:         osID,
				AddDiskSize:  addDisk,
				WithEipNum:   eipNum,
				WithEipType:  eipType,
				WithEipFlags: eipFlags,
				WithCouponID: coupon,
				Try:          try,
				Zone:         zone,
			})
			if err != nil {
				return err
			}
			fmt.Printf("Server created (id %d)\n", resp.Data.ID)
			return nil
		},
	}

	createCmd.Flags().Int("plan-id", 0, "Plan ID (required)")
	createCmd.Flags().Int("duration", 1, "Duration in months")
	createCmd.Flags().Int("os-id", 0, "OS template ID (required)")
	createCmd.Flags().Int("add-disk-size", 0, "Extra disk size in GB")
	createCmd.Flags().Int("with-eip-num", 0, "Number of elastic IPs")
	createCmd.Flags().String("with-eip-type", "ipv4", "IP type (ipv4/ipv6)")
	createCmd.Flags().String("with-eip-flags", "", "IP flags (us_ddosip/nb_ddosip)")
	createCmd.Flags().Int("with-coupon-id", 0, "Coupon ID")
	createCmd.Flags().Bool("try", false, "Create as a trial")
	createCmd.Flags().String("zone", "", "Intranet zone")
	createCmd.MarkFlagRequired("plan-id")
	createCmd.MarkFlagRequired("os-id")

	serverCmd.AddCommand(createCmd)
}
