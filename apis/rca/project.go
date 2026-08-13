package rca

import (
	"fmt"
	"strconv"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 云应用区域信息
type PriceInfo struct {
	Cpu     float64 `json:"cpu"`
	Memory  float64 `json:"memory"`
	Ipv4    float64 `json:"ipv4"`
	Traffic float64 `json:"traffic"`
	Disk    float64 `json:"disk"`
}

type RcaRegion struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	ChineseName          string    `json:"chinese_name"`
	WebsiteServiceDomain string    `json:"website_service_domain"`
	SftpServiceDomain    string    `json:"sftp_service_domain"`
	PublicServiceDomain  string    `json:"public_service_domain"`
	PriceInfo            PriceInfo `json:"price_info"`
}

type GetRcaRegionInfoResponse struct {
	Code int         `json:"code"`
	Data []RcaRegion `json:"data"`
}

// 创建云应用项目请求
type CreateRcaProjectRequest struct {
	ChargeType  string `json:"charge_type"`  // 计费类型: 动态计费: elastic
	CPULimit    int    `json:"cpu_limit"`    // （仅限package模式）CPU限制（毫核），1核心=1000，0.1核=100
	DiskSize    int    `json:"disk_size"`    // 磁盘大小（GiB）
	Ipv4Count   int    `json:"ipv4_count"`   // 要添加的IPv4地址数量
	Ipv6Count   int    `json:"ipv6_count"`   // 要添加的IPv6地址数量
	MemoryLimit int    `json:"memory_limit"` // （仅限package模式）内存限制（MiB）
	Name        string `json:"name"`         // 名称
	RegionID    int    `json:"region_id"`    // 部署区域
}

// 云应用节点信息
type RcaNode struct {
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
	Stock             any    `json:"Stock"`
	StatusData        string `json:"StatusData"`
	ShowMonitorData   string `json:"ShowMonitorData"`
	UpdateTime        string `json:"UpdateTime"`
	GitRepositoryName string `json:"GitRepositoryName"`
	CertifyRequired   bool   `json:"CertifyRequired"`
	IsDisableBackup   bool   `json:"IsDisableBackup"`
	IsHidden          bool   `json:"IsHidden"`
	NodeName          string `json:"NodeName"`
}

type RcaResourceLimits struct {
	MaxCPU    int `json:"max_cpu"`
	MaxMemory int `json:"max_memory"`
}

// 云应用项目资源使用情况
type RcaProjectUsageData struct {
	CPU             int    `json:"cpu"`
	Memory          int    `json:"memory"`
	NetOut          int    `json:"net_out"`
	NetIn           int    `json:"net_in"`
	DiskUsage       int    `json:"disk_usage"`
	TrafficToday    int    `json:"traffic_today"`
	Status          string `json:"status"`
	StatusReason    string `json:"status_reason"`
	AllocatedCPU    int    `json:"allocated_cpu"`
	AllocatedMemory int    `json:"allocated_memory"`
	AppCount        int    `json:"app_count"`
	WebsiteCount    int    `json:"website_count"`
	DatabaseCount   int    `json:"database_count"`
	Ipv4Count       int    `json:"ipv4_count"`
	HealthyPods     int    `json:"healthy_pods"`
	UnhealthyPods   int    `json:"unhealthy_pods"`
}

// 备份目标
type RcaBackupTarget struct {
	Type              string `json:"type"`                // 目标类型，支持项目本地备份或者远程S3存储(local/s3)
	S3Endpoint        string `json:"s3_endpoint"`         // S3的端点（仅支持Virtual Host，不支持Path-Style模式）
	S3Bucket          string `json:"s3_bucket"`           // S3存储桶名
	S3AccessKey       string `json:"s3_access_key"`       // S3的AK
	S3SecretKey       string `json:"s3_secret_key"`       // S3的SK
	S3BackupDirectory string `json:"s3_backup_directory"` // s3备份存储的目录
}

