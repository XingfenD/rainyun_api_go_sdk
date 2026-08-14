package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// RgsMpDstInfo 游戏云MP目标集群信息
type RgsMpDstInfo struct {
	ClusterName string `json:"cluster_name"`
	ClusterPass string `json:"cluster_pass,omitempty"`
	Token       string `json:"token,omitempty"`
}

type ListRgsMpRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type ListRgsMpResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型(线上接口当前返回500)
}

type CreateRgsMpRequest struct {
	Duration     int           `json:"duration"`                 // 时长(月) 1/3/6/12
	EggTypeID    int           `json:"egg_type_id"`              // 游戏类型
	PlanID       int           `json:"plan_id,omitempty"`        // 套餐ID
	WithCouponID int           `json:"with_coupon_id,omitempty"` // 优惠券ID
	DstInfo      *RgsMpDstInfo `json:"dst_info,omitempty"`       // 目标集群信息
}

type CreateRgsMpResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型(付费创建,不做线上探测)
}

type RenewRgsMpRequest struct {
	Duration     int `json:"duration"`                 // 续费时长(月) 1/3/6/12
	WithCouponID int `json:"with_coupon_id,omitempty"` // 优惠券ID
}

type RenewRgsMpResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型(付费续费,不做线上探测)
}

// 获取游戏云MP列表
func (s *RgsService) ListRgsMp(req *ListRgsMpRequest) (*ListRgsMpResponse, error) {
	path := "/product/rgs-mp/"

	var resp ListRgsMpResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 创建游戏云MP
func (s *RgsService) CreateRgsMp(req *CreateRgsMpRequest) (*CreateRgsMpResponse, error) {
	path := "/product/rgs-mp/"

	var resp CreateRgsMpResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 续费游戏云MP
//
// id: 游戏云MP ID
func (s *RgsService) RenewRgsMp(id int, req *RenewRgsMpRequest) (*RenewRgsMpResponse, error) {
	path := fmt.Sprintf("/product/rgs-mp/%d/renew/", id)

	var resp RenewRgsMpResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
