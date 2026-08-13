package cliutil

import (
	"testing"
	"time"
)

func TestParseID(t *testing.T) {
	if got, err := ParseID("370018"); err != nil || got != 370018 {
		t.Errorf("ParseID(370018) = %d, %v; want 370018", got, err)
	}
	for _, in := range []string{"", "abc", "12x", "99999999999999999999"} {
		if _, err := ParseID(in); err == nil {
			t.Errorf("ParseID(%q) expected error, got nil", in)
		}
	}
}

func TestParseDuration(t *testing.T) {
	cases := map[string]time.Duration{
		"90s":  90 * time.Second,
		"30m":  30 * time.Minute,
		"1h":   time.Hour,
		"1.5h": 90 * time.Minute,
		"2d":   48 * time.Hour,
		"7d":   7 * 24 * time.Hour,
	}
	for in, want := range cases {
		got, err := ParseDuration(in)
		if err != nil {
			t.Errorf("ParseDuration(%q) error: %v", in, err)
			continue
		}
		if got != want {
			t.Errorf("ParseDuration(%q) = %v, want %v", in, got, want)
		}
	}

	for _, in := range []string{"", "abc", "1x"} {
		if _, err := ParseDuration(in); err == nil {
			t.Errorf("ParseDuration(%q) expected error, got nil", in)
		}
	}
}

func TestParseTime(t *testing.T) {
	if got, err := ParseTime("2023-11-14T22:13:20Z"); err != nil || got != 1700000000 {
		t.Errorf("RFC3339 = %d, %v; want 1700000000", got, err)
	}
	wantLocal := int(time.Date(2023, 11, 14, 0, 0, 0, 0, time.Local).Unix())
	if got, err := ParseTime("2023-11-14"); err != nil || got != wantLocal {
		t.Errorf("date-only = %d, %v; want %d", got, err, wantLocal)
	}
	wantLocalDT := int(time.Date(2023, 11, 14, 22, 13, 20, 0, time.Local).Unix())
	if got, err := ParseTime("2023-11-14 22:13:20"); err != nil || got != wantLocalDT {
		t.Errorf("date-time = %d, %v; want %d", got, err, wantLocalDT)
	}

	for _, in := range []string{"", "abc", "11/14/2023"} {
		if _, err := ParseTime(in); err == nil {
			t.Errorf("ParseTime(%q) expected error, got nil", in)
		}
	}
}

func TestResolveTimeRange(t *testing.T) {
	start, end, err := ResolveTimeRange("", "", "1h")
	if err != nil {
		t.Fatalf("--last only: %v", err)
	}
	if diff := end - start; diff < 3599 || diff > 3601 {
		t.Errorf("--last 1h span = %d, want ~3600", diff)
	}

	start, end, err = ResolveTimeRange("2023-11-14T00:00:00Z", "2023-11-14T01:00:00Z", "")
	if err != nil {
		t.Fatalf("explicit range: %v", err)
	}
	if start != 1699920000 || end != 1699923600 {
		t.Errorf("explicit range = %d..%d, want 1699920000..1699923600", start, end)
	}

	for _, tc := range []struct{ s, e, l string }{
		{"2023-11-14T00:00:00Z", "", ""},
		{"2023-11-14T01:00:00Z", "2023-11-14T00:00:00Z", ""},
		{"", "", "abc"},
	} {
		if _, _, err := ResolveTimeRange(tc.s, tc.e, tc.l); err == nil {
			t.Errorf("ResolveTimeRange(%q, %q, %q) expected error", tc.s, tc.e, tc.l)
		}
	}
}
