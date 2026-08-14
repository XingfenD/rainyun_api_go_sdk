package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 通用模式切换请求
type RgsSwitchModeRequest struct {
	Mode bool `json:"mode"`
}

// 游戏云CPU充电请求
type ChargeRgsCPURequest struct {
	Mode   string `json:"mode"`   // 支付方式(money/point)
	Money  int    `json:"money"`  // 消耗用户余额(支付方式为point时0)
	Points int    `json:"points"` // 消耗用户积分(支付方式为money时0)
}

// 游戏云更换egg(游戏类型)请求
type ChangeRgsEggRequest struct {
	EggTypeID int      `json:"egg_type_id"` // 蛋ID
	SaveDirs  []string `json:"save_dirs"`   // 要保留的目录
}

// 释放游戏云
//
// id: 游戏云ID
func (s *RgsService) FreeRgs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/free", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 游戏云重启操作
//
// id: RGS ID
func (s *RgsService) RebootRgs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/reboot", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 游戏云重置密码
//
// id: 游戏云ID
//
// newPass: 新密码,留空则自动生成
func (s *RgsService) ResetRgsPassword(id int, newPass string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/reset-password", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, rcs.ResetRcsPasswordRequest{Password: newPass}, &resp)
	return &resp, err
}

// 游戏云开机
//
// id: 游戏云ID
func (s *RgsService) StartRgs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/start", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 游戏云关机
//
// id: 游戏云ID
func (s *RgsService) StopRgs(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/stop", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 设置游戏云标签
//
// id: 游戏云ID
//
// tag: 标签
func (s *RgsService) SetRgsTag(id int, tag string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/tag", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, rcs.SetRcsTagRequest{TagName: tag}, &resp)
	return &resp, err
}

// 游戏云限制模式(是否用余额结算)切换
//
// id: 游戏云 ID
//
// useMoney: 是否用余额结算CPU电量
func (s *RgsService) SwitchRgsBalanceMode(id int, useMoney bool) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/cpu-limit-mode", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, RgsSwitchModeRequest{Mode: useMoney}, &resp)
	return &resp, err
}

// 游戏云日付模式开关
//
// id： 游戏云 ID
//
// dailyMode: true: 开启日付模式，false: 关闭日付模式
func (s *RgsService) SwitchRgsDailyMode(id int, dailyMode bool) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/daily-mode", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, RgsSwitchModeRequest{Mode: dailyMode}, &resp)
	return &resp, err
}

// 游戏云CPU充电
//
// id: 游戏云 ID
func (s *RgsService) ChargeRgsCPU(id int, req *ChargeRgsCPURequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/cpu-charge", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 游戏云升级
//
// id: 游戏云ID
//
// plan: 升级到的套餐ID
//
// coupon: 优惠券ID,默认为0
func (s *RgsService) UpgradeRgs(id, plan, coupon int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/upgrade", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, rcs.UpgradeRcsRequest{DestPlan: plan, WithCouponID: coupon}, &resp)
	return &resp, err
}

// 游戏云更换egg(游戏类型)
//
// 此操作需要二步验证
//
// eggTypeID: 游戏类型ID
//
// saveDirs: 要保留的目录
func (s *RgsService) ChangeRgsEgg(id int, eggTypeID int, saveDirs []string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/change-egg", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, ChangeRgsEggRequest{EggTypeID: eggTypeID, SaveDirs: saveDirs}, &resp)
	return &resp, err
}

// SwitchRgsPanelUserRequest 切换面板用户请求
// TODO: 结构未公开,实测后补强类型(字段为推测)
type SwitchRgsPanelUserRequest struct {
	Subtype  string `json:"subtype,omitempty"`  // 面板类型(ptero/mcsm/k8s_panel)
	Name     string `json:"name,omitempty"`     // 用户名
	Password string `json:"password,omitempty"` // 密码
}

// 游戏云切换面板用户
//
// id: 游戏云 ID
func (s *RgsService) SwitchRgsPanelUser(id int, req *SwitchRgsPanelUserRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/switch-user", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
