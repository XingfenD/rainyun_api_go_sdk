package public

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

// https://api.rainyun.com/#/paths/product~1rgs~1egg_type/get

// 获取蛋(游戏类型)类型列表
type GetEggTypeListResponse struct {
	Code int           `json:"code"`
	Data []EggTypeItem `json:"data"`
}

type EggTypeItem struct {
	ID         int        `json:"id"`       // 蛋(游戏类型)ID
	EggName    string     `json:"egg_name"` // 蛋(游戏类型)名称
	Egg        EggItem    `json:"egg"`
	Docker     string     `json:"docker"` // docker镜像地址
	McsmDocker string     `json:"mcsm_docker"`
	Env        EggTypeEnv `json:"env"`
	Order      int        `json:"order"`
	IsHidden   bool       `json:"is_hidden"`
	UpdateTime int        `json:"update_time"`
	AutoUpdate bool       `json:"auto_update"`
}

type EggTypeEnv struct {
	ServerType    string `json:"SERVER_TYPE"`
	ServerVersion string `json:"SERVER_VERSION"`
	Subver        string `json:"SUBVER"`
	Latest        bool   `json:"LATEST"`
	Prerelease    bool   `json:"PRERELEASE"`
}

func (s *PublicService) GetEggTypeList() (*GetEggTypeListResponse, error) {
	path := "/product/rgs/egg_type"

	var resp GetEggTypeListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)

	return &resp, err
}
