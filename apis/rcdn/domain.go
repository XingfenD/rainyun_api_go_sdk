package rcdn

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 获取加速域名列表
func (s *RcdnService) GetRcdnDomainList(req *GetRcdnDomainListRequest) (*GetRcdnDomainListResponse, error) {
	path := "/product/rcdn/domain"

	var resp GetRcdnDomainListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 创建加速域名
func (s *RcdnService) AddRcdnDomain(req *AddRcdnDomainRequest) (*common.BasicOperationResponse, error) {
	path := "/product/rcdn/domain"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 删除加速域名
func (s *RcdnService) DeleteRcdnDomain(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/domain/%d", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

// 获取加速域名详情
func (s *RcdnService) GetRcdnDomainDetail(id int) (*GetRcdnDomainDetailResponse, error) {
	path := fmt.Sprintf("/product/rcdn/domain/%d", id)

	var resp GetRcdnDomainDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 手动开关防御
func (s *RcdnService) ToggleRcdnDomainWaf(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/domain/%d/toggle_waf", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RCDN缓存清理
func (s *RcdnService) RefreshRcdnCache(instanceID, domainID int, req *RefreshRcdnCacheRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/domain/%d/cache_refresh", instanceID, domainID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// rcdn域名用量
func (s *RcdnService) GetRcdnDomainUsage(id int, req *GetRcdnUsageRequest) (*GetRcdnUsageResponse, error) {
	path := fmt.Sprintf("/product/rcdn/domain/%d/usage", id)

	var resp GetRcdnUsageResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
