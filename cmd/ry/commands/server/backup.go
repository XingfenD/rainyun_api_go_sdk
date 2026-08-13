package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addBackupCommands(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "Manage server backups",
	}

	backupCreateCmd := &cobra.Command{
		Use:   "create <id> <label>",
		Short: "Create a server backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).CreateRcsBackup(id, args[1]); err != nil {
				return err
			}
			fmt.Printf("Backup %q created for server %d\n", args[1], id)
			return nil
		},
	}

	backupDeleteCmd := &cobra.Command{
		Use:   "delete <id> <backup-id>",
		Short: "Delete a server backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			bid, err := parseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteRcsBackup(id, bid); err != nil {
				return err
			}
			fmt.Printf("Backup %d deleted for server %d\n", bid, id)
			return nil
		},
	}

	backupCancelCmd := &cobra.Command{
		Use:   "cancel <id> <backup-id>",
		Short: "Cancel a server backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			bid, err := parseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).CancelRcsBackup(id, bid); err != nil {
				return err
			}
			fmt.Printf("Backup %d cancelled for server %d\n", bid, id)
			return nil
		},
	}

	backupRestoreCmd := &cobra.Command{
		Use:   "restore <id> <backup-id>",
		Short: "Restore a server backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			bid, err := parseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).RestoreRcsBackup(id, bid); err != nil {
				return err
			}
			fmt.Printf("Backup %d restoring for server %d\n", bid, id)
			return nil
		},
	}

	backupAutoCmd := &cobra.Command{
		Use:   "auto <id>",
		Short: "Configure server auto backup",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			hour, _ := cmd.Flags().GetInt("hour")
			minute, _ := cmd.Flags().GetInt("minute")
			keep, _ := cmd.Flags().GetInt("keep")
			if _, err := (*rySDK).EnableRcsAutoBackup(id, &rcs.RcsSetBackupOptionsRequest{
				AutoBackupHour:   hour,
				AutoBackupMinute: minute,
				KeepLast:         keep,
			}); err != nil {
				return err
			}
			fmt.Printf("Auto backup configured for server %d\n", id)
			return nil
		},
	}
	backupAutoCmd.Flags().Int("hour", 0, "Auto backup hour (0-23)")
	backupAutoCmd.Flags().Int("minute", 0, "Auto backup minute (0-59)")
	backupAutoCmd.Flags().Int("keep", 1, "Number of backups to keep (1/3/7)")

	backupCmd.AddCommand(backupCreateCmd, backupDeleteCmd, backupCancelCmd, backupRestoreCmd, backupAutoCmd)
	serverCmd.AddCommand(backupCmd)
}
