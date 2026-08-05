package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// MCSM面板用户
type McsmUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	UserID    int    `json:"user_id,omitempty"`
	PanelUUID string `json:"panel_uuid,omitempty"`
}

// 创建MCSM面板用户
//
// name: 用户名
//
// password: 密码
func (s *RgsService) CreateMcsmUser(name, password string) (*common.BasicOperationResponse, error) {
	path := "/product/rgs/mcsm/panel_user"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, McsmUser{Name: name, Password: password}, &resp)
	return &resp, err
}

// 编辑MCSM面板用户
//
// name: 用户名
//
// password: 密码
func (s *RgsService) EditMcsmUser(name, password string) (*common.BasicOperationResponse, error) {
	path := "/product/rgs/mcsm/panel_user"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, McsmUser{Name: name, Password: password}, &resp)
	return &resp, err
}

// 删除MCSM面板用户
//
// name: 用户名
func (s *RgsService) DeleteMcsmUser(name string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/mcsm/panel_user/%s", name)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}
