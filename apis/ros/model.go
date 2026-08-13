package ros

// RosNode 节点信息
type RosNode struct {
	UUID              string `json:"UUID"`
	AuthKey           string `json:"AuthKey"`
	Region            string `json:"Region"`
	IPRegion          string `json:"IpRegion"`
	Machine           string `json:"Machine"`
	Product           string `json:"Product"`
	Subtype           string `json:"Subtype"`
	ChineseName       string `json:"ChineseName"`
	PhysicalNode      string `json:"PhysicalNode"`
	Config            string `json:"Config"`
	Stock             any    `json:"Stock"` // TODO: 结构未公开,实测后补强类型
	StatusData        string `json:"StatusData"`
	ShowMonitorData   string `json:"ShowMonitorData"`
	UpdateTime        string `json:"UpdateTime"`
	GitRepositoryName string `json:"GitRepositoryName"`
	CertifyRequired   bool   `json:"CertifyRequired"`
	IsDisableBackup   bool   `json:"IsDisableBackup"`
	IsHidden          bool   `json:"IsHidden"`
	NodeName          string `json:"NodeName"`
}

// RosPlan 套餐信息
type RosPlan struct {
	ID                 int     `json:"id"`
	Region             string  `json:"region"`
	Subtype            string  `json:"subtype"`
	PlanName           string  `json:"plan_name"`
	Machine            string  `json:"machine"`
	ChargeType         string  `json:"charge_type"`
	Chinese            string  `json:"chinese"`
	IsFree             bool    `json:"is_free"`
	PointRenewPrice    any     `json:"point_renew_price"` // TODO: 结构未公开,实测后补强类型
	IsSelling          bool    `json:"is_selling"`
	Price              int     `json:"price"`
	StorageSize        int     `json:"storage_size"`
	Bandwidth          int     `json:"bandwidth"`
	ExtraTransferPrice float64 `json:"extra_transfer_price"`
	ExtraStoragePrice  float64 `json:"extra_storage_price"`
}

// RosUsageData 实例使用数据
type RosUsageData struct {
	MetricsUsageTotalBytes int `json:"metrics_usage_total_bytes"` // 存储使用量
	MetricsReceivedBytes   int `json:"metrics_received_bytes"`    // 接收字节数
	MetricsSentBytes       int `json:"metrics_sent_bytes"`        // 发送字节数
	MetricsRequests        int `json:"metrics_requests"`          // 请求数
	SpeedReceivedBytes     int `json:"speed_received_bytes"`      // 接收字节速度
	SpeedSentBytes         int `json:"speed_sent_bytes"`          // 发送字节速度
	SpeedRequests          int `json:"speed_requests"`            // 请求速度
	UpdateTime             int `json:"UpdateTime"`                // 更新时间
}

// RosInstance 对象存储实例信息
type RosInstance struct {
	ID         int     `json:"ID"`
	UID        int     `json:"UID"`
	PlanID     int     `json:"PlanID"`
	CreateDate int     `json:"CreateDate"` // 创建时间
	NodeUUID   string  `json:"NodeUUID"`   // 节点UUID
	Node       RosNode `json:"Node"`       // 节点信息
	Status     string  `json:"Status"`
	StopReason string  `json:"StopReason"`

	RewardPointsToBeCollect int          `json:"RewardPointsToBeCollect"` // 待领取积分奖励
	Tag                     string       `json:"Tag"`
	Plan                    RosPlan      `json:"Plan"`
	AccessKey               string       `json:"access_key"`                 // ⚠️列表接口拿不到AK，响应里面的AK是空的
	SecretKey               string       `json:"secret_key"`                 // ⚠️列表接口拿不到SK，响应里面的SK是空的
	IsPublicAccess          bool         `json:"is_public_access"`           // 是否允许公共访问
	IsEnableExtraAccounting bool         `json:"is_enable_extra_accounting"` // 是否启用额外计费
	LastResetDate           int          `json:"last_reset_date"`            // 最后重置时间
	NextResetDate           int          `json:"next_reset_date"`            // 下次重置时间
	ExpDate                 int          `json:"ExpDate"`                    // 到期时间
	ExpireNotice            int          `json:"ExpireNotice"`
	AutoRenew               bool         `json:"AutoRenew"` // 是否自动续费
	UnsubscribeAble         bool         `json:"UnsubscribeAble"`
	Try                     bool         `json:"Try"` // 是否试用
	PublicAPIURL            string       `json:"public_api_url"`
	Buckets                 any          `json:"buckets"` // TODO: 结构未公开,实测后补强类型
	ExtraPayTime            int          `json:"extra_pay_time"`
	UsageData               RosUsageData `json:"UsageData"`
	Restrictions            string       `json:"restrictions"`
	PublicAccessRestricted  bool         `json:"public_access_restricted"`
}

