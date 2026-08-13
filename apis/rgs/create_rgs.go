package rgs

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

// 创建游戏云请求
type CreateRgsRequest struct {
	AppVars      []RgsAppVar `json:"app_vars"` // 当空数组时,进行单次任务下发(可选)
	Config       RgsConfig   `json:"config"`
	CPULimitMode bool        `json:"cpu_limit_mode"`
	Duration     int         `json:"duration"`    // 创建时长(月)
	EggTypeID    int         `json:"egg_type_id"` // 游戏类型
	NodeUUID     string      `json:"node_uuid"`
	OnlineMode   bool        `json:"online_mode"`
	OsID         int         `json:"os_id"`
	PanelUser    string      `json:"panel_user"` // 游戏云面板用户
	PayMode      string      `json:"pay_mode"`
	PlanID       int         `json:"plan_id"`
	Subtype      string      `json:"subtype"` // kvm/mcsm
	Try          bool        `json:"try"`
	WithCouponID int         `json:"with_coupon_id"`
	WithEipFlags string      `json:"with_eip_flags"` // 是否开启高防，us_ddosip -> 美国高防，nb_ddosip -> 宁波高防
	WithEipNum   int         `json:"with_eip_num"`
	WithEipType  string      `json:"with_eip_type"`
	Zone         string      `json:"zone"`
}

type RgsAppVar struct {
	AppID int  `json:"app_id"`
	Retry bool `json:"retry"` // 重发之前的任务,此项存在时,无需传入参数
	Vars  any  `json:"vars"`
}

// 游戏云配置信息
type RgsConfig struct {
	CPU        int `json:"cpu"`
	Memory     int `json:"memory"`
	NetOut     int `json:"net_out"`
	NetIn      int `json:"net_in"`
	BaseDisk   int `json:"base_disk"`
	DataDisk   int `json:"data_disk"`
	Allocation int `json:"allocation"`
	Database   int `json:"database"`
	Backup     int `json:"backup"`
}

// 创建游戏云响应
type CreateRgsResponse struct {
	Code int           `json:"code"`
	Data RgsCreateData `json:"data"`
}

// RgsCreateData 创建游戏云返回的完整实例数据
type RgsCreateData struct {
	RgsRecord
	Node          RgsNode `json:"Node"` // 节点信息
	Plan          RgsPlan `json:"Plan"`
	EggTypeID     int     `json:"EggTypeId"`
	EggType       any     `json:"EggType"` // TODO: 结构未公开,实测后补强类型
	PteroUserName string  `json:"PteroUserName"`
	PteroUser     any     `json:"PteroUser"` // TODO: 结构未公开,实测后补强类型
	ServerID      int     `json:"ServerID"`
	AllocationID  int     `json:"AllocationID"`
	ServerUUID    string  `json:"ServerUUID"`
	DaemonUUID    string  `json:"DaemonUUID"`
	GameInfo      any     `json:"GameInfo"` // TODO: 结构未公开,实测后补强类型
}

type RgsNode struct {
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

type RgsPlan struct {
	ID               int             `json:"id"`
	Region           string          `json:"region"`
	Subtype          string          `json:"subtype"`
	PlanName         string          `json:"plan_name"`
	Machine          string          `json:"machine"`
	ChargeType       string          `json:"charge_type"`
	Chinese          string          `json:"chinese"`
	IsFree           bool            `json:"is_free"`
	PointRenewPrice  any             `json:"point_renew_price"` // TODO: 结构未公开,实测后补强类型
	IsSelling        bool            `json:"is_selling"`
	StockDiscount    int             `json:"stock_discount"`
	EipStockDiscount int             `json:"eip_stock_discount"`
	IPPrices         any             `json:"ip_prices"`  // TODO: 结构未公开,实测后补强类型
	IPSelling        any             `json:"ip_selling"` // TODO: 结构未公开,实测后补强类型
	CPUPrice         int             `json:"cpu_price"`
	MemoryPrice      int             `json:"memory_price"`
	NetInPrice       int             `json:"net_in_price"`
	NetOutPrice      int             `json:"net_out_price"`
	BaseDiskPrice    int             `json:"base_disk_price"`
	DataDiskPrice    int             `json:"data_disk_price"`
	Config           []RgsPlanConfig `json:"config"` // 这里不知道为什么要返还一堆套餐列表
	AutoRestock      int             `json:"auto_restock"`
	AvailableStock   int             `json:"available_stock"`
	CPUPointDefault  int             `json:"cpu_point_default"`
	CPUPointConsume  int             `json:"cpu_point_consume"`
	CPUPointPrice    float64         `json:"cpu_point_price"`
	CPUBase          float64         `json:"cpu_base"`
	CPUMax           int             `json:"cpu_max"`
	EipPrice         int             `json:"eip_price"`
	DefencePrice     int             `json:"defence_price"`
	AllocationPrice  int             `json:"allocation_price"`
	DatabasePrice    int             `json:"database_price"`
	BackupPrice      int             `json:"backup_price"`
	DailyModeSupport bool            `json:"daily_mode_support"`
	DailyPriceScale  int             `json:"daily_price_scale"`
}

type RgsPlanConfig struct {
	CPU         int `json:"cpu"`
	Memory      int `json:"memory"`
	NetIn       int `json:"net_in,omitempty"`
	CPUMax      int `json:"cpu_max,omitempty"`
	CPUMin      int `json:"cpu_min,omitempty"`
	NetOut      int `json:"net_out"`
	BaseDisk    int `json:"base_disk,omitempty"`
	DataDisk    int `json:"data_disk,omitempty"`
	BasePrice   int `json:"base_price"`
	MemoryMax   int `json:"memory_max"`
	MemoryMin   int `json:"memory_min"`
	NetInMax    int `json:"net_in_max,omitempty"`
	NetInMin    int `json:"net_in_min,omitempty"`
	NetOutMax   int `json:"net_out_max,omitempty"`
	NetOutMin   int `json:"net_out_min"`
	BaseDiskMax int `json:"base_disk_max,omitempty"`
	BaseDiskMin int `json:"base_disk_min,omitempty"`
	DataDiskMax int `json:"data_disk_max,omitempty"`
	DataDiskMin int `json:"data_disk_min,omitempty"`
}

// 创建游戏云
func (s *RgsService) CreateRgs(req *CreateRgsRequest) (*CreateRgsResponse, error) {
	path := "/product/rgs"

	var resp CreateRgsResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
