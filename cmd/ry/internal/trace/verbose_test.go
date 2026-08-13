package trace

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

func TestFormatPreview(t *testing.T) {
	tests := []struct {
		name string
		body []byte
		want string
	}{
		{
			name: "JSON is indented",
			body: []byte(`{"code":0,"data":{"id":1}}`),
			want: "{\n  \"code\": 0,\n  \"data\": {\n    \"id\": 1\n  }\n}",
		},
		{name: "non JSON unchanged", body: []byte("plain text"), want: "plain text"},
		{name: "empty prints empty marker", want: "(empty)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatPreview(tt.body); got != tt.want {
				t.Errorf("formatPreview() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestVerboseRendererHTTPTrace(t *testing.T) {
	var buf bytes.Buffer
	r := NewVerboseRenderer(&buf)
	r.OnHTTPTrace(apis.HTTPTrace{
		Method:        constant.HTTPMethod_GET,
		URL:           "https://api.v2.rainyun.com/test",
		StatusCode:    200,
		Elapsed:       123 * time.Millisecond,
		ContentType:   "application/json",
		ResponseBytes: 100,
		BodyPreview:   []byte(`{"a":1}`),
		BodyTruncated: true,
	})

	got := buf.String()
	for _, want := range []string{
		"[debug] GET https://api.v2.rainyun.com/test status=200",
		"elapsed=123ms",
		"content-type: application/json",
		"response bytes: 100",
		"[debug] response body truncated: shown=7B total=100B",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output missing %q:\n%s", want, got)
		}
	}
}

func TestVerboseRendererEmptyPreview(t *testing.T) {
	var buf bytes.Buffer
	r := NewVerboseRenderer(&buf)
	r.OnHTTPTrace(apis.HTTPTrace{Method: constant.HTTPMethod_GET, URL: "/test", StatusCode: 200})

	if !strings.Contains(buf.String(), "(empty)") {
		t.Errorf("output missing %q:\n%s", "(empty)", buf.String())
	}
}

func TestVerboseRendererHTTPTraceError(t *testing.T) {
	var buf bytes.Buffer
	r := NewVerboseRenderer(&buf)
	r.OnHTTPTrace(apis.HTTPTrace{Method: constant.HTTPMethod_GET, URL: "/test", Err: errors.New("boom")})

	if !strings.Contains(buf.String(), "error=boom") {
		t.Errorf("output missing %q:\n%s", "error=boom", buf.String())
	}
}

type sampleResult struct{ ID int }

func TestVerboseRendererResultTrace(t *testing.T) {
	var buf bytes.Buffer
	r := NewVerboseRenderer(&buf)
	res := &sampleResult{ID: 7}
	r.OnResultTrace(apis.ResultTrace{Result: res})

	want := fmt.Sprintf("[debug] result: type=%T\n", res)
	if buf.String() != want {
		t.Errorf("output = %q, want %q", buf.String(), want)
	}
}
