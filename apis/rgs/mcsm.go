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

type GetMcsmUserListResponse struct {
	Code int        `json:"code"`
	Data []McsmUser `json:"data"` // TODO: 结构未公开,实测后补强类型(按创建请求推断,账号无用户未能验证字段)
}

// 获取MCSM面板用户列表
func (s *RgsService) GetMcsmUserList() (*GetMcsmUserListResponse, error) {
	path := "/product/rgs/mcsm/panel_user/"

	var resp GetMcsmUserListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 初始化/刷新游戏云sftp功能
//
// id: 游戏云 ID
func (s *RgsService) McsmSftpInit(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/sftp_init", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// MCSM面板开服(启动实例)
//
// id: 游戏云 ID
func (s *RgsService) StartMcsmInstance(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/start", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

type GetMcsmStatusResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型
}

// 获取MCSM实例信息和状态
//
// id: 游戏云 ID
func (s *RgsService) GetMcsmStatus(id int) (*GetMcsmStatusResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/status", id)

	var resp GetMcsmStatusResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
