package workorder

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type WorkorderService struct {
	client *apis.RyClient
}

func NewWorkorderService(c *apis.RyClient) *WorkorderService {
	return &WorkorderService{client: c}
}
