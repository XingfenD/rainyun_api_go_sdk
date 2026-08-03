package public

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type PublicService struct {
	client *apis.Client
}

func NewPublicService(c *apis.Client) *PublicService {
	return &PublicService{client: c}
}
