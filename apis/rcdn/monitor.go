package rcdn

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// rcdn基础用量
func (s *RcdnService) GetRcdnInstanceUsage(id int, req *GetRcdnUsageRequest) (*GetRcdnUsageResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/usage", id)

	var resp GetRcdnUsageResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// rcdn监控
func (s *RcdnService) GetRcdnMonitorData(id int, req *GetRcdnMonitorDataRequest) (*GetRcdnMonitorDataResponse, error) {
	path := fmt.Sprintf("/product/rcdn/%d/monitor", id)

	var resp GetRcdnMonitorDataResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取RCDN套餐列表
func (s *RcdnService) GetRcdnPlanList() (*GetRcdnPlanListResponse, error) {
	path := "/product/rcdn/plans"

	var resp GetRcdnPlanListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 获取RCDN折扣比率(响应结构未公开,透传)
func (s *RcdnService) GetRcdnPrice() (*GetRcdnPriceResponse, error) {
	path := "/product/rcdn/price"

	var resp GetRcdnPriceResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
