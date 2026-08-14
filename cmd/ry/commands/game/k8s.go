package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addK8SCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	k8sCmd := &cobra.Command{
		Use:   "k8s",
		Short: "Manage K8S panel settings",
	}

	databaseCmd := &cobra.Command{
		Use:   "database <id>",
		Short: "Set K8S panel database settings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			enable, _ := cmd.Flags().GetBool("enable")
			version, _ := cmd.Flags().GetString("version")
			if _, err := (*rySDK).SetK8SPanelDatabase(id, &rgs.SetK8SPanelDatabaseRequest{
				IsEnabled: enable,
				Version:   version,
			}); err != nil {
				return err
			}
			fmt.Printf("K8S database settings updated for game server %d\n", id)
			return nil
		},
	}
	databaseCmd.Flags().Bool("enable", false, "Enable database")
	databaseCmd.Flags().String("version", "", "Database version")

	sftpCmd := &cobra.Command{
		Use:   "sftp <id>",
		Short: "Set K8S panel SFTP credentials",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).SetK8SPanelSFTP(id, &rgs.SetK8SPanelSFTPRequest{
				Username: mustFlagString(cmd, "username"),
				Password: mustFlagString(cmd, "password"),
			}); err != nil {
				return err
			}
			fmt.Printf("K8S SFTP settings updated for game server %d\n", id)
			return nil
		},
	}
	sftpCmd.Flags().String("username", "", "SFTP username")
	sftpCmd.Flags().String("password", "", "SFTP password")

	startCommandCmd := &cobra.Command{
		Use:   "start-command <id> <command>",
		Short: "Set start command (RainYun panel only)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).SetK8SPanelStartCommand(id, &rgs.SetK8SPanelStartCommandRequest{
				Command: args[1],
			}); err != nil {
				return err
			}
			fmt.Printf("Start command set for game server %d\n", id)
			return nil
		},
	}

	k8sCmd.AddCommand(databaseCmd)
	k8sCmd.AddCommand(sftpCmd)
	k8sCmd.AddCommand(startCommandCmd)

	gameCmd.AddCommand(k8sCmd)
}
