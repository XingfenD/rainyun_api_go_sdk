package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 添加游戏云NAT端口映射
//
// id: 游戏云 ID
func (s *RgsService) AddRgsNatPortMapping(id int, req *rcs.AddRcsNatPortMappingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/nat", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

type DeleteRgsNatPortMappingRequest struct {
	NatID int `json:"nat_id"`
}

// 删除游戏云NAT端口映射
//
// id: 游戏云 ID
func (s *RgsService) DeleteRgsNatPortMapping(id int, req *DeleteRgsNatPortMappingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/nat", id)

	var resp common.BasicOperationResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_DELETE, path, querys, nil, &resp)
	return &resp, err
}
