package rcs

import (
	"fmt"
	"net/url"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// GetRcsVnc 连接VNC
//
// consoleType: 控制台类型,可选值: novnc,xtermjs
func (s *RcsService) GetRcsVnc(id int, consoleType string) (*common.VncConnectionInfo, error) {
	path := fmt.Sprintf("/product/rcs/%d/vnc", id)

	var resp common.VncConnectionInfo
	querys := map[string]string{"console_type": consoleType}
	err := s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// GetRcsPveAddress 获取RCS对应的PVE地址
func (s *RcsService) GetRcsPveAddress(id int) (string, error) {
	v, err := s.GetRcsVnc(id, "novnc")
	if err != nil {
		return "", err
	}

	parsedURL, err := url.Parse(v.Data.RequestURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), nil
}
