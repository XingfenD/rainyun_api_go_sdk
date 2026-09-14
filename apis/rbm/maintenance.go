package rbm

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetRbmMaintenanceResponse struct {
	Code int `json:"code"`
	// TODO: 结构未公开,实测后补强类型
	Data any `json:"data"`
}

// GetRbmMaintenance 查询裸金属维护状态
func (s *RbmService) GetRbmMaintenance(id int) (*GetRbmMaintenanceResponse, error) {
	path := fmt.Sprintf("/product/rbm/%d/maintenance", id)

	var resp GetRbmMaintenanceResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
