package domain

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 添加域名DNS解析
//
// id: 域名ID
func (s *DomainService) AddDomainDNSRecord(id int, req *AddDomainDNSRecordRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dns", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)

	return &resp, err
}

// 修改域名DNS解析
//
// req.RecordID 指定要修改的记录
func (s *DomainService) UpdateDomainDNSRecord(id int, req *AddDomainDNSRecordRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dns", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)

	return &resp, err
}

// 删除域名DNS解析
func (s *DomainService) DeleteDomainDNSRecord(id, recordID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dns", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, DeleteDomainDNSRecordRequest{RecordID: recordID}, &resp)

	return &resp, err
}

// 添加域名DNSSEC
func (s *DomainService) AddDomainDNSSEC(id int, req *AddDomainDNSSECRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dnssec", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)

	return &resp, err
}

// 删除域名DNSSEC
func (s *DomainService) DeleteDomainDNSSEC(id int, req *DeleteDomainDNSSECRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dnssec/delete", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)

	return &resp, err
}

// 同步域名DNSSEC
func (s *DomainService) SyncDomainDNSSEC(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dnssec/sync", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)

	return &resp, err
}

// 获取域名DNS解析记录列表
//
// domainID: 域名ID
func (s *DomainService) GetDomainDNSRecordList(domainID int) (*GetDomainDNSRecordListResponse, error) {
	path := fmt.Sprintf("/product/domain/%d/dns", domainID)

	var resp GetDomainDNSRecordListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// DNS 记录列表响应
type GetDomainDNSRecordListResponse struct {
	Code int               `json:"code"`
	Data []DomainDNSRecord `json:"data"`
}

// 域名DNS解析记录
type DomainDNSRecord struct {
	ID         int    `json:"id"`
	RecordType string `json:"record_type"`
	Host       string `json:"host"`
	Value      string `json:"value"`
	TTL        int    `json:"ttl"`
}
