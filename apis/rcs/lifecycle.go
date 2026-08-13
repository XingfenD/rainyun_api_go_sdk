package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type ResetRcsPasswordRequest struct {
	Password string `json:"password"` // 新密码,留空则自动生成
}

type SetRcsTagRequest struct {
	TagName string `json:"tag_name"`
}

func (s *RcsService) FreeRcs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/free", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

func (s *RcsService) RebootRcs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/reboot", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

func (s *RcsService) ResetRcsPassword(id int, newPass string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/reset-password", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, ResetRcsPasswordRequest{Password: newPass}, &resp)
	return &resp, err
}

func (s *RcsService) StartRcs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/start", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

func (s *RcsService) StopRcs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/stop", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

func (s *RcsService) SetRcsTag(id int, tag string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/tag", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, SetRcsTagRequest{TagName: tag}, &resp)
	return &resp, err
}
