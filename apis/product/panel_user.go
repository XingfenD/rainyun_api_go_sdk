package product

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 增减面板用户产品请求
type ModifyPanelUserProductRequests struct {
	Action      string `json:"action"`       // 操作：add/del
	ProductID   int    `json:"product_id"`   // 产品ID
	ProductType string `json:"product_type"` // 产品类型: rcs/rgs/ros/rbm 只支持这四种产品
	Name        string `json:"name"`         // 子用户名
}

// 面板用户请求
type PanelUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 面板用户产品
type PanelUserProduct struct {
	Name        string `json:"name"`         // 自用户名
	ProductType string `json:"product_type"` // 产品类型: rcs/rgs/ros/rbm 只支持这四种产品
	ProductID   int    `json:"product_id"`   // 产品ID
}

// 面板用户记录
type PanelUserRecord struct {
	Name       string             `json:"name"`        // 面板用户名
	Password   string             `json:"password"`    // 面板用户密码
	UserID     int                `json:"user_id"`     // 你的UID
	CreateDate int                `json:"create_date"` // 创建时间
	Products   []PanelUserProduct `json:"products"`    // 产品列表
}

// 面板用户列表
type PanelUserListData struct {
	TotalRecords int               `json:"TotalRecords"` // 面板用户总数
	Records      []PanelUserRecord `json:"Records"`
}

type GetPanelUserListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type GetPanelUserListResponse struct {
	Code int               `json:"code"`
	Data PanelUserListData `json:"data"`
}

// 获取面板用户列表
func (s *ProductService) GetPanelUserList(req *GetPanelUserListRequest) (*GetPanelUserListResponse, error) {
	path := "/product/panel_users/"

	var resp GetPanelUserListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 增减面板用户产品
func (s *ProductService) ModifyPanelUserProduct(req *ModifyPanelUserProductRequests) (*common.BasicOperationResponse, error) {
	path := "/product/panel_users/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PUT, path, nil, req, &resp)
	return &resp, err
}

// 创建面板用户
//
// name： 子用户用户名
//
// pass： 子用户密码
func (s *ProductService) CreatePanelUser(name string, pass string) (*common.BasicOperationResponse, error) {
	path := "/product/panel_users/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, PanelUserRequest{Name: name, Password: pass}, &resp)
	return &resp, err
}

// 面板用户改密
//
// name： 子用户用户名
//
// pass： 子用户新密码
func (s *ProductService) ChangePanelUserPassword(name string, pass string) (*common.BasicOperationResponse, error) {
	path := "/product/panel_users/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, PanelUserRequest{Name: name, Password: pass}, &resp)
	return &resp, err
}

// 删除面板用户
//
// name： 子用户用户名
func (s *ProductService) DeletePanelUser(name string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/panel_users/%s", name)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}
