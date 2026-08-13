package rcdn

import "github.com/XingfenD/rainyun_api_go_sdk/apis/common"

// 实例记录(rcdnM.RCDN)
type RcdnInstance struct {
	UsageData               any      `json:"UsageData"` // 使用状况(缓存) TODO: 结构未公开,实测后补强类型
	AutoRenew               bool     `json:"autoRenew"`
	CreateDate              int      `json:"createDate"`
	DomainConfig            any      `json:"domainConfig"` // 域名默认CDN配置 TODO: 结构未公开,实测后补强类型
	ExpDate                 int      `json:"expDate"`
	ExpireNotice            int      `json:"expireNotice"`
	ExtraPayTime            int      `json:"extra_pay_time"`
	ExtraTraffic            int      `json:"extra_traffic"`
	ID                      int      `json:"id"`
	IsEnableExtraAccounting bool     `json:"is_enable_extra_accounting"`
	LastResetDate           int      `json:"last_reset_date"`
	NextResetDate           int      `json:"next_reset_date"`
	Node                    RcdnNode `json:"node"`
	NodeUUID                string   `json:"nodeUUID"`
	Plan                    RcdnPlan `json:"plan"`
	PlanID                  int      `json:"planID"`
	RewardPointsToBeCollect int      `json:"rewardPointsToBeCollect"`
	Status                  string   `json:"status"`
	StopReason              string   `json:"stopReason"`
	SyncTime                int      `json:"sync_time"`
	Tag                     string   `json:"tag"`
	Try                     bool     `json:"try"`
	UID                     int      `json:"uid"`
	UnsubscribeAble         bool     `json:"unsubscribeAble"`
	UsageTraffic            int      `json:"usage_traffic"`
}

// 节点信息(productM.Node)
type RcdnNode struct {
	UUID              string `json:"uuid"`
	AuthKey           string `json:"AuthKey"`
	Config            string `json:"Config"`
	Stock             []int  `json:"Stock"`
	CertifyRequired   bool   `json:"certifyRequired"`
	ChineseName       string `json:"chineseName"`
	GitRepositoryName string `json:"gitRepositoryName"`
	IPZone            string `json:"ipZone"`
	IsDisableBackup   bool   `json:"isDisableBackup"`
	IsHidden          bool   `json:"isHidden"`
	Machine           string `json:"machine"`
	NodeName          string `json:"nodeName"`
	PhysicalNode      string `json:"physicalNode"`
	Product           string `json:"product"`
	Region            string `json:"region"`
	ShowMonitorData   string `json:"showMonitorData"`
	StatusData        string `json:"statusData"`
	Subtype           string `json:"subtype"`
	UpdateTime        string `json:"updateTime"`
}

// 套餐(rcdnM.RCDNPlan,字段已实测确认)
type RcdnPlan struct {
	ChargeType        string  `json:"charge_type"`
	Chinese           string  `json:"chinese"`
	DomainLimit       int     `json:"domain_limit"`
	ExtraTrafficPrice float64 `json:"extra_traffic_price"`
	ID                int     `json:"id"`
	IsFree            bool    `json:"is_free"`
	IsSelling         bool    `json:"is_selling"`
	Line              string  `json:"line"`
	Machine           string  `json:"machine"`
	PlanName          string  `json:"plan_name"`
	PointRenewPrice   any     `json:"point_renew_price"` // TODO: 结构未公开,实测后补强类型
	Price             float64 `json:"price"`
	RainDiscount      float64 `json:"rain_discount"`
	Region            string  `json:"region"`
	Subtype           string  `json:"subtype"`
	TrafficInGb       int     `json:"traffic_in_gb"`
}

// 加速域名记录(rcdnM.RCDNDomain)
type RcdnDomain struct {
	UsageData     any    `json:"UsageData"` // TODO: 结构未公开,实测后补强类型
	CNAME         string `json:"cname"`
	Config        any    `json:"config"` // datatypes.JSONType-rcdnM_DomainConfig TODO: 结构未公开,实测后补强类型
	CreateDate    int    `json:"create_date"`
	DefenceReason string `json:"defence_reason"`
	DefenceTime   int    `json:"defence_time"`
	Domain        string `json:"domain"`
	ID            int    `json:"id"`
	Product       any    `json:"product"` // TODO: 结构未公开,实测后补强类型
	ProductID     int    `json:"product_id"`
	ReadyDelete   bool   `json:"ready_delete"`
	Region        string `json:"region"`
	UID           int    `json:"uid"`
	UpdateTime    int    `json:"update_time"`
}

// 回源配置(rcdnM.BaseConfig)
type RcdnBaseConfig struct {
	DefaultMaster []string `json:"default_master"`
	DefaultSlave  []string `json:"default_slave"`
	OriginMode    string   `json:"origin_mode"`
	OriginPort    int      `json:"origin_port"`
}

// 域名配置(rcdnM.DomainConfig)
type RcdnDomainConfig struct {
	ACLConfig      RcdnACLConfig     `json:"acl_config"`
	BaseConfig     RcdnBaseConfig    `json:"baseConfig"`
	CacheConfigs   []RcdnCacheConfig `json:"cache_configs"`
	OriginHeaders  []any             `json:"origin_headers"` // TODO: 结构未公开,实测后补强类型
	OriginHost     string            `json:"origin_host"`
	RespConfig     []RcdnRespConfig  `json:"resp_config"`
	RespRemoveList []string          `json:"resp_remove_list"`
	SpeedLimits    []RcdnSpeedLimit  `json:"speed_limits"`
	SSLConfig      RcdnSSLConfig     `json:"ssl_config"`
	WAFConfig      any               `json:"wafConfig"` // TODO: 结构未公开,实测后补强类型
}

