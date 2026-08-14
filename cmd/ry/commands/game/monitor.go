package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addMonitorCommand(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	monitorCmd := &cobra.Command{
		Use:   "monitor <id>",
		Short: "Get game server monitor data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			hours, _ := cmd.Flags().GetInt("hours")
			now := time.Now()
			resp, err := (*rySDK).GetRgsMonitorData(id, &rgs.GetRgsMonitorDataRequest{
				StartDate: int(now.Add(-time.Duration(hours) * time.Hour).Unix()),
				EndDate:   int(now.Unix()),
			})
			if err != nil {
				return err
			}
			return (*out).Print(toGameMonitorSamples(resp.Data))
		},
	}
	monitorCmd.Flags().Int("hours", 1, "Lookback window in hours")

	gameCmd.AddCommand(monitorCmd)
}

func toGameMonitorSamples(d rcs.RcsMonitorData) []model.MonitorSample {
	samples := make([]model.MonitorSample, 0, len(d.Values))
	for _, row := range d.Values {
		s := model.MonitorSample{}
		if len(d.Columns) > 0 && len(row) > 0 {
			s.Time = time.Unix(int64(row[0]), 0).Format("2006-01-02 15:04:05")
		}
		parts := make([]string, 0, len(d.Columns))
		for j := 1; j < len(d.Columns) && j < len(row); j++ {
			parts = append(parts, fmt.Sprintf("%s=%.2f", d.Columns[j], row[j]))
		}
		s.Metrics = strings.Join(parts, " ")
		samples = append(samples, s)
	}
	return samples
}
