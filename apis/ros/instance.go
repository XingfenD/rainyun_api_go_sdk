package ros

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 创建对象存储实例请求
type CreateRosInstanceRequest struct {
	Duration     int `json:"duration"`       // 购买时长（月）
	PlanID       int `json:"plan_id"`        // 套餐ID
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID
}

// 创建对象存储实例响应
type CreateRosInstanceResponse struct {
	Code int         `json:"code"`
	Data RosInstance `json:"data"`
}

// 获取对象存储实例列表请求
type GetRosInstanceListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// 获取对象存储实例列表响应
type GetRosInstanceListResponse struct {
	Code int             `json:"code"`
	Data RosInstanceList `json:"data"`
}

// 对象存储实例列表
type RosInstanceList struct {
	TotalRecords int           `json:"TotalRecords"`
	Records      []RosInstance `json:"Records"`
}

// 获取对象存储实例详情响应
type GetRosInstanceDetailResponse struct {
	Code int               `json:"code"`
	Data RosInstanceDetail `json:"data"`
}

// 对象存储实例详情
type RosInstanceDetail struct {
	Data            RosInstance        `json:"Data"`
	RenewPointPrice RosRenewPointPrice `json:"RenewPointPrice"` // 积分续费
}

// 对象存储积分续费价格
type RosRenewPointPrice struct {
	Num7  int `json:"7"`  // 积分续费七天
	Num31 int `json:"31"` // 积分续费一个月
}

// ROS实例续费请求
type RenewRosInstanceRequest struct {
	Duration     int `json:"duration"`       // 续费时长（月）
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID
}

// ROS实例自动续费选项
type RosInstanceAutoRenewOption struct {
	AutoRenewOption bool `json:"auto_renew_option"`
}

// ROS实例缩放请求
type ScaleRosInstanceRequest struct {
	DestPlan     int `json:"dest_plan"`      // 目标套餐ID
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID
}

// 设置对象存储实例标签请求
type SetRosInstanceTagRequest struct {
	TagName string `json:"tag_name"`
}

// 开关对象存储实例的弹性计费选项请求
type ToggleRosInstanceExtraAccountingRequest struct {
	IsEnable bool `json:"is_enable"`
}

// 创建对象存储实例
//
// req: 创建参数
func (s *RosService) CreateRosInstance(req CreateRosInstanceRequest) (*CreateRosInstanceResponse, error) {
	path := "/product/ros/instance"

	var resp CreateRosInstanceResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 获取对象存储实例详情
//
// instanceID: 实例ID
func (s *RosService) GetRosInstanceDetail(instanceID int) (*GetRosInstanceDetailResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d", instanceID)

	var resp GetRosInstanceDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 获取对象存储实例列表
//
// ⚠️注意这个接口是拿不到AK和SK的，响应里面的AK和SK都是空的
func (s *RosService) GetRosInstanceList(req *GetRosInstanceListRequest) (*GetRosInstanceListResponse, error) {
	path := "/product/ros/instance/"

	var resp GetRosInstanceListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取对象存储实例监控数据
//
// instanceID: 实例ID
//
// req: 开始时间/结束时间（timestamp）
func (s *RosService) GetRosInstanceMonitorData(instanceID int, req *GetRosMonitorDataRequest) (*GetRosMonitorDataResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/monitor", instanceID)

	var resp GetRosMonitorDataResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 对象存储实例续费
//
// instanceID: 实例ID
//
// req: 续费参数
func (s *RosService) RenewRosInstance(instanceID int, req RenewRosInstanceRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/renew", instanceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 对象存储实例自动续费选项
//
// instanceID: 实例ID
//
// isOpen: 是否开启自动续费
func (s *RosService) SetRosInstanceAutoRenewOption(instanceID int, isOpen bool) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/renew/option", instanceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, RosInstanceAutoRenewOption{AutoRenewOption: isOpen}, &resp)
	return &resp, err
}

// ROS实例缩放
//
// instanceID: 实例ID
//
// req: 缩放参数
func (s *RosService) ScaleRosInstance(instanceID int, req ScaleRosInstanceRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/scale", instanceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 设置对象存储实例标签
//
// instanceID: 实例ID
//
// tag: 标签
func (s *RosService) SetRosInstanceTags(instanceID int, tag string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/tag", instanceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, SetRosInstanceTagRequest{TagName: tag}, &resp)
	return &resp, err
}

// 开关对象存储实例的弹性计费选项
//
// instanceID: 实例ID
//
// isEnable: 是否开启
func (s *RosService) ToggleRosInstanceExtraAccounting(instanceID int, isEnable bool) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/toggle-extra-accounting", instanceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, ToggleRosInstanceExtraAccountingRequest{IsEnable: isEnable}, &resp)
	return &resp, err
}

// 对象存储实例重新生成密钥
//
// instanceID: 实例ID
func (s *RosService) ReGenerateRosInstanceKeys(instanceID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/regenerate-keys", instanceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 开关对象存储实例匿名访问
//
// instanceID: 实例ID
func (s *RosService) SetRosInstancePublicAccess(instanceID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/toggle-public-access", instanceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}