// RosBucket 对象存储桶信息
type RosBucket struct {
	ID                           int          `json:"id"`   // 存储桶ID
	Name                         string       `json:"name"` // 存储桶名称
	UID                          int          `json:"uid"`  // 用户ID
	CreateDate                   int          `json:"create_date"`
	AccessKey                    string       `json:"access_key"`       // ⚠️列表接口拿不到AK，响应里面的AK是空的
	SecretKey                    string       `json:"secret_key"`       // ⚠️列表接口拿不到SK，响应里面的SK是空的
	IsPublicAccess               bool         `json:"is_public_access"` // 是否公开
	InstanceID                   int          `json:"instance_id"`      // 实例ID
	Instance                     RosInstance  `json:"Instance"`
	StopReason                   string       `json:"stop_reason"`
	DomainList                   any          `json:"DomainList"` // TODO: 结构未公开,实测后补强类型
	SslIsEnable                  bool         `json:"ssl_is_enable"`
	SSLCertList                  any          `json:"SSLCertList"` // TODO: 结构未公开,实测后补强类型
	SslAutoRedirect              bool         `json:"ssl_auto_redirect"`
	WafIsEnable                  bool         `json:"waf_is_enable"`
	WafGlobalQPS                 int          `json:"waf_global_qps"`
	WafPerIPQPS                  int          `json:"waf_per_ip_qps"`
	WafBlockTime                 int          `json:"waf_block_time"`
	WafGlobalJsCheck             bool         `json:"waf_global_js_check"`
	WafPerIPJsCheck              bool         `json:"waf_per_ip_js_check"`
	RefererRestrictIsEnable      bool         `json:"referer_restrict_is_enable"`
	RefererRestrictWhitelist     any          `json:"RefererRestrictWhitelist"` // TODO: 结构未公开,实测后补强类型
	RefererRestrictBlacklist     any          `json:"RefererRestrictBlacklist"` // TODO: 结构未公开,实测后补强类型
	RefererRestrictBypassMissing bool         `json:"referer_restrict_bypass_missing"`
	IPRestrictIsEnable           bool         `json:"ip_restrict_is_enable"`
	IPRestrictWhitelist          any          `json:"IPRestrictWhitelist"` // TODO: 结构未公开,实测后补强类型
	IPRestrictBlacklist          any          `json:"IPRestrictBlacklist"` // TODO: 结构未公开,实测后补强类型
	GzipIsEnable                 bool         `json:"gzip_is_enable"`
	GzipCompressLevel            int          `json:"gzip_compress_level"`
	UsageData                    RosUsageData `json:"UsageData"`
	ReadyDelete                  bool         `json:"ready_delete"`
}

// 获取监控数据请求（桶和实例通用）
type GetRosMonitorDataRequest struct {
	StartDate int `json:"start_date"` // 开始时间（timestamp）
	EndDate   int `json:"end_date"`   // 结束时间（timestamp）
}

// 获取监控数据响应（桶和实例通用）
type GetRosMonitorDataResponse struct {
	Code int            `json:"code"`
	Data RosMonitorData `json:"data"`
}

// RosMonitorData 对象存储监控数据（桶和实例通用）
type RosMonitorData struct {
	Columns []string `json:"Columns"`
	Values  [][]any  `json:"Values"` // TODO: 结构未公开,实测后补强类型
}
