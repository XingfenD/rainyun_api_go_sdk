#!/usr/bin/env python3
"""Split docs/openapi.json by category and generate the SDK/CLI progress doc.

Usage: python3 docs/scripts/generate.py

Outputs:
  docs/openapi/<tag>.json   self-contained per-category OpenAPI fragments
  docs/PROGRESS.md          SDK / CLI / openapi.json integration progress
"""
import glob
import json
import os
import re

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

CATEGORY_CLI_DIR = {
    "rcs": "server",
    "ros": "storage",
    "domain": "domain",
    "rcdn": "rcdn",
    "rca": "rca",
    "rvh": "rvh",
}


def norm_path(p):
    p = re.sub(r"\{[^}]+\}", "{}", p)
    p = p.replace("%d", "{}").replace("%s", "{}").replace("%v", "{}")
    return p.rstrip("/")


def scan_sdk():
    rows = []
    for f in sorted(glob.glob(os.path.join(ROOT, "apis", "*", "*.go"))):
        if f.endswith("_test.go"):
            continue
        svc = f.split(os.sep)[-2]
        text = open(f, encoding="utf-8").read()
        for m in re.finditer(r"func \(s \*\w+Service\) (\w+)\([^)]*\)", text):
            funcname = m.group(1)
            body = text[m.end():]
            nxt = body.find("\nfunc ")
            body = body[:nxt] if nxt != -1 else body
            pm = re.search(r'path\s*:?=\s*(?:fmt\.Sprintf\("([^"]+)"|"([^"]+)")', body)
            raw = pm.group(1) or pm.group(2) if pm else None
            dm = re.search(r"Do\(constant\.HTTPMethod_(\w+),", body)
            if dm and raw:
                rows.append((svc, funcname, dm.group(1), raw))
    return rows


def scan_cli_calls():
    calls = set()
    for f in glob.glob(os.path.join(ROOT, "cmd", "ry", "commands", "*", "*.go")):
        if f.endswith("_test.go"):
            continue
        for m in re.finditer(r"\(\*rySDK\)\.(\w+)", open(f, encoding="utf-8").read()):
            calls.add(m.group(1))
    return calls


def scan_cli_uses(cli_dir):
    uses = []
    pat = os.path.join(ROOT, "cmd", "ry", "commands", cli_dir, "*.go")
    for f in glob.glob(pat):
        if f.endswith("_test.go"):
            continue
        for m in re.finditer(r'Use:\s+"([^"]+)"', open(f, encoding="utf-8").read()):
            uses.append(m.group(1))
    return sorted(set(uses))


def collect_refs(obj):
    refs = set()

    def walk(x):
        if isinstance(x, dict):
            for k, v in x.items():
                if k == "$ref" and isinstance(v, str) and v.startswith("#/definitions/"):
                    refs.add(v[len("#/definitions/"):])
                else:
                    walk(v)
        elif isinstance(x, list):
            for i in x:
                walk(i)

    walk(obj)
    return refs


