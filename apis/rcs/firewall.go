package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRcsFirewallRulesRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type GetRcsFirewallRulesResponse struct {
	Code int                 `json:"code"`
	Data RcsFirewallRuleList `json:"data"`
}

type RcsFirewallRuleList struct {
	TotalRecords int               `json:"TotalRecords"`
	Records      []RcsFirewallRule `json:"Records"`
}

type RcsFirewallRule struct {
	ID            int    `json:"ID"`
	VID           int    `json:"v_id"`
	IsEnable      bool   `json:"is_enable"`
	Pos           int    `json:"pos"`
	SourceAddress string `json:"source_address"`
	DestPort      string `json:"dest_port"`
	SourcePort    string `json:"source_port"`
	Protocol      string `json:"protocol"`
	Action        string `json:"action"`
	Description   string `json:"description"`
}

type SetRcsFirewallRuleRequest struct {
	Action        string `json:"action"`         // 动作，accept/drop，接受或者丢弃
	Description   string `json:"description"`    // 备注(可选)
	DestPort      string `json:"dest_port"`      // 代表本机的目的端口，可以用-来链接，空白代表所有端口(可选)
	ID            int    `json:"id"`             // 规则ID(可选)
	IsEnable      bool   `json:"is_enable"`      // 是否启用该规则(可选)
	Protocol      string `json:"protocol"`       // 协议，udp/tcp/icmp，空白代表所有(可选)
	SourceAddress string `json:"source_address"` // 代表来源的地址，可以用-链接范围，或者用逗号来分割多个地址，可以使用网络，CIDR格式，空则代表所有地址(可选)
	SourcePort    string `json:"source_port"`    // 一般不填(防反射)(可选)
}

type MobileRcsFirewallRulePriorityRequest struct {
	NewPos int `json:"newPos"`
}

func (s *RcsService) GetRcsFirewallRules(id int, req *GetRcsFirewallRulesRequest) (*GetRcsFirewallRulesResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule", id)

	var resp GetRcsFirewallRulesResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

func (s *RcsService) SetRcsFirewallRule(id int, req *SetRcsFirewallRuleRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

func (s *RcsService) DeleteRcsFirewallRule(id, ruleID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule/%d", id, ruleID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

func (s *RcsService) MobileRcsFirewallRulePriority(id, ruleID int, req MobileRcsFirewallRulePriorityRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/firewall/rule/%d/pos", id, ruleID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PUT, path, nil, req, &resp)
	return &resp, err
}
