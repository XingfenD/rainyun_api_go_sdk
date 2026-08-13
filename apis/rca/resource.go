package rca

import (
	"fmt"
	"strconv"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 云应用项目设置备份目标请求
type SetRcaProjectBackupTargetRequest struct {
	S3AccessKey       string `json:"s3_access_key"`       // S3的AK
	S3BackupDirectory string `json:"s3_backup_directory"` // s3备份存储的目录
	S3Bucket          string `json:"s3_bucket"`           // S3存储桶名
	S3Endpoint        string `json:"s3_endpoint"`         // S3的端点（仅支持Virtual Host，不支持Path-Style模式）
	S3SecretKey       string `json:"s3_secret_key"`       // S3的SK
	TargetType        string `json:"target_type"`         // 目标类型，支持项目本地备份或者远程S3存储(local/s3)
}

// 云应用项目磁盘扩容请求
type RcaProjectDiskExpansionRequest struct {
	NewDiskSize int `json:"new_disk_size"` // 以GB显示的新项目磁盘大小
}

// 云应用增加IP地址请求
type RcaAddsIpAddressRequest struct {
	Ipv4Count int `json:"ipv4_count"` // 要添加的IPv4地址数量
	Ipv6Count int `json:"ipv6_count"` // 要添加的IPv6地址数量
}

// 云应用移除IP地址请求
type RcaRemoveIPRequest struct {
	IPID int `json:"ip_id"` // 要删除的IP地址ID
}

// 云应用项目修改SFTP设置请求
type RcaProjectSetSftpConfigRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// 云应用IP信息
type RcaIP struct {
	ID            int    `json:"id"`           // ip id
	Region        string `json:"region"`       // 地域
	Type          string `json:"type"`         // ip类型(IPv4/IPv6)
	IP            string `json:"ip"`           // ip地址
	AddressPool   string `json:"address_pool"` // ip池(user-ip-pool)
	Gateway       string `json:"gateway"`      // 网关
	Block         string `json:"block"`        // CIDR(24)
	UID           int    `json:"uid"`          // 用户id
	ProjectID     int    `json:"project_id"`
	Info          string `json:"info"`
	AllocatedDate int    `json:"allocated_date"` // 分配时间
}

type GetRcaProjectIPsResponse struct {
	Code int     `json:"code"`
	Data []RcaIP `json:"data"`
}

// 云应用项目设置备份目标
//
// id: RCA项目ID
func (s *RcaService) SetRcaProjectBackupTarget(id int, req *SetRcaProjectBackupTargetRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/backup_target", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, req, &resp)
	return &resp, err
}

// 云应用项目磁盘扩容
//
// id: RCA项目ID
func (s *RcaService) ExpandRcaProjectDisk(id int, req *RcaProjectDiskExpansionRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/disk_expand", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 云应用增加IP地址
//
// id: RCA项目ID
func (s *RcaService) AddRcaProjectIP(id int, req *RcaAddsIpAddressRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/eip", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 云应用移除IP地址
//
// id: RCA项目ID
func (s *RcaService) RemoveRcaProjectIP(id int, req *RcaRemoveIPRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/eip", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_DELETE, path, nil, req, &resp)
	return &resp, err
}

// 云应用项目修改SFTP设置
//
// id: RCA项目ID
//
// password: 密码
//
// username: 用户名
func (s *RcaService) SetRcaProjectSFTPConfig(id int, password string, username string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/rca/project/%d/sftp", id)

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_PATCH, path, nil, RcaProjectSetSftpConfigRequest{Password: password, Username: username}, &resp)
	return &resp, err
}

// 云应用项目列出IP地址
func (s *RcaService) ListRcaProjectIPs(id int) (*GetRcaProjectIPsResponse, error) {
	path := "/product/rca/project/eip"

	var resp GetRcaProjectIPsResponse
	querys := map[string]string{"id": strconv.Itoa(id)}
	err := s.client.Do(constant.HTTPMethod_GET, path, querys, nil, &resp)
	return &resp, err
}
