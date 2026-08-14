package game

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addPanelCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	mcsmCmd := &cobra.Command{
		Use:   "mcsm",
		Short: "Manage MCSM panel",
	}

	mcsmUserCmd := &cobra.Command{
		Use:   "user",
		Short: "Manage MCSM panel users",
	}

	mcsmUserListCmd := &cobra.Command{
		Use:   "list",
		Short: "List MCSM panel users",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetMcsmUserList()
			if err != nil {
				return err
			}
			return (*out).Print(toGamePanelUsers(resp.Data))
		},
	}

	mcsmUserCreateCmd := &cobra.Command{
		Use:   "create <name> <password>",
		Short: "Create an MCSM panel user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := (*rySDK).CreateMcsmUser(args[0], args[1]); err != nil {
				return err
			}
			fmt.Printf("MCSM user %s created\n", args[0])
			return nil
		},
	}

	mcsmUserEditCmd := &cobra.Command{
		Use:   "edit <name> <password>",
		Short: "Edit an MCSM panel user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := (*rySDK).EditMcsmUser(args[0], args[1]); err != nil {
				return err
			}
			fmt.Printf("MCSM user %s updated\n", args[0])
			return nil
		},
	}

	mcsmUserDeleteCmd := &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete an MCSM panel user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := (*rySDK).DeleteMcsmUser(args[0]); err != nil {
				return err
			}
			fmt.Printf("MCSM user %s deleted\n", args[0])
			return nil
		},
	}

	mcsmUserCmd.AddCommand(mcsmUserListCmd)
	mcsmUserCmd.AddCommand(mcsmUserCreateCmd)
	mcsmUserCmd.AddCommand(mcsmUserEditCmd)
	mcsmUserCmd.AddCommand(mcsmUserDeleteCmd)

	mcsmStartCmd := &cobra.Command{
		Use:   "start <id>",
		Short: "Start the MCSM instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).StartMcsmInstance(id); err != nil {
				return err
			}
			fmt.Printf("MCSM instance %d starting\n", id)
			return nil
		},
	}

	mcsmStatusCmd := &cobra.Command{
		Use:   "status <id>",
		Short: "Get MCSM instance status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetMcsmStatus(id)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	mcsmSftpInitCmd := &cobra.Command{
		Use:   "sftp-init <id>",
		Short: "Initialize or refresh SFTP for the MCSM instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).McsmSftpInit(id); err != nil {
				return err
			}
			fmt.Printf("SFTP initialized for game server %d\n", id)
			return nil
		},
	}

	mcsmCmd.AddCommand(mcsmUserCmd)
	mcsmCmd.AddCommand(mcsmStartCmd)
	mcsmCmd.AddCommand(mcsmStatusCmd)
	mcsmCmd.AddCommand(mcsmSftpInitCmd)

	pteroCmd := &cobra.Command{
		Use:   "ptero",
		Short: "Manage Pterodactyl panel",
	}

	pteroUserCmd := &cobra.Command{
		Use:   "user",
		Short: "Manage Pterodactyl panel users",
	}

	pteroUserListCmd := &cobra.Command{
		Use:   "list",
		Short: "List Pterodactyl panel users",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetPteroUserList()
			if err != nil {
				return err
			}
			return (*out).Print(toGamePanelUsersFromPtero(resp.Data))
		},
	}

	pteroUserCreateCmd := &cobra.Command{
		Use:   "create <name> <password>",
		Short: "Create a Pterodactyl panel user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := (*rySDK).CreatePteroUser(args[0], args[1]); err != nil {
				return err
			}
			fmt.Printf("Ptero user %s created\n", args[0])
			return nil
		},
	}

	pteroUserEditCmd := &cobra.Command{
		Use:   "edit <name> <password>",
		Short: "Edit a Pterodactyl panel user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := (*rySDK).EditPteroUser(args[0], args[1]); err != nil {
				return err
			}
			fmt.Printf("Ptero user %s updated\n", args[0])
			return nil
		},
	}

	pteroUserDeleteCmd := &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete a Pterodactyl panel user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := (*rySDK).DeletePteroUser(args[0]); err != nil {
				return err
			}
			fmt.Printf("Ptero user %s deleted\n", args[0])
			return nil
		},
	}

	pteroUserCmd.AddCommand(pteroUserListCmd)
	pteroUserCmd.AddCommand(pteroUserCreateCmd)
	pteroUserCmd.AddCommand(pteroUserEditCmd)
	pteroUserCmd.AddCommand(pteroUserDeleteCmd)
	pteroCmd.AddCommand(pteroUserCmd)

	switchUserCmd := &cobra.Command{
		Use:   "switch-user <id>",
		Short: "Switch panel user of a game server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).SwitchRgsPanelUser(id, &rgs.SwitchRgsPanelUserRequest{
				Subtype:  mustFlagString(cmd, "subtype"),
				Name:     mustFlagString(cmd, "name"),
				Password: mustFlagString(cmd, "password"),
			}); err != nil {
				return err
			}
			fmt.Printf("Panel user switched for game server %d\n", id)
			return nil
		},
	}
	switchUserCmd.Flags().String("subtype", "", "Panel subtype (ptero/mcsm/k8s_panel)")
	switchUserCmd.Flags().String("name", "", "Panel user name")
	switchUserCmd.Flags().String("password", "", "Panel user password")

	gameCmd.AddCommand(mcsmCmd)
	gameCmd.AddCommand(pteroCmd)
	gameCmd.AddCommand(switchUserCmd)
}

func toGamePanelUsers(users []rgs.McsmUser) []model.GamePanelUser {
	out := make([]model.GamePanelUser, 0, len(users))
	for _, u := range users {
		out = append(out, model.GamePanelUser{Name: u.Name, UserID: u.UserID, PanelUUID: u.PanelUUID})
	}
	return out
}

func toGamePanelUsersFromPtero(users []rgs.PteroUser) []model.GamePanelUser {
	out := make([]model.GamePanelUser, 0, len(users))
	for _, u := range users {
		out = append(out, model.GamePanelUser{Name: u.Name})
	}
	return out
}
