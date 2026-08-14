package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 翼龙面板用户
type PteroUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type GetPteroUserListResponse struct {
	Code int         `json:"code"`
	Data []PteroUser `json:"data"` // TODO: 结构未公开,实测后补强类型(按创建请求推断,账号无用户未能验证字段)
}

// 获取翼龙面板用户列表
func (s *RgsService) GetPteroUserList() (*GetPteroUserListResponse, error) {
	path := "/product/rgs/ptero/panel_user/"

	var resp GetPteroUserListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 创建翼龙面板用户
//
// name: 用户名
//
// password: 密码
func (s *RgsService) CreatePteroUser(name, password string) (*common.BasicOperationResponse, error) {
	path := "/product/rgs/ptero/panel_user"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, PteroUser{Name: name, Password: password}, &resp)
	return &resp, err
}

// 编辑翼龙面板用户
//
// name: 用户名
//
// password: 密码
func (s *RgsService) EditPteroUser(name, password string) (*common.BasicOperationResponse, error) {
	path := "/product/rgs/ptero/panel_user"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, PteroUser{Name: name, Password: password}, &resp)
	return &resp, err
}

// 删除翼龙面板用户
//
// name: 用户名
func (s *RgsService) DeletePteroUser(name string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/ptero/panel_user/%s", name)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}
