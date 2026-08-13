# Repository Guidelines

## Project Structure & Module Organization

This Go module (`github.com/XingfenD/rainyun_api_go_sdk`) provides Rainyun API clients and the `ry` CLI. API services live in `apis/`, grouped by product (`rcs`, `rgs`, `ros`, `domain`, and others); shared request/query helpers are in `apis/common/`. The public SDK facade is in `sdk/`, constants in `constant/`, and CLI code in `cmd/ry/`: commands in `commands/`, implementation-only packages in `internal/`. Keep release notes in `docs/CHANGELOG.md` and user-facing documentation in `docs/`.

## Build, Test, and Development Commands

- `make build` builds `bin/ry` with version and commit metadata.
- `make test` runs all Go tests with the race detector (`go test -race ./...`).
- `make vet` runs static checks with `go vet ./...`.
- `make fmt` applies `gofmt -s` and tidies modules; review resulting `go.mod`/`go.sum` changes.
- `make cover` writes `coverage.out` and an inspectable `coverage.html` report.
- `make install` installs the CLI to `GOPATH/bin`; use `make clean` only to remove generated local artifacts.

Use Go 1.24.1 or later, as declared in `go.mod`.

## Coding Style & Naming Conventions

Format all Go changes with `gofmt -s`; use tabs and standard Go import grouping. Follow existing package naming: concise, lowercase directory and package names such as `apis/workorder`. Exported identifiers use `PascalCase`; unexported identifiers use `camelCase`. Keep one API area per `apis/<service>/` package and place its client setup in `client.go`. Prefer explicit request/response model names and JSON tags that match the API contract.

## Testing Guidelines

Write standard-library `testing` tests in adjacent `*_test.go` files. Name cases `Test<Subject>` (for example, `TestMarshalQueryParams`) and include error paths and boundary/default behavior. Use `t.TempDir()` for file-backed tests; do not depend on user configuration or live API credentials. Run `make test` before submitting, and run `make cover` when changing complex serialization, configuration, or output logic.

## Commit & Pull Request Guidelines

Recent history uses Conventional Commit-style messages such as `feat(cli): add interactive config input`, `fix(cmd/ry): ...`, and `docs: ...`. Use a concise imperative subject and an appropriate scope. PRs should explain the behavior change, list validation performed, link relevant issues, and include CLI output or screenshots when user-visible behavior changes.

Before every commit, run `git branch --show-current`. Commit only on a development branch whose name contains `feat`, `fix`, `refactor`, or `docs`; otherwise ask for a branch name and create one first. Core changes require an entry in `docs/CHANGELOG.md`, with newer versions above older ones.
