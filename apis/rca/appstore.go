package rca

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// 云应用创建App模板请求
type RcaCreateAppTemplateRequest struct {
	CrossVersionUpdate bool     `json:"cross_version_update"` // 是否支持跨版本更新
	Description        string   `json:"description"`          // 描述
	DisableUpdate      bool     `json:"disable_update"`       // 是否禁用版本更新功能，禁用后已安装应用无法检测到更新
	Logo               string   `json:"logo"`                 // logo（base64），最大50KB，可以不传
	Name               string   `json:"name"`                 // 应用名称（英文标识）
	ProjectLink        string   `json:"project_link"`         // 应用项目链接
	Readme             string   `json:"readme"`               // 介绍（markdown）
	Tags               []string `json:"tags"`                 // 标签，多选，目前支持的tag可以参考现有应用
	Title              string   `json:"title"`                // 应用标题
	Website            string   `json:"website"`              // 应用官方网站链接
}

// App模板
type AppTemplate struct {
	ID                 int      `json:"id"` // 模板id
	UID                int      `json:"uid"`
	Name               string   `json:"name"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Logo               string   `json:"logo"`
	Type               string   `json:"type"`
	Tags               []string `json:"tags"`
	Website            string   `json:"website"`
	Github             string   `json:"github"`
	CreateDate         int      `json:"create_date"`
	Readme             string   `json:"readme"`
	Provider           string   `json:"provider"`
	CrossVersionUpdate bool     `json:"cross_version_update"`
	Downloads          int      `json:"downloads"`
	IsPublic           bool     `json:"is_public"`
	ReviewStatus       string   `json:"review_status"`
	ReviewComment      string   `json:"review_comment"`
	StopReason         string   `json:"stop_reason"`
	DisableUpdate      bool     `json:"disable_update"`
}

type RcaCreateAppTemplateResponse struct {
	Code int         `json:"code"`
	Data AppTemplate `json:"data"`
}

// App模板版本
type AppTemplateVersion struct {
	ID       int    `json:"id"`
	Version  string `json:"version"`
	IsPublic bool   `json:"is_public"`
}

type AppTemplateVersionData struct {
	Data     AppTemplate          `json:"data"`
	Versions []AppTemplateVersion `json:"versions"`
}

type CreateAppTemplateVersionResponse struct {
	Code int                    `json:"code"`
	Data AppTemplateVersionData `json:"data"`
}

// 配置文件
type AppTemplateConfigMap struct {
	ContainerPath string `json:"container_path"` // 配置文件存放在容器里面的目录
	Content       string `json:"content"`        // 文件内容，最大1MB
	FileName      string `json:"file_name"`      // 文件名
}

// 环境变量
type AppTemplateEnv struct {
	Key   string `json:"key"`   // 键
	Value string `json:"value"` // 值，允许空白
}

// 安装选项
type AppTemplateOption struct {
	Default  string                   `json:"default"`  // 默认值
	Disabled bool                     `json:"disabled"` // 是否不允许编辑
	EnvKey   string                   `json:"env_key"`  // 绑定到环境变量的Key
	Label    string                   `json:"label"`    // 标签，告知用户这个选项是干什么用的，例如root用户密码
	Random   bool                     `json:"random"`   // 是否进行随机化生成
	Required bool                     `json:"required"` // 是否必选
	Rule     string                   `json:"rule"`     // 规则，可用项目按照前端已实现支持的规则来
	Type     string                   `json:"type"`     // 类型，如：password number等
	Value    string                   `json:"value"`    // 填入值
	Values   []AppTemplateOptionValue `json:"values"`   // 选项，type为select时提供给客户选择
}

// 选项可选值
type AppTemplateOptionValue struct {
	Label string `json:"label"` // 提示显示
	Value string `json:"value"` // 对应值
}

// 资源需求
type AppTemplateResourceRequest struct {
	MinCPU    int `json:"min_cpu"`    // 如1000，单位是m
	MinMemory int `json:"min_memory"` // 如4096，单位是Mi
}

// 脚本钩子
type AppTemplateScripts struct {
	Install      string `json:"install"`       // install（安装应用的时候执行，可以用于下载特定的持久化文件，可以指定运行环境，在底层这个是用initContainer来实现
	InstallImage string `json:"install_image"` // 执行安装脚本所使用的容器镜像，其他阶段会使用运行时镜像
	PostStart    string `json:"post_start"`    // poststart（容器创建后立即执行）
	PreStop      string `json:"pre_stop"`      // prestop（终止前运行）
}

// 服务
type AppTemplateService struct {
	ExternalPort string `json:"external_port"` // 外部端口，之所以是string，因为支持以变量形式传入
	InternalPort string `json:"internal_port"` // 内部端口，对应AppPort，之所以是string，因为支持以变量形式传入
	Label        string `json:"label"`         // 标签，可以是中文
	Name         string `json:"name"`          // 服务名称
	Protocol     string `json:"protocol"`      // 协议
	Type         string `json:"type"`          // 类型，可以是internal或external，对应clusterIP或者lb
}

// 文件目录挂载
type AppTemplateVolumeMount struct {
	MountContentType string `json:"mount_content_type"` // 挂载类型，文件file或目录dir
	MountPath        string `json:"mount_path"`         // 容器内路径
	Name             string `json:"name"`               // 挂载描述
	PreContent       string `json:"pre_content"`        // file类型专用，以base64存储的预先准备二进制数据，这个是用于一些小型二进制文件例如.db或者如/etc/localtime等的预先准备，模拟在docker中如"./data/cloudreve.db:/cloudreve/cloudreve.db"或/etc/timezone:/etc/timezone:ro
	SubPath          string `json:"sub_path"`           // Project卷内路径，不能以.或者/开头
}

// 创建App模板版本请求
//
// 我尝试创建了一个函数用于将docker compose文件转换成这种格式，但在编写了近1000行代码后，我成功的放弃了这个想法
// :(
type CreateAppTemplateVersionRequest struct {
	Args            []string                   `json:"args"`    // 运行参数
	Command         []string                   `json:"command"` // 运行命令
	ConfigMaps      []AppTemplateConfigMap     `json:"config_maps"`
	Env             []AppTemplateEnv           `json:"env"`   // 环境变量
	Image           string                     `json:"image"` // 容器镜像，如nginx:1.21.0
	Options         []AppTemplateOption        `json:"options"`
	ReleaseID       int                        `json:"release_id"`       // 更新用release ID
	ResourceRequest AppTemplateResourceRequest `json:"resource_request"` // 资源需求
	Scripts         AppTemplateScripts         `json:"scripts"`          // 脚本钩子
	Services        []AppTemplateService       `json:"services"`
	Version         string                     `json:"version"`       // 版本号，如1.21.0
	VolumeMounts    []AppTemplateVolumeMount   `json:"volume_mounts"` // 文件目录，如data:/var/lib/mysql
}

// 云应用创建App模板
func (s *RcaService) CreateRcaAppTemplate(req *RcaCreateAppTemplateRequest) (*RcaCreateAppTemplateResponse, error) {
	path := "/product/rca/appstore/"

	var resp RcaCreateAppTemplateResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}

// 创建App模板版本
//
// id: APP模板ID
func (s *RcaService) CreateRcaAppTemplateVersion(id int, req *CreateAppTemplateVersionRequest) (*CreateAppTemplateVersionResponse, error) {
	path := fmt.Sprintf("/product/rca/appstore/%d/release", id)

	var resp CreateAppTemplateVersionResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, req, &resp)
	return &resp, err
}
