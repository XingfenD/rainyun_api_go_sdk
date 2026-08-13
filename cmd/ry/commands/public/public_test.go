package public

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
		"app-config": true, "news": true, "status": true,
	}

	got := map[string]bool{}
	for _, c := range cmd.Commands() {
		got[c.Name()] = true
	}
	for name := range topLevel {
		if !got[name] {
			t.Errorf("missing subcommand %q", name)
		}
	}
	for name := range got {
		if !topLevel[name] {
			t.Errorf("unexpected subcommand %q", name)
		}
	}
}

func TestCmdNestedSubcommands(t *testing.T) {
	var rySDK *sdk.RainyunSDK
	var out *output.Printer
	cmd := Cmd(&rySDK, &out)

	if findCommand(cmd, "status") == nil {
		t.Errorf("status command not found")
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
