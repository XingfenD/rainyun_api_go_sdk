package rgs

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRgsPlanListResponse struct {
	Code int       `json:"code"`
	Data []RgsPlan `json:"data"`
}

// 获取游戏云套餐列表
func (s *RgsService) GetRgsPlanList() (*GetRgsPlanListResponse, error) {
	path := "/product/rgs/plans"

	var resp GetRgsPlanListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// GetRgsDiscountPercentRequest 获取游戏云折扣比率请求
// TODO: 结构未公开,实测后补强类型(线上返回10002输入参数无效,参数未知)
type GetRgsDiscountPercentRequest struct {
	Options map[string]any `json:"options,omitempty"`
}

type GetRgsDiscountPercentResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型
}

// 获取游戏云折扣比率
func (s *RgsService) GetRgsDiscountPercent(req *GetRgsDiscountPercentRequest) (*GetRgsDiscountPercentResponse, error) {
	path := "/product/rgs/discount-percent"

	var resp GetRgsDiscountPercentResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
