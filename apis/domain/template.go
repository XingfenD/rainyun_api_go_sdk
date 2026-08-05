package domain

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 查询域名模板列表
func (s *DomainService) GetDomainTemplateList(req *GetDomainTemplateListRequest) (*GetDomainTemplateListResponse, error) {
	path := "/product/domain/template/"

	var resp GetDomainTemplateListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 删除域名信息模板
func (s *DomainService) DeleteDomainTemplate(sysID string) (*common.BasicOperationResponse, error) {
	path := "/product/domain/template/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, DeleteDomainTemplateRequest{SysID: sysID}, &resp)

	return &resp, err
}

// 编辑域名模板
func (s *DomainService) EditDomainTemplate(req *EditDomainTemplateRequest) (*common.BasicOperationResponse, error) {
	path := "/product/domain/template"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PUT, path, nil, req, &resp)

	return &resp, err
}

// 获取域名模板详情
func (s *DomainService) GetDomainTemplateDetail(req *GetDomainTemplateDetailRequest) (*GetDomainTemplateDetailResponse, error) {
	path := "/product/domain/template/detail/"

	var resp GetDomainTemplateDetailResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
