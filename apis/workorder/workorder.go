package workorder

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 获取工单列表
//
// req.Options 为分页查询参数, 可用 common.StandQueryParameters 构造
func (s *WorkorderService) GetWorkOrderList(req *GetWorkOrderListRequest) (*GetWorkOrderListResponse, error) {
	path := "/workorder/"

	var resp GetWorkOrderListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取工单详情
//
// id: 工单ID
func (s *WorkorderService) GetWorkOrderDetail(id int) (*GetWorkOrderDetailResponse, error) {
	path := fmt.Sprintf("/workorder/%d", id)

	var resp GetWorkOrderDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 创建工单
func (s *WorkorderService) CreateWorkOrder(req *CreateWorkerorderRequest) (*CreateWorkerorderResponse, error) {
	path := "/workorder/"

	var resp CreateWorkerorderResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 产品授权
//
// id:工单ID
//
// productID:产品ID
//
// productType:产品类型(rvh/rcs/rgs/rbm/ros)
func (s *WorkorderService) ProductAuth(id int, productID int, productType string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/auth", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, ProductAuthRequest{ProductID: productID, ProductType: productType}, &resp)
	return &resp, err
}

// 回复工单
//
// id: 工单ID
//
// content: 回复内容
func (s *WorkorderService) ReplyWorkorder(id int, content string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/reply_order", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, ReplyWorkerorderRequest{Content: content}, &resp)
	return &resp, err
}

// 编辑回复工单
//
// id: 工单ID
//
// replyID: 回复ID
//
// content: 编辑后的内容
func (s *WorkorderService) EditWorkorderReply(id int, replyID int, content string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/reply_order/%d", id, replyID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, ReplyWorkerorderRequest{Content: content}, &resp)
	return &resp, err
}

// 获取工单状态
//
// id: 工单ID
func (s *WorkorderService) GetWorkorderStatus(id int) (*GetWorkorderStatusResponse, error) {
	path := fmt.Sprintf("/workorder/%d/status", id)

	var resp GetWorkorderStatusResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 设置工单状态
//
// id: 工单ID
//
// status: 目标状态
func (s *WorkorderService) SetWorkorderStatus(id int, status string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/status", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, SetWorkorderStatusRequest{Status: status}, &resp)
	return &resp, err
}
