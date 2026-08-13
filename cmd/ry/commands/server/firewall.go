package server

import (
	"fmt"
	"strconv"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addFirewallCommands(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	firewallCmd := &cobra.Command{
		Use:   "firewall",
		Short: "Manage server firewall rules",
	}

	firewallListCmd := &cobra.Command{
		Use:   "list <id>",
		Short: "List firewall rules",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRcsFirewallRules(id, &rcs.GetRcsFirewallRulesRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			return (*out).Print(toFirewallRules(resp.Data.Records))
		},
	}

	firewallSetCmd := &cobra.Command{
		Use:   "set <id>",
		Short: "Add or update a firewall rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			action, _ := cmd.Flags().GetString("action")
			protocol, _ := cmd.Flags().GetString("protocol")
			destPort, _ := cmd.Flags().GetString("dest-port")
			source, _ := cmd.Flags().GetString("source")
			sourcePort, _ := cmd.Flags().GetString("source-port")
			desc, _ := cmd.Flags().GetString("description")
			enable, _ := cmd.Flags().GetBool("enable")
			ruleID, _ := cmd.Flags().GetInt("rule-id")
			if _, err := (*rySDK).SetRcsFirewallRule(id, &rcs.SetRcsFirewallRuleRequest{
				Action:        action,
				Protocol:      protocol,
				DestPort:      destPort,
				SourceAddress: source,
				SourcePort:    sourcePort,
				Description:   desc,
				IsEnable:      enable,
				ID:            ruleID,
			}); err != nil {
				return err
			}
			fmt.Printf("Firewall rule set for server %d\n", id)
			return nil
		},
	}
	firewallSetCmd.Flags().String("action", "", "Action (accept/drop) (required)")
	firewallSetCmd.Flags().String("protocol", "", "Protocol (tcp/udp/icmp)")
	firewallSetCmd.Flags().String("dest-port", "", "Destination port(s)")
	firewallSetCmd.Flags().String("source", "", "Source address(es)")
	firewallSetCmd.Flags().String("source-port", "", "Source port(s)")
	firewallSetCmd.Flags().String("description", "", "Rule description")
	firewallSetCmd.Flags().Bool("enable", true, "Whether the rule is enabled")
	firewallSetCmd.Flags().Int("rule-id", 0, "Rule ID (set to update an existing rule)")
	firewallSetCmd.MarkFlagRequired("action")

	firewallDeleteCmd := &cobra.Command{
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
			if _, err := (*rySDK).DeleteRcsFirewallRule(id, ruleID); err != nil {
				return err
			}
			fmt.Printf("Firewall rule %d deleted for server %d\n", ruleID, id)
			return nil
		},
	}

	firewallMoveCmd := &cobra.Command{
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
			pos, _ := cmd.Flags().GetInt("pos")
			if _, err := (*rySDK).MobileRcsFirewallRulePriority(id, ruleID, rcs.MobileRcsFirewallRulePriorityRequest{
				NewPos: pos,
			}); err != nil {
				return err
			}
			fmt.Printf("Firewall rule %d moved to position %d\n", ruleID, pos)
			return nil
		},
	}
	firewallMoveCmd.Flags().Int("pos", 0, "New position (required)")
	firewallMoveCmd.MarkFlagRequired("pos")

	firewallCmd.AddCommand(firewallListCmd, firewallSetCmd, firewallDeleteCmd, firewallMoveCmd)
	serverCmd.AddCommand(firewallCmd)
}

func toFirewallRules(records []rcs.RcsFirewallRule) []model.FirewallRule {
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
