package server

import (
	"fmt"
	"strconv"
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
			startStr, _ := cmd.Flags().GetString("start")
			endStr, _ := cmd.Flags().GetString("end")
			last, _ := cmd.Flags().GetString("last")

			start, end, err := resolveMonitorRange(startStr, endStr, last)
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

// resolveMonitorRange turns the --last/--start/--end flags into Unix second
// timestamps. When --start and --end are both empty, the --last duration is
// used (end = now, start = now - duration). Explicit --start/--end take
// precedence over --last.
func resolveMonitorRange(startStr, endStr, last string) (start, end int, err error) {
	if startStr == "" && endStr == "" {
		d, derr := parseDurationArg(last)
		if derr != nil {
			return 0, 0, fmt.Errorf("invalid --last %q: %w", last, derr)
		}
		now := time.Now()
		return int(now.Add(-d).Unix()), int(now.Unix()), nil
	}
	if startStr == "" || endStr == "" {
		return 0, 0, fmt.Errorf("--start and --end must be specified together")
	}
	if start, err = parseTimeArg(startStr); err != nil {
		return 0, 0, fmt.Errorf("invalid --start %q: %w", startStr, err)
	}
	if end, err = parseTimeArg(endStr); err != nil {
		return 0, 0, fmt.Errorf("invalid --end %q: %w", endStr, err)
	}
	if start >= end {
		return 0, 0, fmt.Errorf("--start must be before --end")
	}
	return start, end, nil
}

// parseDurationArg parses a human-friendly duration such as "30m", "1h", "7d"
// (d = 24h). It extends time.ParseDuration with a day suffix.
func parseDurationArg(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty duration")
	}
	if strings.HasSuffix(s, "d") {
		n, err := strconv.ParseFloat(strings.TrimSuffix(s, "d"), 64)
		if err != nil {
			return 0, err
		}
		return time.Duration(n * float64(24*time.Hour)), nil
	}
	return time.ParseDuration(s)
}

// parseTimeArg parses a human-readable time into a Unix timestamp in seconds.
// Accepted formats: RFC3339 ("2023-11-14T22:13:20Z"),
// "YYYY-MM-DD HH:MM[:SS]", and "YYYY-MM-DD" (local midnight).
func parseTimeArg(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty time")
	}
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
	}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, s, time.Local); err == nil {
			return int(t.Unix()), nil
		}
	}
	return 0, fmt.Errorf("unsupported format (want RFC3339, YYYY-MM-DD HH:MM:SS, or YYYY-MM-DD)")
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
