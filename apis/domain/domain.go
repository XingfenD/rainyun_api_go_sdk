package domain

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 关闭域名锁定
func (s *DomainService) UnlockDomain(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/lock/disable", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PUT, path, nil, nil, &resp)

	return &resp, err
}

// 开启域名锁定
func (s *DomainService) LockDomain(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/lock/enable", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PUT, path, nil, nil, &resp)

	return &resp, err
}

// 修改域名NS服务器
func (s *DomainService) UpdateDomainNS(id int, nss []string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/nameservers", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PUT, path, nil, UpdateDomainNSRequest{NameServers: nss}, &resp)

	return &resp, err
}

// 重置域名NS服务器
func (s *DomainService) ResetDomainNS(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/nameservers/reset", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)

	return &resp, err
}

// 续费域名
func (s *DomainService) RenewDomain(id, years int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/renew", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, RenewDomainRequest{Years: years}, &resp)

	return &resp, err
}

// 域名过户
func (s *DomainService) TransferDomain(id int, req *DomainTransferRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/transfer", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)

	return &resp, err
}

// 更新域名管理密码
func (s *DomainService) UpdateDomainPassword(id int, password string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/password", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, UpdateDomainPasswordRequest{Password: password}, &resp)

	return &resp, err
}

// 获取域名whois信息
func (s *DomainService) GetDomainWhoisInfo(req *GetDomainWhoisInfoRequest) (*GetDomainWhoisInfoResponse, error) {
	path := "/product/domain/whois"

	var resp GetDomainWhoisInfoResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 检查域名能否注册
func (s *DomainService) CheckDomainAvailability(req *CheckDomainAvailableRequest) (*CheckDomainAvailableResponse, error) {
	path := "/product/domain/check"

	var resp CheckDomainAvailableResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
