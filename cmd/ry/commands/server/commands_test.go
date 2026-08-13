package server

import (
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func TestCmdRegistersAllCommands(t *testing.T) {
	var rySDK *sdk.RainyunSDK
	var out *output.Printer
	cmd := Cmd(&rySDK, &out)

	topLevel := map[string]bool{
		"list": true, "get": true, "start": true, "stop": true, "reboot": true,
		"reinstall": true, "reset-password": true, "vnc": true,
		"free": true, "set-tag": true, "create": true,
		"renew-price": true, "renew": true, "auto-renew": true,
		"upgrade": true, "edisk": true, "monitor": true,
		"backup": true, "eip": true, "nat": true, "traffic": true,
		"firewall": true, "pve-address": true,
	}

	got := map[string]bool{}
	for _, c := range cmd.Commands() {
		got[c.Name()] = true
	}
	for name := range topLevel {
		if !got[name] {
			t.Errorf("missing top-level subcommand %q", name)
		}
	}
	for name := range got {
		if !topLevel[name] {
			t.Errorf("unexpected top-level subcommand %q", name)
		}
	}
}

func TestCmdNestedSubcommands(t *testing.T) {
	var rySDK *sdk.RainyunSDK
	var out *output.Printer
	cmd := Cmd(&rySDK, &out)

	cases := map[string][]string{
		"edisk":    {"create", "expand"},
		"backup":   {"create", "delete", "cancel", "restore", "auto"},
		"eip":      {"list", "set-description", "create", "change", "discard"},
		"nat":      {"add", "delete"},
		"traffic":  {"charge", "limit"},
		"firewall": {"list", "set", "delete", "move"},
	}

	for parent, children := range cases {
		p := findCommand(cmd, parent)
		if p == nil {
			t.Errorf("parent command %q not found", parent)
			continue
		}
		for _, child := range children {
			if findCommand(p, child) == nil {
				t.Errorf("missing subcommand %q under %q", child, parent)
			}
		}
	}
}

func findCommand(cmd *cobra.Command, name string) *cobra.Command {
	for _, c := range cmd.Commands() {
		if c.Name() == name {
			return c
		}
	}
	return nil
}
