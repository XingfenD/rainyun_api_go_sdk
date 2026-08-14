package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addNatCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	natCmd := &cobra.Command{
		Use:   "nat",
		Short: "Manage game server NAT port mappings",
	}

	listCmd := &cobra.Command{
		Use:   "list <id>",
		Short: "List NAT port mappings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRgsDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(toGameNatMappings(resp.Data.NatList))
		},
	}

	addCmd := &cobra.Command{
		Use:   "add <id>",
		Short: "Add a NAT port mapping",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			portIn, _ := cmd.Flags().GetInt("port-in")
			portOut, _ := cmd.Flags().GetInt("port-out")
			portType, _ := cmd.Flags().GetString("type")
			tag, _ := cmd.Flags().GetString("tag")
			if _, err := (*rySDK).AddRgsNatPortMapping(id, &rcs.AddRcsNatPortMappingRequest{
				PortIn:   portIn,
				PortOut:  portOut,
				PortType: portType,
				Tag:      tag,
			}); err != nil {
				return err
			}
			fmt.Printf("NAT mapping added for game server %d\n", id)
			return nil
		},
	}
	addCmd.Flags().Int("port-in", 0, "Inner port (required)")
	addCmd.Flags().Int("port-out", 0, "Outer port 10000-60000 (required)")
	addCmd.Flags().String("type", "tcp", "Port type (tcp/udp/tcp_udp)")
	addCmd.Flags().String("tag", "", "Tag")
	addCmd.MarkFlagRequired("port-in")
	addCmd.MarkFlagRequired("port-out")

	deleteCmd := &cobra.Command{
		Use:   "delete <id> <nat-id>",
		Short: "Delete a NAT port mapping",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			natID, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteRgsNatPortMapping(id, &rgs.DeleteRgsNatPortMappingRequest{NatID: natID}); err != nil {
				return err
			}
			fmt.Printf("NAT mapping %d deleted for game server %d\n", natID, id)
			return nil
		},
	}

	natCmd.AddCommand(listCmd)
	natCmd.AddCommand(addCmd)
	natCmd.AddCommand(deleteCmd)

	gameCmd.AddCommand(natCmd)
}

func toGameNatMappings(items []rgs.RgsNatItem) []model.GameNatMapping {
	nats := make([]model.GameNatMapping, 0, len(items))
	for _, n := range items {
		nats = append(nats, model.GameNatMapping{
			ID: n.ID, PortIn: n.PortIn, PortOut: n.PortOut, PortType: n.PortType, Tag: n.Tag,
		})
	}
	return nats
}
