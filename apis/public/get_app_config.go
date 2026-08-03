package public

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

// https://api.rainyun.com/#/paths/app_config/get

type GetAppConfigResponse struct {
	Code int             `json:"code"`
	Data []AppConfigItem `json:"data"`
}

type AppConfigItem struct {
	AgentID int              `json:"agent_id"`
	Type    string           `json:"type"`
	Value   []AppConfigValue `json:"value"`
}

type AppConfigValue struct {
	SenderID      int    `json:"sender_id"`
	Content       string `json:"content"`
	Title         string `json:"title"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	Page          string `json:"page"`
	VgtID         int    `json:"vgt_id"`
	OriginalIndex int    `json:"originalIndex"`
}

func (s *PublicService) GetAppConfig() (*GetAppConfigResponse, error) {
	path := "/app_config"
	var resp GetAppConfigResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
