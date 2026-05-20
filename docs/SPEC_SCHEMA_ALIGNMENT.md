# Spec ↔ schema ↔ snippets alignment (HDGP-Core)

> **Checked**: 2026-05-20 · **Scope**: Meta-only weaving defaults and field names.

## Summary

`spec/HDGP_INTEGRATION_SPEC.md` §2.4, `schemas/hdgp-core-meta.schema.json`, and `snippets/` are **aligned** on:

| Concept | Integration spec §2.4 | JSON Schema | TS / Python snippets |
|---------|----------------------|-------------|----------------------|
| Default `domain` | `general` | Documented in `Scene` description | `hdgpcDefaultScene()` / `hdgpc_default_scene()` |
| Default `intent` | `chat` | same | same |
| Default `risk_level` | `medium` | `RiskLevel` enum includes `medium` | same |
| Required root field | `meta.scene.*` recommended | `scene` required on root object | `scene` required in types |
| `risk_level` values | low / medium / high / critical (routing hints) | `enum` on `RiskLevel` | `HdgpcRiskLevel` union |

## Intentional notes (not conflicts)

- **Integration spec** also mentions Engine minimums (`subject.type`, `candidate.text`) — those are **out of scope** for the Core JSON Schema (Meta-only baseline).
- **Schema** allows custom `domain` / `intent` via `lowercase_snake` string pattern (`anyOf` enum + pattern); integration spec allows extension beyond listed enumerations. (`anyOf` avoids Draft-07 `oneOf` ambiguity when enum values also match the pattern.)
- **Examples** under `examples/` illustrate explicit `scene` values; omitted fields should be filled by application defaults per §2.4.

## Re-check trigger

Re-run alignment review when §2.4, the schema, or snippets change materially (CHIP / patch release).
