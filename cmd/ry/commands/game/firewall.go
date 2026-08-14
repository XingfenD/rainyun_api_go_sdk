package game

import (
	"fmt"
	"strconv"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addFirewallCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	firewallCmd := &cobra.Command{
		Use:   "firewall",
		Short: "Manage game server firewall rules",
	}

	listCmd := &cobra.Command{
		Use:   "list <id>",
		Short: "List firewall rules",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRgsFirewallRules(id, &rgs.GetRgsFirewallRulesRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			return (*out).Print(toGameFirewallRules(resp.Data.Records))
		},
	}

	addCmd := &cobra.Command{
		Use:   "add <id>",
		Short: "Create or update a firewall rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			req := &rcs.SetRcsFirewallRuleRequest{
				Action:        mustFlagString(cmd, "action"),
				Protocol:      mustFlagString(cmd, "protocol"),
				SourceAddress: mustFlagString(cmd, "source"),
				SourcePort:    mustFlagString(cmd, "source-port"),
				DestPort:      mustFlagString(cmd, "dest-port"),
				Description:   mustFlagString(cmd, "desc"),
			}
			if _, err := (*rySDK).SetRgsFirewallRule(id, req); err != nil {
				return err
			}
			fmt.Printf("Firewall rule set for game server %d\n", id)
			return nil
		},
	}
	addCmd.Flags().String("action", "ACCEPT", "Action (ACCEPT/DROP)")
	addCmd.Flags().String("protocol", "", "Protocol (tcp/udp/icmp, empty for all)")
	addCmd.Flags().String("source", "", "Source address(es)")
	addCmd.Flags().String("source-port", "", "Source port(s)")
	addCmd.Flags().String("dest-port", "", "Destination port(s)")
	addCmd.Flags().String("desc", "", "Description")

	deleteCmd := &cobra.Command{
		Use:   "delete <id> <rule-id>",
		Short: "Delete a firewall rule",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			ruleID, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteRgsFirewallRule(id, ruleID); err != nil {
				return err
			}
			fmt.Printf("Firewall rule %d deleted for game server %d\n", ruleID, id)
			return nil
		},
	}

	moveCmd := &cobra.Command{
		Use:   "move <id> <rule-id>",
		Short: "Move a firewall rule priority",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			ruleID, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			newPos, _ := cmd.Flags().GetInt("pos")
			if _, err := (*rySDK).MobileRgsFirewallRulePriority(id, ruleID, rcs.MobileRcsFirewallRulePriorityRequest{
				NewPos: newPos,
			}); err != nil {
				return err
			}
			fmt.Printf("Firewall rule %d moved to pos %d for game server %d\n", ruleID, newPos, id)
			return nil
		},
	}
	moveCmd.Flags().Int("pos", 0, "New position (required)")
	moveCmd.MarkFlagRequired("pos")

	modeCmd := &cobra.Command{
		Use:   "mode <id> <mode>",
		Short: "Set firewall mode (black/white)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).SetRgsFirewallMode(id, args[1]); err != nil {
				return err
			}
			fmt.Printf("Firewall mode set to %s for game server %d\n", args[1], id)
			return nil
		},
	}

	syncTimeCmd := &cobra.Command{
		Use:   "sync-time <id>",
		Short: "Get firewall sync start time",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRgsFirewallSyncTime(id)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	firewallCmd.AddCommand(listCmd)
	firewallCmd.AddCommand(addCmd)
	firewallCmd.AddCommand(deleteCmd)
	firewallCmd.AddCommand(moveCmd)
	firewallCmd.AddCommand(modeCmd)
	firewallCmd.AddCommand(syncTimeCmd)

	gameCmd.AddCommand(firewallCmd)
}

func toGameFirewallRules(records []rcs.RcsFirewallRule) []model.FirewallRule {
	rules := make([]model.FirewallRule, 0, len(records))
	for _, r := range records {
		rules = append(rules, model.FirewallRule{
			ID:       strconv.Itoa(r.ID),
			Protocol: r.Protocol,
			Source:   r.SourceAddress,
			DestPort: r.DestPort,
			Action:   r.Action,
			Enabled:  r.IsEnable,
			Desc:     r.Description,
		})
	}
	return rules
}
