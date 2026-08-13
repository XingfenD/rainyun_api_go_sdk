package rca

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type RcaService struct {
	client *apis.RyClient
}

func NewRcaService(c *apis.RyClient) *RcaService {
	return &RcaService{client: c}
}
