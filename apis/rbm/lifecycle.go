package rbm

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// RBM实例启动KVM代理
func (s *RbmService) StartKVMAgent(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/kvm-proxy", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RBM重新启动KVM
func (s *RbmService) RestartKVM(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/kvm-reboot", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RBM实例关机
func (s *RbmService) ShutdownRBM(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/poweroff", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// RBM实例开机
func (s *RbmService) StartRBM(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/poweron", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}
