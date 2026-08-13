package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type RcsManagesElasticCloudDisksRequest struct {
	Actions []struct {
		Type   string `json:"type"`   // 操作类型: expand: 扩容, create: 创建
		Action any    `json:"action"` // 操作参数, RcsManagesElasticCloudDisksExpand或RcsManagesElasticCloudDisksCreate
	} `json:"actions"`
}

type RcsManagesElasticCloudDisksExpand struct {
	EdiskID  int  `json:"edisk_id"`   // 弹性云盘ID
	SizeInGb int  `json:"size_in_gb"` // 操作容量
	Backup   bool `json:"backup"`     // 支持备份
}

type RcsManagesElasticCloudDisksCreate struct {
	SizeInGb int    `json:"size_in_gb"` // 操作容量
	DiskType string `json:"disk_type"`  // 磁盘类型(ssd/hdd)
	Backup   bool   `json:"backup"`     // 支持备份
	Tag      string `json:"tag"`        // 标签
}

func (s *RcsService) RcsManagesElasticCloudDisks(id int, req *RcsManagesElasticCloudDisksRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/edisk", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
