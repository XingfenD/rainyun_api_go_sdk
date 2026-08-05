package user

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetUserRewardProductsResponse struct {
	Code int                    `json:"code"`
	Data UserRewardProductsData `json:"data"`
}

// UserRewardProductsData 用户可用积分兑换的产品
type UserRewardProductsData struct {
	Rcs []any `json:"rcs"`
	Rvh []any `json:"rvh"`
	Rgs []any `json:"rgs"`
	Ros []any `json:"ros"`
	Rbm []any `json:"rbm"`
}

// 获取可兑换积分产品列表.
func (s *UserService) GetUserRewardProducts() (*GetUserRewardProductsResponse, error) {
	path := "/user/reward/products"

	var resp GetUserRewardProductsResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)

	return &resp, err
}

// 发布优惠券给下级用户
type PublishCouponsToLowerLevelUsersRequest struct {
	BaseLimit      int    `json:"base_limit"`      // 满减条件(满多少才能用)
	Color          string `json:"color"`           // 颜色: waring: 黄，danger: 红，success: 绿...
	Count          int    `json:"count"`           // 发放数量
	ExpDate        int    `json:"exp_date"`        // 过期时间(timestamp)
	FriendlyName   string `json:"friendly_name"`   // 优惠券标题
	Type           string `json:"type"`            // 类型: discount:折扣, normal: 直减
	UID            int    `json:"uid"`             // 要发放到用户ID(如:114514),为空时则返回兑换码
	UsableDuration string `json:"usable_duration"` // unknown
	UsableProduct  string `json:"usable_product"`  // 可用产品(默认全部),","分隔: renew,create,upgrade
	UsableScenes   string `json:"usable_scenes"`   // 适用操作(默认全部),","分隔: rvh,rcs,rgs,ros,rbm
	Value          int    `json:"value"`           // 直减(元)/折扣(折), 折扣时:1~9:一~九折；11~99:一一~九九折
}

// 发布优惠券给下级用户
func (s *UserService) PublishCouponsToLowerLevelUsers(req *PublishCouponsToLowerLevelUsersRequest) (*common.BasicOperationResponse, error) {
	path := "/user/vip/coupon"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)

	return &resp, err
}

// 发送优惠券到积分商城
type PostCouponsToPointsMallRequest struct {
	AvailableDays  int    `json:"available_days"`  // 可用天数
	BaseLimit      int    `json:"base_limit"`      // 满减条件(满多少才能用)
	BuyLimit       int    `json:"buy_limit"`       // 领取次数限制
	Color          string `json:"color"`           // 颜色: waring: 黄, danger: 红, success: 绿, info: 蓝
	Count          int    `json:"count"`           // 发放数量
	EndDate        int    `json:"end_date"`        // 截止日期
	FirstSend      bool   `json:"first_send"`      // 绑定微信后立即自动领取(设置后不可手动领取并且不会显示)
	FriendlyName   string `json:"friendly_name"`   // 优惠券名称
	Name           string `json:"name"`            // 标识名称
	Order          int    `json:"order"`           // 排序,越大越靠前
	Points         int    `json:"points"`          // 领取积分
	Type           string `json:"type"`            // 类型: discount:折扣, normal: 直减
	UsableDuration string `json:"usable_duration"` // unknown
	UsableProduct  string `json:"usable_product"`  // 可用产品(默认全部),","分隔: renew,create,upgrade
	UsableScenes   string `json:"usable_scenes"`   // 适用操作(默认全部),","分隔: rvh,rcs,rgs,ros,rbm
	Value          int    `json:"value"`           // 直减(元)/折扣(折), 折扣时:1~9:一~九折；11~99:一一~九九折
}

// 发布优惠券到积分商城(供下级领取)
func (s *UserService) PostCouponsToPointsMall(req *PostCouponsToPointsMallRequest) (*common.BasicOperationResponse, error) {
	path := "/user/vip/coupon"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)

	return &resp, err
}

