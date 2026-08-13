package domain

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 备案域名过白
//
// domain： 域名
//
// region： 区域：cn-sq1/cn-nb1/cn-xy1/cn-cq1
func (s *DomainService) AddDomainToWhiteList(domain, region string) (*common.BasicOperationResponse, error) {
	path := "/product/domain_white_list"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, AddDomainToWhiteListRequest{
		Domain: domain,
		Region: region,
	}, &resp)

	return &resp, err
}

// 获取域名白名单列表
func (s *DomainService) GetDomainWhiteList(req *GetDomainWhiteListRequest) (*GetDomainWhiteListResponse, error) {
	path := "/product/domain/whitelist"

	var resp GetDomainWhiteListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取已验证域名列表
func (s *DomainService) GetVerifiedDomainList(req *GetVerifiedDomainListRequest) (*GetVerifiedDomainListResponse, error) {
	path := "/product/domain/certify"

	var resp GetVerifiedDomainListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 添加域名认证
//
// domain: 域名
func (s *DomainService) AddDomainCertify(domain string) (*GetDomainVerificationInfoResponse, error) {
	path := "/product/domain/certify"

	var resp GetDomainVerificationInfoResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, AddDomainVerificationRequest{Domain: domain}, &resp)

	return &resp, err
}

// 域名认证校验
//
// domain: 域名
func (s *DomainService) VerifyDomainCertify(domain string) (*common.BasicOperationResponse, error) {
	path := "/product/domain/certify/verify"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, AddDomainVerificationRequest{Domain: domain}, &resp)

	return &resp, err
}
