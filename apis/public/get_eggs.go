package public

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

// https://api.rainyun.com/#/paths/product~1rgs~1egg/get

// 获取蛋(游戏)列表
type GetEggListResponse struct {
	Code int       `json:"code"`
	Data []EggItem `json:"data"`
}

type EggItem struct {
	Name     string   `json:"name"`      // 名称
	EggGroup string   `json:"egg_group"` // 分组
	Title    string   `json:"title"`     // 标题
	Desc     string   `json:"desc"`      // 描述
	IconURL  string   `json:"icon_url"`  // 图标
	Order    int      `json:"order"`     // 排序
	IsHidden bool     `json:"is_hidden"` // 是否隐藏
	SaveDirs []string `json:"save_dirs"` // 重装时建议的需要保留的目录
}

func (s *PublicService) GetEggList() (*GetEggListResponse, error) {
	path := "/product/rgs/egg"

	var resp GetEggListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)

	return &resp, err
}