type GetPointsMallItemsResponse struct {
	Code int              `json:"code"`
	Data []PointsMallItem `json:"data"`
}

// PointsMallItem 积分商城物品
type PointsMallItem struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	Points         int                `json:"points"`
	Type           string             `json:"type"`
	AvailableStock int                `json:"available_stock"`
	FriendlyName   string             `json:"friendly_name"`
	ItemData       PointsMallItemData `json:"item_data"`
	BuyLimit       int                `json:"buy_limit"`
	SenderID       int                `json:"sender_id"`
	FirstSend      bool               `json:"first_send"`
	ByInvite       bool               `json:"by_invite"`
	Color          string             `json:"color"`
	Order          int                `json:"order"`
	PublicTime     int                `json:"public_time"`
	EndDate        int                `json:"end_date"`
	AutoRefresh    int                `json:"auto_refresh"`
	RefreshLimit   int                `json:"refresh_limit"`
	MoneyRequired  int                `json:"money_required"`
}

// PointsMallItemData 积分商城物品数据
type PointsMallItemData struct {
	Color           string                  `json:"color,omitempty"`
	FriendlyName    string                  `json:"friendly_name,omitempty"`
	UsableScenes    string                  `json:"usable_scenes,omitempty"`
	AvailableDays   int                     `json:"available_days,omitempty"`
	UsableProduct   string                  `json:"usable_product,omitempty"`
	BaseLimit       int                     `json:"base_limit,omitempty"`
	PublicPoint     int                     `json:"public_point,omitempty"`
	Type            string                  `json:"type,omitempty"`
	UsableDuration  string                  `json:"usable_duration,omitempty"`
	UsablePlanID    int                     `json:"usable_plan_id,omitempty"`
	Value           any                     `json:"value,omitempty"`
	ProductType     string                  `json:"product_type,omitempty"`
	ProductSubtype  string                  `json:"product_subtype,omitempty"`
	DurationSeconds int                     `json:"duration_seconds,omitempty"`
	ProductConfig   PointsMallProductConfig `json:"product_config,omitempty"`
	Desc            string                  `json:"desc,omitempty"`
	ImgURL          string                  `json:"img_url,omitempty"`
	DescURL         string                  `json:"desc_url,omitempty"`
}

// PointsMallProductConfig 积分商城物品的产品配置
type PointsMallProductConfig struct {
	OsID         int                  `json:"os_id"`
	PlanID       int                  `json:"plan_id"`
	Subtype      string               `json:"subtype"`
	Duration     int                  `json:"duration"`
	PayMode      string               `json:"pay_mode"`
	PanelUser    string               `json:"panel_user"`
	EggTypeID    int                  `json:"egg_type_id"`
	WithCoupon   int                  `json:"with_coupon"`
	CPULimitMode bool                 `json:"cpu_limit_mode"`
	Config       PointsMallItemConfig `json:"config"`
}

// PointsMallItemConfig 积分商城物品配置
type PointsMallItemConfig struct {
	CPU        int `json:"cpu"`
	Backup     int `json:"backup"`
	Memory     int `json:"memory"`
	NetIn      int `json:"net_in"`
	NetOut     int `json:"net_out"`
	Database   int `json:"database"`
	BaseDisk   int `json:"base_disk"`
	DataDisk   int `json:"data_disk"`
	Allocation int `json:"allocation"`
}

// 获取积分商城商品列表
func (s *UserService) GetPointsMallItems() (*GetPointsMallItemsResponse, error) {
	path := "/user/reward/items"

	var resp GetPointsMallItemsResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)

	return &resp, err
}

// 兑换积分物品
type RedeemPointsForItemRequest struct {
	ItemID int `json:"item_id"`
}

// 兑换积分物品
func (s *UserService) RedeemPointsForItem(id int) (*common.BasicOperationResponse, error) {
	path := "/user/reward/items"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, RedeemPointsForItemRequest{ItemID: id}, &resp)

	return &resp, err
}
