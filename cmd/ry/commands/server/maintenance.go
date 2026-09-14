package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addMaintenanceCommand(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	maintenanceCmd := &cobra.Command{
		Use:   "maintenance <id>",
		Short: "Get server maintenance status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRcsMaintenance(id)
			if err != nil {
				return err
			}
			if resp.Data == nil {
				fmt.Printf("Server %s: no active maintenance\n", args[0])
				return nil
			}
			return (*out).Print(resp.Data)
		},
	}

	serverCmd.AddCommand(maintenanceCmd)
}
