package ssl

import "github.com/XingfenD/rainyun_api_go_sdk/apis/common"

// SSL证书
type SslCertificate struct {
	Cert string `json:"cert"` // 证书
	Key  string `json:"key"`  // 私钥
}

// SSL证书列表记录
type SslCertificateRecord struct {
	ID            int    `json:"ID"`
	UID           int    `json:"UID"`
	Domain        string `json:"Domain"`        // 域名(逗号分割)
	Issuer        string `json:"Issuer"`        // 品牌
	StartDate     int    `json:"StartDate"`     // 开始时间
	ExpDate       int    `json:"ExpDate"`       // 结束时间
	UploadTime    int    `json:"UploadTime"`    // 上传时间
	NginxErr      string `json:"NginxErr"`      // ？
	BaishanCertID int    `json:"BaishanCertID"` // 白山云证书ID
	BindDomains   any    `json:"BindDomains"`   // 绑定的域名 TODO: 结构未公开,实测后补强类型
}

// SSL证书列表数据
type SslCertificateListData struct {
	TotalRecords int                    `json:"TotalRecords"`
	Records      []SslCertificateRecord `json:"Records"`
}

type GetSslCertificateListResponse struct {
	Code int                    `json:"code"`
	Data SslCertificateListData `json:"data"`
}

type GetSslCertificateListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// SSL证书详情数据
type SslDetailData struct {
	Cert       string `json:"Cert"`       // 证书
	Key        string `json:"Key"`        // 私钥
	DomainName string `json:"DomainName"` // 域名(逗号分割)
	Issuer     string `json:"Issuer"`     // 品牌
	StartDate  int    `json:"StartDate"`  // 开始时间
	ExpDate    int    `json:"ExpDate"`    // 结束时间
	RemainDays int    `json:"RemainDays"` // 剩余天数
}

type GetSslDetailResponse struct {
	Code int           `json:"code"`
	Data SslDetailData `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type SslPassthroughResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// SSL证书产品(实测 /product/sslcenter/product 响应)
type SslProduct struct {
	ID                int            `json:"id"`
	Name              string         `json:"name"`
	VerifyMethods     []string       `json:"verifyMethods"`
	Algorithms        []string       `json:"algorithms"`
	PriceMap          map[string]int `json:"priceMap"`         // 时长(月)->现价
	OriginalPriceMap  map[string]int `json:"originalPriceMap"` // 时长(月)->原价
	IsOnSale          bool           `json:"isOnSale"`
	Provider          string         `json:"provider"`
	ProviderProductID string         `json:"providerProductId"`
	ProviderCaa       string         `json:"providerCaa"`
	MultiDomain       bool           `json:"multiDomain"`
	Wildcard          bool           `json:"wildcard"`
	Type              string         `json:"type"` // dv/ov
	Brand             string         `json:"brand"`
}

type GetSSLProductListResponse struct {
	Code int          `json:"code"`
	Data []SslProduct `json:"data"`
}

// SSL订单(实测 /product/sslcenter/order 响应)
type SslOrder struct {
	ID                       int                   `json:"id"`
	UID                      int                   `json:"uid"`
	CsrInfo                  SslOrderCsrInfo       `json:"csrInfo"`
	Product                  SslProduct            `json:"product"`
	Duration                 int                   `json:"duration"`
	Price                    float64               `json:"price"`
	Status                   string                `json:"status"`
	ValidationRecords        []SslValidationRecord `json:"validationRecords"`
	Description              any                   `json:"description"` // TODO: 结构未公开,实测后补强类型
	RewardToBeCollect        int                   `json:"rewardToBeCollect"`
	UnsubscribeAble          bool                  `json:"unsubscribeAble"`
	UnsubscribeTime          int                   `json:"unsubscribeTime"`
	UnsubscribeReason        any                   `json:"unsubscribeReason"`        // TODO: 结构未公开,实测后补强类型
	NotUnsubscribeAbleReason any                   `json:"notUnsubscribeAbleReason"` // TODO: 结构未公开,实测后补强类型
	CertIssuedAt             int64                 `json:"certIssuedAt"`
	CertStartAt              int64                 `json:"certStartAt"`
	CertExpireAt             int64                 `json:"certExpireAt"`
	CreatedAt                int64                 `json:"createdAt"`
	UpdatedAt                int64                 `json:"updatedAt"`
}

type SslOrderCsrInfo struct {
	CommonName string   `json:"commonName"`
	DNSNames   []string `json:"dnsNames"`
	KeyAlgo    string   `json:"keyAlgo"`
	KeyLen     int      `json:"keyLen"`
	SignHash   string   `json:"signHash"`
	Country    string   `json:"country"`
}

type SslValidationRecord struct {
	ID      int    `json:"ID"`
	Domain  string `json:"domain"`
	Host    string `json:"host"`
	Value   string `json:"value"`
	Type    string `json:"type"`
	Purpose string `json:"purpose"`
}

type SslOrderList struct {
	TotalRecords int        `json:"totalRecords"`
	Records      []SslOrder `json:"records"`
}

type GetSSLOrderListResponse struct {
	Code int          `json:"code"`
	Data SslOrderList `json:"data"`
}

type GetSSLOrderDetailResponse struct {
	Code int      `json:"code"`
	Data SslOrder `json:"data"`
}

// 订单价格(实测 /product/sslcenter/price 响应)
type SslOrderPrice struct {
	Price        float64 `json:"price"`
	Reward       float64 `json:"reward"`
	RewardPoints int64   `json:"rewardPoints"`
}

type GetSSLOrderPriceResponse struct {
	Code int           `json:"code"`
	Data SslOrderPrice `json:"data"`
}

type ApplyFreeSSLCertRequest struct {
	Domains      string `json:"domains"`       // 域名列表
	VerifyMethod string `json:"verify_method"` // 验证方式:dns/http/auto
}

type VerifyFreeSSLCertRequest struct {
	OrderID int `json:"order_id"`
}

type GetSSLOrderListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

type CreateSSLOrderRequest struct {
	Domains      string  `json:"domains"`      // 域名
	Duration     int     `json:"duration"`     // 购买时长(单位:月)
	Price        float64 `json:"price"`        // 价格,仅用于核验
	ProductID    int     `json:"productId"`    // 欲购买产品ID
	WithCouponID int     `json:"withCouponId"` // 优惠券ID
}

type UpdateSSLOrderDescriptionRequest struct {
	NewDescription string `json:"newDescription"`
}

type RevokeSSLOrderRequest struct {
	Letter string `json:"letter"` // 吊销函内容(Base64编码,非DV必传)
	Reason string `json:"reason"` // 申请吊销原因
}

type VerifySSLOrderRequest struct {
	ForceRefresh bool `json:"forceRefresh"` // 强制刷新证书
}
