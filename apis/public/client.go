package public

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type Service struct {
	client *apis.Client
}

func NewService(c *apis.Client) *Service {
	return &Service{client: c}
}
