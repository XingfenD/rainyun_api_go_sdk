package ssl

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type SslService struct {
	client *apis.RyClient
}

func NewSslService(c *apis.RyClient) *SslService {
	return &SslService{client: c}
}
