package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addEIPCommands(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	eipCmd := &cobra.Command{
		Use:   "eip",
		Short: "Manage server elastic IPs",
	}

	eipSetDescCmd := &cobra.Command{
		Use:   "set-description <id> <ip> <description>",
		Short: "Set elastic IP description",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).SetRcsEipDescription(id, args[1], args[2]); err != nil {
				return err
			}
			fmt.Printf("Description set for IP %s on server %d\n", args[1], id)
			return nil
		},
	}

	eipCreateCmd := &cobra.Command{
		Use:   "create <id>",
		Short: "Create and bind an elastic IP",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			num, _ := cmd.Flags().GetInt("num")
			ipType, _ := cmd.Flags().GetString("type")
			flags, _ := cmd.Flags().GetString("flags")
			if _, err := (*rySDK).CreateAndBindElasticIpToRcs(id, &rcs.CreateAndBindIpToRcsRequest{
				WithIPNum:  num,
				WithIPType: ipType,
				WithFlags:  flags,
			}); err != nil {
				return err
			}
			fmt.Printf("Elastic IP(s) created for server %d\n", id)
			return nil
		},
	}
	eipCreateCmd.Flags().Int("num", 1, "Number of IPs")
	eipCreateCmd.Flags().String("type", "ipv4", "IP type (ipv4/ipv6)")
	eipCreateCmd.Flags().String("flags", "", "IP flags (us_ddosip/nb_ddosip)")

	eipChangeCmd := &cobra.Command{
		Use:   "change <id> <ip>",
		Short: "Change elastic IP",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			toIP, _ := cmd.Flags().GetString("to-ip")
			reason, _ := cmd.Flags().GetString("reason")
			if _, err := (*rySDK).ChangeRcsIP(id, &rcs.ChangeRcsIPRequest{
				IP:                 args[1],
				ToIP:               toIP,
				DisableOldIPReason: reason,
			}); err != nil {
				return err
			}
			fmt.Printf("IP %s changed for server %d\n", args[1], id)
			return nil
		},
	}
	eipChangeCmd.Flags().String("to-ip", "", "Target IP (optional)")
	eipChangeCmd.Flags().String("reason", "", "Reason for disabling old IP (optional)")

	eipDiscardCmd := &cobra.Command{
		Use:   "discard <id> <ip>",
		Short: "Discard an elastic IP",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DisCardRcsIP(id, rcs.DisCardRcsIPRequest{IP: args[1]}); err != nil {
				return err
			}
			fmt.Printf("IP %s discarded for server %d\n", args[1], id)
			return nil
		},
	}

	eipCmd.AddCommand(eipSetDescCmd, eipCreateCmd, eipChangeCmd, eipDiscardCmd)
	serverCmd.AddCommand(eipCmd)
}
