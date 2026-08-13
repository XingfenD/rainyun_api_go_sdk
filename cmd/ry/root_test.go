package main

import (
	"reflect"
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/sdk"
)

func TestResolveOutput(t *testing.T) {
	tests := []struct {
		name                     string
		configFormat, flagFormat string
		raw                      bool
		wantFormat, wantSource   string
	}{
		{name: "config", configFormat: "table", wantFormat: "table", wantSource: "config"},
		{name: "output flag", configFormat: "table", flagFormat: "json", wantFormat: "json", wantSource: "--output"},
		{name: "raw overrides output flag", configFormat: "table", flagFormat: "yaml", raw: true, wantFormat: "raw", wantSource: "--raw"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			format, source := resolveOutput(tt.configFormat, tt.flagFormat, tt.raw)
			if format != tt.wantFormat || source != tt.wantSource {
				t.Errorf("resolveOutput() = (%q, %q), want (%q, %q)", format, source, tt.wantFormat, tt.wantSource)
			}
		})
	}
}

type noopTraceSink struct{}

func (noopTraceSink) OnHTTPTrace(sdk.HTTPTrace)     {}
func (noopTraceSink) OnResultTrace(sdk.ResultTrace) {}

func TestVerboseTraceOptions(t *testing.T) {
	sink := noopTraceSink{}
	wantDefault := sdk.NewTraceOptions(sink)
	wantNoBody := sdk.NewTraceOptions(sink).WithoutBodyPreview()
	wantLimited := sdk.NewTraceOptions(sink).WithBodyPreviewLimit(128)
	wantFull := sdk.NewTraceOptions(sink).WithFullBodyPreview()

	tests := []struct {
		name    string
		enabled bool
		limit   int
		full    bool
		want    *sdk.TraceOptions
	}{
		{name: "disabled", enabled: false, want: nil},
		{name: "default is 65536", enabled: true, limit: 64 * 1024, want: &wantDefault},
		{name: "zero disables preview", enabled: true, limit: 0, want: &wantNoBody},
		{name: "positive limit retained", enabled: true, limit: 128, want: &wantLimited},
		{name: "full overrides numeric limit", enabled: true, limit: 128, full: true, want: &wantFull},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := verboseTraceOptions(tt.enabled, tt.limit, tt.full, sink)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verboseTraceOptions() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
