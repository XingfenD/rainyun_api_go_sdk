package domain

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type DomainService struct {
	client *apis.RyClient
}

func NewDomainService(c *apis.RyClient) *DomainService {
	return &DomainService{client: c}
}
