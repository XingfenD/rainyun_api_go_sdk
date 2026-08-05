package rbm

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type RbmService struct {
	client *apis.RyClient
}

func NewRbmService(c *apis.RyClient) *RbmService {
	return &RbmService{client: c}
}
