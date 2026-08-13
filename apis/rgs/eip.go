package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 创建并绑定弹性IP到游戏云
//
// id: 游戏云 ID
func (s *RgsService) CreateAndBindElasticIpToRgs(id int, req *rcs.CreateAndBindIpToRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/eip", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 更换游戏云IP
//
// id: 游戏云 ID
func (s *RgsService) ChangeRgsIP(id int, req *rcs.ChangeRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/eip/change", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 放弃游戏云IP
//
// id: 游戏云ID
func (s *RgsService) DisCardRgsIP(id int, req rcs.DisCardRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/eip/discard", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
