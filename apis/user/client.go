package user

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type UserService struct {
	client *apis.RyClient
}

func NewUserService(c *apis.RyClient) *UserService {
	return &UserService{client: c}
}
