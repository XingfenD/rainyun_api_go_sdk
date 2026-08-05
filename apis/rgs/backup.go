package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 游戏云创建备份
//
// id： 游戏云 ID
//
// label： 备份标签
func (s *RgsService) CreateRgsBackup(id int, label string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, rcs.CreateRcsBackupRequest{Label: label}, &resp)
	return &resp, err
}

// 游戏云删除备份
//
// id: 游戏云 ID, bid: 备份ID
func (s *RgsService) DeleteRgsBackup(id, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/%d", id, bid)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

// 游戏云取消备份
//
// id: RGS ID, bid: 备份ID
func (s *RgsService) CancelRgsBackup(id, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/%d/cancel", id, bid)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 游戏云还原备份
//
// id: 游戏云 ID, bid: 备份ID
func (s *RgsService) RestoreRgsBackup(id, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/%d/restore", id, bid)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 游戏云设置备份选项
//
// id: RGS ID
//
// 没错就是Rcs,这俩是一样的
func (s *RgsService) EnableRgsAutoBackup(id int, req *rcs.RcsSetBackupOptionsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/backup/setting", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}