def main():
    spec = json.load(open(os.path.join(ROOT, "docs", "openapi.json"), encoding="utf-8"))
    definitions = spec.get("definitions", {})
    paths = spec.get("paths", {})

    # group endpoint rows by tag
    tagged = {}
    for p, ops in paths.items():
        for m, op in ops.items():
            if not isinstance(op, dict):
                continue
            method = m.upper()
            for tag in op.get("tags", []):
                tagged.setdefault(tag, []).append((method, p, op.get("summary", "")))

    # ---- split files ----
    os.makedirs(os.path.join(ROOT, "docs", "openapi"), exist_ok=True)
    meta = {k: v for k, v in spec.items() if k in ("swagger", "info", "host", "schemes", "securityDefinitions")}
    for tag, rows in tagged.items():
        sub_paths = {}
        for method, p, _summary in rows:
            mkey = method.lower()
            sub_paths.setdefault(p, {})[mkey] = paths[p][mkey]
        refs = resolve(definitions, sub_paths)
        fragment = dict(meta)
        fragment["tags"] = [{"name": tag}]
        fragment["paths"] = sub_paths
        fragment["definitions"] = {n: definitions[n] for n in sorted(refs)}
        out = os.path.join(ROOT, "docs", "openapi", tag + ".json")
        json.dump(fragment, open(out, "w", encoding="utf-8"), ensure_ascii=False, indent=2)

    # ---- SDK / CLI scan ----
    sdk_rows = scan_sdk()
    sdk_index = {(norm_path(raw), method): (svc, func) for svc, func, method, raw in sdk_rows}
    cli_calls = scan_cli_calls()

    # ---- progress doc ----
    lines = []
    lines.append("# SDK / CLI / openapi.json 接入进度")
    lines.append("")
    lines.append("本文档追踪 Rainyun API（`docs/openapi.json`）在 Go SDK 与 `ry` CLI 中的接入进度。")
    lines.append("SDK 代码在 `apis/`，CLI 代码在 `cmd/ry/commands/`。")
    lines.append("")
    lines.append("> 重新生成：`python3 docs/scripts/generate.py`")
    lines.append("")

    total_eps = 0
    total_sdk = 0
    total_cli = 0
    summary = []
    summary.append("## 总览")
    summary.append("")
    summary.append("| 分类 | openapi 端点 | SDK 已实现 | CLI 已接入 |")
    summary.append("|---|---|---|---|")

    for tag in sorted(tagged):
        rows = tagged[tag]
        n = len(rows)
        n_sdk = sum(1 for method, p, _ in rows if (norm_path(p), method) in sdk_index)
        n_cli = sum(
            1
            for method, p, _ in rows
            if (norm_path(p), method) in sdk_index and sdk_index[(norm_path(p), method)][1] in cli_calls
        )
        total_eps += n
        total_sdk += n_sdk
        total_cli += n_cli
        summary.append(f"| {tag} | {n} | {n_sdk} | {n_cli} |")

    summary.append(f"| **合计** | **{total_eps}** | **{total_sdk}** | **{total_cli}** |")
    summary.append("")
    lines.extend(summary)

    for tag in sorted(tagged):
        rows = tagged[tag]
        lines.append(f"## {tag}")
        lines.append("")
        cli_dir = CATEGORY_CLI_DIR.get(tag)
        if cli_dir:
            uses = scan_cli_uses(cli_dir)
            if uses:
                lines.append("CLI 命令：`" + "`, `".join(uses) + "`")
                lines.append("")
        lines.append("| 方法 | 路径 | 说明 | SDK | CLI |")
        lines.append("|---|---|---|---|---|")
        for method, p, summary in sorted(rows, key=lambda r: (r[1], r[0])):
            key = (norm_path(p), method)
            hit = sdk_index.get(key)
            if hit:
                svc, func = hit
                sdk_cell = f"`{svc}.{func}`"
                cli_cell = "✓" if func in cli_calls else "—"
            else:
                sdk_cell = "—"
                cli_cell = "—"
            lines.append(f"| {method} | `{p}` | {summary} | {sdk_cell} | {cli_cell} |")
        lines.append("")

    # services in the SDK but not covered by openapi.json
    sdk_svcs = sorted({svc for svc, _, _, _ in sdk_rows})
    spec_tags = set(tagged)
    missing = [s for s in sdk_svcs if s not in spec_tags]
    if missing:
        lines.append("## 未纳入 openapi.json 的 SDK 服务")
        lines.append("")
        lines.append("以下 SDK 服务已存在，但当前 `docs/openapi.json` 未包含对应端点，暂无法核对：")
        lines.append("")
        lines.append(", ".join(f"`{s}`" for s in missing))
        lines.append("")

    open(os.path.join(ROOT, "docs", "PROGRESS.md"), "w", encoding="utf-8").write("\n".join(lines))

    print(f"split files: {sorted(tagged.keys())}")
    print(f"endpoints={total_eps} sdk={total_sdk} cli={total_cli}")
    print(f"SDK services missing from openapi: {missing}")


def resolve(definitions, sub_paths):
    todo = collect_refs(sub_paths)
    done = set()
    while todo:
        name = todo.pop()
        if name in done or name not in definitions:
            continue
        done.add(name)
        todo |= collect_refs(definitions[name])
    return done


if __name__ == "__main__":
    main()
