package rvh

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// RVH域名绑定
func (s *RvhService) BindRvhDomain(id int, req BindRvhDomainRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/domain/", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RVH域名解绑
// ponytail: 旧 spec 的解绑查询参数写作 "id"(应为笔误),解绑语义必传域名,故用 query "domain"
func (s *RvhService) UnbindRvhDomain(id int, domain string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/domain/", id)

	var resp common.BasicOperationResponse
	querys := map[string]string{"domain": domain}
	err := s.client.Do(constant.HTTPMethod_DELETE, path, querys, nil, &resp)
	return &resp, err
}
