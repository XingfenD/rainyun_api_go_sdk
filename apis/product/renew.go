package product

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 积分续费请求
type PointRenewalRequest struct {
	DurationDay int    `json:"duration_day"` // 续费天数
	ProductID   int    `json:"product_id"`   // 产品ID
	ProductType string `json:"product_type"` // 产品类型: rcs/rgs/ros/rvh/rcdn
}

// 积分续费
func (s *ProductService) PointRenewal(req *PointRenewalRequest) (*common.BasicOperationResponse, error) {
	path := "/product/point_renew"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
