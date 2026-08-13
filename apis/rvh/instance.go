package rvh

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 获取虚拟主机列表
func (s *RvhService) GetRvhList(req *GetRvhListRequest) (*GetRvhListResponse, error) {
	path := "/product/rvh/"

	var resp GetRvhListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取RVH虚拟主机详情
func (s *RvhService) GetRvhDetail(id int) (*GetRvhDetailResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/", id)

	var resp GetRvhDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 创建虚拟主机
func (s *RvhService) CreateRvh(req *CreateRvhRequest) (*common.BasicOperationResponse, error) {
	path := "/product/rvh/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 释放
func (s *RvhService) FreeRvh(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/free", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RVH重装操作
func (s *RvhService) ReinstallRvh(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/reinstall", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 续费
func (s *RvhService) RenewRvh(id int, req RenewRvhRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/renew/", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 自动续费选项
func (s *RvhService) EnableRvhAutoRenew(id int, req EnableRvhAutoRenewRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/renew/option", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 升级
func (s *RvhService) UpgradeRvh(id int, req UpgradeRvhRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/upgrade/", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RVH设置维护模式
func (s *RvhService) SetRvhMaintenanceMode(id int, req SetRvhMaintenanceModeRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/maintenance-mode", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 设置虚拟主机标签
func (s *RvhService) SetRvhTag(req *SetRvhTagRequest) (*common.BasicOperationResponse, error) {
	path := "/product/rvh/tag"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// RVH EP主机重置密码操作
func (s *RvhService) ResetRvhEpPassword(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/ep/reset-pass", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}
