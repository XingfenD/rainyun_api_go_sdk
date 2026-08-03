package public

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
	"github.com/bytedance/sonic"
)

// https://api.rainyun.com/#/paths/status/get

type GetStatusRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type GetStatusResponse struct {
	Code int        `json:"code"`
	Data NodeStatus `json:"data"`
}

type NodeStatus struct {
	TotalRecords         int                    `json:"TotalRecords"`
	OngoingServiceEvents []interface{}          `json:"OngoingServiceEvents"`
	Records              []NodeStatusRecordItem `json:"Records"`
}

type NodeStatusRecordItem struct {
	UUID        string  `json:"UUID"`
	ChineseName string  `json:"ChineseName"`
	Product     string  `json:"Product"`
	CPU         float64 `json:"CPU"`
	Memory      float64 `json:"Memory"`
	NetOut      int     `json:"NetOut"`
	UpdateTime  string  `json:"UpdateTime"`
	Status      string  `json:"Status"`
	Data        string  `json:"Data"`
}

func (req *GetStatusRequest) BuildQueryMap() (map[string]string, error) {
	rst := make(map[string]string, 1)
	optionString, err := sonic.Marshal(req.Options)
	if err != nil {
		return nil, err
	}

	rst["options"] = string(optionString)

	return rst, nil
}

func (s *PublicService) GetStatus(req *GetStatusRequest) (*GetStatusResponse, error) {
	path := "/status"

	var resp GetStatusResponse

	querys, err := req.BuildQueryMap()
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
