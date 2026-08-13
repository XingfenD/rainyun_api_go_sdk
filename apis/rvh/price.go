package rvh

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 获取虚拟主机折扣比率(透传)
func (s *RvhService) GetRvhRenewPrice(id int) (*GetRvhRenewPriceResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/renew/", id)

	var resp GetRvhRenewPriceResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 获取虚拟主机套餐列表
func (s *RvhService) GetRvhPlanList() (*GetRvhPlanListResponse, error) {
	path := "/product/rvh/plans"

	var resp GetRvhPlanListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 获取虚拟主机折扣比率(透传)
func (s *RvhService) GetRvhPrice() (*GetRvhPriceResponse, error) {
	path := "/product/rvh/price"

	var resp GetRvhPriceResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
