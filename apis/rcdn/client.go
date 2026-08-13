package rcdn

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type RcdnService struct {
	client *apis.RyClient
}

func NewRcdnService(c *apis.RyClient) *RcdnService {
	return &RcdnService{client: c}
}
