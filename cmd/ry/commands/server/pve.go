package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addPveAddressCommand(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	pveCmd := &cobra.Command{
		Use:   "pve-address <id>",
		Short: "Get server PVE address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			addr, err := (*rySDK).GetRcsPveAddress(id)
			if err != nil {
				return err
			}
			fmt.Println(addr)
			return nil
		},
	}

	serverCmd.AddCommand(pveCmd)
}
