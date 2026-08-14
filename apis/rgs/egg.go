package rgs

import (
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// RgsEggServer 游戏云服务端类型
type RgsEggServer struct {
	Server      string `json:"server"`
	Desc        string `json:"desc"`
	OfficialURL string `json:"official_url"`
	IconURL     string `json:"icon_url"`
	Order       int    `json:"order"`
	EggName     string `json:"egg_name"`
}

type GetRgsEggServerListResponse struct {
	Code int            `json:"code"`
	Data []RgsEggServer `json:"data"`
}

// 获取游戏云服务端类型列表
func (s *RgsService) GetRgsEggServerList() (*GetRgsEggServerListResponse, error) {
	path := "/product/rgs/egg_server"

	var resp GetRgsEggServerListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
