package rbm

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// RBM实例更换系统请求
type RBMChangeOSRequest struct {
	OsID int `json:"os_id"`
}

// RBM实例更换系统
func (s *RbmService) ChangeRBMOS(id int, req *RBMChangeOSRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/changeos", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
