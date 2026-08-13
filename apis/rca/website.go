package rca

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 云应用列出网站(响应结构未公开,透传)
func (s *RcaService) GetRcaWebsiteList() (*GetRcaWebsiteListResponse, error) {
	path := "/product/rca/website/"

	var resp GetRcaWebsiteListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 创建云应用网站
func (s *RcaService) CreateRcaWebsite(req *CreateRcaWebsiteRequest) (*common.BasicOperationResponse, error) {
	path := "/product/rca/website/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 云应用删除网站
func (s *RcaService) DeleteRcaWebsite(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/website/%d/", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

// 获取云应用网站详情(透传)
func (s *RcaService) GetRcaWebsiteDetail(id int) (*GetRcaWebsiteDetailResponse, error) {
	path := fmt.Sprintf("/product/rca/website/%d/", id)

	var resp GetRcaWebsiteDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 云应用网站更新Nginx相关配置
func (s *RcaService) UpdateRcaWebsiteNginx(id int, req *UpdateRcaWebsiteNginxRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/website/%d/config/nginx", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 云应用网站获取重写配置模板(透传)
func (s *RcaService) GetRcaWebsiteRewriteConfig() (*GetRcaWebsiteRewriteConfigResponse, error) {
	path := "/product/rca/website/rewrite_config"

	var resp GetRcaWebsiteRewriteConfigResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
