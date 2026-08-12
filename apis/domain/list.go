package domain

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 获取域名列表查询参数
type GetDomainListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// 域名列表响应
type GetDomainListResponse struct {
	Code int            `json:"code"`
	Data DomainListData `json:"data"`
}

// 域名列表数据
type DomainListData struct {
	TotalRecords int            `json:"TotalRecords"`
	Records      []DomainRecord `json:"Records"`
}

// 域名记录
type DomainRecord struct {
	ID         int    `json:"id"`
	DomainName string `json:"domain_name"`
	Status     string `json:"status"`
	ExpireDate string `json:"expire_date"`
}

// 获取域名列表
func (s *DomainService) GetDomainList(req *GetDomainListRequest) (*GetDomainListResponse, error) {
	path := "/product/domain/"

	var resp GetDomainListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}