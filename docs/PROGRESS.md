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

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/domain/` | 列出域名列表 | `domain.GetDomainList` | typed | ✓ |
| DELETE | `/product/domain/certify` | 删除域名认证 | — |  | — |
| GET | `/product/domain/certify` | 获取已验证域名列表 | `domain.GetVerifiedDomainList` | typed | — |
| POST | `/product/domain/certify` | 添加域名认证 | `domain.AddDomainCertify` | typed | — |
| POST | `/product/domain/certify/verify` | 域名认证校验 | `domain.VerifyDomainCertify` | typed | — |
| POST | `/product/domain/check` | 检查域名能否注册 | — |  | — |
| DELETE | `/product/domain/free_subdomain` | 删除免费二级域名 | — |  | — |
| GET | `/product/domain/free_subdomain` | 获取免费二级域名列表 | — |  | — |
| POST | `/product/domain/free_subdomain` | 创建免费二级域名 | — |  | — |
| POST | `/product/domain/free_subdomain/proxy` | 修改免费二级域名的CDN设置 | — |  | — |
| GET | `/product/domain/free_subdomain/usable` | 获取可用的免费域名列表 | — |  | — |
| POST | `/product/domain/register` | 域名注册 | — |  | — |
| PATCH | `/product/domain/template` | 编辑域名模板 | `domain.EditDomainTemplate` | typed | — |
| DELETE | `/product/domain/template/` | 删除域名信息模板 | `domain.DeleteDomainTemplate` | typed | — |
| GET | `/product/domain/template/` | 查询域名模板列表 | `domain.GetDomainTemplateList` | typed | — |
| GET | `/product/domain/template/detail/` | 获取域名模板详情 | `domain.GetDomainTemplateDetail` | typed | — |
| GET | `/product/domain/whitelist` | 获取域名白名单列表 | `domain.GetDomainWhiteList` | typed | — |
| POST | `/product/domain/whitelist` | 添加域名白名单 | — |  | — |
| GET | `/product/domain/whois` | 获取域名whois信息 | `domain.GetDomainWhoisInfo` | typed | — |
| GET | `/product/domain/{id}` | 获取域名详情 | — |  | — |
| GET | `/product/domain/{id}/cert` | 下载域名证书 | — |  | — |
| PATCH | `/product/domain/{id}/dns` | 修改域名DNS解析 | `domain.UpdateDomainDNSRecord` | typed | — |
| POST | `/product/domain/{id}/dns` | 添加域名DNS解析 | `domain.AddDomainDNSRecord` | typed | ✓ |
| DELETE | `/product/domain/{id}/dns/` | 删除域名DNS解析 | `domain.DeleteDomainDNSRecord` | typed | ✓ |
| GET | `/product/domain/{id}/dns/` | 获取域名DNS解析记录列表 | `domain.GetDomainDNSRecordList` | typed | ✓ |
| GET | `/product/domain/{id}/dnssec` | 获取域名DNSSEC详情 | — |  | — |
| POST | `/product/domain/{id}/dnssec` | 添加域名DNSSEC | `domain.AddDomainDNSSEC` | typed | — |
| POST | `/product/domain/{id}/dnssec/delete` | 删除域名DNSSEC | `domain.DeleteDomainDNSSEC` | typed | — |
| POST | `/product/domain/{id}/dnssec/sync` | 同步域名DNSSEC | `domain.SyncDomainDNSSEC` | typed | — |
| PUT | `/product/domain/{id}/lock/disable` | 关闭域名锁定 | `domain.UnlockDomain` | typed | — |
| PUT | `/product/domain/{id}/lock/enable` | 开启域名锁定 | `domain.LockDomain` | typed | — |
| POST | `/product/domain/{id}/nameservers` | 修改域名NS服务器 | `domain.UpdateDomainNS` | typed | — |
| POST | `/product/domain/{id}/nameservers/reset` | 重置域名NS服务器 | `domain.ResetDomainNS` | typed | — |
| GET | `/product/domain/{id}/password` | 获取域名管理密码 | — |  | — |
| POST | `/product/domain/{id}/password` | 更新域名管理密码 | `domain.UpdateDomainPassword` | typed | — |
| POST | `/product/domain/{id}/renew` | 续费域名 | `domain.RenewDomain` | typed | — |
| GET | `/product/domain/{id}/renew-price` | 获取域名续费价格 | — |  | — |
| POST | `/product/domain/{id}/transfer` | 域名过户 | `domain.TransferDomain` | typed | — |

## public

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/app_config` | 获取页面信息 | `public.GetAppConfig` | typed | ✓ |
| GET | `/discourse` | 获取论坛数据 | `public.GetDiscourse` | typed | ✓ |
| GET | `/news` | 获取论坛公告 | `public.GetNews` | typed | ✓ |
| GET | `/short_params` | 查询参数缩短的具体base64 | — |  | — |
| POST | `/short_params` | 创建参数缩短 | — |  | — |
| GET | `/status` | 获取节点网络状态 | `public.GetStatus` | typed | ✓ |

