#!/usr/bin/env python3
"""Keep only endpoints whose tags intersect TAGS (plus their definitions).

Usage: python3 docs/scripts/filter_openapi.py SPEC.json "tag1,tag2" OUT.json
"""
import json
import sys


def collect_refs(obj):
    refs = set()

    def walk(x):
        if isinstance(x, dict):
            for k, v in x.items():
                if k == '$ref' and isinstance(v, str) and v.startswith('#/definitions/'):
                    refs.add(v[len('#/definitions/'):])
                else:
                    walk(v)
        elif isinstance(x, list):
            for i in x:
                walk(i)

    walk(obj)
    return refs


def main():
    spec = json.load(open(sys.argv[1], encoding='utf-8'))
    tags = set(sys.argv[2].split(','))
    paths = {}
    for p, ops in spec.get('paths', {}).items():
        for m, op in ops.items():
            if isinstance(op, dict) and (set(op.get('tags', [])) & tags):
                paths.setdefault(p, {})[m] = op
    todo = collect_refs(paths)
    done = set()
    defs = spec.get('definitions', {})
    while todo:
        n = todo.pop()
        if n in done or n not in defs:
            continue
        done.add(n)
        todo |= collect_refs(defs[n])
    out = {k: v for k, v in spec.items() if k not in ('paths', 'definitions')}
    out['paths'] = paths
    out['definitions'] = {n: defs[n] for n in sorted(done)}
    json.dump(out, open(sys.argv[3], 'w', encoding='utf-8'), ensure_ascii=False, indent=4)
    print(f'kept endpoints: {sum(len(v) for v in paths.values())}, definitions: {len(done)}')


if __name__ == '__main__':
    main()
