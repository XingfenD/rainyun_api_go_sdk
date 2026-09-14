package rcs

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRcsMaintenanceResponse struct {
	Code int `json:"code"`
	// TODO: 结构未公开,实测后补强类型
	Data any `json:"data"`
}

// GetRcsMaintenance 查询云服务器维护状态
func (s *RcsService) GetRcsMaintenance(id int) (*GetRcsMaintenanceResponse, error) {
	path := fmt.Sprintf("/product/rcs/%d/maintenance", id)

	var resp GetRcsMaintenanceResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
