package rgs

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type RgsService struct {
	client *apis.RyClient
}

func NewRgsService(c *apis.RyClient) *RgsService {
	return &RgsService{client: c}
}
