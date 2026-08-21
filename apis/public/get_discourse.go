package public

import (
	"github.com/bytedance/sonic"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// https://api.rainyun.com/#/paths/discourse/get

// GetDiscourseRequest 获取论坛数据请求参数
type GetDiscourseRequest struct {
	Paths string `json:"paths"` // 论坛路径
}

// GetDiscourseResponse 获取论坛数据响应，Data 为论坛返回的原始数据
type GetDiscourseResponse struct {
	Code int                    `json:"code"`
	Data sonic.NoCopyRawMessage `json:"data"`
}

// GetDiscourse 获取论坛数据
func (s *PublicService) GetDiscourse(req *GetDiscourseRequest) (*GetDiscourseResponse, error) {
	path := "/discourse"

	var resp GetDiscourseResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
