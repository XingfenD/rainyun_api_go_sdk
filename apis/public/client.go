package public

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type PublicService struct {
	client *apis.RyClient
}

func NewPublicService(c *apis.RyClient) *PublicService {
	return &PublicService{client: c}
}
