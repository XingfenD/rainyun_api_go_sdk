package apis

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

const sampleResponseBody = `{"code":0,"data":{"id":1},"list":[1,2,3,4,5,6,7,8,9,10]}`

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type seqSink struct {
	order  []string
	http   HTTPTrace
	result ResultTrace
}

func (s *seqSink) OnHTTPTrace(ev HTTPTrace) {
	s.order = append(s.order, "http")
	s.http = ev
}

func (s *seqSink) OnResultTrace(ev ResultTrace) {
	s.order = append(s.order, "result")
	s.result = ev
}

type countSink struct {
	httpCount   int
	resultCount int
}

func (s *countSink) OnHTTPTrace(HTTPTrace)     { s.httpCount++ }
func (s *countSink) OnResultTrace(ResultTrace) { s.resultCount++ }

type panicSink struct{}

func (panicSink) OnHTTPTrace(HTTPTrace)     { panic("sink panic") }
func (panicSink) OnResultTrace(ResultTrace) { panic("sink panic") }

type testResult struct {
	Code int `json:"code"`
	Data struct {
		ID int `json:"id"`
	} `json:"data"`
}

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func newTracedClient(sink TraceSink, options TraceOptions) *RyClient {
	c := NewRyClientWithTrace("test-key", options)
	c.client.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return jsonResponse(200, sampleResponseBody), nil
	}))
	return c
}

func TestDoEmitsHTTPThenResult(t *testing.T) {
	sink := &seqSink{}
	c := newTracedClient(sink, NewTraceOptions(sink).WithBodyPreviewLimit(4))

	var result testResult
	if err := c.Do(constant.HTTPMethod_GET, "/test", nil, nil, &result); err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if want := []string{"http", "result"}; !reflect.DeepEqual(sink.order, want) {
		t.Errorf("event order = %v, want %v", sink.order, want)
	}
	if sink.http.ResponseBytes != len(sampleResponseBody) {
		t.Errorf("ResponseBytes = %d, want %d", sink.http.ResponseBytes, len(sampleResponseBody))
	}
	if want := sampleResponseBody[:4]; string(sink.http.BodyPreview) != want {
		t.Errorf("BodyPreview = %q, want %q", sink.http.BodyPreview, want)
	}
	if !sink.http.BodyTruncated {
		t.Error("BodyTruncated = false, want true")
	}
	if sink.http.StatusCode != 200 {
		t.Errorf("StatusCode = %d, want 200", sink.http.StatusCode)
	}
	if sink.result.RequestID != sink.http.RequestID {
		t.Errorf("result RequestID = %d, http RequestID = %d", sink.result.RequestID, sink.http.RequestID)
	}
	if sink.result.Result != &result {
		t.Errorf("Result = %v, want %v", sink.result.Result, &result)
	}
}

func TestDoEmitsNoCapture(t *testing.T) {
	sink := &seqSink{}
	c := newTracedClient(sink, NewTraceOptions(sink).WithoutBodyPreview())

	var result testResult
	if err := c.Do(constant.HTTPMethod_GET, "/test", nil, nil, &result); err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if len(sink.http.BodyPreview) != 0 {
		t.Errorf("BodyPreview len = %d, want 0", len(sink.http.BodyPreview))
	}
	if sink.http.BodyTruncated {
		t.Error("BodyTruncated = true, want false")
	}
	if sink.http.ResponseBytes != len(sampleResponseBody) {
		t.Errorf("ResponseBytes = %d, want %d", sink.http.ResponseBytes, len(sampleResponseBody))
	}
}

func TestDoEmitsFullCapture(t *testing.T) {
	sink := &seqSink{}
	c := newTracedClient(sink, NewTraceOptions(sink).WithFullBodyPreview())

	var result testResult
	if err := c.Do(constant.HTTPMethod_GET, "/test", nil, nil, &result); err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if string(sink.http.BodyPreview) != sampleResponseBody {
		t.Errorf("BodyPreview = %q, want full body", sink.http.BodyPreview)
	}
	if sink.http.BodyTruncated {
		t.Error("BodyTruncated = true, want false")
	}
}

func TestDoEmitsNon200OnlyHTTP(t *testing.T) {
	sink := &countSink{}
	c := NewRyClientWithTrace("test-key", NewTraceOptions(sink))
	c.client.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return jsonResponse(500, `{"code":500,"message":"boom"}`), nil
	}))

	var result testResult
	if err := c.Do(constant.HTTPMethod_GET, "/test", nil, nil, &result); err == nil {
		t.Fatal("Do() error = nil, want non-nil")
	}

	if sink.httpCount != 1 {
		t.Errorf("http events = %d, want 1", sink.httpCount)
	}
	if sink.resultCount != 0 {
		t.Errorf("result events = %d, want 0", sink.resultCount)
	}
}

func TestDoEmitsTransportErrorOnlyHTTP(t *testing.T) {
	sink := &seqSink{}
	c := NewRyClientWithTrace("test-key", NewTraceOptions(sink))
	c.client.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("connection refused")
	}))

	var result testResult
	if err := c.Do(constant.HTTPMethod_GET, "/test", nil, nil, &result); err == nil {
		t.Fatal("Do() error = nil, want non-nil")
	}

	if len(sink.order) != 1 || sink.order[0] != "http" {
		t.Errorf("event order = %v, want [http]", sink.order)
	}
	if sink.http.Err == nil {
		t.Error("http event Err = nil, want non-nil")
	}
	if sink.http.StatusCode != 0 {
		t.Errorf("StatusCode = %d, want 0", sink.http.StatusCode)
	}
}

func TestDoEmitsDecodeErrorOnlyHTTP(t *testing.T) {
	sink := &countSink{}
	c := NewRyClientWithTrace("test-key", NewTraceOptions(sink))
	c.client.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return jsonResponse(200, "not-json"), nil
	}))

	var result testResult
	if err := c.Do(constant.HTTPMethod_GET, "/test", nil, nil, &result); err == nil {
		t.Fatal("Do() error = nil, want decode error")
	}

	if sink.httpCount != 1 {
		t.Errorf("http events = %d, want 1", sink.httpCount)
	}
	if sink.resultCount != 0 {
		t.Errorf("result events = %d, want 0", sink.resultCount)
	}
}

func TestDoEmitsSinkPanicDoesNotAlterOutcome(t *testing.T) {
	c := NewRyClientWithTrace("test-key", NewTraceOptions(panicSink{}))
	c.client.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return jsonResponse(200, sampleResponseBody), nil
	}))

	var result testResult
	if err := c.Do(constant.HTTPMethod_GET, "/test", nil, nil, &result); err != nil {
		t.Fatalf("Do() error = %v, want nil despite sink panic", err)
	}
}

func TestDoEmitsNothingWithoutSink(t *testing.T) {
	c := NewRyClient("test-key")
	c.client.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return jsonResponse(200, sampleResponseBody), nil
	}))

	var result testResult
	if err := c.Do(constant.HTTPMethod_GET, "/test", nil, nil, &result); err != nil {
		t.Fatalf("Do() error = %v", err)
	}
}
