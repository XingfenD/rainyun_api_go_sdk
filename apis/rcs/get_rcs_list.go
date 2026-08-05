package rcs

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRcsListResponse struct {
	Code int             `json:"code"`
	Data RcsListResponse `json:"data"`
}

type Node struct {
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
	Stock             any    `json:"Stock"` // NOTE: any?
	StatusData        string `json:"StatusData"`
	ShowMonitorData   string `json:"ShowMonitorData"`
	UpdateTime        string `json:"UpdateTime"`
	GitRepositoryName string `json:"GitRepositoryName"`
	CertifyRequired   bool   `json:"CertifyRequired"`
	IsDisableBackup   bool   `json:"IsDisableBackup"`
	IsHidden          bool   `json:"IsHidden"`
	NodeName          string `json:"NodeName"`
}

// NOTE: strange struct naming, needing checking
type DiskUsageData struct {
	NAMING_FAILED struct {
		Total int64 `json:"Total"`
		Used  int64 `json:"Used"`
	} `json:"/"` // NOTE: strange json tag
}

type UsageData struct {
	CPU         float64       `json:"CPU"`
	MaxMem      int64         `json:"MaxMem"`
	FreeMem     int64         `json:"FreeMem"`
	UsedMem     int           `json:"UsedMem"`
	Disks       DiskUsageData `json:"Disks"`
	DiskRead    int           `json:"DiskRead"`
	DiskWrite   float64       `json:"DiskWrite"`
	NetOut      float64       `json:"NetOut"`
	NetIn       float64       `json:"NetIn"`
	SmartHealth any           `json:"SmartHealth"`
	SmartTemp   int           `json:"SmartTemp"`
	UpdateTime  int           `json:"UpdateTime"`
}

type TrafficPrice struct {
	Num300  int `json:"300"`
	Num1024 int `json:"1024"`
	Num2048 int `json:"2048"`
}

type DiskPrice struct {
	Ssd float64 `json:"ssd"`
	Hdd float64 `json:"hdd"`
}

type Plan struct {
	ID              int          `json:"id"`
	Region          string       `json:"region"`
	Subtype         string       `json:"subtype"`
	PlanName        string       `json:"plan_name"`
	Machine         string       `json:"machine"`
	ChargeType      string       `json:"charge_type"`
	Chinese         string       `json:"chinese"`
	IsFree          bool         `json:"is_free"`
	PointRenewPrice any          `json:"point_renew_price"`
	IsSelling       bool         `json:"is_selling"`
	Price           int          `json:"price"`
	TrafficBaseGb   int          `json:"traffic_base_gb"`
	TrafficPrice    TrafficPrice `json:"traffic_price"`
	CPU             int          `json:"cpu"`
	Memory          int          `json:"memory"`
	NetIn           int          `json:"net_in"`
	NetOut          int          `json:"net_out"`
	IPPrices        any          `json:"ip_prices"`
	IPSelling       any          `json:"ip_selling"`
	AutoRestock     int          `json:"auto_restock"`
	AvailableStock  int          `json:"available_stock"`
	GpuMemorySize   int          `json:"gpu_memory_size"`
	DgpuDevType     string       `json:"dgpu_dev_type"`
	WebbarConfig    any          `json:"webbar_config"`
	NoBackup        bool         `json:"no_backup"`
	DiskPrice       DiskPrice    `json:"disk_price"`
}

type OsInfo struct {
	ID             int    `json:"id"`
	Region         string `json:"region"`
	Subtype        string `json:"subtype"`
	Machine        string `json:"machine"`
	Name           string `json:"name"`
	Version        string `json:"version"`
	SyncStatus     string `json:"sync_status"`
	OsType         string `json:"os_type"`
	ChineseName    string `json:"chinese_name"`
	Icon           string `json:"icon"`
	IsWithBbr      bool   `json:"is_with_bbr"`
	IsEol          bool   `json:"is_eol"`
	IsAvailable    bool   `json:"is_available"`
	Order          int    `json:"order"`
	LatestFilename string `json:"latest_filename"`
	NoVMAgent      bool   `json:"no_vm_agent"`
}

