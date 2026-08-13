package expense

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 订单列表查询参数
type GetOrdersListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// 订单列表响应
type GetOrdersListResponse struct {
	Code int            `json:"code"`
	Data OrdersListData `json:"data"`
}

// 订单列表数据
type OrdersListData struct {
	TotalRecords int     `json:"TotalRecords"`
	Records      []Order `json:"Records"`
}

// 订单
type Order struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	CreateDate  string  `json:"create_date"`
}

// 获取订单列表
func (s *ExpenseService) GetOrdersList(req *GetOrdersListRequest) (*GetOrdersListResponse, error) {
	path := "/expense/orders/list"

	var resp GetOrdersListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
