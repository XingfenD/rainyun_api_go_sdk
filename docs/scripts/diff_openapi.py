#!/usr/bin/env python3
"""Generate a human-readable diff report between two OpenAPI specs.

Usage: python3 docs/scripts/diff_openapi.py OLD.json NEW.json [OUTPUT.md]

Prints the report to stdout, or writes it to OUTPUT.md when given.
Exit code is 0 when the synced parts (paths + definitions) are identical,
1 when they differ.
"""
import json
import sys

MAX_ROWS = 50


def load(path):
    with open(path, encoding="utf-8") as f:
        return json.load(f)


def endpoints(spec):
    out = {}
    for path, ops in spec.get("paths", {}).items():
        for method, op in ops.items():
            if isinstance(op, dict):
                out[(method.upper(), path)] = op
    return out


def definitions(spec):
    return dict(spec.get("definitions", {}))


def short(v, limit=120):
    if v is None:
        return "—"
    s = json.dumps(v, ensure_ascii=False) if isinstance(v, (dict, list)) else str(v)
    return s if len(s) <= limit else s[:limit] + "…"


def summarize_op_change(old, new):
    lines = []
    if old.get("summary") != new.get("summary"):
        lines.append("- `summary`：{} → {}".format(short(old.get("summary")), short(new.get("summary"))))
    if old.get("description") != new.get("description"):
        lines.append("- `description`：已变更")
    for k in sorted(set(old) | set(new)):
        if k in ("summary", "description"):
            continue
        if k not in old:
            lines.append(f"- `{k}`：新增")
        elif k not in new:
            lines.append(f"- `{k}`：移除")
        elif old[k] != new[k]:
            if isinstance(old[k], (dict, list)):
                lines.append(f"- `{k}`：结构变更")
            else:
                lines.append(f"- `{k}`：{short(old[k])} → {short(new[k])}")
    return lines


def report(old_spec, new_spec):
    old_eps = endpoints(old_spec)
    new_eps = endpoints(new_spec)
    old_defs = definitions(old_spec)
    new_defs = definitions(new_spec)

    added_eps = sorted(set(new_eps) - set(old_eps))
    removed_eps = sorted(set(old_eps) - set(new_eps))
    modified_eps = sorted(
        (m, p) for m, p in (set(old_eps) & set(new_eps))
        if old_eps[(m, p)] != new_eps[(m, p)]
    )

    added_defs = sorted(set(new_defs) - set(old_defs))
    removed_defs = sorted(set(old_defs) - set(new_defs))
    modified_defs = sorted(
        n for n in (set(old_defs) & set(new_defs))
        if old_defs[n] != new_defs[n]
    )

    lines = []
    lines.append("## 变更摘要")
    lines.append("")
    lines.append(f"- 新增端点：{len(added_eps)}")
    lines.append(f"- 移除端点：{len(removed_eps)}")
    lines.append(f"- 修改端点：{len(modified_eps)}")
    lines.append(f"- 新增定义：{len(added_defs)}")
    lines.append(f"- 移除定义：{len(removed_defs)}")
    lines.append(f"- 修改定义：{len(modified_defs)}")
    lines.append("")

    def ep_table(rows, source):
        lines.append("| 方法 | 路径 | 说明 |")
        lines.append("|---|---|---|")
        for method, path in rows[:MAX_ROWS]:
            op = source.get((method, path), {})
            lines.append(f"| {method} | `{path}` | {op.get('summary', '')} |")
        if len(rows) > MAX_ROWS:
            lines.append(f"| … | 及另外 {len(rows) - MAX_ROWS} 个端点 | |")
        lines.append("")

    def def_table(rows, source):
        lines.append("| 名称 | 说明 |")
        lines.append("|---|---|")
        for name in rows[:MAX_ROWS]:
            lines.append(f"| `{name}` | {short(source[name].get('title')) if isinstance(source[name], dict) else ''} |")
        if len(rows) > MAX_ROWS:
            lines.append(f"| … | 及另外 {len(rows) - MAX_ROWS} 个定义 |")
        lines.append("")

    if added_eps:
        lines.append("## 新增端点")
        lines.append("")
        ep_table(added_eps, new_eps)

    if removed_eps:
        lines.append("## 移除端点")
        lines.append("")
        ep_table(removed_eps, old_eps)

    if modified_eps:
        lines.append("## 修改端点")
        lines.append("")
        for method, path in modified_eps:
            o, n = old_eps[(method, path)], new_eps[(method, path)]
            title = o.get("summary") or n.get("summary") or ""
            lines.append(f"### `{method} {path}` {title}".rstrip())
            lines.append("")
            lines.extend(summarize_op_change(o, n))
            lines.append("")

    if added_defs:
        lines.append("## 新增定义")
        lines.append("")
        def_table(added_defs, new_defs)

    if removed_defs:
        lines.append("## 移除定义")
        lines.append("")
        def_table(removed_defs, old_defs)

    if modified_defs:
        lines.append("## 修改定义")
        lines.append("")
        for name in modified_defs:
            lines.append(f"### `{name}`")
            lines.append("")
            lines.extend(summarize_op_change(old_defs[name], new_defs[name]))
            lines.append("")

    return "\n".join(lines)


def main():
    if len(sys.argv) < 3:
        print(__doc__)
        sys.exit(2)
    old = load(sys.argv[1])
    new = load(sys.argv[2])
    old_eps, new_eps = endpoints(old), endpoints(new)
    old_defs, new_defs = definitions(old), definitions(new)
    identical = (old_eps == new_eps) and (old_defs == new_defs)
    out = report(old, new)
    if len(sys.argv) > 3:
        with open(sys.argv[3], "w", encoding="utf-8") as f:
            f.write(out)
    else:
        print(out)
    sys.exit(0 if identical else 1)


if __name__ == "__main__":
    main()
