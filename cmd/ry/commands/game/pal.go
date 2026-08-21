package game

import (
	"fmt"

	"github.com/bytedance/sonic"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addPalCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	palCmd := &cobra.Command{
		Use:   "pal",
		Short: "Manage Palworld (pal) panel",
	}

	configGetCmd := &cobra.Command{
		Use:   "config <id>",
		Short: "Get pal configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetPalConfig(id)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	configSetCmd := &cobra.Command{
		Use:   "config-set <id> <json>",
		Short: "Set pal configuration (JSON object)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			var config map[string]any
			if err := sonic.Unmarshal([]byte(args[1]), &config); err != nil {
				return fmt.Errorf("invalid JSON config: %w", err)
			}
			if _, err := (*rySDK).SetPalConfig(id, config); err != nil {
				return err
			}
			fmt.Printf("Pal config updated for game server %d\n", id)
			return nil
		},
	}

	initCmd := &cobra.Command{
		Use:   "init <id>",
		Short: "Initialize pal configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).InitPal(id); err != nil {
				return err
			}
			fmt.Printf("Pal init requested for game server %d\n", id)
			return nil
		},
	}

	langCmd := &cobra.Command{
		Use:   "lang <id>",
		Short: "Get pal configuration descriptions (Chinese)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetPalLang(id)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	rconCmd := &cobra.Command{
		Use:   "rcon <id> <command>",
		Short: "Send a pal RCON command",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).PalRcon(id, args[1]); err != nil {
				return err
			}
			fmt.Printf("RCON command sent to game server %d\n", id)
			return nil
		},
	}

	restartCmd := &cobra.Command{
		Use:   "restart <id>",
		Short: "Restart pal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).RestartPal(id); err != nil {
				return err
			}
			fmt.Printf("Pal restarting for game server %d\n", id)
			return nil
		},
	}

	stopCmd := &cobra.Command{
		Use:   "stop <id>",
		Short: "Stop pal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).StopPal(id); err != nil {
				return err
			}
			fmt.Printf("Pal stopping for game server %d\n", id)
			return nil
		},
	}

	palCmd.AddCommand(configGetCmd)
	palCmd.AddCommand(configSetCmd)
	palCmd.AddCommand(initCmd)
	palCmd.AddCommand(langCmd)
	palCmd.AddCommand(rconCmd)
	palCmd.AddCommand(restartCmd)
	palCmd.AddCommand(stopCmd)

	gameCmd.AddCommand(palCmd)
}
