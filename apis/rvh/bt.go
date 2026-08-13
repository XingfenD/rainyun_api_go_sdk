package rvh

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 附加独立IP地址
func (s *RvhService) RvhBtAttachDedip(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/bt/attach-dedip", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RVH宝塔主机修复操作
func (s *RvhService) RvhBtFix(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/bt/fix", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RVH宝塔主机重启操作
func (s *RvhService) RvhBtReboot(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/bt/reboot", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}
