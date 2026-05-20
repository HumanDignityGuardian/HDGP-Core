# HDGP-Core Meta examples (Meta-only)

Minimal example payloads aligned with `schemas/hdgp-core-meta.schema.json`.  
Defaults when `scene` fields are omitted: `domain=general`, `intent=chat`, `risk_level=medium` (see `spec/HDGP_INTEGRATION_SPEC.md` §2.4).

| File | Scenario |
|------|----------|
| `meta.general.chat.json` | General domain, chat intent, medium risk |
| `meta.medical.advice.json` | Medical domain, advice intent, high risk + optional adopter disclosure |

No verdict, Judge, or audit-chain fields.
