# Changelog / 更新日志

All notable changes to this template should be documented in this file.
本模板的重要变更建议统一记录在此文件中。

The format loosely follows Keep a Changelog and can be adapted to the team's habits.
本文档参考了 Keep a Changelog 的思路，也可以根据团队习惯调整。

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
