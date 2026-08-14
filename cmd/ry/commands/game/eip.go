package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addEIPCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	eipCmd := &cobra.Command{
		Use:   "eip",
		Short: "Manage game server elastic IPs",
	}

	addCmd := &cobra.Command{
		Use:   "add <id>",
		Short: "Create and bind an elastic IP",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			num, _ := cmd.Flags().GetInt("num")
			ipType, _ := cmd.Flags().GetString("type")
			flags, _ := cmd.Flags().GetString("flags")
			if _, err := (*rySDK).CreateAndBindElasticIpToRgs(id, &rcs.CreateAndBindIpToRcsRequest{
				WithIPNum:  num,
				WithIPType: ipType,
				WithFlags:  flags,
			}); err != nil {
				return err
			}
			fmt.Printf("Elastic IP(s) created for game server %d\n", id)
			return nil
		},
	}
	addCmd.Flags().Int("num", 1, "Number of IPs")
	addCmd.Flags().String("type", "ipv4", "IP type (ipv4/ipv6)")
	addCmd.Flags().String("flags", "", "IP flags (us_ddosip/nb_ddosip)")

	changeCmd := &cobra.Command{
		Use:   "change <id> <ip>",
		Short: "Change elastic IP",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			toIP, _ := cmd.Flags().GetString("to-ip")
			reason, _ := cmd.Flags().GetString("reason")
			if _, err := (*rySDK).ChangeRgsIP(id, &rcs.ChangeRcsIPRequest{
				IP:                 args[1],
				ToIP:               toIP,
				DisableOldIPReason: reason,
			}); err != nil {
				return err
			}
			fmt.Printf("IP %s changed for game server %d\n", args[1], id)
			return nil
		},
	}
	changeCmd.Flags().String("to-ip", "", "Target IP")
	changeCmd.Flags().String("reason", "", "Reason for disabling old IP")

	setDescCmd := &cobra.Command{
		Use:   "set-description <id> <ip> <description>",
		Short: "Set elastic IP description",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).SetRgsEipDescription(id, args[1], args[2]); err != nil {
				return err
			}
			fmt.Printf("Description set for IP %s on game server %d\n", args[1], id)
			return nil
		},
	}

	discardCmd := &cobra.Command{
		Use:   "discard <id> <ip>",
		Short: "Discard an elastic IP",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DisCardRgsIP(id, rcs.DisCardRcsIPRequest{IP: args[1]}); err != nil {
				return err
			}
			fmt.Printf("IP %s discarded for game server %d\n", args[1], id)
			return nil
		},
	}

	eipCmd.AddCommand(addCmd)
	eipCmd.AddCommand(changeCmd)
	eipCmd.AddCommand(setDescCmd)
	eipCmd.AddCommand(discardCmd)

	gameCmd.AddCommand(eipCmd)
}
