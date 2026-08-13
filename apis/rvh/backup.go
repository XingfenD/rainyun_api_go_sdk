package rvh

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// RVH创建备份
func (s *RvhService) CreateRvhBackup(id int, req CreateRvhBackupRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/backup/", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RVH删除备份
func (s *RvhService) DeleteRvhBackup(id, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/backup/%d/", id, bid)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

// RVH还原备份
func (s *RvhService) RestoreRvhBackup(id, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/backup/%d/restore", id, bid)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RVH设置备份选项
func (s *RvhService) SetRvhBackupSetting(id int, req SetRvhBackupSettingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/backup/setting", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}
