package rcs

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type RcsService struct {
	client *apis.RyClient
}

func NewRcsService(c *apis.RyClient) *RcsService {
	return &RcsService{client: c}
}
