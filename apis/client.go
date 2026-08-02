package apis

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	APIKey string
	client *resty.Client
}

func New(apiKey string) *Client {
	c := resty.New()
	c.SetBaseURL(constant.BaseURL)
	c.SetHeader("x-api-key", apiKey)
	c.SetHeader("User-Agent", "github.com/XingfenD/rainyun_api_go_sdk/"+constant.RainyunSdkVersion)
	c.SetHeader("Accept", "application/json")

	return &Client{
		APIKey: apiKey,
		client: c,
	}
}

func (c *Client) Do(method constant.HTTPMethod, endpoint string, body, result any) error {
	req := c.client.R()
	if body != nil {
		req.SetBody(body)
	}

	if result != nil {
		req.SetResult(result)
	}
	resp, err := req.Execute(string(method), endpoint)

	if err != nil {
		return err
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
