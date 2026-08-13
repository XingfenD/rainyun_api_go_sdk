package domain

import "github.com/XingfenD/rainyun_api_go_sdk/apis/common"

// 添加域名过白请求
type AddDomainToWhiteListRequest struct {
	Domain string `json:"domain"` // 域名
	Region string `json:"region"` // 区域：cn-sq1/cn-nb1/cn-xy1/cn-cq1
}

// 获取域名白名单列表查询参数
type GetDomainWhiteListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// 域名白名单列表响应
type GetDomainWhiteListResponse struct {
	Code int                 `json:"code"`
	Data DomainWhiteListData `json:"data"`
}

// 域名白名单列表数据
type DomainWhiteListData struct {
	TotalRecords int                     `json:"TotalRecords"`
	Records      []DomainWhiteListRecord `json:"Records"`
}

// 域名白名单记录
type DomainWhiteListRecord struct {
	ID          int    `json:"id"` // ID
	UID         int    `json:"uid"`
	Domain      string `json:"domain"`       // 域名
	SiteLicense string `json:"site_license"` // 备案号（貌似只有早期过白的会有这个）
	Region      string `json:"region"`       // 区域
	AddTime     int    `json:"add_time"`     // 时间
	Status      int    `json:"status"`       // 状态: 1:已处理
}

// 获取已验证域名列表查询参数
type GetVerifiedDomainListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// 已验证域名列表响应
type GetVerifiedDomainListResponse struct {
	Code int                    `json:"code"`
	Data VerifiedDomainListData `json:"data"`
}

// 已验证域名列表数据
type VerifiedDomainListData struct {
	TotalRecords int                    `json:"TotalRecords"`
	Records      []VerifiedDomainRecord `json:"Records"`
}

// 已验证域名记录
type VerifiedDomainRecord struct {
	ID      int    `json:"id"` // id
	UID     int    `json:"uid"`
	Domain  string `json:"domain"`   // 域名
	AddTime int    `json:"add_time"` // 时间
}

// 添加域名认证请求
type AddDomainVerificationRequest struct {
	Domain string `json:"domain"`
}

// 域名验证信息响应
type GetDomainVerificationInfoResponse struct {
	Code int                    `json:"code"`
	Data DomainVerificationData `json:"data"`
}

// 域名验证信息数据
type DomainVerificationData struct {
	Rr        string `json:"rr"`         // 主机名
	TopDomain string `json:"top_domain"` // 主域名
	Record    string `json:"record"`     // 记录值
}

// 添加域名DNS解析请求
type AddDomainDNSRecordRequest struct {
	Host            string `json:"host"`              // 主机名
	Level           int    `json:"level"`             // 优先等级
	Line            string `json:"line"`              // 线路，必须有一个默认的才可以 DEFAULT/LTEL/LCNC/LMOB/LEDU/LSEO/LFOR
	RainProductID   int    `json:"rain_product_id"`   // 雨云产品id(搭配其它产品时使用)
	RainProductType string `json:"rain_product_type"` // 雨云产品类型(搭配其它产品时使用) rcs/rvh/rgs/rbm
	RecordID        int    `json:"record_id"`         // 记录id(修改时使用)
	TTL             int    `json:"ttl"`               // 解析生效时间
	Type            string `json:"type"`              // 解析类型 A/AAAA/CNAME/TXT/MX/SRV
	Value           string `json:"value"`             // 解析值
}

// 删除域名DNS解析请求
type DeleteDomainDNSRecordRequest struct {
	RecordID int `json:"record_id"` // 记录id
}

// 添加域名DNSSEC请求
type AddDomainDNSSECRequest struct {
	Domain    string `json:"domain"`
	Keyalg    string `json:"keyalg"`
	Keydigest string `json:"keydigest"`
	Keytag    string `json:"keytag"`
	Keytype   string `json:"keytype"`
}

// 删除域名DNSSEC请求
type DeleteDomainDNSSECRequest struct {
	Domain    string `json:"domain"`
	Keydigest string `json:"keydigest"`
	Keytag    int    `json:"keytag"`
}

// 修改域名NS服务器请求
type UpdateDomainNSRequest struct {
	NameServers []string `json:"name_servers"`
}

// 续费域名请求
type RenewDomainRequest struct {
	Years int `json:"years"` // 续费年限
}

// 域名过户请求
type DomainTransferRequest struct {
	NewTemplateInfo EditDomainTemplateRequest `json:"new_template_info"` // 和SysID二选一，如果没模板就用这个新建模板
	SysID           string                    `json:"sys_id"`            // 新域名模板ID
}

// 查询域名模板列表查询参数
type GetDomainTemplateListRequest struct {
	Options common.StandQueryParameters `json:"options"`
}

// 域名模板列表响应
type GetDomainTemplateListResponse struct {
	Code int                    `json:"code"`
	Data DomainTemplateListData `json:"data"`
}

// 域名模板列表数据
type DomainTemplateListData struct {
	TotalRecords int                    `json:"TotalRecords"`
	Records      []DomainTemplateDetail `json:"Records"` // 域名模板详情
}

