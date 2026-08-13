package apis

import (
	"bytes"
	"sync/atomic"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// DefaultBodyPreviewLimit is the default maximum number of response body bytes
// captured in an HTTPTrace.
const DefaultBodyPreviewLimit = 64 * 1024

// TraceSink receives structured trace events emitted by the SDK. Events are
// delivered synchronously and in request order.
type TraceSink interface {
	OnHTTPTrace(HTTPTrace)
	OnResultTrace(ResultTrace)
}

// TraceOptions configures structured tracing for a RyClient. A nil Sink
// disables tracing.
type TraceOptions struct {
	Sink             TraceSink
	BodyPreviewLimit int
	CaptureBody      bool
}

// unlimitedPreviewLimit marks a caller-selected full-body capture policy.
const unlimitedPreviewLimit = -1

// NewTraceOptions returns tracing options with a default bounded body preview.
func NewTraceOptions(sink TraceSink) TraceOptions {
	return TraceOptions{
		Sink:             sink,
		BodyPreviewLimit: DefaultBodyPreviewLimit,
		CaptureBody:      true,
	}
}

// WithBodyPreviewLimit bounds the response body preview to limit bytes.
// Non-positive limits disable body capture.
func (o TraceOptions) WithBodyPreviewLimit(limit int) TraceOptions {
	if limit <= 0 {
		return o.WithoutBodyPreview()
	}
	o.CaptureBody = true
	o.BodyPreviewLimit = limit
	return o
}

// WithoutBodyPreview disables response body capture.
func (o TraceOptions) WithoutBodyPreview() TraceOptions {
	o.CaptureBody = false
	o.BodyPreviewLimit = 0
	return o
}

// WithFullBodyPreview captures the complete response body.
func (o TraceOptions) WithFullBodyPreview() TraceOptions {
	o.CaptureBody = true
	o.BodyPreviewLimit = unlimitedPreviewLimit
	return o
}

// bodyPreviewPolicy is the normalized, unambiguous preview policy derived from
// TraceOptions: disabled, bounded, or full-body capture.
type bodyPreviewPolicy struct {
	enabled   bool
	unlimited bool
	limit     int
}

func (o TraceOptions) previewPolicy() bodyPreviewPolicy {
	if o.Sink == nil || !o.CaptureBody {
		return bodyPreviewPolicy{}
	}
	if o.BodyPreviewLimit == unlimitedPreviewLimit {
		return bodyPreviewPolicy{enabled: true, unlimited: true}
	}
	limit := o.BodyPreviewLimit
	if limit <= 0 {
		limit = DefaultBodyPreviewLimit
	}
	return bodyPreviewPolicy{enabled: true, limit: limit}
}

// HTTPTrace represents one completed HTTP attempt. It never contains API keys,
// headers, or request bodies.
type HTTPTrace struct {
	RequestID     uint64
	Method        constant.HTTPMethod
	URL           string
	StatusCode    int
	Elapsed       time.Duration
	ContentType   string
	ResponseBytes int
	BodyPreview   []byte
	BodyTruncated bool
	Err           error
}

// ResultTrace represents a successfully parsed SDK result.
type ResultTrace struct {
	RequestID uint64
	Result    any
}

var traceIDCounter atomic.Uint64

func nextTraceID() uint64 {
	return traceIDCounter.Add(1)
}

// sliceBodyPreview returns a bounded view of body per policy and whether the
// body was truncated. The view may alias body; emit helpers copy it before
// dispatch.
func sliceBodyPreview(body []byte, p bodyPreviewPolicy) (preview []byte, truncated bool) {
	if !p.enabled || len(body) == 0 {
		return nil, false
	}
	if p.unlimited || len(body) <= p.limit {
		return body, false
	}
	return body[:p.limit], true
}

// emitHTTPTrace delivers an HTTP trace event to the sink. The preview is copied
// so the sink owns its bytes, and a panicking sink never alters the API call.
func emitHTTPTrace(sink TraceSink, ev HTTPTrace) {
	if sink == nil {
		return
	}
	ev.BodyPreview = bytes.Clone(ev.BodyPreview)
	defer func() { _ = recover() }()
	sink.OnHTTPTrace(ev)
}

// emitResultTrace delivers a result trace event to the sink. A panicking sink
// never alters the API call.
func emitResultTrace(sink TraceSink, ev ResultTrace) {
	if sink == nil {
		return
	}
	defer func() { _ = recover() }()
	sink.OnResultTrace(ev)
}
