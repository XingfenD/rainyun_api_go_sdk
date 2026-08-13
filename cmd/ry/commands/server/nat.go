package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addNatCommands(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	natCmd := &cobra.Command{
		Use:   "nat",
		Short: "Manage server NAT port mappings",
	}

	natAddCmd := &cobra.Command{
		Use:   "add <id>",
		Short: "Add a NAT port mapping",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			portIn, _ := cmd.Flags().GetInt("port-in")
			portOut, _ := cmd.Flags().GetInt("port-out")
			portType, _ := cmd.Flags().GetString("port-type")
			tag, _ := cmd.Flags().GetString("tag")
			if _, err := (*rySDK).AddRcsNatPortMapping(id, &rcs.AddRcsNatPortMappingRequest{
				PortIn:   portIn,
				PortOut:  portOut,
				PortType: portType,
				Tag:      tag,
			}); err != nil {
				return err
			}
			fmt.Printf("NAT port mapping added for server %d\n", id)
			return nil
		},
	}
	natAddCmd.Flags().Int("port-in", 0, "Internal port (required)")
	natAddCmd.Flags().Int("port-out", 0, "External port (required)")
	natAddCmd.Flags().String("port-type", "tcp", "Port type (tcp/udp/tcp_udp)")
	natAddCmd.Flags().String("tag", "", "Mapping tag")
	natAddCmd.MarkFlagRequired("port-in")
	natAddCmd.MarkFlagRequired("port-out")

	natDeleteCmd := &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a NAT port mapping",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			natID, _ := cmd.Flags().GetInt("nat-id")
			if _, err := (*rySDK).DeleteRcsNatPortMapping(id, &rcs.DeleteRcsNatPortMappingRequest{
				NatID: natID,
			}); err != nil {
				return err
			}
			fmt.Printf("NAT port mapping %d deleted for server %d\n", natID, id)
			return nil
		},
	}
	natDeleteCmd.Flags().Int("nat-id", 0, "NAT mapping ID (required)")
	natDeleteCmd.MarkFlagRequired("nat-id")

	natCmd.AddCommand(natAddCmd, natDeleteCmd)
	serverCmd.AddCommand(natCmd)
}
