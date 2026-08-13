package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 游戏云详情
// 由于响应实在过于庞大，我们只维护部分必要的响应，如有扩展需求，请在项目的data文件夹下寻找响应实例自行解码
type GetRgsDetailResponse struct {
	Code int           `json:"code"`
	Data RgsDetailData `json:"data"`
}

type RgsDetailData struct {
	Data            RgsRecord          `json:"Data"`
	NatList         []RgsNatItem       `json:"NatList"` // 端口映射列表
	EIPList         any                `json:"EIPList"` // TODO: 结构未公开,实测后补强类型
	ConfigPrice     int                `json:"ConfigPrice"`
	RenewPointPrice RgsRenewPointPrice `json:"RenewPointPrice"` // 积分续费
}

type RgsNatItem struct {
	ID       int    `json:"ID"`
	PortIn   int    `json:"PortIn"`
	PortOut  int    `json:"PortOut"`
	PortType string `json:"PortType"`
	Tag      string `json:"Tag"`
}

type RgsRenewPointPrice struct {
	Num7  int `json:"7"`  // 积分续费七天
	Num31 int `json:"31"` // 积分续费31天
}

// 获取游戏云详情
//
// id： 游戏云 ID
func (s *RgsService) GetRgsDetail(id int) (*GetRgsDetailResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d", id)

	var resp GetRgsDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
