package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type CreateRcsBackupRequest struct {
	Label string `json:"label"` // 备份名称
}

type RcsSetBackupOptionsRequest struct {
	AutoBackupHour   int `json:"auto_backup_hour"`   // 自动备份时间的小时
	AutoBackupMinute int `json:"auto_backup_minute"` // 自动备份时间的分钟
	KeepLast         int `json:"keep_last"`          // 保留份数(1/3/7)
}

func (s *RcsService) CreateRcsBackup(id int, label string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, CreateRcsBackupRequest{Label: label}, &resp)
	return &resp, err
}

func (s *RcsService) DeleteRcsBackup(id, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d", id, bid)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

func (s *RcsService) CancelRcsBackup(id, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d/cancel", id, bid)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

func (s *RcsService) RestoreRcsBackup(id, bid int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/%d/restore", id, bid)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

func (s *RcsService) EnableRcsAutoBackup(id int, req *RcsSetBackupOptionsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/backup/setting", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}
