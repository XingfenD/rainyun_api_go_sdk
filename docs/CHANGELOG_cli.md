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

### Changed / 变更

- `-o raw`（原 `--raw` 标志移除）改为原样透传 API 响应体，对全部读取命令统一生效，便于用户自行解析。
- `ry --verbose` 现在显示输出格式及其来源（config/`--output`）。
- 命令 ID 参数统一为 `int`，显示层 model ID 保持 `string`。
- `server reinstall` flag 从 `--os` 改为 `--os-id int`。
