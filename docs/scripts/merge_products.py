#!/usr/bin/env python3
"""Merge rcdn/rca/rvh domains from an old spec snapshot into docs/openapi.json.

Usage: python3 docs/scripts/merge_products.py OLD_SNAPSHOT.json
"""
import json
import re
import sys

OLD_TAGS = {'rcdn', 'rca', 'rca/app', 'rca/appstore', 'rca/project',
            'rca/service', 'rca/website', 'rvh'}


def norm_path(p):
    return re.sub(r':([A-Za-z_]+)', r'{\1}', p)


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
    if len(sys.argv) < 2:
        print(__doc__)
        sys.exit(1)
    old = json.load(open(sys.argv[1], encoding='utf-8'))
    spec = json.load(open('docs/openapi.json', encoding='utf-8'))

    sub_paths = {}
    for p, ops in old.get('paths', {}).items():
        np = norm_path(p)
        for m, op in ops.items():
            if not isinstance(op, dict):
                continue
            tags = set(op.get('tags', []))
            if tags & OLD_TAGS:
                op = dict(op)
                if tags & {'rcdn', 'rvh'}:
                    op['tags'] = ['rcdn' if 'rcdn' in tags else 'rvh']
                else:
                    op['tags'] = ['rca']
                sub_paths.setdefault(np, {})[m] = op

    todo = collect_refs(sub_paths)
    done = set()
    while todo:
        name = todo.pop()
        if name in done or name not in old.get('definitions', {}):
            continue
        done.add(name)
        todo |= collect_refs(old['definitions'][name])

    added_paths, added_defs, skipped_defs = 0, 0, 0
    for p, ops in sub_paths.items():
        if p in spec['paths']:
            raise SystemExit(f'path conflict: {p}')
        spec['paths'][p] = ops
        added_paths += 1
    for name in sorted(done):
        if name in spec['definitions']:
            skipped_defs += 1
            continue
        spec['definitions'][name] = old['definitions'][name]
        added_defs += 1

    with open('docs/openapi.json', 'w', encoding='utf-8') as f:
        json.dump(spec, f, ensure_ascii=False, indent=4)

    total = sum(len(v) for v in sub_paths.values())
    print(f'added paths: {added_paths}, endpoints: {total}')
    print(f'added definitions: {added_defs}, skipped (existing): {skipped_defs}')


if __name__ == '__main__':
    main()
