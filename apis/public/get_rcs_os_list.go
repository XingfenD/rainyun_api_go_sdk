package public

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

// https://api.rainyun.com/#/paths/product~1rcs~1os-templates/get

// 获取RCS操作系统列表
type GetRcsOSListResponse struct {
	Code int             `json:"code"`
	Data []RcsOSListItem `json:"data"`
}

type RcsOSListItem struct {
	ID             int    `json:"id"`      // 系统ID
	Region         string `json:"region"`  // 地域
	Subtype        string `json:"subtype"` // 类型(kvm)
	Machine        string `json:"machine"` // unknown
	Name           string `json:"name"`    // 英文名
	Version        string `json:"version"` // 版本
	SyncStatus     string `json:"sync_status"`
	OsType         string `json:"os_type"`         // 系统类型(windows/linux)
	ChineseName    string `json:"chinese_name"`    // 中文名
	Icon           string `json:"icon"`            // 图标
	IsWithBbr      bool   `json:"is_with_bbr"`     // 是否支持BBR
	IsEol          bool   `json:"is_eol"`          // 是否已过时
	IsAvailable    bool   `json:"is_available"`    // 是否可用
	Order          int    `json:"order"`           // 排序
	LatestFilename string `json:"latest_filename"` // 最新文件名
	NoVMAgent      bool   `json:"no_vm_agent"`     // 是否无虚拟机Agent
}

func (s *PublicService) GetRcsOSList() (*GetRcsOSListResponse, error) {
	path := "/product/rcs/os-templates"

	var resp GetRcsOSListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)

	return &resp, err
}