## rbm

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/rbm` | 列出RBM实例 | — |  | — |
| POST | `/product/rbm/` | 创建RBM实例 | — |  | — |
| GET | `/product/rbm/price` | 价格计算 | — |  | — |
| GET | `/product/rbm/usage` | 获取使用情况列表 | — |  | — |
| POST | `/product/rbm/{id}/bios-flash` | 裸金属刷bios | — |  | — |
| POST | `/product/rbm/{id}/changeos` | RBM实例更换系统 | `rbm.ChangeRBMOS` | typed | — |
| POST | `/product/rbm/{id}/commission` | RBM清点配置 | — |  | — |
| POST | `/product/rbm/{id}/eip/` | 创建并绑定弹性IP到RBM | `rbm.AssociateEIP` | typed | — |
| POST | `/product/rbm/{id}/eip/change` | 更换IP | — |  | — |
| POST | `/product/rbm/{id}/eip/description` | 设置IP描述 | `rbm.SetIPDescription` | typed | — |
| POST | `/product/rbm/{id}/eip/discard` | 放弃IP | `rbm.ReleaseIP` | typed | — |
| POST | `/product/rbm/{id}/free` | 释放 | — |  | — |
| POST | `/product/rbm/{id}/kvm-proxy` | RBM实例启动KVM代理 | `rbm.StartKVMAgent` | typed | — |
| POST | `/product/rbm/{id}/kvm-reboot` | RBM重新启动KVM | `rbm.RestartKVM` | typed | — |
| GET | `/product/rbm/{id}/monitor` | 获取监控数据 | — |  | — |
| POST | `/product/rbm/{id}/poweroff` | RBM实例关机 | `rbm.ShutdownRBM` | typed | — |
| POST | `/product/rbm/{id}/poweron` | RBM实例开机 | `rbm.StartRBM` | typed | — |
| POST | `/product/rbm/{id}/rescue` | RBM救援模式切换 | — |  | — |
| POST | `/product/rbm/{id}/reset-password` | 重置RBM实例IPMI密码 | — |  | — |
| POST | `/product/rbm/{id}/traffic/auto` | 充流量 | — |  | — |
| POST | `/product/rbm/{id}/traffic/charge` | 充流量 | `rbm.ChargeTraffic` | typed | — |
| POST | `/product/rbm/{id}/traffic/limit` | 限流 | — |  | — |
| POST | `/product/rbm/{id}/traffic/switch` | 切换流量套餐 | — |  | — |

## rca

CLI 命令：`app`, `balance`, `buy`, `create`, `delete <id>`, `get <id>`, `list`, `list <project-id>`, `log`, `plans`, `project`, `raindrop`, `rca`, `regions`, `restart <id>`, `start <id>`, `stop <id>`, `website`

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/rca/app/` | 云应用列出App | `rca.GetRcaAppList` | typed | ✓ |
| POST | `/product/rca/app/` | 安装云应用App | `rca.InstallRcaApp` | typed | ✓ |
| DELETE | `/product/rca/app/{id}/` | 卸载云应用App | `rca.UninstallRcaApp` | typed | ✓ |
| GET | `/product/rca/app/{id}/` | 获取云应用App详情 | `rca.GetRcaAppDetail` | typed | ✓ |
| PATCH | `/product/rca/app/{id}/` | 更新云应用App设定 | `rca.UpdateRcaApp` | typed | — |
| POST | `/product/rca/app/{id}/restart` | 云应用重启App | `rca.RestartRcaApp` | typed | ✓ |
| POST | `/product/rca/app/{id}/start` | 云应用启动App | `rca.StartRcaApp` | typed | ✓ |
| POST | `/product/rca/app/{id}/stop` | 云应用停止App | `rca.StopRcaApp` | typed | ✓ |
| POST | `/product/rca/app/{id}/upgrade` | 升级云应用App | `rca.UpgradeRcaApp` | typed | — |
| PATCH | `/product/rca/app/{id}/{container_id}/` | 更新云应用App容器设定 | `rca.UpdateRcaAppContainer` | typed | — |
| GET | `/product/rca/app/{id}/{container_id}/config_map` | 获取云应用App配置文件 | `rca.GetRcaAppConfigMap` | typed | — |
| GET | `/product/rca/app/{id}/{container_id}/metrics` | 获取App容器的指标信息 | `rca.GetRcaAppContainerMetrics` | typed | — |
| GET | `/product/rca/app/{id}/{container_id}/php_setting` | 云应用获取PHP相关配置 | `rca.GetRcaAppPHPSetting` | typed | — |
| POST | `/product/rca/app/{id}/{container_id}/php_setting` | 云应用更新PHP相关配置 | `rca.UpdateRcaAppPHPSetting` | typed | — |
| GET | `/product/rca/app/{id}/{container_id}/service/` | 云应用列出服务 | `rca.GetRcaAppServiceList` | typed | — |
| POST | `/product/rca/app/{id}/{container_id}/service/` | 创建云应用服务 | `rca.CreateRcaAppService` | typed | — |
| DELETE | `/product/rca/app/{id}/{container_id}/service/{service_id}` | 删除云应用服务 | `rca.DeleteRcaAppService` | typed | — |
| PATCH | `/product/rca/app/{id}/{container_id}/service/{service_id}` | 更新云应用服务 | `rca.UpdateRcaAppService` | typed | — |
| POST | `/product/rca/app/{id}/{container_id}/webserver_access` | 云应用web服务器更新访问设定 | `rca.UpdateRcaAppWebserverAccess` | typed | — |
| GET | `/product/rca/appstore/` | 云应用列出App商店 | — |  | — |
| POST | `/product/rca/appstore/` | 云应用创建App模板 | — |  | — |
| GET | `/product/rca/appstore/{id}` | 云应用列出App商店应用详情 | — |  | — |
| DELETE | `/product/rca/appstore/{id}/` | 删除云应用App模板 | — |  | — |
| PATCH | `/product/rca/appstore/{id}/` | 云应用更新App模板 | — |  | — |
| POST | `/product/rca/appstore/{id}/import_docker` | 云应用从docker导入容器模版 | — |  | — |
| DELETE | `/product/rca/appstore/{id}/release` | 删除App模板版本 | — |  | — |
| GET | `/product/rca/appstore/{id}/release` | 云应用列出App商店应用版本详情 | — |  | — |
| PATCH | `/product/rca/appstore/{id}/release` | 云应用更新App模板版本 | — |  | — |
| POST | `/product/rca/appstore/{id}/release` | 创建App模板版本 | — |  | — |
| POST | `/product/rca/appstore/{id}/release/clone` | 克隆App模板版本 | — |  | — |
| POST | `/product/rca/appstore/{id}/release/public` | 开关App模板版本公开访问 | — |  | — |
| POST | `/product/rca/appstore/{id}/submit` | 云应用提交App到商店 | — |  | — |
| POST | `/product/rca/appstore/{id}/unsubmit` | 云应用取消提交App到商店 | — |  | — |
| GET | `/product/rca/project/` | 云应用列出项目 | `rca.ListRcaProjects` | typed | ✓ |
| POST | `/product/rca/project/` | 创建云应用项目 | `rca.CreateRcaProject` | typed | ✓ |
| GET | `/product/rca/project/eip` | 云应用项目列出IP地址 | `rca.ListRcaProjectIPs` | typed | — |
| DELETE | `/product/rca/project/{id}/` | 销毁云应用项目 | `rca.DestroyRcaProject` | typed | ✓ |
| GET | `/product/rca/project/{id}/` | 获取云应用项目详情 | `rca.GetRcaProjectDetail` | typed | ✓ |
| PATCH | `/product/rca/project/{id}/backup_target` | 云应用项目设置备份目标 | `rca.SetRcaProjectBackupTarget` | typed | — |
| POST | `/product/rca/project/{id}/disk_expand` | 云应用项目磁盘扩容 | `rca.ExpandRcaProjectDisk` | typed | — |
| DELETE | `/product/rca/project/{id}/eip` | 云应用移除IP地址 | `rca.RemoveRcaProjectIP` | typed | — |
| POST | `/product/rca/project/{id}/eip` | 云应用增加IP地址 | `rca.AddRcaProjectIP` | typed | — |
| GET | `/product/rca/project/{id}/metrics` | 获取项目的指标信息 | `rca.GetRcaProjectMetrics` | typed | — |
| PATCH | `/product/rca/project/{id}/sftp` | 云应用项目修改SFTP设置 | `rca.SetRcaProjectSFTPConfig` | typed | — |
| GET | `/product/rca/raindrop` | 云应用获取雨点余额 | `rca.GetRcaRaindropBalance` | typed | ✓ |
| POST | `/product/rca/raindrop` | 云应用购买雨点 | `rca.BuyRaindrop` | typed | ✓ |
| GET | `/product/rca/raindrop/consume_log` | 云应用获取雨点消费历史 | `rca.GetRaindropConsumeLog` | typed | ✓ |
| GET | `/product/rca/raindrop/plans` | 云应用获取雨点套餐列表 | `rca.GetRcaRaindropPlansList` | typed | ✓ |
| GET | `/product/rca/raindrop/usage` | 云应用获取雨点余额使用情况 | `rca.GetRcaRaindropUsage` | typed | — |
| GET | `/product/rca/region` | 云应用获取区域信息 | `rca.GetRcaRegionInfo` | typed | ✓ |
| GET | `/product/rca/website/` | 云应用列出网站 | `rca.GetRcaWebsiteList` | typed | ✓ |
| POST | `/product/rca/website/` | 创建云应用网站 | `rca.CreateRcaWebsite` | typed | ✓ |
| GET | `/product/rca/website/rewrite_config` | 云应用网站获取重写配置模板 | `rca.GetRcaWebsiteRewriteConfig` | typed | — |
| DELETE | `/product/rca/website/{id}/` | 云应用删除网站 | `rca.DeleteRcaWebsite` | typed | ✓ |
| GET | `/product/rca/website/{id}/` | 获取云应用网站详情 | `rca.GetRcaWebsiteDetail` | typed | ✓ |
| POST | `/product/rca/website/{id}/config/nginx` | 云应用网站更新Nginx相关配置 | `rca.UpdateRcaWebsiteNginx` | typed | — |

