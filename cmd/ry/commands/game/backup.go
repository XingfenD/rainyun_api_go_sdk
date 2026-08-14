package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addBackupCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "Manage game server backups",
	}

	createCmd := &cobra.Command{
		Use:   "create <id> <label>",
		Short: "Create a game server backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).CreateRgsBackup(id, args[1]); err != nil {
				return err
			}
			fmt.Printf("Backup %q created for game server %d\n", args[1], id)
			return nil
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete <id> <backup-id>",
		Short: "Delete a game server backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			bid, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteRgsBackup(id, bid); err != nil {
				return err
			}
			fmt.Printf("Backup %d deleted for game server %d\n", bid, id)
			return nil
		},
	}

	cancelCmd := &cobra.Command{
		Use:   "cancel <id> <backup-id>",
		Short: "Cancel a game server backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			bid, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).CancelRgsBackup(id, bid); err != nil {
				return err
			}
			fmt.Printf("Backup %d cancelled for game server %d\n", bid, id)
			return nil
		},
	}

	restoreCmd := &cobra.Command{
		Use:   "restore <id> <backup-id>",
		Short: "Restore a game server backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			bid, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).RestoreRgsBackup(id, bid); err != nil {
				return err
			}
			fmt.Printf("Backup %d restoring for game server %d\n", bid, id)
			return nil
		},
	}

	settingCmd := &cobra.Command{
		Use:   "setting <id>",
		Short: "Configure game server auto backup",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			hour, _ := cmd.Flags().GetInt("hour")
			minute, _ := cmd.Flags().GetInt("minute")
			keep, _ := cmd.Flags().GetInt("keep")
			if _, err := (*rySDK).EnableRgsAutoBackup(id, &rcs.RcsSetBackupOptionsRequest{
				AutoBackupHour:   hour,
				AutoBackupMinute: minute,
				KeepLast:         keep,
			}); err != nil {
				return err
			}
			fmt.Printf("Auto backup configured for game server %d\n", id)
			return nil
		},
	}
	settingCmd.Flags().Int("hour", -1, "Auto backup hour (-1 disables)")
	settingCmd.Flags().Int("minute", -1, "Auto backup minute (-1 disables)")
	settingCmd.Flags().Int("keep", 1, "Backups to keep (0/1/3/7/15/31)")

	backupCmd.AddCommand(createCmd)
	backupCmd.AddCommand(deleteCmd)
	backupCmd.AddCommand(cancelCmd)
	backupCmd.AddCommand(restoreCmd)
	backupCmd.AddCommand(settingCmd)

	gameCmd.AddCommand(backupCmd)
}
