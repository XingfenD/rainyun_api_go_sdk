# SDK / CLI / openapi.json 接入进度

本文档追踪 Rainyun API（`docs/openapi.json`）在 Go SDK 与 `ry` CLI 中的接入进度。
SDK 代码在 `apis/`，CLI 代码在 `cmd/ry/commands/`。

> 重新生成：`python3 docs/scripts/generate.py`

## 总览

| 分类 | openapi 端点 | SDK 已实现 | CLI 已接入 |
|---|---|---|---|
| domain | 38 | 24 | 4 |
| public | 6 | 4 | 4 |
| rbm | 23 | 9 | 0 |
| rca | 56 | 42 | 20 |
| rcdn | 22 | 21 | 9 |
| rcs | 52 | 35 | 34 |
| rgs | 75 | 75 | 75 |
| ros | 36 | 19 | 3 |
| rvh | 25 | 25 | 11 |
| ssl | 18 | 18 | 18 |
| **合计** | **351** | **272** | **178** |

## domain

CLI 命令：`add <domain-id>`, `delete <domain-id> <record-id>`, `dns`, `domain`, `list`, `list <domain-id>`

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/domain/` | 列出域名列表 | `domain.GetDomainList` | ✓ |
| DELETE | `/product/domain/certify` | 删除域名认证 | — | — |
| GET | `/product/domain/certify` | 获取已验证域名列表 | `domain.GetVerifiedDomainList` | — |
| POST | `/product/domain/certify` | 添加域名认证 | `domain.AddDomainCertify` | — |
| POST | `/product/domain/certify/verify` | 域名认证校验 | `domain.VerifyDomainCertify` | — |
| POST | `/product/domain/check` | 检查域名能否注册 | — | — |
| DELETE | `/product/domain/free_subdomain` | 删除免费二级域名 | — | — |
| GET | `/product/domain/free_subdomain` | 获取免费二级域名列表 | — | — |
| POST | `/product/domain/free_subdomain` | 创建免费二级域名 | — | — |
| POST | `/product/domain/free_subdomain/proxy` | 修改免费二级域名的CDN设置 | — | — |
| GET | `/product/domain/free_subdomain/usable` | 获取可用的免费域名列表 | — | — |
| POST | `/product/domain/register` | 域名注册 | — | — |
| PATCH | `/product/domain/template` | 编辑域名模板 | `domain.EditDomainTemplate` | — |
| DELETE | `/product/domain/template/` | 删除域名信息模板 | `domain.DeleteDomainTemplate` | — |
| GET | `/product/domain/template/` | 查询域名模板列表 | `domain.GetDomainTemplateList` | — |
| GET | `/product/domain/template/detail/` | 获取域名模板详情 | `domain.GetDomainTemplateDetail` | — |
| GET | `/product/domain/whitelist` | 获取域名白名单列表 | `domain.GetDomainWhiteList` | — |
| POST | `/product/domain/whitelist` | 添加域名白名单 | — | — |
| GET | `/product/domain/whois` | 获取域名whois信息 | `domain.GetDomainWhoisInfo` | — |
| GET | `/product/domain/{id}` | 获取域名详情 | — | — |
| GET | `/product/domain/{id}/cert` | 下载域名证书 | — | — |
| PATCH | `/product/domain/{id}/dns` | 修改域名DNS解析 | `domain.UpdateDomainDNSRecord` | — |
| POST | `/product/domain/{id}/dns` | 添加域名DNS解析 | `domain.AddDomainDNSRecord` | ✓ |
| DELETE | `/product/domain/{id}/dns/` | 删除域名DNS解析 | `domain.DeleteDomainDNSRecord` | ✓ |
| GET | `/product/domain/{id}/dns/` | 获取域名DNS解析记录列表 | `domain.GetDomainDNSRecordList` | ✓ |
| GET | `/product/domain/{id}/dnssec` | 获取域名DNSSEC详情 | — | — |
| POST | `/product/domain/{id}/dnssec` | 添加域名DNSSEC | `domain.AddDomainDNSSEC` | — |
| POST | `/product/domain/{id}/dnssec/delete` | 删除域名DNSSEC | `domain.DeleteDomainDNSSEC` | — |
| POST | `/product/domain/{id}/dnssec/sync` | 同步域名DNSSEC | `domain.SyncDomainDNSSEC` | — |
| PUT | `/product/domain/{id}/lock/disable` | 关闭域名锁定 | `domain.UnlockDomain` | — |
| PUT | `/product/domain/{id}/lock/enable` | 开启域名锁定 | `domain.LockDomain` | — |
| POST | `/product/domain/{id}/nameservers` | 修改域名NS服务器 | `domain.UpdateDomainNS` | — |
| POST | `/product/domain/{id}/nameservers/reset` | 重置域名NS服务器 | `domain.ResetDomainNS` | — |
| GET | `/product/domain/{id}/password` | 获取域名管理密码 | — | — |
| POST | `/product/domain/{id}/password` | 更新域名管理密码 | `domain.UpdateDomainPassword` | — |
| POST | `/product/domain/{id}/renew` | 续费域名 | `domain.RenewDomain` | — |
| GET | `/product/domain/{id}/renew-price` | 获取域名续费价格 | — | — |
| POST | `/product/domain/{id}/transfer` | 域名过户 | `domain.TransferDomain` | — |

## public

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/app_config` | 获取页面信息 | `public.GetAppConfig` | ✓ |
| GET | `/discourse` | 获取论坛数据 | `public.GetDiscourse` | ✓ |
| GET | `/news` | 获取论坛公告 | `public.GetNews` | ✓ |
| GET | `/short_params` | 查询参数缩短的具体base64 | — | — |
| POST | `/short_params` | 创建参数缩短 | — | — |
| GET | `/status` | 获取节点网络状态 | `public.GetStatus` | ✓ |

