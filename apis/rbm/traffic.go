package rbm

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 充流量请求
type RBMChargeTrafficRequest struct {
	TrafficInGb int `json:"traffic_in_gb"`
}

// 充流量
func (s *RbmService) ChargeTraffic(id int, req *RBMChargeTrafficRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/traffic/charge", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
