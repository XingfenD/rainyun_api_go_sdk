package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addLifecycleCommands(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	freeCmd := &cobra.Command{
		Use:   "free <id>",
		Short: "Release (free) a server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).FreeRcs(id); err != nil {
				return err
			}
			fmt.Printf("Server %d released\n", id)
			return nil
		},
	}

	setTagCmd := &cobra.Command{
		Use:   "set-tag <id> <tag>",
		Short: "Set server tag",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).SetRcsTag(id, args[1]); err != nil {
				return err
			}
			fmt.Printf("Tag %q set for server %d\n", args[1], id)
			return nil
		},
	}

	serverCmd.AddCommand(freeCmd, setTagCmd)
}
