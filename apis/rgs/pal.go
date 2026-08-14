package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetPalConfigResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型(帕鲁Palworld面板配置)
}

// 获取帕鲁面板配置
//
// id: 游戏云 ID
func (s *RgsService) GetPalConfig(id int) (*GetPalConfigResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/pal/config", id)

	var resp GetPalConfigResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 设置帕鲁面板配置
//
// id: 游戏云 ID
//
// config: 配置项(字段未公开) TODO: 结构未公开,实测后补强类型
func (s *RgsService) SetPalConfig(id int, config map[string]any) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/pal/config", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, config, &resp)
	return &resp, err
}

// 初始化帕鲁配置
//
// id: 游戏云 ID
func (s *RgsService) InitPal(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/pal/init", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

type GetPalLangResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // TODO: 结构未公开,实测后补强类型
}

// 获取帕鲁配置中文说明
//
// id: 游戏云 ID
func (s *RgsService) GetPalLang(id int) (*GetPalLangResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/pal/lang", id)

	var resp GetPalLangResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 帕鲁RCON命令
//
// id: 游戏云 ID
//
// command: RCON命令 TODO: 结构未公开,实测后补强类型(请求体字段为推测)
func (s *RgsService) PalRcon(id int, command string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/pal/rcon", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, map[string]string{"command": command}, &resp)
	return &resp, err
}

// 重启帕鲁
//
// id: 游戏云 ID
func (s *RgsService) RestartPal(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/pal/restart", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 关闭帕鲁
//
// id: 游戏云 ID
func (s *RgsService) StopPal(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/mcsm/pal/stop", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}
