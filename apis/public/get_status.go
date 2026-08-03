package public

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

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

func (s *PublicService) GetStatus() (*GetStatusResponse, error) {
	path := "/status"

	var resp GetStatusResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, &resp)
	return &resp, err
}
