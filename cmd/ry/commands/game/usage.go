package game

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addUsageCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	usageCmd := &cobra.Command{
		Use:   "usage [id]",
		Short: "Get game server usage (all instances or one)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				id, err := cliutil.ParseID(args[0])
				if err != nil {
					return err
				}
				resp, err := (*rySDK).GetRgsUsage(id)
				if err != nil {
					return err
				}
				return (*out).Print(model.RawData{Data: resp.Data})
			}
			resp, err := (*rySDK).GetRgsUsageList(&rgs.GetRgsUsageListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	gameCmd.AddCommand(usageCmd)
}
