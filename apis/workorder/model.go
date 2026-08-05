package workorder

import "github.com/XingfenD/rainyun_api_go_sdk/apis/common"

// 获取工单列表请求
type GetWorkOrderListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// 获取工单列表响应
type GetWorkOrderListResponse struct {
	Code int               `json:"code"`
	Data WorkorderListData `json:"data"`
}

// 工单列表数据
type WorkorderListData struct {
	TotalRecords int               `json:"TotalRecords"`
	Records      []WorkorderRecord `json:"Records"`
}

// 工单
type WorkorderRecord struct {
	ID                 int    `json:"ID"`  // 工单ID
	UID                int    `json:"UID"` // 用户ID
	ExtUserInfo        string `json:"ExtUserInfo"`
	UserName           string `json:"UserName"` // 用户名
	UserEmail          string `json:"UserEmail"`
	UserVip            string `json:"UserVip"`            // 用户vip等级
	UserIcon           string `json:"UserIcon"`           // 用户头像
	Title              string `json:"Title"`              // 工单标题
	Type               string `json:"Type"`               // 工单类型
	RelatedProductType string `json:"RelatedProductType"` // 关联产品类型
	RelatedProductID   int    `json:"RelatedProductID"`   // 关联产品id
	IsUrgent           int    `json:"IsUrgent"`           // 是否为紧急工单
	Status             string `json:"Status"`             // 状态(finished/answered/waiting)
	Time               int    `json:"Time"`               // 工单创建时间
	LastTime           int    `json:"LastTime"`
	WaitBeginTime      int    `json:"WaitBeginTime"`
	AssistID           int    `json:"AssistID"` // 客服id
	AuthStatus         string `json:"AuthStatus"`
	AuthTime           int    `json:"AuthTime"`
	AuthID             int    `json:"AuthID"`
}

// 工单回复
type WorkorderDiscuss struct {
	ID             int    `json:"ID"`
	IsAssist       bool   `json:"IsAssist"`
	UID            int    `json:"UID"`
	UserName       string `json:"UserName"`
	UserEmail      string `json:"UserEmail"`
	UserVip        string `json:"UserVip"`
	UserIcon       string `json:"UserIcon"`
	Content        string `json:"Content"`
	Time           int    `json:"Time"`
	WaitTime       int    `json:"WaitTime"`
	LastEditedTime int    `json:"LastEditedTime"`
	IsScored       bool   `json:"IsScored"`
}

// 工单详情数据
type WorkorderDetailData struct {
	WorkorderRecord
	Content string             `json:"Content"`
	Discuss []WorkorderDiscuss `json:"Discuss"`
}

// 获取工单详情响应
type GetWorkOrderDetailResponse struct {
	Code int                 `json:"code"`
	Data WorkorderDetailData `json:"data"`
}

// 创建工单
type CreateWorkerorderRequest struct {
	Content            string `json:"content"`
	IsAuthed           bool   `json:"is_authed"`            // 是否需要授权(可选)
	IsUrgent           int    `json:"is_urgent"`            // 是否为紧急工单
	RelatedProductID   int    `json:"related_product_id"`   // 关联产品类型(可选)
	RelatedProductType string `json:"related_product_type"` // 关联产品id(可选)
	Title              string `json:"title"`
	Type               string `json:"type"` // 工单类型
}

// 创建工单数据
type CreateWorkorderData struct {
	WorkorderRecord
	Content string `json:"Content"`
}

// 创建工单响应
type CreateWorkerorderResponse struct {
	Code int                 `json:"code"`
	Data CreateWorkorderData `json:"data"`
}

// 产品授权请求
type ProductAuthRequest struct {
	ProductID   int    `json:"product_id"`   // 产品ID
	ProductType string `json:"product_type"` // 产品类型(rvh/rcs/rgs/rbm/ros)
}

// 回复工单请求
type ReplyWorkerorderRequest struct {
	Content string `json:"content"` // 回复内容
}

// 工单打分
type ScoreWorkerorderRequest struct {
	Score     int    `json:"score"`      // 分数(1-5)
	Reason    string `json:"reason"`     // (可选)
	IsSolved  bool   `json:"is_solved"`  // 是否解决(可选)
	DiscussID int    `json:"discuss_id"` // 回复ID(可选)
	Aid       int    `json:"aid"`        // 客服id(可选)
}

// 工单打分详情数据
type ScoreWorkorderData struct {
	UID         int    `json:"uid"`
	Aid         int    `json:"aid"`
	OrderID     int    `json:"order_id"`
	DiscussID   int    `json:"discuss_id"`
	Score       int    `json:"score"`
	Reason      string `json:"reason"`
	IsSolved    bool   `json:"is_solved"`
	DiscussTime int    `json:"DiscussTime"`
	ScoreTime   int    `json:"score_time"`
}

// 获取工单打分响应
type GetScoreWorkorderResponse struct {
	Code int                `json:"code"`
	Data ScoreWorkorderData `json:"data"`
}

// 工单状态数据
type WorkorderStatusData struct {
	LastTime int    `json:"LastTime"`
	Status   string `json:"Status"`
}

// 获取工单状态响应
type GetWorkorderStatusResponse struct {
	Code int                 `json:"code"`
	Data WorkorderStatusData `json:"data"`
}

// 设置工单状态请求
type SetWorkorderStatusRequest struct {
	Status string `json:"status"`
}
