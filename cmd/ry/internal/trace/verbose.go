package trace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/sdk"
)

// VerboseRenderer renders SDK trace events as human-readable diagnostics.
// It writes to stderr only; normal command output stays on stdout.
type VerboseRenderer struct {
	Writer io.Writer
}

func NewVerboseRenderer(w io.Writer) *VerboseRenderer {
	return &VerboseRenderer{Writer: w}
}

func (r *VerboseRenderer) OnHTTPTrace(ev sdk.HTTPTrace) {
	if ev.Err != nil {
		fmt.Fprintf(r.Writer, "[debug] %s %s error=%v\n", ev.Method, ev.URL, ev.Err)
		return
	}
	fmt.Fprintf(r.Writer, "[debug] %s %s status=%d elapsed=%s\n", ev.Method, ev.URL, ev.StatusCode, ev.Elapsed.Round(time.Millisecond))
	if ev.ContentType != "" {
		fmt.Fprintf(r.Writer, "[debug] content-type: %s\n", ev.ContentType)
	}
	fmt.Fprintf(r.Writer, "[debug] response bytes: %d\n", ev.ResponseBytes)
	fmt.Fprintf(r.Writer, "[debug] response:\n%s\n", formatPreview(ev.BodyPreview))
	if ev.BodyTruncated {
		fmt.Fprintf(r.Writer, "[debug] response body truncated: shown=%dB total=%dB\n", len(ev.BodyPreview), ev.ResponseBytes)
	}
}

func (r *VerboseRenderer) OnResultTrace(ev sdk.ResultTrace) {
	fmt.Fprintf(r.Writer, "[debug] result: type=%T\n", ev.Result)
}

func formatPreview(body []byte) string {
	if len(body) == 0 {
		return "(empty)"
	}

	var pretty bytes.Buffer
	if err := json.Indent(&pretty, body, "", "  "); err == nil {
		return pretty.String()
	}
	return string(body)
}
