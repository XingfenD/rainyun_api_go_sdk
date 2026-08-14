package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addNetworkCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	toBridgeCmd := &cobra.Command{
		Use:   "to-bridge <id>",
		Short: "Switch a game server to bridge mode",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).RgsToBridge(id); err != nil {
				return err
			}
			fmt.Printf("Game server %d switching to bridge mode\n", id)
			return nil
		},
	}

	setIntIPCmd := &cobra.Command{
		Use:   "set-int-ip <id> <ip>",
		Short: "Set intranet IP in bridge mode",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).RgsBridgeSetIntIP(id, &rgs.RgsBridgeSetIntIPRequest{IP: args[1]}); err != nil {
				return err
			}
			fmt.Printf("Intranet IP set for game server %d\n", id)
			return nil
		},
	}

	vnetCmd := &cobra.Command{
		Use:   "vnet",
		Short: "Manage intranet subnets (vnet)",
	}

	vnetCreateCmd := &cobra.Command{
		Use:   "create <id> <name>",
		Short: "Create an intranet subnet",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).CreateRgsVnet(id, &rgs.RgsVnetRequest{Name: args[1]}); err != nil {
				return err
			}
			fmt.Printf("Subnet %q created for game server %d\n", args[1], id)
			return nil
		},
	}

	vnetRenameCmd := &cobra.Command{
		Use:   "rename <id> <new-name>",
		Short: "Rename an intranet subnet",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).RenameRgsVnet(id, &rgs.RgsVnetRequest{NewName: args[1]}); err != nil {
				return err
			}
			fmt.Printf("Subnet renamed to %q for game server %d\n", args[1], id)
			return nil
		},
	}

	vnetCmd.AddCommand(vnetCreateCmd)
	vnetCmd.AddCommand(vnetRenameCmd)

	gameCmd.AddCommand(toBridgeCmd)
	gameCmd.AddCommand(setIntIPCmd)
	gameCmd.AddCommand(vnetCmd)
}