## rbm

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/rbm` | 列出RBM实例 | — | — |
| POST | `/product/rbm/` | 创建RBM实例 | — | — |
| GET | `/product/rbm/price` | 价格计算 | — | — |
| GET | `/product/rbm/usage` | 获取使用情况列表 | — | — |
| POST | `/product/rbm/{id}/bios-flash` | 裸金属刷bios | — | — |
| POST | `/product/rbm/{id}/changeos` | RBM实例更换系统 | `rbm.ChangeRBMOS` | — |
| POST | `/product/rbm/{id}/commission` | RBM清点配置 | — | — |
| POST | `/product/rbm/{id}/eip/` | 创建并绑定弹性IP到RBM | `rbm.AssociateEIP` | — |
| POST | `/product/rbm/{id}/eip/change` | 更换IP | — | — |
| POST | `/product/rbm/{id}/eip/description` | 设置IP描述 | `rbm.SetIPDescription` | — |
| POST | `/product/rbm/{id}/eip/discard` | 放弃IP | `rbm.ReleaseIP` | — |
| POST | `/product/rbm/{id}/free` | 释放 | — | — |
| POST | `/product/rbm/{id}/kvm-proxy` | RBM实例启动KVM代理 | `rbm.StartKVMAgent` | — |
| POST | `/product/rbm/{id}/kvm-reboot` | RBM重新启动KVM | `rbm.RestartKVM` | — |
| GET | `/product/rbm/{id}/monitor` | 获取监控数据 | — | — |
| POST | `/product/rbm/{id}/poweroff` | RBM实例关机 | `rbm.ShutdownRBM` | — |
| POST | `/product/rbm/{id}/poweron` | RBM实例开机 | `rbm.StartRBM` | — |
| POST | `/product/rbm/{id}/rescue` | RBM救援模式切换 | — | — |
| POST | `/product/rbm/{id}/reset-password` | 重置RBM实例IPMI密码 | — | — |
| POST | `/product/rbm/{id}/traffic/auto` | 充流量 | — | — |
| POST | `/product/rbm/{id}/traffic/charge` | 充流量 | `rbm.ChargeTraffic` | — |
| POST | `/product/rbm/{id}/traffic/limit` | 限流 | — | — |
| POST | `/product/rbm/{id}/traffic/switch` | 切换流量套餐 | — | — |

## rca

CLI 命令：`app`, `balance`, `buy`, `create`, `delete <id>`, `get <id>`, `list`, `list <project-id>`, `log`, `plans`, `project`, `raindrop`, `rca`, `regions`, `restart <id>`, `start <id>`, `stop <id>`, `website`

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/rca/app/` | 云应用列出App | `rca.GetRcaAppList` | ✓ |
| POST | `/product/rca/app/` | 安装云应用App | `rca.InstallRcaApp` | ✓ |
| DELETE | `/product/rca/app/{id}/` | 卸载云应用App | `rca.UninstallRcaApp` | ✓ |
| GET | `/product/rca/app/{id}/` | 获取云应用App详情 | `rca.GetRcaAppDetail` | ✓ |
| PATCH | `/product/rca/app/{id}/` | 更新云应用App设定 | `rca.UpdateRcaApp` | — |
| POST | `/product/rca/app/{id}/restart` | 云应用重启App | `rca.RestartRcaApp` | ✓ |
| POST | `/product/rca/app/{id}/start` | 云应用启动App | `rca.StartRcaApp` | ✓ |
| POST | `/product/rca/app/{id}/stop` | 云应用停止App | `rca.StopRcaApp` | ✓ |
| POST | `/product/rca/app/{id}/upgrade` | 升级云应用App | `rca.UpgradeRcaApp` | — |
| PATCH | `/product/rca/app/{id}/{container_id}/` | 更新云应用App容器设定 | `rca.UpdateRcaAppContainer` | — |
| GET | `/product/rca/app/{id}/{container_id}/config_map` | 获取云应用App配置文件 | `rca.GetRcaAppConfigMap` | — |
| GET | `/product/rca/app/{id}/{container_id}/metrics` | 获取App容器的指标信息 | `rca.GetRcaAppContainerMetrics` | — |
| GET | `/product/rca/app/{id}/{container_id}/php_setting` | 云应用获取PHP相关配置 | `rca.GetRcaAppPHPSetting` | — |
| POST | `/product/rca/app/{id}/{container_id}/php_setting` | 云应用更新PHP相关配置 | `rca.UpdateRcaAppPHPSetting` | — |
| GET | `/product/rca/app/{id}/{container_id}/service/` | 云应用列出服务 | `rca.GetRcaAppServiceList` | — |
| POST | `/product/rca/app/{id}/{container_id}/service/` | 创建云应用服务 | `rca.CreateRcaAppService` | — |
| DELETE | `/product/rca/app/{id}/{container_id}/service/{service_id}` | 删除云应用服务 | `rca.DeleteRcaAppService` | — |
| PATCH | `/product/rca/app/{id}/{container_id}/service/{service_id}` | 更新云应用服务 | `rca.UpdateRcaAppService` | — |
| POST | `/product/rca/app/{id}/{container_id}/webserver_access` | 云应用web服务器更新访问设定 | `rca.UpdateRcaAppWebserverAccess` | — |
| GET | `/product/rca/appstore/` | 云应用列出App商店 | — | — |
| POST | `/product/rca/appstore/` | 云应用创建App模板 | — | — |
| GET | `/product/rca/appstore/{id}` | 云应用列出App商店应用详情 | — | — |
| DELETE | `/product/rca/appstore/{id}/` | 删除云应用App模板 | — | — |
| PATCH | `/product/rca/appstore/{id}/` | 云应用更新App模板 | — | — |
| POST | `/product/rca/appstore/{id}/import_docker` | 云应用从docker导入容器模版 | — | — |
| DELETE | `/product/rca/appstore/{id}/release` | 删除App模板版本 | — | — |
| GET | `/product/rca/appstore/{id}/release` | 云应用列出App商店应用版本详情 | — | — |
| PATCH | `/product/rca/appstore/{id}/release` | 云应用更新App模板版本 | — | — |
| POST | `/product/rca/appstore/{id}/release` | 创建App模板版本 | — | — |
| POST | `/product/rca/appstore/{id}/release/clone` | 克隆App模板版本 | — | — |
| POST | `/product/rca/appstore/{id}/release/public` | 开关App模板版本公开访问 | — | — |
| POST | `/product/rca/appstore/{id}/submit` | 云应用提交App到商店 | — | — |
| POST | `/product/rca/appstore/{id}/unsubmit` | 云应用取消提交App到商店 | — | — |
| GET | `/product/rca/project/` | 云应用列出项目 | `rca.ListRcaProjects` | ✓ |
| POST | `/product/rca/project/` | 创建云应用项目 | `rca.CreateRcaProject` | ✓ |
| GET | `/product/rca/project/eip` | 云应用项目列出IP地址 | `rca.ListRcaProjectIPs` | — |
| DELETE | `/product/rca/project/{id}/` | 销毁云应用项目 | `rca.DestroyRcaProject` | ✓ |
| GET | `/product/rca/project/{id}/` | 获取云应用项目详情 | `rca.GetRcaProjectDetail` | ✓ |
| PATCH | `/product/rca/project/{id}/backup_target` | 云应用项目设置备份目标 | `rca.SetRcaProjectBackupTarget` | — |
| POST | `/product/rca/project/{id}/disk_expand` | 云应用项目磁盘扩容 | `rca.ExpandRcaProjectDisk` | — |
| DELETE | `/product/rca/project/{id}/eip` | 云应用移除IP地址 | `rca.RemoveRcaProjectIP` | — |
| POST | `/product/rca/project/{id}/eip` | 云应用增加IP地址 | `rca.AddRcaProjectIP` | — |
| GET | `/product/rca/project/{id}/metrics` | 获取项目的指标信息 | `rca.GetRcaProjectMetrics` | — |
| PATCH | `/product/rca/project/{id}/sftp` | 云应用项目修改SFTP设置 | `rca.SetRcaProjectSFTPConfig` | — |
| GET | `/product/rca/raindrop` | 云应用获取雨点余额 | `rca.GetRcaRaindropBalance` | ✓ |
| POST | `/product/rca/raindrop` | 云应用购买雨点 | `rca.BuyRaindrop` | ✓ |
| GET | `/product/rca/raindrop/consume_log` | 云应用获取雨点消费历史 | `rca.GetRaindropConsumeLog` | ✓ |
| GET | `/product/rca/raindrop/plans` | 云应用获取雨点套餐列表 | `rca.GetRcaRaindropPlansList` | ✓ |
| GET | `/product/rca/raindrop/usage` | 云应用获取雨点余额使用情况 | `rca.GetRcaRaindropUsage` | — |
| GET | `/product/rca/region` | 云应用获取区域信息 | `rca.GetRcaRegionInfo` | ✓ |
| GET | `/product/rca/website/` | 云应用列出网站 | `rca.GetRcaWebsiteList` | ✓ |
| POST | `/product/rca/website/` | 创建云应用网站 | `rca.CreateRcaWebsite` | ✓ |
| GET | `/product/rca/website/rewrite_config` | 云应用网站获取重写配置模板 | `rca.GetRcaWebsiteRewriteConfig` | — |
| DELETE | `/product/rca/website/{id}/` | 云应用删除网站 | `rca.DeleteRcaWebsite` | ✓ |
| GET | `/product/rca/website/{id}/` | 获取云应用网站详情 | `rca.GetRcaWebsiteDetail` | ✓ |
| POST | `/product/rca/website/{id}/config/nginx` | 云应用网站更新Nginx相关配置 | `rca.UpdateRcaWebsiteNginx` | — |

