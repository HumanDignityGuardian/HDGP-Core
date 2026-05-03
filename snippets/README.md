# HDGP-Core snippets (optional)

Small **copy-paste** starters aligned with **`schemas/hdgp-core-meta.schema.json`**. They are **not** a published SDK: no package name, no runtime, no Engine/Judge calls.

| File | Notes |
|------|------|
| [`hdgp-meta.types.ts`](hdgp-meta.types.ts) | TypeScript interfaces + default scene helper |
| [`hdgp_meta.py`](hdgp_meta.py) | Python 3.11+ `TypedDict` + default scene helper |

Validate JSON payloads against the schema in CI if you need stricter checks than these types provide.
