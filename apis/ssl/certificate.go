package ssl

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

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
