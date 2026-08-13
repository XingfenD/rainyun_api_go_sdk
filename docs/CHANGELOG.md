# Changelog / 更新日志

All notable changes to this template should be documented in this file.
本模板的重要变更建议统一记录在此文件中。

The format loosely follows Keep a Changelog and can be adapted to the team's habits.
本文档参考了 Keep a Changelog 的思路，也可以根据团队习惯调整。

## [Unreleased] - 2026-08-13

### Added / 新增

- `config set <key>` 支持交互式输入：省略 value 时交互式提示，`apikey` 使用隐藏输入避免敏感信息暴露在终端/历史中。
- 新增 `Makefile`（build/test/vet/fmt/cover/tidy/clean/help）。
- `ry server get` 输出细化为 `ServerDetail`：补充内网 IP、网络模式/带宽、实时 CPU/内存/网络、流量包用量与上限、自动续费/套餐/续费积分，以及扩展盘、弹性 IP、备份、可升级套餐等区块；`-o json` 输出完整结构化字段。
- `-o raw`（原 `--raw` 标志移除）改为原样透传 API 响应体（`apis.RyClient.RawBody()` / `sdk.RainyunSDK.RawResponseBody()`），对全部读取命令统一生效，便于用户自行解析。

- 迁移 rainyun-cli 到仓库内 `cmd/ry/`：新增 `main.go`、`root.go`、`commands/{server,domain,storage,billing,configcmd}`、`internal/{config,model,output}`。
- CLI 直接调用 `sdk.New(apiKey)`，删除 CLI 自带 provider/HTTP 层。
- 配置简化为单一 `apikey` + `output`（table/json/yaml），落盘到 `~/.config/ry/config.toml`。
- 补齐 SDK 4 个缺口端点：`domain list`、`dns list`、`bucket-by-instance`、`expense orders`。
- 新增 `cmd/ry/internal/config` 与 `cmd/ry/internal/output` 测试（6/6 + 7/7）。

### Fixed / 修复

- 修复 `output --raw` 输出 bug：新增 raw case，原样透传 SDK Response。

### Changed / 变更

- `ry --verbose` 重构为结构化链路追踪：SDK 通过 `TraceSink` 发出 HTTP/Result 事件，CLI 渲染到 stderr，不影响标准输出；默认预览 64 KiB 响应体（JSON 自动格式化），新增 `--verbose-body-limit int`（0 关闭预览）与 `--verbose-full-body` 标志；不输出 API key、请求头或请求体。
- `ry --verbose` 现在显示输出格式、其来源（config/`--output`/`--raw`）及 raw 状态。
- SDK 调试接口重构：移除 `sdk.NewWithDebug` 与 `apis.RyClient.SetDebugWriter`，改为统一的 Builder 链式构造：`sdk.NewBuilder(apiKey).WithTrace(...).Build()`、`apis.NewBuilder(apiKey).WithTrace(...).Build()`，sdk 仅保留最基础的 `New()`；追踪类型（`TraceSink`/`TraceOptions`/`HTTPTrace`/`ResultTrace`）定义于 `apis` 包。
- 命令 ID 参数统一为 `int`，显示层 model ID 保持 `string`。
- `server reinstall` flag 从 `--os` 改为 `--os-id int`。
- `.gitignore` 新增 `bin/`、`coverage.out`、`coverage.html`。
- 新增 `golang.org/x/term` 依赖（交互式隐藏输入）。

## [0.2.0] - 2026-08-05

### Added / 新增

- 迁移剩余全部模块：domain(24)、product(8)、rbm(10)、rca(21)、rgs(38)、ros(19)、ssl(5)、user(10)、workorder(11)，并补全 public 剩余 5 个接口。
- 全部响应结构体嵌套拆分命名复用；rgs 复用 rcs 的请求/响应类型。
- `sdk.New` 注册全部 11 个服务。

### Fixed / 修复

- 修正源 SDK 中的路径错误（`{id}` 占位符、缺少 `/`、`/product/rcs/` 误写为 rgs 路径等）与参数传递错误（`resp` vs `&resp`、漏传 body）。

## [0.1.0] - 2026-08-05

### Added / 新增

- 从 `rainyun-go-sdk` 迁移 RCS 全部 35 个接口（列表/详情/创建/备份/防火墙/监控/NAT/续费/升降级等），响应结构体嵌套已拆分复用（`RcsRecord`、`Node`、`Plan`、`OsInfo`、`UsageData` 等）。
- 新增 `apis/common` 基础类型：`BasicOperationResponse`、`VncConnectionInfo`。
- `constant` 新增 `HTTPMethod_PATCH`。

## [0.0.0] - 2026-08-02

### Notes / 说明

- The repository currently provides structure and placeholder files rather than a finished implementation.
- 当前仓库主要提供目录结构和占位文件，尚不是一个已完成功能实现的成品项目。

- Future updates should record framework selection, startup steps, deployment workflow, and major documentation changes.
- 后续若补充了技术栈、启动流程、部署方式或重要文档内容，建议继续记录在本文件中。
