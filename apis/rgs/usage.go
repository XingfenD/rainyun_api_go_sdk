package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRgsUsageListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type RgsUsageListData struct {
	TotalRecords int `json:"TotalRecords"`
	Records      any `json:"Records"` // TODO: 结构未公开,实测后补强类型(账号无实例,实测为null)
}

type GetRgsUsageListResponse struct {
	Code int              `json:"code"`
	Data RgsUsageListData `json:"data"`
}

type GetRgsUsageResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型
}

// 获取游戏云使用情况列表
func (s *RgsService) GetRgsUsageList(req *GetRgsUsageListRequest) (*GetRgsUsageListResponse, error) {
	path := "/product/rgs/usage"

	var resp GetRgsUsageListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取游戏云使用情况
//
// id: 游戏云 ID
func (s *RgsService) GetRgsUsage(id int) (*GetRgsUsageResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/usage", id)

	var resp GetRgsUsageResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
