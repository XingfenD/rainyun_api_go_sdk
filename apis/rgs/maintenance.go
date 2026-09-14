package rgs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRgsMaintenanceResponse struct {
	Code int `json:"code"`
	// TODO: 结构未公开,实测后补强类型
	Data any `json:"data"`
}

// GetRgsMaintenance 查询游戏云维护状态
func (s *RgsService) GetRgsMaintenance(id int) (*GetRgsMaintenanceResponse, error) {
	path := fmt.Sprintf("/product/rgs/%d/maintenance", id)

	var resp GetRgsMaintenanceResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
