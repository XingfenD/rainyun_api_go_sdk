package public

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

// https://api.rainyun.com/#/paths/product~1rgs~1os-templates/get

// 获取游戏云系统列表
type GetRgsOSListResponse struct {
	Code int             `json:"code"`
	Data []RgsOSListItem `json:"data"`
}

type RgsOSListItem struct {
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

func (s *PublicService) GetRgsOSList() (*GetRgsOSListResponse, error) {
	path := "/product/rgs/os-templates"

	var resp GetRgsOSListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)

	return &resp, err
}
