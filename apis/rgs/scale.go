package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// ScaleRgsRequest 游戏云升级(缩放配置)请求
type ScaleRgsRequest struct {
	DestConfig   RgsConfig `json:"dest_config"`
	DestPlan     int       `json:"dest_plan,omitempty"`
	WithCouponID int       `json:"with_coupon_id,omitempty"`
}

type ScaleRgsResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型(计费操作,不做线上探测)
}

// 游戏云升级(缩放配置)
//
// id: 游戏云 ID
func (s *RgsService) ScaleRgs(id int, req *ScaleRgsRequest) (*ScaleRgsResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/scale", id)

	var resp ScaleRgsResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
