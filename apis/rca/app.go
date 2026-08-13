package rca

import (
	"fmt"
	"strconv"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 云应用列出App
func (s *RcaService) GetRcaAppList(projectID int) (*GetRcaAppListResponse, error) {
	path := "/product/rca/app/"

	var resp GetRcaAppListResponse
	querys := map[string]string{"project_id": strconv.Itoa(projectID)}
	err := s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 安装云应用App
func (s *RcaService) InstallRcaApp(req *InstallRcaAppRequest) (*common.BasicOperationResponse, error) {
	path := "/product/rca/app/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 卸载云应用App
func (s *RcaService) UninstallRcaApp(id int, req *UninstallRcaAppRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, req, &resp)
	return &resp, err
}

// 获取云应用App详情(响应结构未公开,透传)
func (s *RcaService) GetRcaAppDetail(id int) (*GetRcaAppDetailResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/", id)

	var resp GetRcaAppDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 更新云应用App设定
func (s *RcaService) UpdateRcaApp(id int, req *UpdateRcaAppRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// 更新云应用App容器设定
func (s *RcaService) UpdateRcaAppContainer(id, containerID int, req *UpdateRcaAppContainerRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/", id, containerID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// 获取云应用App配置文件(透传)
func (s *RcaService) GetRcaAppConfigMap(id, containerID int) (*GetRcaAppConfigMapResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/config_map", id, containerID)

	var resp GetRcaAppConfigMapResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 获取App容器的指标信息(透传)
func (s *RcaService) GetRcaAppContainerMetrics(id, containerID int) (*GetRcaAppContainerMetricsResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/metrics", id, containerID)

	var resp GetRcaAppContainerMetricsResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 云应用获取PHP相关配置(透传)
func (s *RcaService) GetRcaAppPHPSetting(id, containerID int) (*GetRcaAppPHPSettingResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/php_setting", id, containerID)

	var resp GetRcaAppPHPSettingResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 云应用更新PHP相关配置
func (s *RcaService) UpdateRcaAppPHPSetting(id, containerID int, req *UpdateRcaAppPHPSettingRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/php_setting", id, containerID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 云应用列出服务(透传)
func (s *RcaService) GetRcaAppServiceList(id, containerID int) (*GetRcaAppServiceListResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/service/", id, containerID)

	var resp GetRcaAppServiceListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 创建云应用服务
func (s *RcaService) CreateRcaAppService(id, containerID int, req *CreateRcaAppServiceRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/service/", id, containerID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 删除云应用服务
func (s *RcaService) DeleteRcaAppService(id, containerID, serviceID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/service/%d", id, containerID, serviceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

// 更新云应用服务
func (s *RcaService) UpdateRcaAppService(id, containerID, serviceID int, req *CreateRcaAppServiceRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/service/%d", id, containerID, serviceID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// 云应用web服务器更新访问设定
func (s *RcaService) UpdateRcaAppWebserverAccess(id, containerID int, req *UpdateRcaAppWebserverAccessRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/%d/webserver_access", id, containerID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 云应用重启App
func (s *RcaService) RestartRcaApp(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/restart", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 云应用启动App
func (s *RcaService) StartRcaApp(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/start", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 云应用停止App
func (s *RcaService) StopRcaApp(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/stop", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 升级云应用App
func (s *RcaService) UpgradeRcaApp(id int, req *UpgradeRcaAppRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/app/%d/upgrade", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
