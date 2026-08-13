package rgs

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRgsListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type GetRgsListResponse struct {
	Code int         `json:"code"`
	Data RgsListData `json:"data"`
}

type RgsListData struct {
	TotalRecords int         `json:"TotalRecords"`
	Records      []RgsRecord `json:"Records"`
}

// 游戏云记录信息
// 由于响应实在过于庞大，我们只维护部分必要的响应，如有扩展需求，请在项目的data文件夹下寻找响应实例自行解码
type RgsRecord struct {
	ExpDate                 int          `json:"ExpDate"`
	ExpireNotice            int          `json:"ExpireNotice"`
	AutoRenew               bool         `json:"AutoRenew"`
	UnsubscribeAble         bool         `json:"UnsubscribeAble"`
	Try                     bool         `json:"Try"`
	ID                      int          `json:"ID"`
	UID                     int          `json:"UID"`
	PlanID                  int          `json:"PlanID"`
	CreateDate              int          `json:"CreateDate"`
	Status                  string       `json:"Status"`
	StopReason              string       `json:"StopReason"`
	RewardPointsToBeCollect int          `json:"RewardPointsToBeCollect"` // 待领取的积分
	Tag                     string       `json:"Tag"`
	OsID                    int          `json:"OsID"`
	OsName                  string       `json:"OsName"`
	HostName                string       `json:"HostName"`
	DefaultPass             string       `json:"DefaultPass"`
	MainIPv4                string       `json:"MainIPv4"`
	IntIPv4                 string       `json:"IntIPv4"`
	UsageData               RgsUsageData `json:"UsageData"`
	Zone                    string       `json:"Zone"`
	NatPublicIP             string       `json:"NatPublicIP"`
	NatPublicDomain         string       `json:"NatPublicDomain"`
	NATSpareDomain          string       `json:"NATSpareDomain"`
	NetIn                   int          `json:"NetIn"`
	NetOut                  int          `json:"NetOut"`
	NowNetIn                int          `json:"NowNetIn"`
	NowNetOut               int          `json:"NowNetOut"`
	NetMode                 string       `json:"NetMode"`
	BridgeSyncing           bool         `json:"BridgeSyncing"`
	VnetID                  int          `json:"VnetID"`
	UpdateTime              int          `json:"UpdateTime"`
	FwSyncTime              int          `json:"FwSyncTime"`
	FwMode                  string       `json:"FwMode"`
	AbCPULimit              int          `json:"AbCpuLimit"`
	AbNetLimit              int          `json:"AbNetLimit"`
	AbWhiteReason           string       `json:"AbWhiteReason"`
	OsInfo                  RgsOsInfo    `json:"OsInfo"`
	CPU                     int          `json:"CPU"`
	Memory                  int          `json:"Memory"`
	BaseDisk                int          `json:"BaseDisk"`
	DataDisk                int          `json:"DataDisk"`
	InitedDate              int          `json:"InitedDate"`
	Allocation              int          `json:"Allocation"`
	Database                int          `json:"Database"`
	Backup                  int          `json:"Backup"`
	CPULimitMode            bool         `json:"CpuLimitMode"`
	CPULimitStatus          bool         `json:"CpuLimitStatus"`
	CPUPoint                int          `json:"CpuPoint"`
	DailyMode               bool         `json:"DailyMode"` // 是否日付模式
	RBSKeepLast             int          `json:"RBSKeepLast"`
	RBSAutoBackup           string       `json:"RBSAutoBackup"`
	RBSLastAutoBackupDate   int          `json:"RBSLastAutoBackupDate"`
	McsmUserName            string       `json:"McsmUserName"`
	McsmUser                McsmUser     `json:"McsmUser"`
}

type RgsUsageData struct {
	CPU        int `json:"CPU"`
	Mem        int `json:"Mem"`
	MemUsage   int `json:"MemUsage"`
	DiskRead   int `json:"DiskRead"`
	DiskWrite  int `json:"DiskWrite"`
	Disk       int `json:"Disk"`
	NetOut     int `json:"NetOut"`
	NetIn      int `json:"NetIn"`
	UpdateTime int `json:"UpdateTime"`
}

type RgsOsInfo struct {
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

// 获取游戏云列表
//
// options: 查询参数 可以用 common.StandQueryParameters 获取
func (s *RgsService) GetRgsList(req *GetRgsListRequest) (*GetRgsListResponse, error) {
	path := "/product/rgs"

	var resp GetRgsListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
