package ssl

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// TODO: 响应结构未公开,透传;实测后补强类型
//
// 创建SSL证书申请
//
// req: domains 域名列表,verify_method 验证方式(auto代表从雨云注册的域名可以自动完成)
func (s *SslService) ApplyFreeSSLCertificate(req *ApplyFreeSSLCertRequest) (*SslPassthroughResponse, error) {
	path := "/product/sslcenter/cert/order"

	var resp SslPassthroughResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// TODO: 响应结构未公开,透传;实测后补强类型
//
// 验证SSL证书申请
//
// order_id: 订单ID
func (s *SslService) VerifyFreeSSLCertificate(orderID int) (*SslPassthroughResponse, error) {
	path := "/product/sslcenter/cert/order_verify"

	var resp SslPassthroughResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, VerifyFreeSSLCertRequest{OrderID: orderID}, &resp)
	return &resp, err
}

// TODO: 响应结构未公开,透传;实测后补强类型
//
// 获取SSL证书申请列表
func (s *SslService) GetSSLCertApplyList(req *GetSSLOrderListRequest) (*SslPassthroughResponse, error) {
	path := "/product/sslcenter/cert/orders"

	var resp SslPassthroughResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
