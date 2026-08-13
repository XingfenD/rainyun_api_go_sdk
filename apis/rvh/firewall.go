package rvh

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// RVH防火墙设置选项
func (s *RvhService) SetRvhFirewallOption(id int, req SetRvhFirewallOptionRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/firewall/option", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// RVH防火墙设置规则
func (s *RvhService) SetRvhFirewallRule(id int, req SetRvhFirewallRuleRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rvh/%d/firewall/rule", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
