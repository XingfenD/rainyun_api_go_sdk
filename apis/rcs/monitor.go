package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRcsMonitorDataRequest struct {
	StartDate int `json:"start_date"` // 开始时间(时间戳)
	EndDate   int `json:"end_date"`   // 结束时间(时间戳)
}

type GetRcsMonitorDataResponse struct {
	Code int            `json:"code"`
	Data RcsMonitorData `json:"data"`
}

// RcsMonitorData 监控数据, Columns 与 Values 一一对应
type RcsMonitorData struct {
	Columns []string    `json:"Columns"`
	Values  [][]float64 `json:"Values"`
}

func (s *RcsService) GetRcsMonitorData(id int, req *GetRcsMonitorDataRequest) (*GetRcsMonitorDataResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/monitor", id)

	var resp GetRcsMonitorDataResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
