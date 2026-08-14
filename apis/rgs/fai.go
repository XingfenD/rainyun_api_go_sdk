package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// SendRgsFaiTaskRequest 发布快速app安装任务请求
// TODO: 结构未公开,实测后补强类型(按changeos的app_vars推断)
type SendRgsFaiTaskRequest struct {
	AppVars []RgsAppVar `json:"app_vars"`
}

// 发布快速app安装任务
//
// id: 游戏云 ID
func (s *RgsService) SendRgsFaiTask(id int, req *SendRgsFaiTaskRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/fai-send", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
