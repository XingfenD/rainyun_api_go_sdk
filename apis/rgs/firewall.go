package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRgsFirewallRulesRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type GetRgsFirewallRulesResponse struct {
	Code int                     `json:"code"`
	Data rcs.RcsFirewallRuleList `json:"data"`
}

// 获取游戏云防火墙规则列表
//
// id: 游戏云 ID
func (s *RgsService) GetRgsFirewallRules(id int, req *GetRgsFirewallRulesRequest) (*GetRgsFirewallRulesResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/firewall/rule", id)

	var resp GetRgsFirewallRulesResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 创建/设置游戏云防火墙规则
//
// id: 游戏云 ID
func (s *RgsService) SetRgsFirewallRule(id int, req *rcs.SetRcsFirewallRuleRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/firewall/rule", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 删除游戏云防火墙规则
//
// id: 游戏云 ID
func (s *RgsService) DeleteRgsFirewallRule(id, ruleID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/firewall/rule/%d", id, ruleID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

// 移动游戏云防火墙规则优先级
//
// id: 游戏云 ID
func (s *RgsService) MobileRgsFirewallRulePriority(id, ruleID int, req rcs.MobileRcsFirewallRulePriorityRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/firewall/rule/%d/pos", id, ruleID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PUT, path, nil, req, &resp)
	return &resp, err
}
