package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addBillingCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	cpuChargeCmd := &cobra.Command{
		Use:   "cpu-charge <id>",
		Short: "Charge CPU points for a game server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			mode, _ := cmd.Flags().GetString("mode")
			money, _ := cmd.Flags().GetInt("money")
			points, _ := cmd.Flags().GetInt("points")
			if _, err := (*rySDK).ChargeRgsCPU(id, &rgs.ChargeRgsCPURequest{
				Mode:   mode,
				Money:  money,
				Points: points,
			}); err != nil {
				return err
			}
			fmt.Printf("CPU charge submitted for game server %d\n", id)
			return nil
		},
	}
	cpuChargeCmd.Flags().String("mode", "money", "Pay mode (money/point)")
	cpuChargeCmd.Flags().Int("money", 0, "Money to spend (mode=money)")
	cpuChargeCmd.Flags().Int("points", 0, "Points to spend (mode=point)")

	dailyModeCmd := &cobra.Command{
		Use:   "daily-mode <id>",
		Short: "Toggle daily billing mode",
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
			if _, err := (*rySDK).SwitchRgsDailyMode(id, enable); err != nil {
				return err
			}
			state := "disabled"
			if enable {
				state = "enabled"
			}
			fmt.Printf("Daily mode %s for game server %d\n", state, id)
			return nil
		},
	}
	dailyModeCmd.Flags().Bool("enable", false, "Enable daily mode")
	dailyModeCmd.Flags().Bool("disable", false, "Disable daily mode")

	cpuLimitModeCmd := &cobra.Command{
		Use:   "cpu-limit-mode <id>",
		Short: "Toggle CPU balance-settlement mode",
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
			if _, err := (*rySDK).SwitchRgsBalanceMode(id, enable); err != nil {
				return err
			}
			state := "disabled"
			if enable {
				state = "enabled"
			}
			fmt.Printf("CPU limit mode %s for game server %d\n", state, id)
			return nil
		},
	}
	cpuLimitModeCmd.Flags().Bool("enable", false, "Enable balance settlement")
	cpuLimitModeCmd.Flags().Bool("disable", false, "Disable balance settlement")

	gameCmd.AddCommand(cpuChargeCmd)
	gameCmd.AddCommand(dailyModeCmd)
	gameCmd.AddCommand(cpuLimitModeCmd)
}
