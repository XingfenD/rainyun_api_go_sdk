package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
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
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			start, _ := cmd.Flags().GetInt("start")
			end, _ := cmd.Flags().GetInt("end")
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
	monitorCmd.Flags().Int("start", 0, "Start timestamp (required)")
	monitorCmd.Flags().Int("end", 0, "End timestamp (required)")
	monitorCmd.MarkFlagRequired("start")
	monitorCmd.MarkFlagRequired("end")

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