## rcdn

CLI 命令：`add <domain>`, `create`, `delete <id>`, `domain`, `get <id>`, `list`, `rcdn`, `renew <id>`, `scale <id>`

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/rcdn/discount-percent` | 获取RCDN实例折扣比率 | — | — |
| GET | `/product/rcdn/domain` | 获取域名列表 | `rcdn.GetRcdnDomainList` | ✓ |
| POST | `/product/rcdn/domain` | 创建加速域名 | `rcdn.AddRcdnDomain` | ✓ |
| DELETE | `/product/rcdn/domain/{id}` | 删除加速域名 | `rcdn.DeleteRcdnDomain` | ✓ |
| GET | `/product/rcdn/domain/{id}` | 获取加速域名详情 | `rcdn.GetRcdnDomainDetail` | ✓ |
| POST | `/product/rcdn/domain/{id}/toggle_waf` | 手动开关防御 | `rcdn.ToggleRcdnDomainWaf` | — |
| GET | `/product/rcdn/domain/{id}/usage` | rcdn域名用量 | `rcdn.GetRcdnDomainUsage` | — |
| GET | `/product/rcdn/instance` | 获取RCDN实例列表 | `rcdn.GetRcdnInstanceList` | ✓ |
| POST | `/product/rcdn/instance` | 创建RCDN实例 | `rcdn.CreateRcdnInstance` | ✓ |
| GET | `/product/rcdn/instance/{id}` | 获取RCDN实例详情 | `rcdn.GetRcdnInstanceDetail` | ✓ |
| POST | `/product/rcdn/instance/{id}/domain/{domain_id}/cache_refresh` | RCDN缓存清理 | `rcdn.RefreshRcdnCache` | — |
| POST | `/product/rcdn/instance/{id}/renew` | RCDN实例续费 | `rcdn.RenewRcdnInstance` | ✓ |
| POST | `/product/rcdn/instance/{id}/renew/option` | 自动续费选项 | `rcdn.EnableRcdnInstanceAutoRenew` | — |
| POST | `/product/rcdn/instance/{id}/scale` | RCDN实例缩放 | `rcdn.ScaleRcdnInstance` | ✓ |
| POST | `/product/rcdn/instance/{id}/setting` | RCDN实例设置 | `rcdn.SetRcdnInstanceSetting` | — |
| POST | `/product/rcdn/instance/{id}/ssl_bind` | RCDNSSL绑定域名 | `rcdn.BindRcdnSSLDomains` | — |
| PATCH | `/product/rcdn/instance/{id}/tag` | 设置RCDN实例标签 | `rcdn.SetRcdnInstanceTag` | — |
| POST | `/product/rcdn/instance/{id}/toggle-extra-accounting` | 开关RCDN实例的弹性计费选项 | `rcdn.ToggleRcdnInstanceExtraAccounting` | — |
| GET | `/product/rcdn/instance/{id}/usage` | rcdn基础用量 | `rcdn.GetRcdnInstanceUsage` | — |
| GET | `/product/rcdn/plans` | 获取RCDN套餐列表 | `rcdn.GetRcdnPlanList` | — |
| GET | `/product/rcdn/price` | 获取RCDN折扣比率 | `rcdn.GetRcdnPrice` | — |
| GET | `/product/rcdn/{id}/monitor` | rcdn监控 | `rcdn.GetRcdnMonitorData` | — |

## rcs

CLI 命令：`add <id>`, `auto <id>`, `auto-renew <id>`, `backup`, `cancel <id> <backup-id>`, `change <id> <ip>`, `charge <id>`, `create`, `create <id>`, `create <id> <label>`, `delete <id>`, `delete <id> <backup-id>`, `delete <id> <rule-id>`, `discard <id> <ip>`, `edisk`, `eip`, `expand <id>`, `firewall`, `free <id>`, `get <id>`, `limit <id>`, `list`, `list <id>`, `monitor <id>`, `move <id> <rule-id>`, `nat`, `pve-address <id>`, `reboot <id>`, `reinstall <id>`, `renew <id>`, `renew-price <id>`, `reset-password <id>`, `restore <id> <backup-id>`, `server`, `set <id>`, `set-description <id> <ip> <description>`, `set-tag <id> <tag>`, `start <id>`, `stop <id>`, `traffic`, `upgrade <id>`, `vnc <id>`

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/rcs/` | 获取云服务器列表 | `rcs.GetRcsList` | ✓ |
| POST | `/product/rcs/` | 创建云服务器 | `rcs.CreateRcs` | ✓ |
| GET | `/product/rcs/discount-percent` | 获取云服务器折扣比率 | — | — |
| GET | `/product/rcs/os-templates` | 获取RCS操作系统列表 | `public.GetRcsOSList` | — |
| GET | `/product/rcs/plans` | 云服务器获取套餐列表 | — | — |
| GET | `/product/rcs/price` | 获取云服务器价格 | `rcs.GetRcsRenewPrice` | ✓ |
| GET | `/product/rcs/usage` | 获取使用情况列表 | — | — |
| GET | `/product/rcs/{id}/` | 获取RCS详情 | `rcs.GetRcsDetail` | ✓ |
| POST | `/product/rcs/{id}/backup/` | RCS创建备份 | `rcs.CreateRcsBackup` | ✓ |
| PATCH | `/product/rcs/{id}/backup/setting` | RCS设置备份选项 | `rcs.EnableRcsAutoBackup` | ✓ |
| DELETE | `/product/rcs/{id}/backup/{bid}/` | RCS删除备份 | `rcs.DeleteRcsBackup` | ✓ |
| POST | `/product/rcs/{id}/backup/{bid}/cancel` | RCS取消备份 | `rcs.CancelRcsBackup` | ✓ |
| POST | `/product/rcs/{id}/backup/{bid}/restore` | RCS还原备份 | `rcs.RestoreRcsBackup` | ✓ |
| POST | `/product/rcs/{id}/bridge_setintip` | 桥接模式下设置内网 | — | — |
| POST | `/product/rcs/{id}/changeos` | RCS重装系统 | `rcs.ReinstallRcs` | ✓ |
| POST | `/product/rcs/{id}/edisk/` | RCS管理弹性云盘 | `rcs.RcsManagesElasticCloudDisks` | ✓ |
| POST | `/product/rcs/{id}/eip/` | 创建并绑定弹性IP到RCS | `rcs.CreateAndBindElasticIpToRcs` | ✓ |
| POST | `/product/rcs/{id}/eip/change` | 更换IP | `rcs.ChangeRcsIP` | ✓ |
| POST | `/product/rcs/{id}/eip/description` | 设置IP描述 | `rcs.SetRcsEipDescription` | ✓ |
| POST | `/product/rcs/{id}/eip/discard` | 放弃IP | `rcs.DisCardRcsIP` | ✓ |
| POST | `/product/rcs/{id}/fai-send` | 发布快速app安装任务 | — | — |
| POST | `/product/rcs/{id}/firewall/mode` | 创建/设置防火墙规则 | — | — |
| GET | `/product/rcs/{id}/firewall/rule` | 获取防火墙规则列表 | `rcs.GetRcsFirewallRules` | ✓ |
| POST | `/product/rcs/{id}/firewall/rule` | 创建/设置防火墙规则 | `rcs.SetRcsFirewallRule` | ✓ |
| DELETE | `/product/rcs/{id}/firewall/rule/{ruleId}` | 删除防火墙规则 | `rcs.DeleteRcsFirewallRule` | ✓ |
| PUT | `/product/rcs/{id}/firewall/rule/{ruleId}/pos` | 移动防火墙规则优先级 | `rcs.MobileRcsFirewallRulePriority` | ✓ |
| GET | `/product/rcs/{id}/firewall/sync_time` | 获取防火墙同步开始时间 | — | — |
| POST | `/product/rcs/{id}/free` | 释放 | `rcs.FreeRcs` | ✓ |
| GET | `/product/rcs/{id}/monitor` | 获取监控数据 | `rcs.GetRcsMonitorData` | ✓ |
| DELETE | `/product/rcs/{id}/nat` | 删除NAT端口映射 | `rcs.DeleteRcsNatPortMapping` | ✓ |
| GET | `/product/rcs/{id}/nat` | 添加NAT端口映射 | `rcs.AddRcsNatPortMapping` | ✓ |
| POST | `/product/rcs/{id}/reboot` | 云服务器重启操作 | `rcs.RebootRcs` | ✓ |
| GET | `/product/rcs/{id}/renew/` | 获取续费价格 | — | — |
| POST | `/product/rcs/{id}/renew/` | 续费 | `rcs.RenewRcs` | ✓ |
| POST | `/product/rcs/{id}/renew/option` | 自动续费选项 | `rcs.EnableRcsAutoRenew` | ✓ |
| POST | `/product/rcs/{id}/reset-password` | 云服务器重置密码操作 | `rcs.ResetRcsPassword` | ✓ |
| POST | `/product/rcs/{id}/start` | 云服务器开机操作 | `rcs.StartRcs` | ✓ |
| POST | `/product/rcs/{id}/stop` | 云服务器关机操作 | `rcs.StopRcs` | ✓ |
| PATCH | `/product/rcs/{id}/tag` | 设置云服务器标签 | `rcs.SetRcsTag` | ✓ |
| POST | `/product/rcs/{id}/to-bridge` | 转成桥接 | — | — |
| POST | `/product/rcs/{id}/toggle_dgpu` | RCS充流量 | — | — |
| POST | `/product/rcs/{id}/toggle_primary_gpu` | RCS充流量 | — | — |
| POST | `/product/rcs/{id}/traffic/auto` | RCS设置自动充流量 | — | — |
| POST | `/product/rcs/{id}/traffic/charge` | RCS充流量 | `rcs.ChargeRcsTrafic` | ✓ |
| POST | `/product/rcs/{id}/traffic/limit` | RCS限流 | `rcs.LimitRcsTrafic` | ✓ |
| POST | `/product/rcs/{id}/upgrade` | 升级 | `rcs.UpgradeRcs` | ✓ |
| GET | `/product/rcs/{id}/usage` | 获取使用情况 | — | — |
| GET | `/product/rcs/{id}/vnc` | 连接VNC | `rcs.GetRcsVnc` | ✓ |
| PATCH | `/product/rcs/{id}/vnet` | 子网改名 | — | — |
| POST | `/product/rcs/{id}/vnet` | 创建虚拟机内网子网 | — | — |
| POST | `/product/rcs/{id}/webbar/charge` | RCS充流量 | — | — |
| POST | `/product/ros/{id}/free` | 释放 | — | — |

