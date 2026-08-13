# CLI Changelog / CLI 更新日志

所有与 `ry` 命令行工具（`cmd/ry/`）相关的变更记录在此文件。
SDK（`apis/`、`sdk/`）与模块级变更记录在 [`CHANGELOG.md`](./CHANGELOG.md)。

格式沿用 Keep a Changelog。

## [0.1.0] - 2026-08-13

### Added / 新增

- `config set <key>` 支持交互式输入：省略 value 时交互式提示，`apikey` 使用隐藏输入避免敏感信息暴露在终端/历史中。
- 迁移 rainyun-cli 到仓库内 `cmd/ry/`：新增 `main.go`、`root.go`、`commands/{server,domain,storage,billing,configcmd}`、`internal/{config,model,output}`。
- CLI 直接调用 `sdk.New(apiKey)`，删除 CLI 自带 provider/HTTP 层。
- 配置简化为单一 `apikey` + `output`（table/json/yaml），落盘到 `~/.config/ry/config.toml`。
- `ry server get` 输出细化为 `ServerDetail`：补充内网 IP、网络模式/带宽、实时 CPU/内存/网络、流量包用量与上限、自动续费/套餐/续费积分，以及扩展盘、弹性 IP、备份、可升级套餐等区块；`-o json` 输出完整结构化字段。
- 新增 `cmd/ry/internal/config` 与 `cmd/ry/internal/output` 测试（6/6 + 7/7）。
- `ry --verbose` 渲染结构化链路追踪到 stderr，不影响标准输出；新增 `--verbose-body-limit int`（0 关闭预览）与 `--verbose-full-body` 标志。
- 新增 `cmd/ry/internal/constant` 维护 CLI 版本号，支持 `ry version` 与 `ry --version`。
- `ry server` 补全剩余 SDK 接口：`free`、`set-tag`、`create`、`upgrade`、`renew-price`、`renew`、`auto-renew`、`edisk (create/expand)`、`monitor`、`backup (create/delete/cancel/restore/auto)`、`eip (set-description/create/change/discard)`、`nat (add/delete)`、`traffic (charge/limit)`、`firewall (list/set/delete/move)`、`pve-address`。
- `server monitor` 时间参数人类可读：`--last 30m/1h/7d`（默认 1h），或 `--start`/`--end` 接受 RFC3339 / `YYYY-MM-DD[ HH:MM[:SS]]`，不再要求手输 Unix 时间戳。
- `server get` 的扩展盘与备份区块显示各自 ID（`#id`），便于 `edisk expand --edisk-id` 与 `backup delete/cancel/restore <backup-id>` 使用。
- 新增 `server eip list <id>`：SDK 无独立 EIP 列表接口，封装 `server get` 的 `EIPList` 提供列表视图。

### Changed / 变更

- `-o raw`（原 `--raw` 标志移除）改为原样透传 API 响应体，对全部读取命令统一生效，便于用户自行解析。
- `ry --verbose` 现在显示输出格式及其来源（config/`--output`）。
- 命令 ID 参数统一为 `int`，显示层 model ID 保持 `string`。
- `server reinstall` flag 从 `--os` 改为 `--os-id int`。
- 抽取通用参数解析到 `cmd/ry/internal/cliutil`（`ParseID`/`ParseDuration`/`ParseTime`/`ResolveTimeRange`），server/storage/domain 命令统一复用，消除内联 `strconv.Atoi` 与时间解析重复。
