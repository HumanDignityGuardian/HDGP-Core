#!/usr/bin/env python3
"""Validate JSON files and hdgp-core-meta.schema.json examples (Meta-only)."""

from __future__ import annotations

import json
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SCHEMA_PATH = ROOT / "schemas" / "hdgp-core-meta.schema.json"
EXAMPLES_DIR = ROOT / "examples"


def load_json(path: Path) -> object:
    with path.open(encoding="utf-8") as f:
        return json.load(f)


def main() -> int:
    try:
        from jsonschema import Draft7Validator
    except ImportError:
        print("jsonschema package required (pip install jsonschema)", file=sys.stderr)
        return 1

    schema = load_json(SCHEMA_PATH)
    validator = Draft7Validator(schema)

    errors: list[str] = []

    for path in sorted(ROOT.rglob("*.json")):
        if ".git" in path.parts:
            continue
        try:
            load_json(path)
        except json.JSONDecodeError as exc:
            errors.append(f"{path.relative_to(ROOT)}: invalid JSON — {exc}")

    for path in sorted(EXAMPLES_DIR.glob("*.json")):
        instance = load_json(path)
        for err in sorted(validator.iter_errors(instance), key=str):
            errors.append(f"{path.relative_to(ROOT)}: schema — {err.message}")

    if errors:
        for msg in errors:
            print(msg, file=sys.stderr)
        return 1

    print(
        f"OK: JSON parse check + {len(list(EXAMPLES_DIR.glob('*.json')))} example(s) validated against schema"
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