## rgs

CLI 命令：`add <id>`, `auto-renew <id>`, `backup`, `cancel <id> <backup-id>`, `change <id>`, `change <id> <ip>`, `config <id>`, `config-set <id> <json>`, `cpu-charge <id>`, `cpu-limit-mode <id>`, `create`, `create <id> <label>`, `create <id> <name>`, `create <name> <password>`, `daily-mode <id>`, `database <id>`, `delete <id> <backup-id>`, `delete <id> <nat-id>`, `delete <id> <rule-id>`, `delete <name>`, `discard <id> <ip>`, `discount-percent`, `edit <name> <password>`, `egg`, `eip`, `fai-send <id>`, `firewall`, `free <id>`, `game`, `get <id>`, `init <id>`, `k8s`, `lang <id>`, `list`, `list <id>`, `mcsm`, `mode <id> <mode>`, `monitor <id>`, `move <id> <rule-id>`, `mp`, `nat`, `os`, `pal`, `plans`, `ptero`, `rcon <id> <command>`, `reboot <id>`, `reinstall <id>`, `rename <id> <new-name>`, `renew <id>`, `renew-price <id>`, `reset-password <id>`, `restart <id>`, `restore <id> <backup-id>`, `scale <id>`, `server`, `set-description <id> <ip> <description>`, `set-int-ip <id> <ip>`, `set-tag <id> <tag>`, `setting <id>`, `sftp <id>`, `sftp-init <id>`, `start <id>`, `start-command <id> <command>`, `status <id>`, `stop <id>`, `switch-user <id>`, `sync-time <id>`, `to-bridge <id>`, `type`, `upgrade-price <id>`, `usage [id]`, `user`, `vnc <id>`, `vnet`

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/rgs-mp/` | 获取列表 | `rgs.ListRgsMp` | ✓ |
| POST | `/product/rgs-mp/` | 创建游戏云MP | `rgs.CreateRgsMp` | ✓ |
| POST | `/product/rgs-mp/{id}/renew/` | 续费游戏云MP | `rgs.RenewRgsMp` | ✓ |
| GET | `/product/rgs/` | 获取列表 | `rgs.GetRgsList` | ✓ |
| POST | `/product/rgs/` | 创建游戏云 | `rgs.CreateRgs` | ✓ |
| POST | `/product/rgs/change-egg` | RGS切换egg(游戏类型) | `rgs.ChangeRgsEgg` | ✓ |
| GET | `/product/rgs/discount-percent` | 获取游戏云折扣比率 | `rgs.GetRgsDiscountPercent` | ✓ |
| GET | `/product/rgs/egg` | 蛋(游戏)列表 | `public.GetEggList` | ✓ |
| GET | `/product/rgs/egg_server` | 服务端类型列表 | `rgs.GetRgsEggServerList` | ✓ |
| GET | `/product/rgs/egg_type` | 蛋(游戏类型)类型列表 | `public.GetEggTypeList` | ✓ |
| GET | `/product/rgs/mcsm/pal/config` | pal配置 | `rgs.GetPalConfig` | ✓ |
| POST | `/product/rgs/mcsm/pal/config` | pal配置 | `rgs.SetPalConfig` | ✓ |
| POST | `/product/rgs/mcsm/pal/init` | pal配置 | `rgs.InitPal` | ✓ |
| GET | `/product/rgs/mcsm/pal/lang` | pal配置中文 | `rgs.GetPalLang` | ✓ |
| POST | `/product/rgs/mcsm/pal/rcon` | 关闭pal | `rgs.PalRcon` | ✓ |
| POST | `/product/rgs/mcsm/pal/restart` | 关闭pal | `rgs.RestartPal` | ✓ |
| POST | `/product/rgs/mcsm/pal/stop` | 关闭pal | `rgs.StopPal` | ✓ |
| GET | `/product/rgs/mcsm/panel_user/` | 翼龙面板用户列表 | `rgs.GetMcsmUserList` | ✓ |
| PATCH | `/product/rgs/mcsm/panel_user/` | 编辑面板用户 | `rgs.EditMcsmUser` | ✓ |
| POST | `/product/rgs/mcsm/panel_user/` | 创建面板用户 | `rgs.CreateMcsmUser` | ✓ |
| DELETE | `/product/rgs/mcsm/panel_user/{name}` | 删除面板用户 | `rgs.DeleteMcsmUser` | ✓ |
| POST | `/product/rgs/mcsm/sftp_init` | 初始化/刷新sftp功能 | `rgs.McsmSftpInit` | ✓ |
| POST | `/product/rgs/mcsm/start/` | 开服 | `rgs.StartMcsmInstance` | ✓ |
| GET | `/product/rgs/mcsm/status` | 实例信息和状态 | `rgs.GetMcsmStatus` | ✓ |
| GET | `/product/rgs/os-templates` | 系统列表 | `public.GetRgsOSList` | ✓ |
| GET | `/product/rgs/plans` | 获取套餐列表 | `rgs.GetRgsPlanList` | ✓ |
| GET | `/product/rgs/price` | 获取游戏云价格 | `rgs.GetRgsUpgradePrice` | ✓ |
| GET | `/product/rgs/ptero/panel_user/` | 翼龙面板用户列表 | `rgs.GetPteroUserList` | ✓ |
| PATCH | `/product/rgs/ptero/panel_user/` | 编辑面板用户 | `rgs.EditPteroUser` | ✓ |
| POST | `/product/rgs/ptero/panel_user/` | 创建翼龙面板用户 | `rgs.CreatePteroUser` | ✓ |
| DELETE | `/product/rgs/ptero/panel_user/{name}` | 删除面板用户 | `rgs.DeletePteroUser` | ✓ |
| POST | `/product/rgs/switch-user` | RGS切换面板用户 | `rgs.SwitchRgsPanelUser` | ✓ |
| GET | `/product/rgs/usage` | 获取使用情况列表 | `rgs.GetRgsUsageList` | ✓ |
| GET | `/product/rgs/{id}/` | 获取RGS详情 | `rgs.GetRgsDetail` | ✓ |
| POST | `/product/rgs/{id}/backup/` | RGS创建备份 | `rgs.CreateRgsBackup` | ✓ |
| PATCH | `/product/rgs/{id}/backup/setting` | RGS设置备份选项 | `rgs.EnableRgsAutoBackup` | ✓ |
| DELETE | `/product/rgs/{id}/backup/{bid}/` | RGS删除备份 | `rgs.DeleteRgsBackup` | ✓ |
| POST | `/product/rgs/{id}/backup/{bid}/cancel` | RGS取消备份 | `rgs.CancelRgsBackup` | ✓ |
| POST | `/product/rgs/{id}/backup/{bid}/restore` | RGS还原备份 | `rgs.RestoreRgsBackup` | ✓ |
| POST | `/product/rgs/{id}/bridge_setintip` | 桥接模式下设置内网 | `rgs.RgsBridgeSetIntIP` | ✓ |
| POST | `/product/rgs/{id}/changeos` | RGS重装系统 | `rgs.ReinstallRgs` | ✓ |
| POST | `/product/rgs/{id}/cpu-charge` | cpu充电 | `rgs.ChargeRgsCPU` | ✓ |
| POST | `/product/rgs/{id}/cpu-limit-mode` | 游戏云限制模式(是否用余额结算)切换 | `rgs.SwitchRgsBalanceMode` | ✓ |
| POST | `/product/rgs/{id}/daily-mode` | 游戏云日付模式开关 | `rgs.SwitchRgsDailyMode` | ✓ |
| POST | `/product/rgs/{id}/eip` | 创建并绑定弹性IP到RGS | `rgs.CreateAndBindElasticIpToRgs` | ✓ |
| POST | `/product/rgs/{id}/eip/change` | 更换IP | `rgs.ChangeRgsIP` | ✓ |
| POST | `/product/rgs/{id}/eip/description` | 设置IP描述 | `rgs.SetRgsEipDescription` | ✓ |
| POST | `/product/rgs/{id}/eip/discard` | 放弃IP | `rgs.DisCardRgsIP` | ✓ |
| POST | `/product/rgs/{id}/fai-send` | 发布快速app安装任务 | `rgs.SendRgsFaiTask` | ✓ |
| POST | `/product/rgs/{id}/firewall/mode` | 创建/设置防火墙规则 | `rgs.SetRgsFirewallMode` | ✓ |
| GET | `/product/rgs/{id}/firewall/rule` | 获取防火墙规则列表 | `rgs.GetRgsFirewallRules` | ✓ |
| POST | `/product/rgs/{id}/firewall/rule` | 创建/设置防火墙规则 | `rgs.SetRgsFirewallRule` | ✓ |
| DELETE | `/product/rgs/{id}/firewall/rule/{ruleId}` | 删除防火墙规则 | `rgs.DeleteRgsFirewallRule` | ✓ |
| PUT | `/product/rgs/{id}/firewall/rule/{ruleId}/pos` | 移动防火墙规则优先级 | `rgs.MobileRgsFirewallRulePriority` | ✓ |
| GET | `/product/rgs/{id}/firewall/sync_time` | 获取防火墙同步开始时间 | `rgs.GetRgsFirewallSyncTime` | ✓ |
| POST | `/product/rgs/{id}/free` | 释放游戏云 | `rgs.FreeRgs` | ✓ |
| PATCH | `/product/rgs/{id}/k8s-panel/database` | K8S面板修改数据库设置 | `rgs.SetK8SPanelDatabase` | ✓ |
| POST | `/product/rgs/{id}/k8s-panel/set-start-command` | 游戏云设置启动命令（仅支持雨云面板） | `rgs.SetK8SPanelStartCommand` | ✓ |
| PATCH | `/product/rgs/{id}/k8s-panel/sftp` | K8S面板修改SFTP设置 | `rgs.SetK8SPanelSFTP` | ✓ |
| POST | `/product/rgs/{id}/monitor` | 获取监控数据 | `rgs.GetRgsMonitorData` | ✓ |
| DELETE | `/product/rgs/{id}/nat` | 删除NAT端口映射 | `rgs.DeleteRgsNatPortMapping` | ✓ |
| POST | `/product/rgs/{id}/nat` | 添加NAT端口映射 | `rgs.AddRgsNatPortMapping` | ✓ |
| POST | `/product/rgs/{id}/reboot` | 游戏云重启操作 | `rgs.RebootRgs` | ✓ |
| POST | `/product/rgs/{id}/renew/` | 续费 | `rgs.RenewRgs` | ✓ |
| POST | `/product/rgs/{id}/renew/option` | 自动续费选项 | `rgs.EnableRgsAutoRenew` | ✓ |
| POST | `/product/rgs/{id}/reset-password` | 游戏云重置密码操作 | `rgs.ResetRgsPassword` | ✓ |
| POST | `/product/rgs/{id}/scale` | 升级 | `rgs.ScaleRgs` | ✓ |
| POST | `/product/rgs/{id}/start` | 游戏云开机操作 | `rgs.StartRgs` | ✓ |
| POST | `/product/rgs/{id}/stop` | 游戏云关机操作 | `rgs.StopRgs` | ✓ |
| PATCH | `/product/rgs/{id}/tag` | 设置游戏云标签 | `rgs.SetRgsTag` | ✓ |
| POST | `/product/rgs/{id}/to-bridge` | 转成桥接 | `rgs.RgsToBridge` | ✓ |
| GET | `/product/rgs/{id}/usage` | 获取使用情况 | `rgs.GetRgsUsage` | ✓ |
| GET | `/product/rgs/{id}/vnc` | 连接VNC | `rgs.GetRgsVnc` | ✓ |
| PATCH | `/product/rgs/{id}/vnet` | 子网改名 | `rgs.RenameRgsVnet` | ✓ |
| POST | `/product/rgs/{id}/vnet` | 创建虚拟机内网子网 | `rgs.CreateRgsVnet` | ✓ |

## ros

CLI 命令：`bucket`, `create <instance-id> <name>`, `list`, `list <instance-id>`, `storage`

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/ros/bucket` | 获取存储桶列表 | `ros.GetRosBucketList` | — |
| POST | `/product/ros/bucket` | 创建对象存储桶 | `ros.CreateRosBucket` | ✓ |
| DELETE | `/product/ros/bucket/{id}` | 删除对象存储桶 | `ros.DeleteRosBucket` | — |
| GET | `/product/ros/bucket/{id}` | 获取对象存储桶详情 | `ros.GetRosBucketListByInstance` | ✓ |
| GET | `/product/ros/bucket/{id}/lifecycle` | 查询生命周期规则列表 | — | — |
| POST | `/product/ros/bucket/{id}/lifecycle` | 创建生命周期规则 | — | — |
| DELETE | `/product/ros/bucket/{id}/lifecycle/{rule_id}` | 删除生命周期规则 | — | — |
| GET | `/product/ros/bucket/{id}/monitor` | 实例监控 | `ros.GetRosBucketMonitorData` | — |
| GET | `/product/ros/bucket/{id}/offline-download` | 列出离线下载任务 | — | — |
| POST | `/product/ros/bucket/{id}/offline-download` | 创建离线下载任务 | — | — |
| DELETE | `/product/ros/bucket/{id}/offline-download/{task_id}` | 取消离线下载任务 | — | — |
| GET | `/product/ros/bucket/{id}/offline-download/{task_id}` | 查询离线下载任务详情 | — | — |
| PATCH | `/product/ros/bucket/{id}/proxy` | 修改存储桶Proxy设置 | `ros.ModifyRosBucketProxySettings` | — |
| POST | `/product/ros/bucket/{id}/regenerate-keys` | 对象存储桶重新生成密钥 | `ros.ReGenerateRosBucketKeys` | — |
| GET | `/product/ros/bucket/{id}/request-log` | 获取请求日志生成任务列表 | — | — |
| POST | `/product/ros/bucket/{id}/request-log` | 生成请求日志 | — | — |
| GET | `/product/ros/bucket/{id}/request-log/{task_id}/link` | 获取请求日志下载链接 | — | — |
| POST | `/product/ros/bucket/{id}/request-log/{task_id}/retrigger` | 重新触发请求日志生成任务 | — | — |
| GET | `/product/ros/bucket/{id}/statistics` | 获取存储桶热点数据 | — | — |
| DELETE | `/product/ros/bucket/{id}/sync` | 取消主动同步 | — | — |
| GET | `/product/ros/bucket/{id}/sync` | 查询主动同步状态 | — | — |
| POST | `/product/ros/bucket/{id}/sync` | 创建或重新运行主动同步 | — | — |
| POST | `/product/ros/bucket/{id}/toggle-public-access` | 开关对象存储桶匿名访问 | `ros.SetRosBucketPublicAccess` | — |
| GET | `/product/ros/discount-percent` | 获取对象存储实例折扣比率 | — | — |
| GET | `/product/ros/instance` | 获取对象存储实例列表 | `ros.GetRosInstanceList` | ✓ |
| POST | `/product/ros/instance` | 创建对象存储实例 | `ros.CreateRosInstance` | — |
| GET | `/product/ros/instance/{id}` | 获取对象存储实例详情 | `ros.GetRosInstanceDetail` | — |
| POST | `/product/ros/instance/{id}/regenerate-keys` | 对象存储实例重新生成密钥 | `ros.ReGenerateRosInstanceKeys` | — |
| POST | `/product/ros/instance/{id}/renew` | ROS实例续费 | `ros.RenewRosInstance` | — |
| POST | `/product/ros/instance/{id}/renew/option` | 自动续费选项 | `ros.SetRosInstanceAutoRenewOption` | — |
| POST | `/product/ros/instance/{id}/scale` | ROS实例缩放 | `ros.ScaleRosInstance` | — |
| PATCH | `/product/ros/instance/{id}/tag` | 设置对象存储实例标签 | `ros.SetRosInstanceTags` | — |
| POST | `/product/ros/instance/{id}/toggle-extra-accounting` | 开关对象存储实例的弹性计费选项 | `ros.ToggleRosInstanceExtraAccounting` | — |
| POST | `/product/ros/instance/{id}/toggle-public-access` | 开关对象存储实例的匿名访问 | `ros.SetRosInstancePublicAccess` | — |
| GET | `/product/ros/plans` | 获取对象存储套餐列表 | `public.GetRosPlanList` | — |
| GET | `/product/ros/price` | 获取云服务器折扣比率 | — | — |

