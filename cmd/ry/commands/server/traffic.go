package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addTrafficCommands(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	trafficCmd := &cobra.Command{
		Use:   "traffic",
		Short: "Manage server traffic",
	}

	trafficChargeCmd := &cobra.Command{
		Use:   "charge <id>",
		Short: "Charge additional traffic",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			gb, _ := cmd.Flags().GetInt("gb")
			if _, err := (*rySDK).ChargeRcsTrafic(id, gb); err != nil {
				return err
			}
			fmt.Printf("Charged %d GB traffic for server %d\n", gb, id)
			return nil
		},
	}
	trafficChargeCmd.Flags().Int("gb", 0, "Traffic amount in GB (required)")
	trafficChargeCmd.MarkFlagRequired("gb")

	trafficLimitCmd := &cobra.Command{
		Use:   "limit <id>",
		Short: "Set server traffic limit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			dayGb, _ := cmd.Flags().GetInt("day-gb")
			bandwidth, _ := cmd.Flags().GetInt("bandwidth-mbps")
			if _, err := (*rySDK).LimitRcsTrafic(id, dayGb, bandwidth); err != nil {
				return err
			}
			fmt.Printf("Traffic limit set for server %d\n", id)
			return nil
		},
	}
	trafficLimitCmd.Flags().Int("day-gb", 0, "Daily traffic threshold in GB (required)")
	trafficLimitCmd.Flags().Int("bandwidth-mbps", 0, "Bandwidth limit in Mbps (required)")
	trafficLimitCmd.MarkFlagRequired("day-gb")
	trafficLimitCmd.MarkFlagRequired("bandwidth-mbps")

	trafficCmd.AddCommand(trafficChargeCmd, trafficLimitCmd)
	serverCmd.AddCommand(trafficCmd)
}
