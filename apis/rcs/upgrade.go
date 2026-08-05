package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type UpgradeRcsRequest struct {
	DestPlan     int `json:"dest_plan"`      // 升级到的套餐ID
	WithCouponID int `json:"with_coupon_id"` // 优惠券ID,默认为0
}

func (s *RcsService) UpgradeRcs(id, plan, coupon int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/upgrade", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, UpgradeRcsRequest{DestPlan: plan, WithCouponID: coupon}, &resp)
	return &resp, err
}