## rcdn

CLI 命令：`add <domain>`, `create`, `delete <id>`, `domain`, `get <id>`, `list`, `rcdn`, `renew <id>`, `scale <id>`

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/rcdn/discount-percent` | 获取RCDN实例折扣比率 | — |  | — |
| GET | `/product/rcdn/domain` | 获取域名列表 | `rcdn.GetRcdnDomainList` | typed | ✓ |
| POST | `/product/rcdn/domain` | 创建加速域名 | `rcdn.AddRcdnDomain` | typed | ✓ |
| DELETE | `/product/rcdn/domain/{id}` | 删除加速域名 | `rcdn.DeleteRcdnDomain` | typed | ✓ |
| GET | `/product/rcdn/domain/{id}` | 获取加速域名详情 | `rcdn.GetRcdnDomainDetail` | typed | ✓ |
| POST | `/product/rcdn/domain/{id}/toggle_waf` | 手动开关防御 | `rcdn.ToggleRcdnDomainWaf` | typed | — |
| GET | `/product/rcdn/domain/{id}/usage` | rcdn域名用量 | `rcdn.GetRcdnDomainUsage` | typed | — |
| GET | `/product/rcdn/instance` | 获取RCDN实例列表 | `rcdn.GetRcdnInstanceList` | typed | ✓ |
| POST | `/product/rcdn/instance` | 创建RCDN实例 | `rcdn.CreateRcdnInstance` | typed | ✓ |
| GET | `/product/rcdn/instance/{id}` | 获取RCDN实例详情 | `rcdn.GetRcdnInstanceDetail` | typed | ✓ |
| POST | `/product/rcdn/instance/{id}/domain/{domain_id}/cache_refresh` | RCDN缓存清理 | `rcdn.RefreshRcdnCache` | typed | — |
| POST | `/product/rcdn/instance/{id}/renew` | RCDN实例续费 | `rcdn.RenewRcdnInstance` | typed | ✓ |
| POST | `/product/rcdn/instance/{id}/renew/option` | 自动续费选项 | `rcdn.EnableRcdnInstanceAutoRenew` | typed | — |
| POST | `/product/rcdn/instance/{id}/scale` | RCDN实例缩放 | `rcdn.ScaleRcdnInstance` | typed | ✓ |
| POST | `/product/rcdn/instance/{id}/setting` | RCDN实例设置 | `rcdn.SetRcdnInstanceSetting` | typed | — |
| POST | `/product/rcdn/instance/{id}/ssl_bind` | RCDNSSL绑定域名 | `rcdn.BindRcdnSSLDomains` | typed | — |
| PATCH | `/product/rcdn/instance/{id}/tag` | 设置RCDN实例标签 | `rcdn.SetRcdnInstanceTag` | typed | — |
| POST | `/product/rcdn/instance/{id}/toggle-extra-accounting` | 开关RCDN实例的弹性计费选项 | `rcdn.ToggleRcdnInstanceExtraAccounting` | typed | — |
| GET | `/product/rcdn/instance/{id}/usage` | rcdn基础用量 | `rcdn.GetRcdnInstanceUsage` | typed | — |
| GET | `/product/rcdn/plans` | 获取RCDN套餐列表 | `rcdn.GetRcdnPlanList` | typed | — |
| GET | `/product/rcdn/price` | 获取RCDN折扣比率 | `rcdn.GetRcdnPrice` | typed | — |
| GET | `/product/rcdn/{id}/monitor` | rcdn监控 | `rcdn.GetRcdnMonitorData` | typed | — |

## rcs

CLI 命令：`add <id>`, `auto <id>`, `auto-renew <id>`, `backup`, `cancel <id> <backup-id>`, `change <id> <ip>`, `charge <id>`, `create`, `create <id>`, `create <id> <label>`, `delete <id>`, `delete <id> <backup-id>`, `delete <id> <rule-id>`, `discard <id> <ip>`, `edisk`, `eip`, `expand <id>`, `firewall`, `free <id>`, `get <id>`, `limit <id>`, `list`, `list <id>`, `monitor <id>`, `move <id> <rule-id>`, `nat`, `pve-address <id>`, `reboot <id>`, `reinstall <id>`, `renew <id>`, `renew-price <id>`, `reset-password <id>`, `restore <id> <backup-id>`, `server`, `set <id>`, `set-description <id> <ip> <description>`, `set-tag <id> <tag>`, `start <id>`, `stop <id>`, `traffic`, `upgrade <id>`, `vnc <id>`

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/rcs/` | 获取云服务器列表 | `rcs.GetRcsList` | typed | ✓ |
| POST | `/product/rcs/` | 创建云服务器 | `rcs.CreateRcs` | typed | ✓ |
| GET | `/product/rcs/discount-percent` | 获取云服务器折扣比率 | — |  | — |
| GET | `/product/rcs/os-templates` | 获取RCS操作系统列表 | `public.GetRcsOSList` | typed | — |
| GET | `/product/rcs/plans` | 云服务器获取套餐列表 | — |  | — |
| GET | `/product/rcs/price` | 获取云服务器价格 | `rcs.GetRcsRenewPrice` | typed | ✓ |
| GET | `/product/rcs/usage` | 获取使用情况列表 | — |  | — |
| GET | `/product/rcs/{id}/` | 获取RCS详情 | `rcs.GetRcsDetail` | typed | ✓ |
| POST | `/product/rcs/{id}/backup/` | RCS创建备份 | `rcs.CreateRcsBackup` | typed | ✓ |
| PATCH | `/product/rcs/{id}/backup/setting` | RCS设置备份选项 | `rcs.EnableRcsAutoBackup` | typed | ✓ |
| DELETE | `/product/rcs/{id}/backup/{bid}/` | RCS删除备份 | `rcs.DeleteRcsBackup` | typed | ✓ |
| POST | `/product/rcs/{id}/backup/{bid}/cancel` | RCS取消备份 | `rcs.CancelRcsBackup` | typed | ✓ |
| POST | `/product/rcs/{id}/backup/{bid}/restore` | RCS还原备份 | `rcs.RestoreRcsBackup` | typed | ✓ |
| POST | `/product/rcs/{id}/bridge_setintip` | 桥接模式下设置内网 | — |  | — |
| POST | `/product/rcs/{id}/changeos` | RCS重装系统 | `rcs.ReinstallRcs` | typed | ✓ |
| POST | `/product/rcs/{id}/edisk/` | RCS管理弹性云盘 | `rcs.RcsManagesElasticCloudDisks` | typed | ✓ |
| POST | `/product/rcs/{id}/eip/` | 创建并绑定弹性IP到RCS | `rcs.CreateAndBindElasticIpToRcs` | typed | ✓ |
| POST | `/product/rcs/{id}/eip/change` | 更换IP | `rcs.ChangeRcsIP` | typed | ✓ |
| POST | `/product/rcs/{id}/eip/description` | 设置IP描述 | `rcs.SetRcsEipDescription` | typed | ✓ |
| POST | `/product/rcs/{id}/eip/discard` | 放弃IP | `rcs.DisCardRcsIP` | typed | ✓ |
| POST | `/product/rcs/{id}/fai-send` | 发布快速app安装任务 | — |  | — |
| POST | `/product/rcs/{id}/firewall/mode` | 创建/设置防火墙规则 | — |  | — |
| GET | `/product/rcs/{id}/firewall/rule` | 获取防火墙规则列表 | `rcs.GetRcsFirewallRules` | typed | ✓ |
| POST | `/product/rcs/{id}/firewall/rule` | 创建/设置防火墙规则 | `rcs.SetRcsFirewallRule` | typed | ✓ |
| DELETE | `/product/rcs/{id}/firewall/rule/{ruleId}` | 删除防火墙规则 | `rcs.DeleteRcsFirewallRule` | typed | ✓ |
| PUT | `/product/rcs/{id}/firewall/rule/{ruleId}/pos` | 移动防火墙规则优先级 | `rcs.MobileRcsFirewallRulePriority` | typed | ✓ |
| GET | `/product/rcs/{id}/firewall/sync_time` | 获取防火墙同步开始时间 | — |  | — |
| POST | `/product/rcs/{id}/free` | 释放 | `rcs.FreeRcs` | typed | ✓ |
| GET | `/product/rcs/{id}/monitor` | 获取监控数据 | `rcs.GetRcsMonitorData` | typed | ✓ |
| DELETE | `/product/rcs/{id}/nat` | 删除NAT端口映射 | `rcs.DeleteRcsNatPortMapping` | typed | ✓ |
| GET | `/product/rcs/{id}/nat` | 添加NAT端口映射 | `rcs.AddRcsNatPortMapping` | typed | ✓ |
| POST | `/product/rcs/{id}/reboot` | 云服务器重启操作 | `rcs.RebootRcs` | typed | ✓ |
| GET | `/product/rcs/{id}/renew/` | 获取续费价格 | — |  | — |
| POST | `/product/rcs/{id}/renew/` | 续费 | `rcs.RenewRcs` | typed | ✓ |
| POST | `/product/rcs/{id}/renew/option` | 自动续费选项 | `rcs.EnableRcsAutoRenew` | typed | ✓ |
| POST | `/product/rcs/{id}/reset-password` | 云服务器重置密码操作 | `rcs.ResetRcsPassword` | typed | ✓ |
| POST | `/product/rcs/{id}/start` | 云服务器开机操作 | `rcs.StartRcs` | typed | ✓ |
| POST | `/product/rcs/{id}/stop` | 云服务器关机操作 | `rcs.StopRcs` | typed | ✓ |
| PATCH | `/product/rcs/{id}/tag` | 设置云服务器标签 | `rcs.SetRcsTag` | typed | ✓ |
| POST | `/product/rcs/{id}/to-bridge` | 转成桥接 | — |  | — |
| POST | `/product/rcs/{id}/toggle_dgpu` | RCS充流量 | — |  | — |
| POST | `/product/rcs/{id}/toggle_primary_gpu` | RCS充流量 | — |  | — |
| POST | `/product/rcs/{id}/traffic/auto` | RCS设置自动充流量 | — |  | — |
| POST | `/product/rcs/{id}/traffic/charge` | RCS充流量 | `rcs.ChargeRcsTrafic` | typed | ✓ |
| POST | `/product/rcs/{id}/traffic/limit` | RCS限流 | `rcs.LimitRcsTrafic` | typed | ✓ |
| POST | `/product/rcs/{id}/upgrade` | 升级 | `rcs.UpgradeRcs` | typed | ✓ |
| GET | `/product/rcs/{id}/usage` | 获取使用情况 | — |  | — |
| GET | `/product/rcs/{id}/vnc` | 连接VNC | `rcs.GetRcsVnc` | typed | ✓ |
| PATCH | `/product/rcs/{id}/vnet` | 子网改名 | — |  | — |
| POST | `/product/rcs/{id}/vnet` | 创建虚拟机内网子网 | — |  | — |
| POST | `/product/rcs/{id}/webbar/charge` | RCS充流量 | — |  | — |
| POST | `/product/ros/{id}/free` | 释放 | — |  | — |

