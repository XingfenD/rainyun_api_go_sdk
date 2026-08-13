package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addMonitorCommand(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	monitorCmd := &cobra.Command{
		Use:   "monitor <id>",
		Short: "Get server monitor data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			startStr, _ := cmd.Flags().GetString("start")
			endStr, _ := cmd.Flags().GetString("end")
			last, _ := cmd.Flags().GetString("last")

			start, end, err := cliutil.ResolveTimeRange(startStr, endStr, last)
			if err != nil {
				return err
			}

			resp, err := (*rySDK).GetRcsMonitorData(id, &rcs.GetRcsMonitorDataRequest{
				StartDate: start,
				EndDate:   end,
			})
			if err != nil {
				return err
			}
			return (*out).Print(toMonitorSamples(resp.Data))
		},
	}
	monitorCmd.Flags().String("last", "1h", "Time range to fetch, e.g. 30m, 1h, 6h, 24h, 7d")
	monitorCmd.Flags().String("start", "", "Start time (RFC3339 or YYYY-MM-DD[ HH:MM[:SS]]); overrides --last")
	monitorCmd.Flags().String("end", "", "End time (RFC3339 or YYYY-MM-DD[ HH:MM[:SS]]); defaults to now")

	serverCmd.AddCommand(monitorCmd)
}

func toMonitorSamples(d rcs.RcsMonitorData) []model.MonitorSample {
	samples := make([]model.MonitorSample, 0, len(d.Values))
	for _, row := range d.Values {
		s := model.MonitorSample{}
		if len(d.Columns) > 0 && len(row) > 0 {
			s.Time = formatMonitorTime(row[0])
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

func formatMonitorTime(v float64) string {
	sec := int64(v)
	if v > 1e11 {
		sec = int64(v / 1000)
	}
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}
