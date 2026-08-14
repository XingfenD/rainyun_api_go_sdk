package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// K8S面板数据库设置请求
type SetK8SPanelDatabaseRequest struct {
	IsEnabled bool   `json:"is_enabled"`
	Version   string `json:"version,omitempty"`
}

// K8S面板SFTP设置请求
type SetK8SPanelSFTPRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// SetK8SPanelStartCommandRequest K8S面板启动命令设置请求
// TODO: 结构未公开,实测后补强类型(字段为推测)
type SetK8SPanelStartCommandRequest struct {
	Command string `json:"command,omitempty"`
}

// K8S面板修改数据库设置
//
// id: 游戏云 ID
func (s *RgsService) SetK8SPanelDatabase(id int, req *SetK8SPanelDatabaseRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/k8s-panel/database", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// 游戏云设置启动命令(仅支持雨云面板)
//
// id: 游戏云 ID
func (s *RgsService) SetK8SPanelStartCommand(id int, req *SetK8SPanelStartCommandRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/k8s-panel/set-start-command", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// K8S面板修改SFTP设置
//
// id: 游戏云 ID
func (s *RgsService) SetK8SPanelSFTP(id int, req *SetK8SPanelSFTPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/k8s-panel/sftp", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}
