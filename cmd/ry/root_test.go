package main

import (
	"reflect"
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis"
)

func TestResolveOutput(t *testing.T) {
	tests := []struct {
		name                     string
		configFormat, flagFormat string
		wantFormat, wantSource   string
	}{
		{name: "config", configFormat: "table", wantFormat: "table", wantSource: "config"},
		{name: "output flag", configFormat: "table", flagFormat: "json", wantFormat: "json", wantSource: "--output"},
		{name: "raw via output flag", configFormat: "table", flagFormat: "raw", wantFormat: "raw", wantSource: "--output"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			format, source := resolveOutput(tt.configFormat, tt.flagFormat)
			if format != tt.wantFormat || source != tt.wantSource {
				t.Errorf("resolveOutput() = (%q, %q), want (%q, %q)", format, source, tt.wantFormat, tt.wantSource)
			}
		})
	}
}

type noopTraceSink struct{}

func (noopTraceSink) OnHTTPTrace(apis.HTTPTrace)     {}
func (noopTraceSink) OnResultTrace(apis.ResultTrace) {}

func TestVerboseTraceOptions(t *testing.T) {
	sink := noopTraceSink{}
	wantDefault := apis.NewTraceOptions(sink)
	wantNoBody := apis.NewTraceOptions(sink).WithoutBodyPreview()
	wantLimited := apis.NewTraceOptions(sink).WithBodyPreviewLimit(128)
	wantFull := apis.NewTraceOptions(sink).WithFullBodyPreview()

	tests := []struct {
		name    string
		enabled bool
		limit   int
		full    bool
		want    *apis.TraceOptions
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
