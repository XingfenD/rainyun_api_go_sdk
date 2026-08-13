package rvh

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type RvhService struct {
	client *apis.RyClient
}

func NewRvhService(c *apis.RyClient) *RvhService {
	return &RvhService{client: c}
}