type RcdnACLConfig struct {
	IPBlackList       []string `json:"ip_black_list"`
	IPBlackMode       string   `json:"ip_black_mode"`
	RefererAllowEmpty bool     `json:"referer_allow_empty"`
	RefererList       []string `json:"referer_list"`
	RefererMode       string   `json:"referer_mode"`
	UAAllowEmpty      bool     `json:"ua_allow_empty"`
	UAList            []string `json:"ua_list"`
	UAMode            string   `json:"ua_mode"`
}

type RcdnCacheConfig struct {
	CaseIgnore           string `json:"case_ignore"`
	Expire               int    `json:"expire"`
	ExpireUnit           string `json:"expire_unit"`
	FollowExpired        string `json:"follow_expired"`
	IgnoreNoCacheHeaders string `json:"ignore_no_cache_headers"`
	MatchMethod          string `json:"match_method"`
	Params               string `json:"params"`
	Pattern              string `json:"pattern"`
	Priority             int    `json:"priority"`
	QueryParamsOp        string `json:"query_params_op"`
	QueryParamsOpWay     string `json:"query_params_op_way"`
	QueryParamsOpWhen    string `json:"query_params_op_when"`
}

type RcdnRespConfig struct {
	Cover   string `json:"cover"`
	Name    string `json:"name"`
	OnlyHit string `json:"only_hit"`
	Value   string `json:"value"`
}

type RcdnSpeedLimit struct {
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
	MatchType string `json:"match_type"`
	Pattern   string `json:"pattern"`
	Speed     int    `json:"speed"`
}

type RcdnSSLConfig struct {
	CertID   int  `json:"cert_id"`
	ForceSSL bool `json:"force_ssl"`
}

// 列表响应
type GetRcdnInstanceListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type GetRcdnInstanceListResponse struct {
	Code int              `json:"code"`
	Data RcdnInstanceList `json:"data"`
}

type RcdnInstanceList struct {
	TotalRecords int            `json:"TotalRecords"`
	Records      []RcdnInstance `json:"Records"`
}

type GetRcdnInstanceDetailResponse struct {
	Code int          `json:"code"`
	Data RcdnInstance `json:"data"`
}

type GetRcdnDomainListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type GetRcdnDomainListResponse struct {
	Code int            `json:"code"`
	Data RcdnDomainList `json:"data"`
}

type RcdnDomainList struct {
	TotalRecords int          `json:"TotalRecords"`
	Records      []RcdnDomain `json:"Records"`
}

type GetRcdnDomainDetailResponse struct {
	Code int        `json:"code"`
	Data RcdnDomain `json:"data"`
}

// 请求类型(以旧 spec definitions 为准;renew/auto-renew/tag 沿用 rcs 同构约定)
type CreateRcdnInstanceRequest struct {
	Config       *RcdnBaseConfig `json:"config"`
	Domains      []string        `json:"domains"`
	Duration     int             `json:"duration"`
	PlanID       int             `json:"plan_id"`
	WithCouponID int             `json:"with_coupon_id"`
}

type RenewRcdnInstanceRequest struct {
	Duration     int `json:"duration"`
	WithCouponID int `json:"with_coupon_id"`
}

type EnableRcdnInstanceAutoRenewRequest struct {
	AutoRenewOption bool `json:"auto_renew_option"`
}

type ScaleRcdnInstanceRequest struct {
	DestPlan     int `json:"dest_plan"`
	WithCouponID int `json:"with_coupon_id"`
}

type SetRcdnInstanceTagRequest struct {
	TagName string `json:"tag_name"`
}

type SetRcdnInstanceSettingRequest struct {
	AsDefault bool              `json:"as_default"`
	Config    *RcdnDomainConfig `json:"config"`
	Domains   []string          `json:"domains"`
	Type      string            `json:"type"`
}

type BindRcdnSSLDomainsRequest struct {
	CertID  int      `json:"cert_id"`
	Domains []string `json:"domains"`
}

type AddRcdnDomainRequest struct {
	CopyFromDomain string `json:"copy_from_domain"`
	Domain         string `json:"domain"`
	InstanceID     int    `json:"instance_id"`
}

type RefreshRcdnCacheRequest struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}

type GetRcdnUsageRequest struct {
	Begin   int    `json:"begin"`
	Domains string `json:"domains"`
	Each    string `json:"each"`
	End     int    `json:"end"`
	Type    string `json:"type"`
	Unit    string `json:"unit"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcdnUsageResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type GetRcdnMonitorDataRequest struct {
	Begin   int    `json:"begin"`
	Domains string `json:"domains"`
	Each    string `json:"each"`
	End     int    `json:"end"`
	Type    string `json:"type"`
	Unit    string `json:"unit"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcdnMonitorDataResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type GetRcdnPlanListResponse struct {
	Code int        `json:"code"`
	Data []RcdnPlan `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcdnPriceResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}
