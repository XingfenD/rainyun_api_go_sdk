package public

import (
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetNewsResponse struct {
	Code int               `json:"code"`
	Data []RainyunNewsItem `json:"data"`
}

type RainyunNewsItem struct {
	Type      string    `json:"Type"`
	Title     string    `json:"Title"`
	TimeStamp time.Time `json:"TimeStamp"`
	URL       string    `json:"URL"`
}

func (s *PublicService) GetNews() (*GetNewsResponse, error) {
	path := "/news"

	var resp GetNewsResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, &resp)

	return &resp, err
}
