package common

import "strconv"

type VncConnectionInfo struct {
	Code int               `json:"code"`
	Data VncConnectionData `json:"data"`
}

type VncConnectionData struct {
	RequestURL  string `json:"RequestURL"`  // 空
	RedirectURL string `json:"RedirectURL"` // 空
	PVEAuth     string `json:"PVEAuth"`     // 空
	VNCProxyURL string `json:"VNCProxyURL"` // PVE代理地址(连接地址)，里面会有一个unicode码(\u0026)(&),转码后就是可以直接在浏览器上使用的VNC连接地址
}

func GetVncConnectURL(v *VncConnectionInfo) (string, error) {
	return strconv.Unquote(v.Data.VNCProxyURL)
}