## rgs

CLI 命令：`add <id>`, `auto-renew <id>`, `backup`, `cancel <id> <backup-id>`, `change <id>`, `change <id> <ip>`, `config <id>`, `config-set <id> <json>`, `cpu-charge <id>`, `cpu-limit-mode <id>`, `create`, `create <id> <label>`, `create <id> <name>`, `create <name> <password>`, `daily-mode <id>`, `database <id>`, `delete <id> <backup-id>`, `delete <id> <nat-id>`, `delete <id> <rule-id>`, `delete <name>`, `discard <id> <ip>`, `discount-percent`, `edit <name> <password>`, `egg`, `eip`, `fai-send <id>`, `firewall`, `free <id>`, `game`, `get <id>`, `init <id>`, `k8s`, `lang <id>`, `list`, `list <id>`, `mcsm`, `mode <id> <mode>`, `monitor <id>`, `move <id> <rule-id>`, `mp`, `nat`, `os`, `pal`, `plans`, `ptero`, `rcon <id> <command>`, `reboot <id>`, `reinstall <id>`, `rename <id> <new-name>`, `renew <id>`, `renew-price <id>`, `reset-password <id>`, `restart <id>`, `restore <id> <backup-id>`, `scale <id>`, `server`, `set-description <id> <ip> <description>`, `set-int-ip <id> <ip>`, `set-tag <id> <tag>`, `setting <id>`, `sftp <id>`, `sftp-init <id>`, `start <id>`, `start-command <id> <command>`, `status <id>`, `stop <id>`, `switch-user <id>`, `sync-time <id>`, `to-bridge <id>`, `type`, `upgrade-price <id>`, `usage [id]`, `user`, `vnc <id>`, `vnet`

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/rgs-mp/` | 获取列表 | `rgs.ListRgsMp` | passthrough | ✓ |
| POST | `/product/rgs-mp/` | 创建游戏云MP | `rgs.CreateRgsMp` | passthrough | ✓ |
| POST | `/product/rgs-mp/{id}/renew/` | 续费游戏云MP | `rgs.RenewRgsMp` | passthrough | ✓ |
| GET | `/product/rgs/` | 获取列表 | `rgs.GetRgsList` | typed | ✓ |
| POST | `/product/rgs/` | 创建游戏云 | `rgs.CreateRgs` | typed | ✓ |
| POST | `/product/rgs/change-egg` | RGS切换egg(游戏类型) | `rgs.ChangeRgsEgg` | typed | ✓ |
| GET | `/product/rgs/discount-percent` | 获取游戏云折扣比率 | `rgs.GetRgsDiscountPercent` | passthrough | ✓ |
| GET | `/product/rgs/egg` | 蛋(游戏)列表 | `public.GetEggList` | typed | ✓ |
| GET | `/product/rgs/egg_server` | 服务端类型列表 | `rgs.GetRgsEggServerList` | typed | ✓ |
| GET | `/product/rgs/egg_type` | 蛋(游戏类型)类型列表 | `public.GetEggTypeList` | typed | ✓ |
| GET | `/product/rgs/mcsm/pal/config` | pal配置 | `rgs.GetPalConfig` | passthrough | ✓ |
| POST | `/product/rgs/mcsm/pal/config` | pal配置 | `rgs.SetPalConfig` | typed | ✓ |
| POST | `/product/rgs/mcsm/pal/init` | pal配置 | `rgs.InitPal` | typed | ✓ |
| GET | `/product/rgs/mcsm/pal/lang` | pal配置中文 | `rgs.GetPalLang` | passthrough | ✓ |
| POST | `/product/rgs/mcsm/pal/rcon` | 关闭pal | `rgs.PalRcon` | typed | ✓ |
| POST | `/product/rgs/mcsm/pal/restart` | 关闭pal | `rgs.RestartPal` | typed | ✓ |
| POST | `/product/rgs/mcsm/pal/stop` | 关闭pal | `rgs.StopPal` | typed | ✓ |
| GET | `/product/rgs/mcsm/panel_user/` | 翼龙面板用户列表 | `rgs.GetMcsmUserList` | typed | ✓ |
| PATCH | `/product/rgs/mcsm/panel_user/` | 编辑面板用户 | `rgs.EditMcsmUser` | typed | ✓ |
| POST | `/product/rgs/mcsm/panel_user/` | 创建面板用户 | `rgs.CreateMcsmUser` | typed | ✓ |
| DELETE | `/product/rgs/mcsm/panel_user/{name}` | 删除面板用户 | `rgs.DeleteMcsmUser` | typed | ✓ |
| POST | `/product/rgs/mcsm/sftp_init` | 初始化/刷新sftp功能 | `rgs.McsmSftpInit` | typed | ✓ |
| POST | `/product/rgs/mcsm/start/` | 开服 | `rgs.StartMcsmInstance` | typed | ✓ |
| GET | `/product/rgs/mcsm/status` | 实例信息和状态 | `rgs.GetMcsmStatus` | passthrough | ✓ |
| GET | `/product/rgs/os-templates` | 系统列表 | `public.GetRgsOSList` | typed | ✓ |
| GET | `/product/rgs/plans` | 获取套餐列表 | `rgs.GetRgsPlanList` | typed | ✓ |
| GET | `/product/rgs/price` | 获取游戏云价格 | `rgs.GetRgsUpgradePrice` | typed | ✓ |
| GET | `/product/rgs/ptero/panel_user/` | 翼龙面板用户列表 | `rgs.GetPteroUserList` | typed | ✓ |
| PATCH | `/product/rgs/ptero/panel_user/` | 编辑面板用户 | `rgs.EditPteroUser` | typed | ✓ |
| POST | `/product/rgs/ptero/panel_user/` | 创建翼龙面板用户 | `rgs.CreatePteroUser` | typed | ✓ |
| DELETE | `/product/rgs/ptero/panel_user/{name}` | 删除面板用户 | `rgs.DeletePteroUser` | typed | ✓ |
| POST | `/product/rgs/switch-user` | RGS切换面板用户 | `rgs.SwitchRgsPanelUser` | typed | ✓ |
| GET | `/product/rgs/usage` | 获取使用情况列表 | `rgs.GetRgsUsageList` | typed | ✓ |
| GET | `/product/rgs/{id}/` | 获取RGS详情 | `rgs.GetRgsDetail` | typed | ✓ |
| POST | `/product/rgs/{id}/backup/` | RGS创建备份 | `rgs.CreateRgsBackup` | typed | ✓ |
| PATCH | `/product/rgs/{id}/backup/setting` | RGS设置备份选项 | `rgs.EnableRgsAutoBackup` | typed | ✓ |
| DELETE | `/product/rgs/{id}/backup/{bid}/` | RGS删除备份 | `rgs.DeleteRgsBackup` | typed | ✓ |
| POST | `/product/rgs/{id}/backup/{bid}/cancel` | RGS取消备份 | `rgs.CancelRgsBackup` | typed | ✓ |
| POST | `/product/rgs/{id}/backup/{bid}/restore` | RGS还原备份 | `rgs.RestoreRgsBackup` | typed | ✓ |
| POST | `/product/rgs/{id}/bridge_setintip` | 桥接模式下设置内网 | `rgs.RgsBridgeSetIntIP` | typed | ✓ |
| POST | `/product/rgs/{id}/changeos` | RGS重装系统 | `rgs.ReinstallRgs` | typed | ✓ |
| POST | `/product/rgs/{id}/cpu-charge` | cpu充电 | `rgs.ChargeRgsCPU` | typed | ✓ |
| POST | `/product/rgs/{id}/cpu-limit-mode` | 游戏云限制模式(是否用余额结算)切换 | `rgs.SwitchRgsBalanceMode` | typed | ✓ |
| POST | `/product/rgs/{id}/daily-mode` | 游戏云日付模式开关 | `rgs.SwitchRgsDailyMode` | typed | ✓ |
| POST | `/product/rgs/{id}/eip` | 创建并绑定弹性IP到RGS | `rgs.CreateAndBindElasticIpToRgs` | typed | ✓ |
| POST | `/product/rgs/{id}/eip/change` | 更换IP | `rgs.ChangeRgsIP` | typed | ✓ |
| POST | `/product/rgs/{id}/eip/description` | 设置IP描述 | `rgs.SetRgsEipDescription` | typed | ✓ |
| POST | `/product/rgs/{id}/eip/discard` | 放弃IP | `rgs.DisCardRgsIP` | typed | ✓ |
| POST | `/product/rgs/{id}/fai-send` | 发布快速app安装任务 | `rgs.SendRgsFaiTask` | typed | ✓ |
| POST | `/product/rgs/{id}/firewall/mode` | 创建/设置防火墙规则 | `rgs.SetRgsFirewallMode` | typed | ✓ |
| GET | `/product/rgs/{id}/firewall/rule` | 获取防火墙规则列表 | `rgs.GetRgsFirewallRules` | typed | ✓ |
| POST | `/product/rgs/{id}/firewall/rule` | 创建/设置防火墙规则 | `rgs.SetRgsFirewallRule` | typed | ✓ |
| DELETE | `/product/rgs/{id}/firewall/rule/{ruleId}` | 删除防火墙规则 | `rgs.DeleteRgsFirewallRule` | typed | ✓ |
| PUT | `/product/rgs/{id}/firewall/rule/{ruleId}/pos` | 移动防火墙规则优先级 | `rgs.MobileRgsFirewallRulePriority` | typed | ✓ |
| GET | `/product/rgs/{id}/firewall/sync_time` | 获取防火墙同步开始时间 | `rgs.GetRgsFirewallSyncTime` | passthrough | ✓ |
| POST | `/product/rgs/{id}/free` | 释放游戏云 | `rgs.FreeRgs` | typed | ✓ |
| PATCH | `/product/rgs/{id}/k8s-panel/database` | K8S面板修改数据库设置 | `rgs.SetK8SPanelDatabase` | typed | ✓ |
| POST | `/product/rgs/{id}/k8s-panel/set-start-command` | 游戏云设置启动命令（仅支持雨云面板） | `rgs.SetK8SPanelStartCommand` | typed | ✓ |
| PATCH | `/product/rgs/{id}/k8s-panel/sftp` | K8S面板修改SFTP设置 | `rgs.SetK8SPanelSFTP` | typed | ✓ |
| POST | `/product/rgs/{id}/monitor` | 获取监控数据 | `rgs.GetRgsMonitorData` | typed | ✓ |
| DELETE | `/product/rgs/{id}/nat` | 删除NAT端口映射 | `rgs.DeleteRgsNatPortMapping` | typed | ✓ |
| POST | `/product/rgs/{id}/nat` | 添加NAT端口映射 | `rgs.AddRgsNatPortMapping` | typed | ✓ |
| POST | `/product/rgs/{id}/reboot` | 游戏云重启操作 | `rgs.RebootRgs` | typed | ✓ |
| POST | `/product/rgs/{id}/renew/` | 续费 | `rgs.RenewRgs` | typed | ✓ |
| POST | `/product/rgs/{id}/renew/option` | 自动续费选项 | `rgs.EnableRgsAutoRenew` | typed | ✓ |
| POST | `/product/rgs/{id}/reset-password` | 游戏云重置密码操作 | `rgs.ResetRgsPassword` | typed | ✓ |
| POST | `/product/rgs/{id}/scale` | 升级 | `rgs.ScaleRgs` | passthrough | ✓ |
| POST | `/product/rgs/{id}/start` | 游戏云开机操作 | `rgs.StartRgs` | typed | ✓ |
| POST | `/product/rgs/{id}/stop` | 游戏云关机操作 | `rgs.StopRgs` | typed | ✓ |
| PATCH | `/product/rgs/{id}/tag` | 设置游戏云标签 | `rgs.SetRgsTag` | typed | ✓ |
| POST | `/product/rgs/{id}/to-bridge` | 转成桥接 | `rgs.RgsToBridge` | typed | ✓ |
| GET | `/product/rgs/{id}/usage` | 获取使用情况 | `rgs.GetRgsUsage` | passthrough | ✓ |
| GET | `/product/rgs/{id}/vnc` | 连接VNC | `rgs.GetRgsVnc` | typed | ✓ |
| PATCH | `/product/rgs/{id}/vnet` | 子网改名 | `rgs.RenameRgsVnet` | typed | ✓ |
| POST | `/product/rgs/{id}/vnet` | 创建虚拟机内网子网 | `rgs.CreateRgsVnet` | typed | ✓ |

## ros

CLI 命令：`bucket`, `create <instance-id> <name>`, `list`, `list <instance-id>`, `storage`

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/ros/bucket` | 获取存储桶列表 | `ros.GetRosBucketList` | typed | — |
| POST | `/product/ros/bucket` | 创建对象存储桶 | `ros.CreateRosBucket` | typed | ✓ |
| DELETE | `/product/ros/bucket/{id}` | 删除对象存储桶 | `ros.DeleteRosBucket` | typed | — |
| GET | `/product/ros/bucket/{id}` | 获取对象存储桶详情 | `ros.GetRosBucketListByInstance` | typed | ✓ |
| GET | `/product/ros/bucket/{id}/lifecycle` | 查询生命周期规则列表 | — |  | — |
| POST | `/product/ros/bucket/{id}/lifecycle` | 创建生命周期规则 | — |  | — |
| DELETE | `/product/ros/bucket/{id}/lifecycle/{rule_id}` | 删除生命周期规则 | — |  | — |
| GET | `/product/ros/bucket/{id}/monitor` | 实例监控 | `ros.GetRosBucketMonitorData` | typed | — |
| GET | `/product/ros/bucket/{id}/offline-download` | 列出离线下载任务 | — |  | — |
| POST | `/product/ros/bucket/{id}/offline-download` | 创建离线下载任务 | — |  | — |
| DELETE | `/product/ros/bucket/{id}/offline-download/{task_id}` | 取消离线下载任务 | — |  | — |
| GET | `/product/ros/bucket/{id}/offline-download/{task_id}` | 查询离线下载任务详情 | — |  | — |
| PATCH | `/product/ros/bucket/{id}/proxy` | 修改存储桶Proxy设置 | `ros.ModifyRosBucketProxySettings` | typed | — |
| POST | `/product/ros/bucket/{id}/regenerate-keys` | 对象存储桶重新生成密钥 | `ros.ReGenerateRosBucketKeys` | typed | — |
| GET | `/product/ros/bucket/{id}/request-log` | 获取请求日志生成任务列表 | — |  | — |
| POST | `/product/ros/bucket/{id}/request-log` | 生成请求日志 | — |  | — |
| GET | `/product/ros/bucket/{id}/request-log/{task_id}/link` | 获取请求日志下载链接 | — |  | — |
| POST | `/product/ros/bucket/{id}/request-log/{task_id}/retrigger` | 重新触发请求日志生成任务 | — |  | — |
| GET | `/product/ros/bucket/{id}/statistics` | 获取存储桶热点数据 | — |  | — |
| DELETE | `/product/ros/bucket/{id}/sync` | 取消主动同步 | — |  | — |
| GET | `/product/ros/bucket/{id}/sync` | 查询主动同步状态 | — |  | — |
| POST | `/product/ros/bucket/{id}/sync` | 创建或重新运行主动同步 | — |  | — |
| POST | `/product/ros/bucket/{id}/toggle-public-access` | 开关对象存储桶匿名访问 | `ros.SetRosBucketPublicAccess` | typed | — |
| GET | `/product/ros/discount-percent` | 获取对象存储实例折扣比率 | — |  | — |
| GET | `/product/ros/instance` | 获取对象存储实例列表 | `ros.GetRosInstanceList` | typed | ✓ |
| POST | `/product/ros/instance` | 创建对象存储实例 | `ros.CreateRosInstance` | typed | — |
| GET | `/product/ros/instance/{id}` | 获取对象存储实例详情 | `ros.GetRosInstanceDetail` | typed | — |
| POST | `/product/ros/instance/{id}/regenerate-keys` | 对象存储实例重新生成密钥 | `ros.ReGenerateRosInstanceKeys` | typed | — |
| POST | `/product/ros/instance/{id}/renew` | ROS实例续费 | `ros.RenewRosInstance` | typed | — |
| POST | `/product/ros/instance/{id}/renew/option` | 自动续费选项 | `ros.SetRosInstanceAutoRenewOption` | typed | — |
| POST | `/product/ros/instance/{id}/scale` | ROS实例缩放 | `ros.ScaleRosInstance` | typed | — |
| PATCH | `/product/ros/instance/{id}/tag` | 设置对象存储实例标签 | `ros.SetRosInstanceTags` | typed | — |
| POST | `/product/ros/instance/{id}/toggle-extra-accounting` | 开关对象存储实例的弹性计费选项 | `ros.ToggleRosInstanceExtraAccounting` | typed | — |
| POST | `/product/ros/instance/{id}/toggle-public-access` | 开关对象存储实例的匿名访问 | `ros.SetRosInstancePublicAccess` | typed | — |
| GET | `/product/ros/plans` | 获取对象存储套餐列表 | `public.GetRosPlanList` | typed | — |
| GET | `/product/ros/price` | 获取云服务器折扣比率 | — |  | — |

