package ros

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type RosService struct {
	client *apis.RyClient
}

func NewRosService(c *apis.RyClient) *RosService {
	return &RosService{client: c}
}
