#!/usr/bin/env python3
"""Merge upstream live spec with repo-only domains into one spec.

Usage: python3 docs/scripts/merge_openapi.py LIVE.json EXTRA.json OUT.json
"""
import json
import sys


def main():
    live = json.load(open(sys.argv[1], encoding='utf-8'))
    extra = json.load(open(sys.argv[2], encoding='utf-8'))
    for p, ops in extra.get('paths', {}).items():
        if p in live['paths']:
            raise SystemExit(f'path conflict with upstream: {p}')
        live['paths'][p] = ops
    for name, d in extra.get('definitions', {}).items():
        live['definitions'].setdefault(name, d)
    json.dump(live, open(sys.argv[3], 'w', encoding='utf-8'), ensure_ascii=False, indent=4)


if __name__ == '__main__':
    main()