type RcsRecord struct {
	ExpDate                 int       `json:"ExpDate"`
	ExpireNotice            int       `json:"ExpireNotice"`
	AutoRenew               bool      `json:"AutoRenew"`
	UnsubscribeAble         bool      `json:"UnsubscribeAble"`
	Try                     bool      `json:"Try"`
	ID                      int       `json:"ID"`
	UID                     int       `json:"UID"`
	PlanID                  int       `json:"PlanID"`
	CreateDate              int       `json:"CreateDate"`
	NodeUUID                string    `json:"NodeUUID"`
	Node                    Node      `json:"Node"`
	Status                  string    `json:"Status"`
	StopReason              string    `json:"StopReason"`
	RewardPointsToBeCollect int       `json:"RewardPointsToBeCollect"`
	Tag                     string    `json:"Tag"`
	OsID                    int       `json:"OsID"`
	OsName                  string    `json:"OsName"`
	HostName                string    `json:"HostName"`
	DefaultPass             string    `json:"DefaultPass"`
	MainIPv4                string    `json:"MainIPv4"`
	IntIPv4                 string    `json:"IntIPv4"`
	UsageData               UsageData `json:"UsageData"`
	Zone                    string    `json:"Zone"`
	NatPublicIP             string    `json:"NatPublicIP"`
	NatPublicDomain         string    `json:"NatPublicDomain"`
	NATSpareDomain          string    `json:"NATSpareDomain"`
	NetIn                   int       `json:"NetIn"`
	NetOut                  int       `json:"NetOut"`
	NowNetIn                int       `json:"NowNetIn"`
	NowNetOut               int       `json:"NowNetOut"`
	NetMode                 string    `json:"NetMode"`
	BridgeSyncing           bool      `json:"BridgeSyncing"`
	VnetID                  int       `json:"VnetID"`
	UpdateTime              int       `json:"UpdateTime"`
	FwSyncTime              int       `json:"FwSyncTime"`
	FwMode                  string    `json:"FwMode"`
	AbCPULimit              int       `json:"AbCpuLimit"`
	AbNetLimit              int       `json:"AbNetLimit"`
	AbWhiteReason           string    `json:"AbWhiteReason"`
	TrafficBytes            int64     `json:"TrafficBytes"`
	TrafficResetDate        int       `json:"TrafficResetDate"`
	TrafficBytesToday       int       `json:"TrafficBytesToday"`
	TrafficBytesDayLimit    int64     `json:"TrafficBytesDayLimit"`
	TrafficOnLimit          int       `json:"TrafficOnLimit"`
	Plan                    Plan      `json:"Plan"`
	OsInfo                  OsInfo    `json:"OsInfo"`
	CPU                     int       `json:"CPU"`
	Memory                  int       `json:"Memory"`
	Disk                    int       `json:"Disk"`
	RBSKeepLast             int       `json:"RBSKeepLast"`
	RBSAutoBackup           string    `json:"RBSAutoBackup"`
	RBSLastAutoBackupDate   int       `json:"RBSLastAutoBackupDate"`
	FastAppInstallTaskKey   string    `json:"FastAppInstallTaskKey"`
	GPUDevice               string    `json:"GPUDevice"`
	GPUMemorySize           int       `json:"GPUMemorySize"`
	DGPUEnable              bool      `json:"DGPUEnable"`
	NoPrimaryGPU            bool      `json:"NoPrimaryGPU"`
	WebbarMinutes           int       `json:"WebbarMinutes"`
	WebbarResetDate         int       `json:"WebbarResetDate"`
}

type RcsListResponse struct {
	TotalRecords int         `json:"TotalRecords"`
	Records      []RcsRecord `json:"Records"`
}

type GetRcsListRequest struct {
	IsRGpu  *bool                       `json:"is_rgpu"` // optional
	Options common.StandQueryParameters `json:"options"`
}

func (req *GetRcsListRequest) BuildQueryMap() map[string]string {
	m, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil
	}
	return m
}

func (s *RcsService) GetRcsList(req *GetRcsListRequest) (*GetRcsListResponse, error) {
	path := "/product/rcs"

	var resp GetRcsListResponse

	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
