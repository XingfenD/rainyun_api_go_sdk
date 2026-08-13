package rvh

import "github.com/XingfenD/rainyun_api_go_sdk/apis/common"

// 虚拟主机套餐(实测 /product/rvh/plans 响应)
type RvhPlan struct {
	ID              int            `json:"id"`
	Region          string         `json:"region"`
	Line            string         `json:"line"`
	Subtype         string         `json:"subtype"`
	PlanName        string         `json:"plan_name"`
	Machine         string         `json:"machine"`
	ChargeType      string         `json:"charge_type"`
	Chinese         string         `json:"chinese"`
	IsFree          bool           `json:"is_free"`
	PointRenewPrice any            `json:"point_renew_price"`
	IsSelling       bool           `json:"is_selling"`
	Price           int            `json:"price"`
	Tools           int            `json:"tools"`
	Disk            int            `json:"disk"`
	Epdb            int            `json:"epdb"`
	BtCPU           int            `json:"bt_cpu"`
	EpBandwidth     int            `json:"ep_bandwidth"`
	BtRAM           int            `json:"bt_ram"`
	BtNetIn         int            `json:"bt_net_in"`
	BtNetOut        int            `json:"bt_net_out"`
	IsSupportBackup bool           `json:"is_support_backup"`
	AutoRestock     int            `json:"auto_restock"`
	AvailableStock  int            `json:"available_stock"`
	IPPrices        map[string]int `json:"ip_prices"`
}

// 请求类型(以旧 spec definitions 为准)
type GetRvhListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type GetRvhListResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // 记录结构未公开,透传
}

type GetRvhDetailResponse struct {
	Code int `json:"code"`
	Data any `json:"data"` // 记录结构未公开,透传
}

type CreateRvhRequest struct {
	Duration     int    `json:"duration"`
	NodeUUID     string `json:"node_uuid"`
	PlanID       int    `json:"plan_id"`
	Try          bool   `json:"try"`
	WithCouponID int    `json:"with_coupon_id"`
}

type RenewRvhRequest struct {
	Duration     int `json:"duration"`
	WithCouponID int `json:"with_coupon_id"`
}

type EnableRvhAutoRenewRequest struct {
	AutoRenewOption bool `json:"auto_renew_option"`
}

type UpgradeRvhRequest struct {
	DestPlan     int `json:"dest_plan"`
	WithCouponID int `json:"with_coupon_id"`
}

type SetRvhMaintenanceModeRequest struct {
	Value int `json:"value"`
}

type SetRvhTagRequest struct {
	TagName string `json:"tag_name"`
}

type CreateRvhBackupRequest struct {
	Label string `json:"label"`
}

type SetRvhBackupSettingRequest struct {
	AutoBackupHour   int `json:"auto_backup_hour"`
	AutoBackupMinute int `json:"auto_backup_minute"`
	KeepLast         int `json:"keep_last"`
}

type BindRvhDomainRequest struct {
	Domain    string `json:"domain"`
	SSLCertID int    `json:"ssl_cert_id"`
	SSLForce  int    `json:"ssl_force"`
}

type SetRvhFirewallOptionRequest struct {
	Option string `json:"option"`
	Value  int    `json:"value"`
}

type SetRvhFirewallRuleRequest struct {
	Option string `json:"option"`
	Value  string `json:"value"`
}

type GetRvhRenewPriceResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type GetRvhPlanListResponse struct {
	Code int       `json:"code"`
	Data []RvhPlan `json:"data"`
}

type GetRvhPriceResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}
