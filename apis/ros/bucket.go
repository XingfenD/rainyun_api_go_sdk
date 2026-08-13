package ros

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 创建对象存储桶请求
type CreateRosBucketRequest struct {
	BucketName string `json:"bucket_name"` // 存储桶名称
	InstanceID int    `json:"instance_id"` // 实例ID
}

// 创建对象存储桶响应
type CreateRosBucketResponse struct {
	Code int       `json:"code"`
	Data RosBucket `json:"data"`
}

// 获取对象存储桶列表请求
type GetRosBucketListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// 获取对象存储桶列表响应
type GetRosBucketListResponse struct {
	Code int           `json:"code"`
	Data RosBucketList `json:"data"`
}

// 对象存储桶列表
type RosBucketList struct {
	TotalRecords int         `json:"TotalRecords"`
	Records      []RosBucket `json:"Records"`
}

// 获取对象存储桶详情响应
type GetRosBucketDetailResponse struct {
	Code int             `json:"code"`
	Data RosBucketDetail `json:"data"`
}

// 对象存储桶详情
type RosBucketDetail struct {
	Data RosBucket `json:"Data"`
}

// 修改存储桶Proxy设置请求
// 一个请求只能设置一项设置，要设置的设置项需要放在 Setting 中，只需要写需要修改的设置项，其他不填
type ModifyRosBucketProxySettingsRequest struct {
	DomainList              []string                   `json:"domain_list,omitempty"` // 域名列表
	GzipSettings            RosGzipSettings            `json:"gzip_settings,omitempty"`
	IPRestrictSettings      RosIPRestrictSettings      `json:"ip_restrict_settings,omitempty"`
	RefererRestrictSettings RosRefererRestrictSettings `json:"referer_restrict_settings,omitempty"`
	Setting                 string                     `json:"setting"` // 必选项，要设置的选项(domain/ssl/waf/gzip/ip_restrict/referer_restrict)
	SslSettings             RosSslSettings             `json:"ssl_settings,omitempty"`
	WafSettings             RosWafSettings             `json:"waf_settings,omitempty"`
}

// Gzip设置
type RosGzipSettings struct {
	CompressLevel int      `json:"compress_level,omitempty"` // 压缩等级（1-5,默认为2）
	FileTypes     []string `json:"file_types,omitempty"`     // 文件MIME类型，默认["*"]
	IsEnable      bool     `json:"is_enable,omitempty"`      // 是否启用GZIP压缩，默认启用
}

// IP访问限制设置
type RosIPRestrictSettings struct {
	Blacklist []string `json:"blacklist,omitempty"` // 黑名单IP或CIDR范围列表，和白名单列表只能二选一不能共存
	IsEnable  bool     `json:"is_enable,omitempty"` // 是否启用IP访问限制
	Whitelist []string `json:"whitelist,omitempty"` // 白名单IP或CIDR范围列表，和黑名单列表只能二选一不能共存
}

// Referer防盗链设置
type RosRefererRestrictSettings struct {
	Blacklist     []string `json:"blacklist,omitempty"`      // 黑名单域名列表，和白名单列表只能二选一不能共存
	BypassMissing bool     `json:"bypass_missing,omitempty"` // 是否允许Referer请求头不存在或者格式有误
	IsEnable      bool     `json:"is_enable,omitempty"`      // 是否启用Referer防盗链
	Whitelist     []string `json:"whitelist,omitempty"`      // 白名单域名列表，和黑名单列表只能二选一不能共存
}

// SSL设置，SSL证书会自动匹配
type RosSslSettings struct {
	IsEnable bool `json:"is_enable,omitempty"` // 是否启用SSL
	IsForce  bool `json:"is_force,omitempty"`  // 是否强制使用SSL
}

// WAF防火墙设置
type RosWafSettings struct {
	BlockTime     int  `json:"block_time,omitempty"`      // 未通过JS验证时的封禁时间，秒
	GlobalJsCheck bool `json:"global_js_check,omitempty"` // 是否启用全局JS检查，只有超过阈值才会生效
	GlobalQPS     int  `json:"global_qps,omitempty"`      // 全局访问速率限制，超过会被拒绝
	IsEnable      bool `json:"is_enable,omitempty"`       // 是否启用WAF防火墙
	PerIPJsCheck  bool `json:"per_ip_js_check,omitempty"` // 是否启用单IP JS检查，只有超过阈值才会生效
	PerIPQPS      int  `json:"per_ip_qps,omitempty"`      // 单个IP访问速率限制，超过会被拒绝
}

// 获取对象存储桶列表
//
// ⚠️注意这个接口是拿不到AK和SK的，响应里面的AK和SK都是空的
func (s *RosService) GetRosBucketList(req *GetRosBucketListRequest) (*GetRosBucketListResponse, error) {
	path := "/product/ros/bucket/"

	var resp GetRosBucketListResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 创建对象存储桶
//
// req: 创建参数（bucketName 存储桶名 >= 3 字符 <= 63 字符）
func (s *RosService) CreateRosBucket(req CreateRosBucketRequest) (*CreateRosBucketResponse, error) {
	path := "/product/ros/bucket"

	var resp CreateRosBucketResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 获取对象存储桶详情
//
// bucketID: 存储桶ID
func (s *RosService) GetRosBucketDetail(bucketID int) (*GetRosBucketDetailResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d", bucketID)

	var resp GetRosBucketDetailResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}

// 删除对象存储桶
//
// bucketID: 存储桶ID
func (s *RosService) DeleteRosBucket(bucketID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d", bucketID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, nil, &resp)
	return &resp, err
}

// 获取对象存储桶监控数据
//
// bucketID: 存储桶ID
//
// req: 开始时间/结束时间（timestamp）
func (s *RosService) GetRosBucketMonitorData(bucketID int, req *GetRosMonitorDataRequest) (*GetRosMonitorDataResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d/monitor", bucketID)

	var resp GetRosMonitorDataResponse
	querys, err := common.MarshalQueryParams(req)
	if err != nil {
		return nil, err
	}
	err = s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}

// 修改存储桶Proxy设置
//
// bucketID: 存储桶ID
func (s *RosService) ModifyRosBucketProxySettings(bucketID int, req *ModifyRosBucketProxySettingsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d/proxy", bucketID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// 对象存储桶重新生成密钥
//
// bucketID: 存储桶ID
func (s *RosService) ReGenerateRosBucketKeys(bucketID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d/regenerate-keys", bucketID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 开关对象存储桶匿名访问
//
// bucketID: 存储桶ID
func (s *RosService) SetRosBucketPublicAccess(bucketID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d/toggle-public-access", bucketID)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, nil, &resp)
	return &resp, err
}

// 按实例获取对象存储桶列表响应
type GetRosBucketListByInstanceResponse struct {
	Code int           `json:"code"`
	Data RosBucketList `json:"data"`
}

// 按实例获取对象存储桶列表
//
// instanceID: 实例ID
func (s *RosService) GetRosBucketListByInstance(instanceID int) (*GetRosBucketListByInstanceResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d", instanceID)

	var resp GetRosBucketListByInstanceResponse
	err := s.client.Do(constant.HTTPMethod_GET, path, nil, nil, &resp)
	return &resp, err
}