## rvh

CLI 命令：`backup`, `bind <id> <domain>`, `create`, `create <id> <label>`, `delete <id>`, `delete <id> <backup-id>`, `domain`, `get <id>`, `list`, `plans`, `renew <id>`, `restore <id> <backup-id>`, `rvh`, `unbind <id> <domain>`

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/rvh/` | 获取虚拟主机列表 | `rvh.GetRvhList` | ✓ |
| POST | `/product/rvh/` | 创建虚拟主机 | `rvh.CreateRvh` | ✓ |
| GET | `/product/rvh/plans` | 获取虚拟主机套餐列表 | `rvh.GetRvhPlanList` | ✓ |
| GET | `/product/rvh/price` | 获取虚拟主机折扣比率 | `rvh.GetRvhPrice` | — |
| PATCH | `/product/rvh/tag` | 设置虚拟主机标签 | `rvh.SetRvhTag` | — |
| GET | `/product/rvh/{id}/` | 获取RVH虚拟主机详情 | `rvh.GetRvhDetail` | ✓ |
| POST | `/product/rvh/{id}/backup/` | RVH创建备份 | `rvh.CreateRvhBackup` | ✓ |
| PATCH | `/product/rvh/{id}/backup/setting` | RVH设置备份选项 | `rvh.SetRvhBackupSetting` | — |
| DELETE | `/product/rvh/{id}/backup/{bid}/` | RVH删除备份 | `rvh.DeleteRvhBackup` | ✓ |
| POST | `/product/rvh/{id}/backup/{bid}/restore` | RVH还原备份 | `rvh.RestoreRvhBackup` | ✓ |
| POST | `/product/rvh/{id}/bt/attach-dedip` | 附加独立IP地址 | `rvh.RvhBtAttachDedip` | — |
| POST | `/product/rvh/{id}/bt/fix` | RVH宝塔主机修复操作 | `rvh.RvhBtFix` | — |
| POST | `/product/rvh/{id}/bt/reboot` | RVH宝塔主机重启操作 | `rvh.RvhBtReboot` | — |
| DELETE | `/product/rvh/{id}/domain/` | RVH域名解绑 | `rvh.UnbindRvhDomain` | ✓ |
| POST | `/product/rvh/{id}/domain/` | RVH域名绑定 | `rvh.BindRvhDomain` | ✓ |
| POST | `/product/rvh/{id}/ep/reset-pass` | RVH EP主机重置密码操作 | `rvh.ResetRvhEpPassword` | — |
| POST | `/product/rvh/{id}/firewall/option` | RVH防火墙设置选项 | `rvh.SetRvhFirewallOption` | — |
| POST | `/product/rvh/{id}/firewall/rule` | RVH防火墙设置规则 | `rvh.SetRvhFirewallRule` | — |
| POST | `/product/rvh/{id}/free` | 释放 | `rvh.FreeRvh` | ✓ |
| POST | `/product/rvh/{id}/maintenance-mode` | RVH设置维护模式 | `rvh.SetRvhMaintenanceMode` | — |
| POST | `/product/rvh/{id}/reinstall` | RVH重装操作 | `rvh.ReinstallRvh` | — |
| GET | `/product/rvh/{id}/renew/` | 获取虚拟主机折扣比率 | `rvh.GetRvhRenewPrice` | — |
| POST | `/product/rvh/{id}/renew/` | 续费 | `rvh.RenewRvh` | ✓ |
| POST | `/product/rvh/{id}/renew/option` | 自动续费选项 | `rvh.EnableRvhAutoRenew` | — |
| POST | `/product/rvh/{id}/upgrade/` | 升级 | `rvh.UpgradeRvh` | — |

## ssl

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/sslcenter/` | SSL证书获取列表 | `ssl.GetSSLCertificateList` | ✓ |
| POST | `/product/sslcenter/` | SSL证书上传操作 | `ssl.UploadSSLCertificate` | ✓ |
| POST | `/product/sslcenter/cert/order` | 创建SSL证书申请 | `ssl.ApplyFreeSSLCertificate` | ✓ |
| POST | `/product/sslcenter/cert/order_verify` | 验证SSL证书申请 | `ssl.VerifyFreeSSLCertificate` | ✓ |
| GET | `/product/sslcenter/cert/orders` | 获取SSL证书申请列表 | `ssl.GetSSLCertApplyList` | ✓ |
| GET | `/product/sslcenter/order` | 获取SSL证书订单列表 | `ssl.GetSSLOrderList` | ✓ |
| POST | `/product/sslcenter/order` | 创建SSL证书订单 | `ssl.CreateSSLOrder` | ✓ |
| GET | `/product/sslcenter/order/{id}` | 获取SSL证书订单信息 | `ssl.GetSSLOrderDetail` | ✓ |
| POST | `/product/sslcenter/order/{id}/assign` | 将SSL证书添加到证书列表 | `ssl.AssignSSLOrder` | ✓ |
| GET | `/product/sslcenter/order/{id}/cert` | 获取SSL证书 | `ssl.GetSSLOrderCert` | ✓ |
| POST | `/product/sslcenter/order/{id}/description` | 更新SSL订单描述 | `ssl.UpdateSSLOrderDescription` | ✓ |
| POST | `/product/sslcenter/order/{id}/revoke` | 申请吊销SSL证书 | `ssl.RevokeSSLOrder` | ✓ |
| POST | `/product/sslcenter/order/{id}/verify` | 验证SSL证书订单 | `ssl.VerifySSLOrder` | ✓ |
| POST | `/product/sslcenter/price` | 获取SSL证书订单价格 | `ssl.GetSSLOrderPrice` | ✓ |
| GET | `/product/sslcenter/product` | 获取SSL证书产品列表 | `ssl.GetSSLProductList` | ✓ |
| DELETE | `/product/sslcenter/{id}` | SSL证书删除操作 | `ssl.DeleteSsl` | ✓ |
| GET | `/product/sslcenter/{id}` | SSL证书查看操作 | `ssl.GetSslDetail` | ✓ |
| PUT | `/product/sslcenter/{id}` | SSL证书替换操作 | `ssl.ReplaceSsl` | ✓ |
