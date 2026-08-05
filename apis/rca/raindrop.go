package rca

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 云应用雨点余额使用情况
type RcaRaindropUsageData struct {
	ExpectedRemainDays     int     `json:"expected_remain_days"`      // 预计剩余天数
	LastMonthUsage         float64 `json:"last_month_usage"`          // 上月使用量
	ExpectedNextMonthUsage float64 `json:"expected_next_month_usage"` // 预计下月使用量
	FreeTrialRemainDays    int     `json:"free_trial_remain_days"`    // 剩余免费试用天数
	IsBeforeFirstPayment   bool    `json:"is_before_first_payment"`
}

type GetRcaRaindropUsageResponse struct {
	Code int                  `json:"code"`
	Data RcaRaindropUsageData `json:"data"`
}

// 雨点套餐
type RaindropPlan struct {
	ID        int    `json:"id"`
	Amount    int    `json:"amount"`
	Price     int    `json:"price"`
	IsSelling bool   `json:"is_selling"`
	PlanName  string `json:"plan_name"`
	Chinese   string `json:"chinese"`
}

type GetRcaRaindropPlansListResponse struct {
	Code int            `json:"code"`
	Data []RaindropPlan `json:"data"`
}

// 雨点消费明细
type RaindropConsumeData struct {
	BasicPrice   float64 `json:"basic_price"`
	TrafficBytes int     `json:"traffic_bytes"`
	TrafficPrice float64 `json:"traffic_price"`
}

// 雨点消费记录
type RaindropConsumeRecord struct {
	ID        int                 `json:"id"`
	UID       int                 `json:"uid"`
	Time      int                 `json:"time"`
	Type      string              `json:"type"`
	ProductID int                 `json:"product_id"`
	Amount    float64             `json:"amount"`
	Data      RaindropConsumeData `json:"data"`
}

// 雨点消费历史
type RaindropConsumeLogData struct {
	TotalRecords int                     `json:"TotalRecords"`
	Records      []RaindropConsumeRecord `json:"Records"`
}

type GetRaindropConsumeLogResponse struct {
	Code int                    `json:"code"`
	Data RaindropConsumeLogData `json:"data"`
}

// 云应用购买雨点请求
type BuyRaindropRequest struct {
	PlanID       int `json:"plan_id"`
	WithCouponID int `json:"with_coupon_id"`
}

// 雨点余额
type GetRcaRaindropBalanceResponse struct {
	Code int     `json:"code"`
	Data float64 `json:"data"`
}

// 云应用获取雨点余额使用情况
func (s *RcaService) GetRcaRaindropUsage() (*GetRcaRaindropUsageResponse, error) {
	path := "/product/rca/raindrop/usage"

	var resp GetRcaRaindropUsageResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 云应用获取雨点套餐列表
func (s *RcaService) GetRcaRaindropPlansList() (*GetRcaRaindropPlansListResponse, error) {
	path := "/product/rca/raindrop/plans"

	var resp GetRcaRaindropPlansListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 云应用获取雨点消费历史
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (s *RcaService) GetRaindropConsumeLog(options string) (*GetRaindropConsumeLogResponse, error) {
	path := "/product/rca/raindrop/consume_log"
	querys := map[string]string{"options": options}

	var resp GetRaindropConsumeLogResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 云应用购买雨点
//
// planID: 雨点套餐ID
//
// couponID: 优惠券ID
func (s *RcaService) BuyRaindrop(planID int, couponID int) (*common.BasicOperationResponse, error) {
	path := "/product/rca/raindrop"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, BuyRaindropRequest{PlanID: planID, WithCouponID: couponID}, &resp)
	return &resp, err
}

// 云应用获取雨点余额
func (s *RcaService) GetRcaRaindropBalance() (*GetRcaRaindropBalanceResponse, error) {
	path := "/product/rca/raindrop"

	var resp GetRcaRaindropBalanceResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
