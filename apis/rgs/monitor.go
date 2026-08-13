package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRgsMonitorDataRequest struct {
	StartDate int `json:"start_date"` // 开始时间(时间戳)
	EndDate   int `json:"end_date"`   // 结束时间(时间戳)
}

type GetRgsMonitorDataResponse struct {
	Code int                `json:"code"`
	Data rcs.RcsMonitorData `json:"data"`
}

// 获取游戏云监控数据
//
// id: 游戏云ID
//
// startDate: 开始时间
//
// endDate: 结束时间
func (s *RgsService) GetRgsMonitorData(id int, req *GetRgsMonitorDataRequest) (*GetRgsMonitorDataResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/monitor", id)

	var resp GetRgsMonitorDataResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
