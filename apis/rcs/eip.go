package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type SetRcsEipDescriptionRequest struct {
	IP          string `json:"ip"`
	Description string `json:"description"`
}

type CreateAndBindIpToRcsRequest struct {
	WithFlags  string `json:"with_flags"`   // IP特征(可选): 应该是高防: us_ddosip -> 美国高防，nb_ddosip -> 宁波高防
	WithIPNum  int    `json:"with_ip_num"`  // IP数量
	WithIPType string `json:"with_ip_type"` // ipv4/ipv6
}

type ChangeRcsIPRequest struct {
	DisableOldIPReason string `json:"disable_old_ip_reason"` // 可选
	IP                 string `json:"ip"`                    // IP地址
	ToIP               string `json:"to_ip"`                 // 可选
}

type DisCardRcsIPRequest struct {
	IP string `json:"ip"`
}

func (s *RcsService) SetRcsEipDescription(id int, ip, desc string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/description", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, SetRcsEipDescriptionRequest{IP: ip, Description: desc}, &resp)
	return &resp, err
}

func (s *RcsService) CreateAndBindElasticIpToRcs(id int, req *CreateAndBindIpToRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

func (s *RcsService) ChangeRcsIP(id int, req *ChangeRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/change", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

func (s *RcsService) DisCardRcsIP(id int, req DisCardRcsIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/eip/discard", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
