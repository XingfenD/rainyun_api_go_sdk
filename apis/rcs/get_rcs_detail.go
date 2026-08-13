package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRcsDetailResponse struct {
	Code int       `json:"code"`
	Data RcsDetail `json:"data"`
}

type RcsDetail struct {
	Data               RcsRecord         `json:"Data"`
	UpgradeablePlans   []UpgradeablePlan `json:"UpgradeablePlans"`
	RBSList            []RBSItem         `json:"RBSList"`
	NatList            []any             `json:"NatList"`
	EDiskList          []EDiskItem       `json:"EDiskList"`
	EIPList            []EIPItem         `json:"EIPList"`
	RenewPointPrice    RenewPointPrice   `json:"RenewPointPrice"`
	FastInstallAppTask []any             `json:"FastInstallAppTask"`
	VNets              []any             `json:"VNets"`
}

type UpgradeablePlan struct {
	ID              int       `json:"id"`
	Region          string    `json:"region"`
	Subtype         string    `json:"subtype"`
	PlanName        string    `json:"plan_name"`
	Machine         string    `json:"machine"`
	ChargeType      string    `json:"charge_type"`
	Chinese         string    `json:"chinese"`
	IsFree          bool      `json:"is_free"`
	PointRenewPrice any       `json:"point_renew_price"`
	IsSelling       bool      `json:"is_selling"`
	Price           int       `json:"price"`
	TrafficBaseGb   int       `json:"traffic_base_gb"`
	TrafficPrice    any       `json:"traffic_price"`
	CPU             int       `json:"cpu"`
	Memory          int       `json:"memory"`
	NetIn           int       `json:"net_in"`
	NetOut          int       `json:"net_out"`
	IPPrices        any       `json:"ip_prices"`
	IPSelling       any       `json:"ip_selling"`
	AutoRestock     int       `json:"auto_restock"`
	AvailableStock  int       `json:"available_stock"`
	GpuMemorySize   int       `json:"gpu_memory_size"`
	DgpuDevType     string    `json:"dgpu_dev_type"`
	WebbarConfig    any       `json:"webbar_config"`
	NoBackup        bool      `json:"no_backup"`
	DiskPrice       DiskPrice `json:"disk_price"`
}

type RBSItem struct {
	ID             int               `json:"ID"`
	UID            int               `json:"UID"`
	ProductID      int               `json:"ProductID"`
	NodeUUID       string            `json:"NodeUUID"`
	Node           Node              `json:"Node"`
	Label          string            `json:"Label"`
	FileName       string            `json:"FileName"`
	PackSize       int64             `json:"PackSize"`
	CreateTime     int               `json:"CreateTime"`
	FinishTime     int               `json:"FinishTime"`
	Retry          int               `json:"Retry"`
	AdditionalInfo RBSAdditionalInfo `json:"AdditionalInfo"`
	Status         string            `json:"Status"`
}

type RBSAdditionalInfo struct {
	Osid           int                `json:"OSID"`
	OSName         string             `json:"OSName"`
	OSBaseDiskSize int                `json:"OSBaseDiskSize"`
	OSDataDiskSize int                `json:"OSDataDiskSize"`
	Slots          map[string]RBSSlot `json:"Slots"`
}

type RBSSlot struct {
	DiskType string `json:"DiskType"`
	Backup   bool   `json:"Backup"`
	Size     int    `json:"Size"`
}

type EDiskItem struct {
	ID       int    `json:"ID"`
	Slot     int    `json:"Slot"`
	UID      int    `json:"UID"`
	DiskType string `json:"DiskType"`
	Tag      string `json:"Tag"`
	OSName   string `json:"OSName"`
	Vid      int    `json:"VID"`
	Size     int    `json:"Size"`
	Backup   bool   `json:"Backup"`
}

type EIPItem struct {
	ID          int    `json:"ID"`
	IPRegion    string `json:"IpRegion"`
	Region      string `json:"Region"`
	Type        string `json:"Type"`
	IP          string `json:"IP"`
	Gateway     string `json:"Gateway"`
	Block       string `json:"Block"`
	UID         int    `json:"UID"`
	Vid         int    `json:"VID"`
	CreateDate  int    `json:"CreateDate"`
	Flags       string `json:"Flags"`
	VlanID      int    `json:"VlanID"`
	Description string `json:"Description"`
}

type RenewPointPrice struct {
	Num7  int `json:"7"`
	Num31 int `json:"31"`
}

func (s *RcsService) GetRcsDetail(id int) (*GetRcsDetailResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d", id)

	var resp GetRcsDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
