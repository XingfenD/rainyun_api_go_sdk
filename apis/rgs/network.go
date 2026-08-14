package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 游戏云转成桥接模式
//
// id: 游戏云 ID
func (s *RgsService) RgsToBridge(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/to-bridge", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RgsBridgeSetIntIPRequest 桥接模式下设置内网IP请求
// TODO: 结构未公开,实测后补强类型(字段为推测)
type RgsBridgeSetIntIPRequest struct {
	IP string `json:"ip,omitempty"`
}

// 桥接模式下设置内网
//
// id: 游戏云 ID
func (s *RgsService) RgsBridgeSetIntIP(id int, req *RgsBridgeSetIntIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/bridge_setintip", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RgsVnetRequest 虚拟机内网子网请求
// TODO: 结构未公开,实测后补强类型(字段为推测)
type RgsVnetRequest struct {
	Name    string `json:"name,omitempty"`
	NewName string `json:"new_name,omitempty"`
}

// 创建虚拟机内网子网
//
// id: 游戏云 ID
func (s *RgsService) CreateRgsVnet(id int, req *RgsVnetRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/vnet", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 子网改名
//
// id: 游戏云 ID
func (s *RgsService) RenameRgsVnet(id int, req *RgsVnetRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/vnet", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}
