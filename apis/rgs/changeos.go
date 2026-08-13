package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 游戏云重装系统
//
// id: RGS ID
func (s *RgsService) ReinstallRgs(id int, req *rcs.ReinstallRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/changeos", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
