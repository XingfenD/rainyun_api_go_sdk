package workorder

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 工单打分
//
// id: 工单ID
func (s *WorkorderService) ScoreWorkorder(id int, req *ScoreWorkerorderRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/score", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 编辑工单打分
//
// id: 工单ID
func (s *WorkorderService) EditScoreWorkorder(id int, req *ScoreWorkerorderRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/workorder/%d/score", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// 获取工单打分
//
// id: 工单ID
//
// discussID: 客服回复ID
func (s *WorkorderService) GetScoreWorkorder(id int, discussID int) (*GetScoreWorkorderResponse, error) {
	path := fmt.Sprintf("/workorder/%d/score/%d", id, discussID)

	var resp GetScoreWorkorderResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
