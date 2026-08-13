package apis

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
)

type RyClient struct {
	APIKey      string
	client      *resty.Client
	traceSink   TraceSink
	tracePolicy bodyPreviewPolicy
}

func NewRyClient(apiKey string) *RyClient {
	return newRyClient(apiKey)
}

// NewRyClientWithTrace creates a client that emits structured trace events to
// options.Sink. A nil sink disables tracing.
func NewRyClientWithTrace(apiKey string, options TraceOptions) *RyClient {
	c := newRyClient(apiKey)
	c.traceSink = options.Sink
	c.tracePolicy = options.previewPolicy()
	return c
}

func newRyClient(apiKey string) *RyClient {
	c := resty.New()
	c.SetBaseURL(constant.BaseURL)
	c.SetHeader("x-api-key", apiKey)
	c.SetHeader("User-Agent", "github.com/XingfenD/rainyun_api_go_sdk/"+constant.RainyunSdkVersion)
	c.SetHeader("Accept", "application/json")

	return &RyClient{
		APIKey: apiKey,
		client: c,
	}
}

func (c *RyClient) Do(method constant.HTTPMethod, endpoint string, querys map[string]string, body, result any) error {
	requestID := nextTraceID()
	req := c.client.R()
	if querys != nil {
		req.SetQueryParams(querys)
	}

	if body != nil {
		req.SetBody(body)
	}

	if result != nil {
		req.SetResult(result)
	}
	resp, err := req.Execute(string(method), endpoint)

	ev := c.httpTrace(requestID, method, endpoint, resp, err)

	if err == nil && resp.StatusCode() != 200 {
		var baseResp struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		if derr := sonic.Unmarshal(resp.Body(), &baseResp); derr != nil {
			err = derr
		} else {
			err = fmt.Errorf("unexpected response: %+v", baseResp)
		}
		ev.Err = err
		emitHTTPTrace(c.traceSink, ev)
		return err
	}

	emitHTTPTrace(c.traceSink, ev)
	if err != nil {
		return err
	}

	if result != nil {
		emitResultTrace(c.traceSink, ResultTrace{RequestID: requestID, Result: result})
	}

	return nil
}

func (c *RyClient) httpTrace(requestID uint64, method constant.HTTPMethod, endpoint string, resp *resty.Response, err error) HTTPTrace {
	ev := HTTPTrace{
		RequestID: requestID,
		Method:    method,
		URL:       endpoint,
		Err:       err,
	}
	if resp == nil {
		return ev
	}
	ev.URL = resp.Request.URL
	ev.StatusCode = resp.StatusCode()
	ev.Elapsed = resp.Time()
	ev.ContentType = resp.Header().Get("Content-Type")
	ev.ResponseBytes = len(resp.Body())
	ev.BodyPreview, ev.BodyTruncated = sliceBodyPreview(resp.Body(), c.tracePolicy)
	return ev
}
