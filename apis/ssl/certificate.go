package ssl

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

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
	BindDomains   any    `json:"BindDomains"`   // 绑定的域名
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

// 获取SSL证书列表
func (s *SslService) GetSSLCertificateList(req *GetSslCertificateListRequest) (*GetSslCertificateListResponse, error) {
	path := "/product/sslcenter/"

	var resp GetSslCertificateListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 上传SSL证书
//
// cert: 证书
//
// key： 私钥
func (s *SslService) UploadSSLCertificate(cert string, key string) (*common.BasicOperationResponse, error) {
	path := "/product/sslcenter/"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, SslCertificate{
		Cert: cert,
		Key:  key,
	}, &resp)
	return &resp, err
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

// 获取SSL证书详情
//
// id: SSL证书ID
func (s *SslService) GetSslDetail(id int) (*GetSslDetailResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/%d", id)

	var resp GetSslDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 替换SSL证书
//
// id: SSL证书ID
//
// cert: 证书
//
// key： 私钥
func (s *SslService) ReplaceSsl(id int, cert, key string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/%d", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PUT, path, nil, SslCertificate{
		Cert: cert,
		Key:  key,
	}, &resp)
	return &resp, err
}

// 删除SSL证书
//
// id: SSL证书ID
func (s *SslService) DeleteSsl(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/%d", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}