## rvh

CLI 命令：`backup`, `bind <id> <domain>`, `create`, `create <id> <label>`, `delete <id>`, `delete <id> <backup-id>`, `domain`, `get <id>`, `list`, `plans`, `renew <id>`, `restore <id> <backup-id>`, `rvh`, `unbind <id> <domain>`

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/rvh/` | 获取虚拟主机列表 | `rvh.GetRvhList` | typed | ✓ |
| POST | `/product/rvh/` | 创建虚拟主机 | `rvh.CreateRvh` | typed | ✓ |
| GET | `/product/rvh/plans` | 获取虚拟主机套餐列表 | `rvh.GetRvhPlanList` | typed | ✓ |
| GET | `/product/rvh/price` | 获取虚拟主机折扣比率 | `rvh.GetRvhPrice` | typed | — |
| PATCH | `/product/rvh/tag` | 设置虚拟主机标签 | `rvh.SetRvhTag` | typed | — |
| GET | `/product/rvh/{id}/` | 获取RVH虚拟主机详情 | `rvh.GetRvhDetail` | typed | ✓ |
| POST | `/product/rvh/{id}/backup/` | RVH创建备份 | `rvh.CreateRvhBackup` | typed | ✓ |
| PATCH | `/product/rvh/{id}/backup/setting` | RVH设置备份选项 | `rvh.SetRvhBackupSetting` | typed | — |
| DELETE | `/product/rvh/{id}/backup/{bid}/` | RVH删除备份 | `rvh.DeleteRvhBackup` | typed | ✓ |
| POST | `/product/rvh/{id}/backup/{bid}/restore` | RVH还原备份 | `rvh.RestoreRvhBackup` | typed | ✓ |
| POST | `/product/rvh/{id}/bt/attach-dedip` | 附加独立IP地址 | `rvh.RvhBtAttachDedip` | typed | — |
| POST | `/product/rvh/{id}/bt/fix` | RVH宝塔主机修复操作 | `rvh.RvhBtFix` | typed | — |
| POST | `/product/rvh/{id}/bt/reboot` | RVH宝塔主机重启操作 | `rvh.RvhBtReboot` | typed | — |
| DELETE | `/product/rvh/{id}/domain/` | RVH域名解绑 | `rvh.UnbindRvhDomain` | typed | ✓ |
| POST | `/product/rvh/{id}/domain/` | RVH域名绑定 | `rvh.BindRvhDomain` | typed | ✓ |
| POST | `/product/rvh/{id}/ep/reset-pass` | RVH EP主机重置密码操作 | `rvh.ResetRvhEpPassword` | typed | — |
| POST | `/product/rvh/{id}/firewall/option` | RVH防火墙设置选项 | `rvh.SetRvhFirewallOption` | typed | — |
| POST | `/product/rvh/{id}/firewall/rule` | RVH防火墙设置规则 | `rvh.SetRvhFirewallRule` | typed | — |
| POST | `/product/rvh/{id}/free` | 释放 | `rvh.FreeRvh` | typed | ✓ |
| POST | `/product/rvh/{id}/maintenance-mode` | RVH设置维护模式 | `rvh.SetRvhMaintenanceMode` | typed | — |
| POST | `/product/rvh/{id}/reinstall` | RVH重装操作 | `rvh.ReinstallRvh` | typed | — |
| GET | `/product/rvh/{id}/renew/` | 获取虚拟主机折扣比率 | `rvh.GetRvhRenewPrice` | typed | — |
| POST | `/product/rvh/{id}/renew/` | 续费 | `rvh.RenewRvh` | typed | ✓ |
| POST | `/product/rvh/{id}/renew/option` | 自动续费选项 | `rvh.EnableRvhAutoRenew` | typed | — |
| POST | `/product/rvh/{id}/upgrade/` | 升级 | `rvh.UpgradeRvh` | typed | — |

## ssl

| 方法 | 路径 | 说明 | SDK | 响应类型 | CLI |
|---|---|---|---|---|---|
| GET | `/product/sslcenter/` | SSL证书获取列表 | `ssl.GetSSLCertificateList` | typed | ✓ |
| POST | `/product/sslcenter/` | SSL证书上传操作 | `ssl.UploadSSLCertificate` | typed | ✓ |
| POST | `/product/sslcenter/cert/order` | 创建SSL证书申请 | `ssl.ApplyFreeSSLCertificate` | typed | ✓ |
| POST | `/product/sslcenter/cert/order_verify` | 验证SSL证书申请 | `ssl.VerifyFreeSSLCertificate` | typed | ✓ |
| GET | `/product/sslcenter/cert/orders` | 获取SSL证书申请列表 | `ssl.GetSSLCertApplyList` | typed | ✓ |
| GET | `/product/sslcenter/order` | 获取SSL证书订单列表 | `ssl.GetSSLOrderList` | typed | ✓ |
| POST | `/product/sslcenter/order` | 创建SSL证书订单 | `ssl.CreateSSLOrder` | typed | ✓ |
| GET | `/product/sslcenter/order/{id}` | 获取SSL证书订单信息 | `ssl.GetSSLOrderDetail` | typed | ✓ |
| POST | `/product/sslcenter/order/{id}/assign` | 将SSL证书添加到证书列表 | `ssl.AssignSSLOrder` | typed | ✓ |
| GET | `/product/sslcenter/order/{id}/cert` | 获取SSL证书 | `ssl.GetSSLOrderCert` | typed | ✓ |
| POST | `/product/sslcenter/order/{id}/description` | 更新SSL订单描述 | `ssl.UpdateSSLOrderDescription` | typed | ✓ |
| POST | `/product/sslcenter/order/{id}/revoke` | 申请吊销SSL证书 | `ssl.RevokeSSLOrder` | typed | ✓ |
| POST | `/product/sslcenter/order/{id}/verify` | 验证SSL证书订单 | `ssl.VerifySSLOrder` | typed | ✓ |
| POST | `/product/sslcenter/price` | 获取SSL证书订单价格 | `ssl.GetSSLOrderPrice` | typed | ✓ |
| GET | `/product/sslcenter/product` | 获取SSL证书产品列表 | `ssl.GetSSLProductList` | typed | ✓ |
| DELETE | `/product/sslcenter/{id}` | SSL证书删除操作 | `ssl.DeleteSsl` | typed | ✓ |
| GET | `/product/sslcenter/{id}` | SSL证书查看操作 | `ssl.GetSslDetail` | typed | ✓ |
| PUT | `/product/sslcenter/{id}` | SSL证书替换操作 | `ssl.ReplaceSsl` | typed | ✓ |
