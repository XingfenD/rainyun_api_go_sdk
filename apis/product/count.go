package product

import (
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 产品数量
type ProductCount struct {
	TotalCount int `json:"TotalCount"`
}

// RbmProductCount 域名转售(RBM)产品数量
type RbmProductCount struct {
	TotalCount    int `json:"TotalCount"`
	AboutToExpire any `json:"AboutToExpire"`
}

// 产品数量列表
type ProductCountData struct {
	Ros    ProductCount    `json:"ros"`
	Rvh    ProductCount    `json:"rvh"`
	Rcs    ProductCount    `json:"rcs"`
	Rgpu   ProductCount    `json:"rgpu"`
	Rgs    ProductCount    `json:"rgs"`
	Ssl    ProductCount    `json:"ssl"`
	Rbm    RbmProductCount `json:"rbm"`
	Domain ProductCount    `json:"domain"`
	Rcdn   ProductCount    `json:"rcdn"`
	Rca    ProductCount    `json:"rca"`
}

type GetProductCountListResponse struct {
	Code int              `json:"code"`
	Data ProductCountData `json:"data"`
}

// 产品ID列表
type ProductIDListData struct {
	Rca    []int `json:"rca"`
	Rcs    []int `json:"rcs"`
	Ros    []int `json:"ros"`
	Rvh    []int `json:"rvh"`
	Rgpu   []int `json:"rgpu"`
	Rgs    []int `json:"rgs"`
	Ssl    []int `json:"ssl"`
	Rbm    []int `json:"rbm"`
	Domain []int `json:"domain"`
	Rcdn   []int `json:"rcdn"`
}

type GetProductIDListResponse struct {
	Code int               `json:"code"`
	Data ProductIDListData `json:"data"`
}

// 获取各产品数量
func (s *ProductService) GetProductCountList() (*GetProductCountListResponse, error) {
	path := "/product/"

	var resp GetProductCountListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 获取产品ID列表
func (s *ProductService) GetProductIDList() (*GetProductIDListResponse, error) {
	path := "/product/id_list"

	var resp GetProductIDListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
