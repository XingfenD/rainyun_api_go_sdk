package public

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

// https://api.rainyun.com/#/paths/product~1ros~1plans/get

// 获取对象存储套餐列表
type GetRosPlanListResponse struct {
	Code int           `json:"code"`
	Data []RosPlanItem `json:"data"`
}

type RosPlanItem struct {
	ID                 int     `json:"id"`
	Region             string  `json:"region"`
	Subtype            string  `json:"subtype"`
	PlanName           string  `json:"plan_name"`
	Machine            string  `json:"machine"`
	ChargeType         string  `json:"charge_type"`
	Chinese            string  `json:"chinese"`
	IsFree             bool    `json:"is_free"`
	PointRenewPrice    any     `json:"point_renew_price"` // 积分续费价格
	IsSelling          bool    `json:"is_selling"`
	Price              int     `json:"price"`
	StorageSize        int     `json:"storage_size"`
	Bandwidth          int     `json:"bandwidth"`
	ExtraTransferPrice float64 `json:"extra_transfer_price"`
	ExtraStoragePrice  float64 `json:"extra_storage_price"`
}

func (s *PublicService) GetRosPlanList() (*GetRosPlanListResponse, error) {
	path := "/product/ros/plans"

	var resp GetRosPlanListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)

	return &resp, err
}
