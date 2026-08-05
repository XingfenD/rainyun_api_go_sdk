package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRcsRenewPriceRequest struct {
	ProductID    int `json:"product_id"`
	Duration     int `json:"duration"`
	WithCouponID int `json:"with_coupon_id"`
}

type GetRcsRenewPriceResponse struct {
	Code int           `json:"code"`
	Data RcsRenewPrice `json:"data"`
}

type RcsRenewPrice struct {
	Detail RcsRenewPriceDetail `json:"detail"`
	Price  float64             `json:"price"`
}

type RcsRenewPriceDetail struct {
	Price        float64          `json:"price"`
	AgentPrice   float64          `json:"agent_price"`
	StockPrice   float64          `json:"stock_price"`
	DefaultPrice int              `json:"default_price"`
	CouponValue  int              `json:"coupon_value"`
	SaleReward   int              `json:"sale_reward"`
	AgentReward  int              `json:"agent_reward"`
	AgentID      int              `json:"agent_id"`
	IgnoreAgent  bool             `json:"ignore_agent"`
	PerScene     RcsRenewPerScene `json:"per_scene"`
}

type RcsRenewPerScene struct {
	Eip      int     `json:"eip"`
	Renew    float64 `json:"renew"`     // 配置价格
	RenewEip float64 `json:"renew_eip"` // IP价格
}

type RenewRcsRequest struct {
	Duration     int `json:"duration"`       // 续费时长(月)
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID
}

type EnableRcsAutoRenewRequest struct {
	AutoRenewOption bool `json:"auto_renew_option"`
}

func (s *RcsService) GetRcsRenewPrice(req *GetRcsRenewPriceRequest) (*GetRcsRenewPriceResponse, error) {
	path := "/product/rcs/price"

	var resp GetRcsRenewPriceResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	querys["scene"] = "renew"
	querys["is_old"] = "true"
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

func (s *RcsService) RenewRcs(id int, req RenewRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/renew", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

func (s *RcsService) EnableRcsAutoRenew(id int, req EnableRcsAutoRenewRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/renew/option", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
