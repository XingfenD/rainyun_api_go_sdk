package rca

// ---- app ----
type InstallRcaAppRequest struct {
	AppstoreReleaseID int                   `json:"appstore_release_id"`
	ContainerSettings []RcaContainerSetting `json:"container_settings"`
	Options           any                   `json:"options"`
	ProjectID         int                   `json:"project_id"`
}

type RcaContainerSetting struct {
	ContainerID            int         `json:"container_id"`
	CPULimit               int         `json:"cpu_limit"`
	EnvList                []RcaAppEnv `json:"env_list"`
	MemoryLimit            int         `json:"memory_limit"`
	ServiceHTTPSDomainType string      `json:"service_https_domain_type"`
	ServiceHTTPSDomains    []string    `json:"service_https_domains"`
	ServiceIP              string      `json:"service_ip"`
	ServiceType            string      `json:"service_type"`
	Services               any         `json:"services"`
}

type RcaAppEnv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UninstallRcaAppRequest struct {
	RemoveData bool `json:"remove_data"`
}

type UpdateRcaAppRequest struct {
	ConfigScope string `json:"config_scope"`
	Options     any    `json:"options"`
}

type UpdateRcaAppContainerRequest struct {
	ConfigMap      any         `json:"config_map"`
	ConfigScope    string      `json:"config_scope"`
	EnvList        []RcaAppEnv `json:"env_list"`
	ResourceLimits any         `json:"resource_limits"`
}

type UpdateRcaAppPHPSettingRequest struct {
	DisableFunctions []string `json:"disable_functions"`
	Params           any      `json:"params"`
	Scope            string   `json:"scope"`
	UploadMaxSize    string   `json:"upload_max_size"`
}

type CreateRcaAppServiceRequest struct {
	DryRun          bool          `json:"dry_run"`
	HTTPSDomainType string        `json:"https_domain_type"`
	HTTPSDomains    []string      `json:"https_domains"`
	IP              string        `json:"ip"`
	Label           string        `json:"label"`
	Name            string        `json:"name"`
	Ports           []RcaPortInfo `json:"ports"`
	Type            string        `json:"type"`
}

type RcaPortInfo struct {
	ExternalPort int    `json:"external_port"`
	InternalPort int    `json:"internal_port"`
	Protocol     string `json:"protocol"`
}

type UpdateRcaAppWebserverAccessRequest struct {
	IP string `json:"ip"`
}

type UpgradeRcaAppRequest struct {
	ReleaseID int `json:"release_id"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaAppListResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaAppDetailResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaAppConfigMapResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaAppContainerMetricsResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaAppPHPSettingResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaAppServiceListResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// ---- website ----
type CreateRcaWebsiteRequest struct {
	AppProxyAppContainerID           int      `json:"app_proxy_app_container_id"`
	AppProxyAppID                    int      `json:"app_proxy_app_id"`
	AppProxyService                  string   `json:"app_proxy_service"`
	AppProxyServicePort              int      `json:"app_proxy_service_port"`
	AppProxyServiceToInternal        bool     `json:"app_proxy_service_to_internal"`
	DomainType                       string   `json:"domain_type"`
	Domains                          []string `json:"domains"`
	Name                             string   `json:"name"`
	PHPLimits                        any      `json:"php_limits"`
	PHPStoreReleaseContainerID       int      `json:"php_store_release_container_id"`
	PHPStoreReleaseID                int      `json:"php_store_release_id"`
	ProjectID                        int      `json:"project_id"`
	ReverseProxyIsEnableCache        bool     `json:"reverse_proxy_is_enable_cache"`
	ReverseProxyURL                  string   `json:"reverse_proxy_url"`
	Type                             string   `json:"type"`
	WebserverLimits                  any      `json:"webserver_limits"`
	WebserverStoreReleaseContainerID int      `json:"webserver_store_release_container_id"`
	WebserverStoreReleaseID          int      `json:"webserver_store_release_id"`
}

type UpdateRcaWebsiteNginxRequest struct {
	AppProxyAppContainerID    int      `json:"app_proxy_app_container_id"`
	AppProxyAppID             int      `json:"app_proxy_app_id"`
	AppProxyService           string   `json:"app_proxy_service"`
	AppProxyServicePort       int      `json:"app_proxy_service_port"`
	ConfigScope               string   `json:"config_scope"`
	DomainList                []string `json:"domain_list"`
	Index                     string   `json:"index"`
	ReverseProxyIsEnableCache bool     `json:"reverse_proxy_is_enable_cache"`
	ReverseProxyURL           string   `json:"reverse_proxy_url"`
	RewriteRule               string   `json:"rewrite_rule"`
	RuntimeDir                string   `json:"runtime_dir"`
	SSLAccessMode             string   `json:"ssl_access_mode"`
	SSLAlgorithm              string   `json:"ssl_algorithm"`
	SSLCertID                 int      `json:"ssl_cert_id"`
	SSLIsEnable               bool     `json:"ssl_is_enable"`
	SSLProtocol               []string `json:"ssl_protocol"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaWebsiteListResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaWebsiteDetailResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// TODO: 响应结构未公开,透传;实测后补强类型
type GetRcaWebsiteRewriteConfigResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}
