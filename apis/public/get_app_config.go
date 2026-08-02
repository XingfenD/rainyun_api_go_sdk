package public

import "github.com/XingfenD/rainyun_api_go_sdk/constant"

type AppConfigItemResponse struct {
	Code     int             `json:"code"`
	Data     []AppConfigItem `json:"data"`
	Messages []string        `json:"messages,omitempty"`
	Msg      string          `json:"msg,omitempty"`
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

func (s *Service) GetAppConfig() (*AppConfigItemResponse, error) {
	path := "/app_config"
	var resp AppConfigItemResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, &resp)
	return &resp, err
}
