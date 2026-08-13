package apis

import (
	"bytes"
	"testing"
)

type captureSink struct {
	httpEvents   []HTTPTrace
	resultEvents []ResultTrace
	panicOnHTTP  bool
}

func (s *captureSink) OnHTTPTrace(ev HTTPTrace) {
	if s.panicOnHTTP {
		panic("sink panic")
	}
	s.httpEvents = append(s.httpEvents, ev)
}

func (s *captureSink) OnResultTrace(ev ResultTrace) {
	s.resultEvents = append(s.resultEvents, ev)
}

func TestTraceOptionsDefaults(t *testing.T) {
	opts := NewTraceOptions(&captureSink{})
	if !opts.CaptureBody {
		t.Error("CaptureBody = false, want true")
	}
	if opts.BodyPreviewLimit != DefaultBodyPreviewLimit {
		t.Errorf("BodyPreviewLimit = %d, want %d", opts.BodyPreviewLimit, DefaultBodyPreviewLimit)
	}
	if opts.Sink == nil {
		t.Error("Sink = nil, want non-nil")
	}
}

func TestTraceOptionsPolicies(t *testing.T) {
	sink := &captureSink{}
	tests := []struct {
		name    string
		options TraceOptions
		want    bodyPreviewPolicy
	}{
		{
			name:    "default bounded at 64 KiB",
			options: NewTraceOptions(sink),
			want:    bodyPreviewPolicy{enabled: true, limit: DefaultBodyPreviewLimit},
		},
		{
			name:    "limited",
			options: NewTraceOptions(sink).WithBodyPreviewLimit(128),
			want:    bodyPreviewPolicy{enabled: true, limit: 128},
		},
		{
			name:    "disabled via WithoutBodyPreview",
			options: NewTraceOptions(sink).WithoutBodyPreview(),
			want:    bodyPreviewPolicy{},
		},
		{
			name:    "disabled via non-positive limit",
			options: NewTraceOptions(sink).WithBodyPreviewLimit(0),
			want:    bodyPreviewPolicy{},
		},
		{
			name:    "full body",
			options: NewTraceOptions(sink).WithFullBodyPreview(),
			want:    bodyPreviewPolicy{enabled: true, unlimited: true},
		},
		{
			name:    "zero value disables",
			options: TraceOptions{},
			want:    bodyPreviewPolicy{},
		},
		{
			name:    "nil sink disables",
			options: TraceOptions{CaptureBody: true, BodyPreviewLimit: 1024},
			want:    bodyPreviewPolicy{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.options.previewPolicy(); got != tt.want {
				t.Errorf("previewPolicy() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestTraceOptionsBoundedPreview(t *testing.T) {
	body := []byte(`{"list":[1,2,3,4,5,6,7,8,9,10]}`)

	tests := []struct {
		name      string
		policy    bodyPreviewPolicy
		wantLen   int
		wantTrunc bool
	}{
		{name: "disabled", policy: bodyPreviewPolicy{}, wantLen: 0},
		{name: "bounded", policy: bodyPreviewPolicy{enabled: true, limit: 4}, wantLen: 4, wantTrunc: true},
		{name: "limit larger than body", policy: bodyPreviewPolicy{enabled: true, limit: 1 << 20}, wantLen: len(body)},
		{name: "full", policy: bodyPreviewPolicy{enabled: true, unlimited: true}, wantLen: len(body)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preview, truncated := sliceBodyPreview(body, tt.policy)
			if len(preview) != tt.wantLen {
				t.Errorf("preview len = %d, want %d", len(preview), tt.wantLen)
			}
			if truncated != tt.wantTrunc {
				t.Errorf("truncated = %t, want %t", truncated, tt.wantTrunc)
			}
		})
	}
}

func TestTraceOptionsIDsUnique(t *testing.T) {
	seen := map[uint64]bool{}
	for i := 0; i < 100; i++ {
		id := nextTraceID()
		if seen[id] {
			t.Fatalf("duplicate trace id %d", id)
		}
		seen[id] = true
	}
}

func TestEmitHTTPTraceCopiesPreview(t *testing.T) {
	sink := &captureSink{}
	original := []byte(`{"code":0}`)
	emitHTTPTrace(sink, HTTPTrace{BodyPreview: original})

	original[0] = 'X'
	if got := sink.httpEvents[0].BodyPreview; !bytes.Equal(got, []byte(`{"code":0}`)) {
		t.Errorf("sink preview = %q, want independent copy", got)
	}
}

func TestEmitPanicRecovery(t *testing.T) {
	sink := &captureSink{panicOnHTTP: true}
	emitHTTPTrace(sink, HTTPTrace{})
	emitResultTrace(sink, ResultTrace{})
}

func TestEmitNilSinkNoOp(t *testing.T) {
	emitHTTPTrace(nil, HTTPTrace{BodyPreview: []byte("x")})
	emitResultTrace(nil, ResultTrace{})
}
