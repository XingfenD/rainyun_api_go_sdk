package rbm

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 创建并绑定弹性IP到RBM请求
type RBMAssociateEIPRequest struct {
	WithFlags  string `json:"with_flags"`   // IP特征(可选)
	WithIPNum  int    `json:"with_ip_num"`  // 数量
	WithIPType string `json:"with_ip_type"` // IPv4/IPv6
}

// 更换IP请求
type RBMChangeIPRequest struct {
	DisableOldIPReason string `json:"disable_old_ip_reason"` // 禁用旧IP的原因(可选)
	IP                 string `json:"ip"`                    // 旧IP地址
	ToIP               string `json:"to_ip"`                 // 不明，可选
}

// 设置IP描述请求
type RBMSetIPDescriptionRequest struct {
	Description string `json:"description"`
	IP          string `json:"ip"`
}

// 放弃IP请求
type RBMReleaseIPRequest struct {
	IP string `json:"ip"`
}

// 创建并绑定弹性IP到RBM
func (s *RbmService) AssociateEIP(id int, req *RBMAssociateEIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/eip", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 设置IP描述
func (s *RbmService) SetIPDescription(id int, req *RBMSetIPDescriptionRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/eip/description", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 放弃IP
func (s *RbmService) ReleaseIP(id int, req *RBMReleaseIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/eip/discard", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
