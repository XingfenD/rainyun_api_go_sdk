package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addMpCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	mpCmd := &cobra.Command{
		Use:   "mp",
		Short: "Manage game server MP instances",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List MP instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).ListRgsMp(&rgs.ListRgsMpRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create an MP instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := &rgs.CreateRgsMpRequest{
				Duration:     mustFlagInt(cmd, "duration"),
				EggTypeID:    mustFlagInt(cmd, "egg-type-id"),
				PlanID:       mustFlagInt(cmd, "plan-id"),
				WithCouponID: mustFlagInt(cmd, "coupon"),
			}
			clusterName, _ := cmd.Flags().GetString("cluster-name")
			if clusterName != "" {
				req.DstInfo = &rgs.RgsMpDstInfo{
					ClusterName: clusterName,
					ClusterPass: mustFlagString(cmd, "cluster-pass"),
					Token:       mustFlagString(cmd, "token"),
				}
			}
			resp, err := (*rySDK).CreateRgsMp(req)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}
	createCmd.Flags().Int("duration", 1, "Duration in months (1/3/6/12)")
	createCmd.Flags().Int("egg-type-id", 0, "Game type ID (required)")
	createCmd.Flags().Int("plan-id", 0, "Plan ID")
	createCmd.Flags().Int("coupon", 0, "Coupon ID")
	createCmd.Flags().String("cluster-name", "", "Destination cluster name")
	createCmd.Flags().String("cluster-pass", "", "Destination cluster password")
	createCmd.Flags().String("token", "", "Destination cluster token")
	createCmd.MarkFlagRequired("egg-type-id")

	renewCmd := &cobra.Command{
		Use:   "renew <id>",
		Short: "Renew an MP instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			duration, _ := cmd.Flags().GetInt("duration")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if _, err := (*rySDK).RenewRgsMp(id, &rgs.RenewRgsMpRequest{
				Duration:     duration,
				WithCouponID: coupon,
			}); err != nil {
				return err
			}
			fmt.Printf("MP instance %d renewed for %d month(s)\n", id, duration)
			return nil
		},
	}
	renewCmd.Flags().Int("duration", 1, "Renewal duration in months")
	renewCmd.Flags().Int("coupon", 0, "Coupon ID")
	renewCmd.MarkFlagRequired("duration")

	mpCmd.AddCommand(listCmd)
	mpCmd.AddCommand(createCmd)
	mpCmd.AddCommand(renewCmd)

	gameCmd.AddCommand(mpCmd)
}
