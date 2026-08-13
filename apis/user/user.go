package user

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type GetUserInfoResponse struct {
	Code int      `json:"code"`
	Data UserInfo `json:"data"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID                  int         `json:"ID"`
	Name                string      `json:"Name"`
	Email               string      `json:"Email"`
	Phone               string      `json:"Phone"`
	Money               float64     `json:"Money"`
	RegisterTime        int         `json:"RegisterTime"`
	QQOpenID            string      `json:"QQOpenID"`
	QQ                  any         `json:"QQ"`
	WechatOpenID        string      `json:"WechatOpenID"`
	CardCode            any         `json:"CardCode"`
	IconURL             string      `json:"IconUrl"`
	Points              int         `json:"Points"`
	MoneyFromPoint      int         `json:"MoneyFromPoint"`
	Inviter             int         `json:"Inviter"`
	APIKey              string      `json:"APIKey"`
	LastIP              string      `json:"LastIP"`
	BanReason           string      `json:"BanReason"`
	UnsubscribeCount    int         `json:"UnsubscribeCount"`
	AlipayAccount       string      `json:"AlipayAccount"`
	AlipayName          string      `json:"AlipayName"`
	LastLogin           string      `json:"LastLogin"`
	LastLoginArea       string      `json:"LastLoginArea"`
	LoginCount          int         `json:"LoginCount"`
	DLWallet            int         `json:"DLWallet"`
	DLLevel             int         `json:"DLLevel"`
	AdminGroup          string      `json:"AdminGroup"`
	TOTPSecret          string      `json:"TOTPSecret"`
	IsLoginEnableTFA    bool        `json:"IsLoginEnableTFA"`
	IsAllowPointUse     int         `json:"IsAllowPointUse"`
	ShareCode           string      `json:"ShareCode"`
	VipLevel            int         `json:"VipLevel"`
	IsAgent             bool        `json:"IsAgent"`
	Valid               bool        `json:"Valid"`
	Source              string      `json:"Source"`
	ConsumeMonthly      int         `json:"ConsumeMonthly"`
	ConsumeAll          int         `json:"ConsumeAll"`
	ConsumeQuarter      int         `json:"ConsumeQuarter"`
	ResellDaily         int         `json:"ResellDaily"`
	ResellMonthly       int         `json:"ResellMonthly"`
	ResellBeforeMonth   float64     `json:"ResellBeforeMonth"`
	ResellQuarter       float64     `json:"ResellQuarter"`
	ResellAll           float64     `json:"ResellAll"`
	StockDaily          int         `json:"StockDaily"`
	StockMonthly        int         `json:"StockMonthly"`
	StockQuarter        float64     `json:"StockQuarter"`
	StockAll            float64     `json:"StockAll"`
	SecondStockQuarter  int         `json:"SecondStockQuarter"`
	SecondStockAll      int         `json:"SecondStockAll"`
	SubUserMonthly      int         `json:"SubUserMonthly"`
	SubUserAll          int         `json:"SubUserAll"`
	ResellPointsMonthly int         `json:"ResellPointsMonthly"`
	ResellPointsAll     int         `json:"ResellPointsAll"`
	Vip                 UserVipInfo `json:"VIP"`
	Certify             int         `json:"Certify"`
	LockPoints          int         `json:"LockPoints"`
	CertifyStatus       string      `json:"CertifyStatus"`
	CertifyType         string      `json:"CertifyType"`
	CertifyAuditNote    string      `json:"CertifyAuditNote"`
}

// UserVipInfo 用户 VIP 等级信息
type UserVipInfo struct {
	Title              string  `json:"Title"`
	SaleRequire        int     `json:"SaleRequire"`
	ResellRequire      int     `json:"ResellRequire"`
	CertifyRequired    bool    `json:"CertifyRequired"`
	SaleProfit         float64 `json:"SaleProfit"`
	ResellProfit       float64 `json:"ResellProfit"`
	SecondResellProfit float64 `json:"SecondResellProfit"`
	CanSendCoupons     bool    `json:"CanSendCoupons"`
	CanCustomCode      bool    `json:"CanCustomCode"`
	CanSendMsg         bool    `json:"CanSendMsg"`
	CanTryUsual        bool    `json:"CanTryUsual"`
	FreeDomainCount    int     `json:"FreeDomainCount"`
	FreeSSLCount       int     `json:"FreeSSLCount"`
	CanBeAgent         bool    `json:"CanBeAgent"`
	AgentTitle         string  `json:"AgentTitle"`
	StockRequire       int     `json:"StockRequire"`
	SecondStockRequire int     `json:"SecondStockRequire"`
	StockDiscount      float64 `json:"StockDiscount"`
	SecondStockProfit  float64 `json:"SecondStockProfit"`
}

// 获取用户信息.
func (s *UserService) GetUserInfo() (*GetUserInfoResponse, error) {
	path := "/user/"

	var resp GetUserInfoResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)

	return &resp, err
}

// RefreshApikeyRequest 刷新 API 密钥请求
type RefreshApikeyRequest struct {
	Option string `json:"option"`
}

// 刷新 api 密钥.
//
// ⚠️不要轻易使用此方法
func (s *UserService) RefreshApikey() (*common.BasicOperationResponse, error) {
	path := "/user/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, RefreshApikeyRequest{Option: "apikey"}, &resp)

	return &resp, err
}
