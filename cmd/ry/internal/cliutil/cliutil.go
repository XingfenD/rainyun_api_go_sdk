// Package cliutil provides shared helpers for parsing CLI arguments.
package cliutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseID parses a command-line argument as an integer ID.
func ParseID(arg string) (int, error) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("invalid id %q", arg)
	}
	return id, nil
}

// ParseDuration parses a human-friendly duration such as "30m", "1h", "7d"
// (d = 24h). It extends time.ParseDuration with a day suffix.
func ParseDuration(s string) (time.Duration, error) {
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

// ParseTime parses a human-readable time into a Unix timestamp in seconds.
// Accepted formats: RFC3339 ("2023-11-14T22:13:20Z"),
// "YYYY-MM-DD HH:MM[:SS]", and "YYYY-MM-DD" (local midnight).
func ParseTime(s string) (int, error) {
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

// ResolveTimeRange turns --last/--start/--end flags into Unix second
// timestamps. When start and end are both empty, the last duration is used
// (end = now, start = now - duration). Explicit start/end take precedence
// over last.
func ResolveTimeRange(startStr, endStr, last string) (start, end int, err error) {
	if startStr == "" && endStr == "" {
		d, derr := ParseDuration(last)
		if derr != nil {
			return 0, 0, fmt.Errorf("invalid --last %q: %w", last, derr)
		}
		now := time.Now()
		return int(now.Add(-d).Unix()), int(now.Unix()), nil
	}
	if startStr == "" || endStr == "" {
		return 0, 0, fmt.Errorf("--start and --end must be specified together")
	}
	if start, err = ParseTime(startStr); err != nil {
		return 0, 0, fmt.Errorf("invalid --start %q: %w", startStr, err)
	}
	if end, err = ParseTime(endStr); err != nil {
		return 0, 0, fmt.Errorf("invalid --end %q: %w", endStr, err)
	}
	if start >= end {
		return 0, 0, fmt.Errorf("--start must be before --end")
	}
	return start, end, nil
}