// 云应用项目
type RcaProject struct {
	ID                      int                 `json:"ID"`
	UID                     int                 `json:"UID"`
	PlanID                  int                 `json:"PlanID"`
	CreateDate              int                 `json:"CreateDate"` // 创建时间
	NodeUUID                string              `json:"NodeUUID"`
	Node                    RcaNode             `json:"Node"`
	Status                  string              `json:"Status"`
	StopReason              string              `json:"StopReason"`
	RewardPointsToBeCollect int                 `json:"RewardPointsToBeCollect"`
	Tag                     string              `json:"Tag"`
	ExpDate                 int                 `json:"ExpDate"`
	ExpireNotice            int                 `json:"ExpireNotice"`
	AutoRenew               bool                `json:"AutoRenew"`
	UnsubscribeAble         bool                `json:"UnsubscribeAble"`
	Try                     bool                `json:"Try"`
	Name                    string              `json:"name"`
	RegionID                int                 `json:"region_id"`
	Region                  *RcaRegion          `json:"region"`
	Namespace               string              `json:"namespace"`
	APIToken                string              `json:"APIToken"`
	ResourceLimits          RcaResourceLimits   `json:"resource_limits"`
	UsageData               RcaProjectUsageData `json:"usage_data"`
	VolumeSize              int                 `json:"volume_size"`
	ChargeType              string              `json:"charge_type"`  // 计费类型: elastic
	HourlyPrice             float64             `json:"hourly_price"` // 小时价格
	NextChargeTime          int                 `json:"next_charge_time"`
	BackupTarget            RcaBackupTarget     `json:"backup_target"` // 备份目标
	SftpSetting             any                 `json:"sftp_setting"`
	IdleAlarmFlag           bool                `json:"idle_alarm_flag"`
	PaymentDueEnd           int                 `json:"payment_due_end"`
}

type CreateRcaProjectResponse struct {
	Code int        `json:"code"`
	Data RcaProject `json:"data"`
}

type RcaProjectListData struct {
	TotalRecords int          `json:"TotalRecords"`
	Records      []RcaProject `json:"Records"`
}

type GetRcaProjectListResponse struct {
	Code int                `json:"code"`
	Data RcaProjectListData `json:"data"`
}

// 云应用项目的指标信息,示例如下:
type RcaProjectMetricsData struct {
	Columns []string    `json:"Columns"`
	Values  [][]float64 `json:"Values"`
}

type GetRcaProjectMetricsResponse struct {
	Code int                   `json:"code"`
	Data RcaProjectMetricsData `json:"data"`
}

type RcaProjectDetailData struct {
	Data RcaProject `json:"Data"`
}

type GetRcaProjectDetailResponse struct {
	Code int                  `json:"code"`
	Data RcaProjectDetailData `json:"data"`
}

// 云应用获取区域信息
func (s *RcaService) GetRcaRegionInfo() (*GetRcaRegionInfoResponse, error) {
	path := "/product/rca/region"

	var resp GetRcaRegionInfoResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 创建云应用项目
func (s *RcaService) CreateRcaProject(req *CreateRcaProjectRequest) (*CreateRcaProjectResponse, error) {
	path := "/product/rca/project/"

	var resp CreateRcaProjectResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 云应用列出项目
//
// options: 查询参数, 可用 common.StandQueryParameters 结合 common.MarshalQueryParams 获取.
func (s *RcaService) ListRcaProjects(options string) (*GetRcaProjectListResponse, error) {
	path := "/product/rca/project/"
	querys := map[string]string{"no_metrics": "false", "options": options} // no_metrics 含义不明

	var resp GetRcaProjectListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 销毁云应用项目
//
// id: RCA项目ID
func (s *RcaService) DestroyRcaProject(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

// 获取Rca项目的指标信息
//
// id: RCA项目ID
//
// starttime: 开始时间(timestamp)
//
// endtime: 结束时间(timestamp)
func (s *RcaService) GetRcaProjectMetrics(id int, startTime int, endTime int) (*GetRcaProjectMetricsResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/metrics", id)
	querys := map[string]string{
		"start_time": strconv.Itoa(startTime),
		"end_time":   strconv.Itoa(endTime),
	}

	var resp GetRcaProjectMetricsResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取云应用项目详情
//
// id: RCA项目ID
func (s *RcaService) GetRcaProjectDetail(id int) (*GetRcaProjectDetailResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/", id)

	var resp GetRcaProjectDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
