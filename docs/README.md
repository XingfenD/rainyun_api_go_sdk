![rainyun-go-sdk](https://www.rainyun.com/img/logo.d193755d.png)

# rainyun_api_go_sdk

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/XingfenD/rainyun_api_go_sdk)
![GitHub](https://img.shields.io/github/license/XingfenD/rainyun_api_go_sdk)

> [!WARNING]
> 本项目由第三方开发，通过 github action cron 监听[雨云api](https://api.rainyun.com/openapi.json)的变更，但不保证及时的维护。

> [!TIP]
> 如果您想支持 sdk 的发展，欢迎从[这个链接](https://www.rainyun.com/MTE3NTg3Mw==_)注册雨云。

[雨云 Rainyun](https://www.rainyun.com/MTE3NTg3Mw==_) API 的 Go SDK 与 `ry` 命令行工具。

- **Go SDK**（`apis/`、`sdk/`）：封装 Rainyun 各产品域 API。
- **`ry` CLI**（`cmd/ry/`）：终端管理云资源，支持 `table`/`json`/`yaml`/`raw` 输出。
- **进度追踪**：见 [`docs/PROGRESS.md`](docs/PROGRESS.md)，追踪 SDK/CLI 对 `docs/openapi.json` 的接入情况。
- **CI 监控**：自动监测上游 `openapi.json` 变化（`.github/workflows/`）。

## 安装

### SDK

```bash
go get github.com/XingfenD/rainyun_api_go_sdk@latest
```

### CLI

```bash
make build        # 构建 bin/ry
# 或
go install ./cmd/ry
```

## 快速开始

### SDK

```go
package main

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"
)

func main() {
	client := sdk.New("你的-API-Key")

	resp, err := client.GetRcsList(&rcs.GetRcsListRequest{
		Options: common.StandQueryParameters{Page: 1, PerPage: 50},
	})
	if err != nil {
		panic(err)
	}
	for _, s := range resp.Data.Records {
		fmt.Println(s.ID, s.Tag, s.Status)
	}
}
```

### CLI

```bash
ry config set apikey <你的-API-Key>   # 首次配置，写入 ~/.config/ry/config.toml
ry server list
ry server get 370018
ry public news
```

输出格式通过 `-o` 指定：`table`（默认）/`json`/`yaml`/`raw`。

## SDK 服务域

| 服务   | 包            | 说明                                           |
| ------ | ------------- | ---------------------------------------------- |
| rcs    | `apis/rcs`    | 云服务器                                       |
| rgs    | `apis/rgs`    | 游戏云                                         |
| ros    | `apis/ros`    | 对象存储                                       |
| domain | `apis/domain` | 域名                                           |
| rbm    | `apis/rbm`    | 裸金属                                         |
| ssl    | `apis/ssl`    | SSL 证书                                       |
| public | `apis/public` | 公共信息（页面配置、公告、节点状态、论坛数据） |

SDK 门面 `sdk.RainyunSDK` 内嵌以上服务，通过 `sdk.New(apiKey)` 或 `sdk.NewBuilder(apiKey).WithTrace(...).Build()` 构建。

## CLI 命令

| 命令         | 说明            | 子命令                                                                                                                                                                                                                   |
| ------------ | --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `server`     | 云服务器        | `list` `get` `start` `stop` `reboot` `reinstall` `reset-password` `vnc` `free` `set-tag` `create` `renew-price` `renew` `auto-renew` `upgrade` `edisk` `monitor` `backup` `eip` `nat` `traffic` `firewall` `pve-address` |
| `storage`    | 对象存储        | `list` `bucket`                                                                                                                                                                                                          |
| `domain`     | 域名            | `list` `dns`                                                                                                                                                                                                             |
| `public`     | 公共信息        | `app-config` `news` `status` `discourse`                                                                                                                                                                                 |
| `config`     | 配置管理        | `set` `show` `path`                                                                                                                                                                                                      |
| `completion` | 生成 shell 补全 | —                                                                                                                                                                                                                        |
| `version`    | 查看版本        | —                                                                                                                                                                                                                        |

## 开发

```bash
make test     # 运行全部测试（-race）
make vet      # 静态检查
make fmt      # gofmt -s + go mod tidy
make cover    # 生成覆盖率报告
make install  # 安装 CLI 到 GOPATH/bin
```

## 文档

- [`docs/CHANGELOG.md`](docs/CHANGELOG.md)：SDK 更新日志
- [`docs/CHANGELOG_cli.md`](docs/CHANGELOG_cli.md)：CLI 更新日志
- [`docs/PROGRESS.md`](docs/PROGRESS.md)：SDK/CLI 接入进度（`python3 docs/scripts/generate.py` 生成）

## License

[Mozilla Public License 2.0](LICENSE)