// 删除域名信息模板请求
type DeleteDomainTemplateRequest struct {
	SysID string `json:"sys_id"` // 模板标识
}

// 获取域名模板详情请求
type GetDomainTemplateDetailRequest struct {
	SysID string `json:"sys_id"` // 模板标识
}

// 获取域名模板详情响应
type GetDomainTemplateDetailResponse struct {
	Code int                  `json:"code"`
	Data DomainTemplateDetail `json:"data"`
}

// 域名模板详情
type DomainTemplateDetail struct {
	ID          int    `json:"id"`
	UID         int    `json:"uid"`
	AddTime     int    `json:"add_time"`     // 添加时间
	SysID       string `json:"sys_id"`       // 模板标识
	Type        string `json:"type"`         // 所有者类型，I是个人，E是企业
	Owner       string `json:"owner"`        // 模板所有者
	ContactName string `json:"contact_name"` // 联系人姓名
	CompanyName string `json:"company_name"` // 企业名称，如果是企业模板就有这个
	Email       string `json:"email"`        // 联系人邮箱
	Country     string `json:"country"`      // 所在国家
	Province    string `json:"province"`     // 中文，省
	City        string `json:"city"`         // 中文，市
	ZipCode     string `json:"zip_code"`     // 中文，邮编
	Phone       string `json:"phone"`        // 联系电话
	Address     string `json:"address"`      // 中文，通讯地址
	IDType      string `json:"id_type"`      // 证件类型，如SFZ
	IDNum       string `json:"id_num"`       // 证件值
	CStatus     string `json:"c_status"`     // 不明， 1
	CFailInfo   string `json:"c_fail_info"`  // 模板失败信息
	RStatus     string `json:"r_status"`     // 不明, 1
	RFailInfo   string `json:"r_fail_info"`  // 不明
}

// 更新域名管理密码请求
type UpdateDomainPasswordRequest struct {
	Password string `json:"password"`
}

// 编辑域名模板请求
type EditDomainTemplateRequest struct {
	Address       string `json:"address"`         // 中文，通讯地址
	City          string `json:"city"`            // 中文，市
	CompanyEnName string `json:"company_en_name"` // 企业英文名称，如果类型是企业就要传入
	CompanyName   string `json:"company_name"`    // 企业名称，如果类型是企业就要传入
	Country       string `json:"country"`         // 所在国家
	Email         string `json:"email"`           // 联系人邮箱
	IDImg         string `json:"id_img"`          // base64编码的证件照片
	IDNum         string `json:"id_num"`          // 证件值
	IDType        string `json:"id_type"`         // 证件类型，如SFZ
	Name          string `json:"name"`            // 联系人名
	Phone         string `json:"phone"`           // 联系电话
	Province      string `json:"province"`        // 中文，省
	SysID         string `json:"sys_id"`          // 模板标识(修改时使用)
	Type          string `json:"type"`            // 所有者类型，I是个人，E是企业
	ZipCode       string `json:"zip_code"`        // 中文，邮编
}

// 获取域名whois信息请求
type GetDomainWhoisInfoRequest struct {
	Domain string `json:"domain"`
}

// 域名WHOIS信息响应
type GetDomainWhoisInfoResponse struct {
	Code int                 `json:"code"`
	Data DomainWhoisInfoData `json:"data"`
}

// 域名WHOIS信息数据
type DomainWhoisInfoData struct {
	RegDate    string `json:"reg_date"`   // 注册时间
	ExpDate    string `json:"exp_date"`   // 到期时间
	Status     string `json:"status"`     // 域名状态
	Owner      string `json:"owner"`      // 域名所有者
	Registrar  string `json:"registrar"`  // 域名注册商
	Email      string `json:"email"`      // 域名所有者邮箱
	Nameserver string `json:"nameserver"` // 域名服务器
	BizServer  string `json:"biz_server"`
	Body       string `json:"body"`    // 域名WHOIS信息原文
	Updated    string `json:"updated"` // WHOIS信息更新时间
}

// 检查域名能否注册请求
type CheckDomainAvailableRequest struct {
	Domain string `json:"domain"` // 要检查的域名（中文域名同时支持汉字及punycode转码）
	Suffix string `json:"suffix"` // 要检查的后缀，多个的话，用英文逗号分割
}

// 检查域名能否注册响应
type CheckDomainAvailableResponse struct {
	Code int                        `json:"code"`
	Data []CheckDomainAvailableItem `json:"data"`
}

// 检查域名能否注册数据
type CheckDomainAvailableItem struct {
	IsAvail     bool        `json:"is_avail"` // 是否可注册
	SpecialType string      `json:"special_type"`
	Name        string      `json:"name"` // 域名
	Price       DomainPrice `json:"price"`
}

// 域名价格
type DomainPrice struct {
	FirstYearPrice int `json:"first_year_price"` // 首年价格
	BuyPrice       int `json:"buy_price"`        // 购买价格
	RenewPrice     int `json:"renew_price"`      // 续费价格
}
