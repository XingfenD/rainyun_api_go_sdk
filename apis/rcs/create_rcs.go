package rcs

import (
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

type CreateRcsRequest struct {
	AddDiskSize  int    `json:"add_disk_size"`  // 额外硬盘容量GB
	PlanID       int    `json:"plan_id"`        // 套餐ID
	Duration     int    `json:"duration"`       // 创建时长(月)
	OsID         int    `json:"os_id"`          // 系统ID
	WithEipNum   int    `json:"with_eip_num"`   // 创建IP数量
	WithEipFlags string `json:"with_eip_flags"` // 是否开启高防，us_ddosip -> 美国高防，nb_ddosip -> 宁波高防
	WithEipType  string `json:"with_eip_type"`  // ipv4(默认)/ipv6
	WithCouponID int    `json:"with_coupon_id"` // 优惠券ID
	Try          bool   `json:"try"`            // 是否为试用
	NodeUUID     string `json:"node_uuid"`      // 指定节点(管理员可用，用户不可用)
	AppVars      []struct {
		AppID int `json:"app_id"`
		Vars  any `json:"vars,omitempty"`
	} `json:"app_vars"` // 预装应用
	Zone string `json:"zone"` // 内网可用区
}

type CreateRcsResponse struct {
	Code int       `json:"code"`
	Data RcsRecord `json:"data"`
}

func (s *RcsService) CreateRcs(req *CreateRcsRequest) (*CreateRcsResponse, error) {
	path := "/product/rcs"

	var resp CreateRcsResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
