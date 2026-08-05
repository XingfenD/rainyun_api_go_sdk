package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type ReinstallRcsRequest struct {
	AppVars []struct {
		AppID int  `json:"app_id"`
		Retry bool `json:"retry"` // 重发之前的任务,此项存在时,无需传入参数
		Vars  any  `json:"vars"`
	} `json:"app_vars"` // 当空数组时,进行单次任务下发
	OsID     int  `json:"os_id"`     // 系统ID
	ResetOsd bool `json:"reset_osd"` // 重置系统盘容量
}

func (s *RcsService) ReinstallRcs(id int, req *ReinstallRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/changeos", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
