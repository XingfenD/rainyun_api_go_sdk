package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type AddRcsNatPortMappingRequest struct {
	PortIn   int    `json:"port_in"`   // >= 1 <= 65535
	PortOut  int    `json:"port_out"`  // >= 10000 <= 60000
	PortType string `json:"port_type"` // tcp/udp/tcp_udp
	Tag      string `json:"tag"`       // 可选
}

type DeleteRcsNatPortMappingRequest struct {
	NatID int `json:"nat_id"`
}

func (s *RcsService) AddRcsNatPortMapping(id int, req *AddRcsNatPortMappingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/nat", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

func (s *RcsService) DeleteRcsNatPortMapping(id int, req *DeleteRcsNatPortMappingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/nat", id)

	var resp common.BasicOperationResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_DELETE, path, querys, nil, &resp)
	return &resp, err
}
