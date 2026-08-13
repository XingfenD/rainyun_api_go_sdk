#!/usr/bin/env python3
"""Generate a human-readable diff report between two OpenAPI specs.

Usage: python3 docs/scripts/diff_openapi.py OLD.json NEW.json [OUTPUT.md]

Prints the report to stdout, or writes it to OUTPUT.md when given.
Exit code is 0 when the specs are identical, 1 when they differ.
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
    old = endpoints(old_spec)
    new = endpoints(new_spec)
    added = sorted(set(new) - set(old))
    removed = sorted(set(old) - set(new))
    modified = sorted((m, p) for m, p in (set(old) & set(new)) if old[(m, p)] != new[(m, p)])

    lines = []
    lines.append("## 变更摘要")
    lines.append("")
    lines.append(f"- 新增端点：{len(added)}")
    lines.append(f"- 移除端点：{len(removed)}")
    lines.append(f"- 修改端点：{len(modified)}")
    lines.append("")

    def table(rows, source):
        lines.append("| 方法 | 路径 | 说明 |")
        lines.append("|---|---|---|")
        for method, path in rows[:MAX_ROWS]:
            op = source.get((method, path), {})
            lines.append(f"| {method} | `{path}` | {op.get('summary', '')} |")
        if len(rows) > MAX_ROWS:
            lines.append(f"| … | 及另外 {len(rows) - MAX_ROWS} 个端点 | |")
        lines.append("")

    if added:
        lines.append("## 新增端点")
        lines.append("")
        table(added, new)

    if removed:
        lines.append("## 移除端点")
        lines.append("")
        table(removed, old)

    if modified:
        lines.append("## 修改端点")
        lines.append("")
        for method, path in modified:
            o, n = old[(method, path)], new[(method, path)]
            title = o.get("summary") or n.get("summary") or ""
            lines.append(f"### `{method} {path}` {title}".rstrip())
            lines.append("")
            lines.extend(summarize_op_change(o, n))
            lines.append("")

    return "\n".join(lines)


def main():
    if len(sys.argv) < 3:
        print(__doc__)
        sys.exit(2)
    old = load(sys.argv[1])
    new = load(sys.argv[2])
    identical = old == new
    out = report(old, new)
    if len(sys.argv) > 3:
        with open(sys.argv[3], "w", encoding="utf-8") as f:
            f.write(out)
    else:
        print(out)
    sys.exit(0 if identical else 1)


if __name__ == "__main__":
    main()
