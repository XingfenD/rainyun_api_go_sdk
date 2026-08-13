package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRgsRenewPriceRequest struct {
	ProductID    int `json:"product_id"`
	Duration     int `json:"duration"`
	WithCouponID int `json:"with_coupon_id"`
}

type GetRgsRenewPriceResponse struct {
	Code int               `json:"code"`
	Data rcs.RcsRenewPrice `json:"data"`
}

type GetRgsUpgradePriceRequest struct {
	ProductID    int       `json:"product_id"`
	Duration     int       `json:"duration"`
	WithCouponID int       `json:"with_coupon_id"`
	Config       RgsConfig `json:"config"`
}

// 游戏云升级价格
type GetRgsUpgradePriceResponse struct {
	Code int                 `json:"code"`
	Data RgsUpgradePriceData `json:"data"`
}

type RgsUpgradePriceData struct {
	Detail RgsUpgradePriceDetail `json:"detail"`
	Price  int                   `json:"price"`
}

type RgsUpgradePriceDetail struct {
	Price        int                `json:"price"`
	AgentPrice   int                `json:"agent_price"`
	StockPrice   int                `json:"stock_price"`
	DefaultPrice int                `json:"default_price"`
	CouponValue  int                `json:"coupon_value"`
	SaleReward   int                `json:"sale_reward"`
	AgentReward  int                `json:"agent_reward"`
	AgentID      int                `json:"agent_id"`
	IgnoreAgent  bool               `json:"ignore_agent"`
	PerScene     RgsUpgradePerScene `json:"per_scene"`
}

type RgsUpgradePerScene struct {
	Upgrade int `json:"upgrade"`
}

// 获取游戏云续费价格
//
// id: RGS ID, duration: 续费时长(月), coupon: 优惠券ID
func (s *RgsService) GetRgsRenewPrice(req *GetRgsRenewPriceRequest) (*GetRgsRenewPriceResponse, error) {
	path := "/product/rgs/price"

	var resp GetRgsRenewPriceResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	querys["scene"] = "renew"
	querys["is_old"] = "true"
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 游戏云续费
//
// id: 游戏云ID
func (s *RgsService) RenewRgs(id int, req rcs.RenewRcsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/renew", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 游戏云自动续费选项
//
// id: 游戏云ID
func (s *RgsService) EnableRgsAutoRenew(id int, req rcs.EnableRcsAutoRenewRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/renew/option", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 获取游戏云升级价格
//
// id: RGS ID, duration: 续费时长(月), coupon: 优惠券ID
func (s *RgsService) GetRgsUpgradePrice(req *GetRgsUpgradePriceRequest) (*GetRgsUpgradePriceResponse, error) {
	path := "/product/rgs/price"

	var resp GetRgsUpgradePriceResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	querys["scene"] = "upgrade"
	querys["is_old"] = "true"
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
