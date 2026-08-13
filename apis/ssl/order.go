package ssl

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 获取SSL证书订单列表
func (s *SslService) GetSSLOrderList(req *GetSSLOrderListRequest) (*GetSSLOrderListResponse, error) {
	path := "/product/sslcenter/order"

	var resp GetSSLOrderListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// TODO: 响应结构未公开,透传;实测后补强类型
//
// 创建SSL证书订单
func (s *SslService) CreateSSLOrder(req *CreateSSLOrderRequest) (*SslPassthroughResponse, error) {
	path := "/product/sslcenter/order"

	var resp SslPassthroughResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 获取SSL证书订单信息
//
// id: 订单ID
func (s *SslService) GetSSLOrderDetail(id int) (*GetSSLOrderDetailResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/order/%d", id)

	var resp GetSSLOrderDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// TODO: 响应结构未公开,透传;实测后补强类型
//
// 将SSL证书添加到证书列表
//
// id: 订单ID
func (s *SslService) AssignSSLOrder(id int) (*SslPassthroughResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/order/%d/assign", id)

	var resp SslPassthroughResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 获取SSL证书
//
// id: 订单ID
func (s *SslService) GetSSLOrderCert(id int) (*GetSslDetailResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/order/%d/cert", id)

	var resp GetSslDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// TODO: 响应结构未公开,透传;实测后补强类型
//
// 更新SSL订单描述
//
// id: 订单ID
//
// newDescription: 新描述
func (s *SslService) UpdateSSLOrderDescription(id int, newDescription string) (*SslPassthroughResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/order/%d/description", id)

	var resp SslPassthroughResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, UpdateSSLOrderDescriptionRequest{NewDescription: newDescription}, &resp)
	return &resp, err
}

// TODO: 响应结构未公开,透传;实测后补强类型
//
// 申请吊销SSL证书
//
// id: 订单ID
//
// reason: 申请吊销原因
//
// letter: 吊销函内容(Base64编码,非DV必传)
func (s *SslService) RevokeSSLOrder(id int, reason, letter string) (*SslPassthroughResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/order/%d/revoke", id)

	var resp SslPassthroughResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, RevokeSSLOrderRequest{Letter: letter, Reason: reason}, &resp)
	return &resp, err
}

// TODO: 响应结构未公开,透传;实测后补强类型
//
// 验证SSL证书订单
//
// id: 订单ID
//
// forceRefresh: 强制刷新证书
func (s *SslService) VerifySSLOrder(id int, forceRefresh bool) (*SslPassthroughResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/order/%d/verify", id)

	var resp SslPassthroughResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, VerifySSLOrderRequest{ForceRefresh: forceRefresh}, &resp)
	return &resp, err
}

// 获取SSL证书订单价格
func (s *SslService) GetSSLOrderPrice(req *CreateSSLOrderRequest) (*GetSSLOrderPriceResponse, error) {
	path := "/product/sslcenter/price"

	var resp GetSSLOrderPriceResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 获取SSL证书产品列表
func (s *SslService) GetSSLProductList() (*GetSSLProductListResponse, error) {
	path := "/product/sslcenter/product"

	var resp GetSSLProductListResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
