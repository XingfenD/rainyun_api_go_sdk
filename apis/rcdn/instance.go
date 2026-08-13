package rcdn

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 获取RCDN实例列表
func (s *RcdnService) GetRcdnInstanceList(req *GetRcdnInstanceListRequest) (*GetRcdnInstanceListResponse, error) {
	path := "/product/rcdn/instance"

	var resp GetRcdnInstanceListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取RCDN实例详情
func (s *RcdnService) GetRcdnInstanceDetail(id int) (*GetRcdnInstanceDetailResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d", id)

	var resp GetRcdnInstanceDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 创建RCDN实例
func (s *RcdnService) CreateRcdnInstance(req *CreateRcdnInstanceRequest) (*common.BasicOperationResponse, error) {
	path := "/product/rcdn/instance"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RCDN实例续费
func (s *RcdnService) RenewRcdnInstance(id int, req RenewRcdnInstanceRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/renew", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RCDN实例自动续费选项
func (s *RcdnService) EnableRcdnInstanceAutoRenew(id int, req EnableRcdnInstanceAutoRenewRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/renew/option", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RCDN实例缩放
func (s *RcdnService) ScaleRcdnInstance(id int, req ScaleRcdnInstanceRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/scale", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RCDN实例设置
func (s *RcdnService) SetRcdnInstanceSetting(id int, req *SetRcdnInstanceSettingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/setting", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RCDN SSL绑定域名
func (s *RcdnService) BindRcdnSSLDomains(id int, req *BindRcdnSSLDomainsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/ssl_bind", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 设置RCDN实例标签
func (s *RcdnService) SetRcdnInstanceTag(id int, req SetRcdnInstanceTagRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/tag", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// 开关RCDN实例的弹性计费选项
func (s *RcdnService) ToggleRcdnInstanceExtraAccounting(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcdn/instance/%d/toggle-extra-accounting", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}
