# SDK / CLI / openapi.json 接入进度

本文档追踪 Rainyun API（`docs/openapi.json`）在 Go SDK 与 `ry` CLI 中的接入进度。
SDK 代码在 `apis/`，CLI 代码在 `cmd/ry/commands/`。

> 重新生成：`python3 docs/scripts/generate.py`

## 总览

| 分类 | openapi 端点 | SDK 已实现 | CLI 已接入 |
|---|---|---|---|
| domain | 38 | 24 | 4 |
| public | 6 | 3 | 3 |
| rbm | 23 | 9 | 0 |
| rcs | 52 | 35 | 34 |
| rgs | 75 | 36 | 0 |
| ros | 36 | 19 | 3 |
| ssl | 18 | 5 | 0 |
| **合计** | **248** | **131** | **44** |

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
| GET | `/discourse` | 获取论坛数据 | — | — |
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

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/rgs-mp/` | 获取列表 | — | — |
| POST | `/product/rgs-mp/` | 创建游戏云MP | — | — |
| POST | `/product/rgs-mp/{id}/renew/` | 续费游戏云MP | — | — |
| GET | `/product/rgs/` | 获取列表 | `rgs.GetRgsList` | — |
| POST | `/product/rgs/` | 创建游戏云 | `rgs.CreateRgs` | — |
| POST | `/product/rgs/change-egg` | RGS切换egg(游戏类型) | — | — |
| GET | `/product/rgs/discount-percent` | 获取游戏云折扣比率 | — | — |
| GET | `/product/rgs/egg` | 蛋(游戏)列表 | `public.GetEggList` | — |
| GET | `/product/rgs/egg_server` | 服务端类型列表 | — | — |
| GET | `/product/rgs/egg_type` | 蛋(游戏类型)类型列表 | `public.GetEggTypeList` | — |
| GET | `/product/rgs/mcsm/pal/config` | pal配置 | — | — |
| POST | `/product/rgs/mcsm/pal/config` | pal配置 | — | — |
| POST | `/product/rgs/mcsm/pal/init` | pal配置 | — | — |
| GET | `/product/rgs/mcsm/pal/lang` | pal配置中文 | — | — |
| POST | `/product/rgs/mcsm/pal/rcon` | 关闭pal | — | — |
| POST | `/product/rgs/mcsm/pal/restart` | 关闭pal | — | — |
| POST | `/product/rgs/mcsm/pal/stop` | 关闭pal | — | — |
| GET | `/product/rgs/mcsm/panel_user/` | 翼龙面板用户列表 | — | — |
| PATCH | `/product/rgs/mcsm/panel_user/` | 编辑面板用户 | `rgs.EditMcsmUser` | — |
| POST | `/product/rgs/mcsm/panel_user/` | 创建面板用户 | `rgs.CreateMcsmUser` | — |
| DELETE | `/product/rgs/mcsm/panel_user/{name}` | 删除面板用户 | `rgs.DeleteMcsmUser` | — |
| POST | `/product/rgs/mcsm/sftp_init` | 初始化/刷新sftp功能 | — | — |
| POST | `/product/rgs/mcsm/start/` | 开服 | — | — |
| GET | `/product/rgs/mcsm/status` | 实例信息和状态 | — | — |
| GET | `/product/rgs/os-templates` | 系统列表 | `public.GetRgsOSList` | — |
| GET | `/product/rgs/plans` | 获取套餐列表 | — | — |
| GET | `/product/rgs/price` | 获取游戏云价格 | `rgs.GetRgsUpgradePrice` | — |
| GET | `/product/rgs/ptero/panel_user/` | 翼龙面板用户列表 | — | — |
| PATCH | `/product/rgs/ptero/panel_user/` | 编辑面板用户 | — | — |
| POST | `/product/rgs/ptero/panel_user/` | 创建翼龙面板用户 | — | — |
| DELETE | `/product/rgs/ptero/panel_user/{name}` | 删除面板用户 | — | — |
| POST | `/product/rgs/switch-user` | RGS切换面板用户 | — | — |
| GET | `/product/rgs/usage` | 获取使用情况列表 | — | — |
| GET | `/product/rgs/{id}/` | 获取RGS详情 | `rgs.GetRgsDetail` | — |
| POST | `/product/rgs/{id}/backup/` | RGS创建备份 | `rgs.CreateRgsBackup` | — |
| PATCH | `/product/rgs/{id}/backup/setting` | RGS设置备份选项 | `rgs.EnableRgsAutoBackup` | — |
| DELETE | `/product/rgs/{id}/backup/{bid}/` | RGS删除备份 | `rgs.DeleteRgsBackup` | — |
| POST | `/product/rgs/{id}/backup/{bid}/cancel` | RGS取消备份 | `rgs.CancelRgsBackup` | — |
| POST | `/product/rgs/{id}/backup/{bid}/restore` | RGS还原备份 | `rgs.RestoreRgsBackup` | — |
| POST | `/product/rgs/{id}/bridge_setintip` | 桥接模式下设置内网 | — | — |
| POST | `/product/rgs/{id}/changeos` | RGS重装系统 | `rgs.ReinstallRgs` | — |
| POST | `/product/rgs/{id}/cpu-charge` | cpu充电 | `rgs.ChargeRgsCPU` | — |
| POST | `/product/rgs/{id}/cpu-limit-mode` | 游戏云限制模式(是否用余额结算)切换 | `rgs.SwitchRgsBalanceMode` | — |
| POST | `/product/rgs/{id}/daily-mode` | 游戏云日付模式开关 | `rgs.SwitchRgsDailyMode` | — |
| POST | `/product/rgs/{id}/eip` | 创建并绑定弹性IP到RGS | `rgs.CreateAndBindElasticIpToRgs` | — |
| POST | `/product/rgs/{id}/eip/change` | 更换IP | `rgs.ChangeRgsIP` | — |
| POST | `/product/rgs/{id}/eip/description` | 设置IP描述 | — | — |
| POST | `/product/rgs/{id}/eip/discard` | 放弃IP | `rgs.DisCardRgsIP` | — |
| POST | `/product/rgs/{id}/fai-send` | 发布快速app安装任务 | — | — |
| POST | `/product/rgs/{id}/firewall/mode` | 创建/设置防火墙规则 | — | — |
| GET | `/product/rgs/{id}/firewall/rule` | 获取防火墙规则列表 | `rgs.GetRgsFirewallRules` | — |
| POST | `/product/rgs/{id}/firewall/rule` | 创建/设置防火墙规则 | `rgs.SetRgsFirewallRule` | — |
| DELETE | `/product/rgs/{id}/firewall/rule/{ruleId}` | 删除防火墙规则 | `rgs.DeleteRgsFirewallRule` | — |
| PUT | `/product/rgs/{id}/firewall/rule/{ruleId}/pos` | 移动防火墙规则优先级 | `rgs.MobileRgsFirewallRulePriority` | — |
| GET | `/product/rgs/{id}/firewall/sync_time` | 获取防火墙同步开始时间 | — | — |
| POST | `/product/rgs/{id}/free` | 释放游戏云 | `rgs.FreeRgs` | — |
| PATCH | `/product/rgs/{id}/k8s-panel/database` | K8S面板修改数据库设置 | — | — |
| POST | `/product/rgs/{id}/k8s-panel/set-start-command` | 游戏云设置启动命令（仅支持雨云面板） | — | — |
| PATCH | `/product/rgs/{id}/k8s-panel/sftp` | K8S面板修改SFTP设置 | — | — |
| POST | `/product/rgs/{id}/monitor` | 获取监控数据 | — | — |
| DELETE | `/product/rgs/{id}/nat` | 删除NAT端口映射 | — | — |
| POST | `/product/rgs/{id}/nat` | 添加NAT端口映射 | `rgs.AddRgsNatPortMapping` | — |
| POST | `/product/rgs/{id}/reboot` | 游戏云重启操作 | `rgs.RebootRgs` | — |
| POST | `/product/rgs/{id}/renew/` | 续费 | `rgs.RenewRgs` | — |
| POST | `/product/rgs/{id}/renew/option` | 自动续费选项 | `rgs.EnableRgsAutoRenew` | — |
| POST | `/product/rgs/{id}/reset-password` | 游戏云重置密码操作 | `rgs.ResetRgsPassword` | — |
| POST | `/product/rgs/{id}/scale` | 升级 | — | — |
| POST | `/product/rgs/{id}/start` | 游戏云开机操作 | `rgs.StartRgs` | — |
| POST | `/product/rgs/{id}/stop` | 游戏云关机操作 | `rgs.StopRgs` | — |
| PATCH | `/product/rgs/{id}/tag` | 设置游戏云标签 | `rgs.SetRgsTag` | — |
| POST | `/product/rgs/{id}/to-bridge` | 转成桥接 | — | — |
| GET | `/product/rgs/{id}/usage` | 获取使用情况 | — | — |
| GET | `/product/rgs/{id}/vnc` | 连接VNC | `rgs.GetRgsVnc` | — |
| PATCH | `/product/rgs/{id}/vnet` | 子网改名 | — | — |
| POST | `/product/rgs/{id}/vnet` | 创建虚拟机内网子网 | — | — |

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

## ssl

| 方法 | 路径 | 说明 | SDK | CLI |
|---|---|---|---|---|
| GET | `/product/sslcenter/` | SSL证书获取列表 | `ssl.GetSSLCertificateList` | — |
| POST | `/product/sslcenter/` | SSL证书上传操作 | `ssl.UploadSSLCertificate` | — |
| POST | `/product/sslcenter/cert/order` | 创建SSL证书申请 | — | — |
| POST | `/product/sslcenter/cert/order_verify` | 验证SSL证书申请 | — | — |
| GET | `/product/sslcenter/cert/orders` | 获取SSL证书申请列表 | — | — |
| GET | `/product/sslcenter/order` | 获取SSL证书订单列表 | — | — |
| POST | `/product/sslcenter/order` | 创建SSL证书订单 | — | — |
| GET | `/product/sslcenter/order/{id}` | 获取SSL证书订单信息 | — | — |
| POST | `/product/sslcenter/order/{id}/assign` | 将SSL证书添加到证书列表 | — | — |
| GET | `/product/sslcenter/order/{id}/cert` | 获取SSL证书 | — | — |
| POST | `/product/sslcenter/order/{id}/description` | 更新SSL订单描述 | — | — |
| POST | `/product/sslcenter/order/{id}/revoke` | 申请吊销SSL证书 | — | — |
| POST | `/product/sslcenter/order/{id}/verify` | 验证SSL证书订单 | — | — |
| POST | `/product/sslcenter/price` | 获取SSL证书订单价格 | — | — |
| GET | `/product/sslcenter/product` | 获取SSL证书产品列表 | — | — |
| DELETE | `/product/sslcenter/{id}` | SSL证书删除操作 | `ssl.DeleteSsl` | — |
| GET | `/product/sslcenter/{id}` | SSL证书查看操作 | `ssl.GetSslDetail` | — |
| PUT | `/product/sslcenter/{id}` | SSL证书替换操作 | `ssl.ReplaceSsl` | — |
