package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type ChargeRcsTraficRequest struct {
	TrafficInGb int `json:"traffic_in_gb"` // 充多少G
}

type LimitRcsTrafficRequest struct {
	DayTrafficInGb int `json:"day_traffic_in_gb"` // 日流量阈值(G)
	TrafficLimit   int `json:"traffic_limit"`     // 限制带宽(M)
}

func (s *RcsService) ChargeRcsTrafic(id, count int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/traffic/charge", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, ChargeRcsTraficRequest{TrafficInGb: count}, &resp)
	return &resp, err
}

func (s *RcsService) LimitRcsTrafic(id, threshold, limit int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/traffic/limit", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, LimitRcsTrafficRequest{DayTrafficInGb: threshold, TrafficLimit: limit}, &resp)
	return &resp, err
}
