package rgs

import (
	"fmt"
	"net/url"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 游戏云连接VNC
//
// id: 游戏云ID
//
// consoleType: 控制台类型,可选值: novnc,xtermjs
func (s *RgsService) GetRgsVnc(id int, consoleType string) (*common.VncConnectionInfo, error) {
	path := fmt.Sprintf("/product/rgs/%d/vnc", id)

	var resp common.VncConnectionInfo
	querys := map[string]string{"console_type": consoleType}
	err := s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 获取游戏云对应的PVE地址
//
// id: 游戏云ID
func (s *RgsService) GetRgsPveAddress(id int) (string, error) {
	v, err := s.GetRgsVnc(id, "novnc")
	if err != nil {
		return "", err
	}

	parsedURL, err := url.Parse(v.Data.RequestURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), nil
}
