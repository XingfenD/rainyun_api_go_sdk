package apis

import (
	"fmt"
	"io"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
)

type RyClient struct {
	APIKey string
	client *resty.Client
	debug  io.Writer
}

func NewRyClient(apiKey string) *RyClient {
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

// SetDebugWriter enables request/response tracing without exposing headers or
// response bodies. Passing nil disables tracing.
func (c *RyClient) SetDebugWriter(w io.Writer) {
	c.debug = w
}

func (c *RyClient) Do(method constant.HTTPMethod, endpoint string, querys map[string]string, body, result any) error {
	started := time.Now()
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

	if err != nil {
		if c.debug != nil {
			fmt.Fprintf(c.debug, "[debug] %s %s error=%v elapsed=%s\n", method, endpoint, err, time.Since(started).Round(time.Millisecond))
		}
		return err
	}
	if c.debug != nil {
		fmt.Fprintf(c.debug, "[debug] %s %s status=%d elapsed=%s\n", method, resp.Request.URL, resp.StatusCode(), time.Since(started).Round(time.Millisecond))
	}

	if resp.StatusCode() != 200 {
		var baseResp struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		if err := sonic.Unmarshal(resp.Body(), &baseResp); err != nil {
			return err
		}
		return fmt.Errorf("unexpected response: %+v", baseResp)
	}

	return nil
}
