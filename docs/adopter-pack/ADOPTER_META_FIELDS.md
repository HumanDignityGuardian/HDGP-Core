# Recommended Adopter Meta Fields (HDGP-Core · Meta-only)

This document suggests a **minimal, comparable** set of Meta fields for external adopters. It does **not** change HDGP-Core’s boundary: Meta-only semantics.

> **Design rule**: keep fields **small and stable**, so different adopters remain comparable. Put project-specific data under namespaced keys, but keep the core set consistent.

## 1) Minimal required (MUST)

These are the smallest fields that make Meta useful for routing, logging, and governance discussions.

- **`scene.domain`** (MUST)  
  - **Meaning**: which domain the output path belongs to.  
  - **Default** (if missing): `general`  
  - **Reference**: `schemas/hdgp-core-meta.schema.json` and `spec/HDGP_INTEGRATION_SPEC.md` §2.4

- **`scene.intent`** (MUST)  
  - **Meaning**: what the system is trying to do on this output path (chat/advice/notification/decision_support…).  
  - **Default**: `chat`

- **`scene.risk_level`** (MUST)  
  - **Meaning**: routing tier for human review hints and conservative behavior in high stakes contexts.  
  - **Enum**: `low | medium | high | critical`  
  - **Default**: `medium`

## 2) Strongly recommended (SHOULD)

- **`request_id`** (SHOULD)  
  - Correlation id for logs, human review and later dispute handling.

- **`channel`** (SHOULD)  
  - Output surface label (e.g. `web`, `api`, `batch`, `email`, `mobile`).

- **`locale`** (SHOULD)  
  - BCP-47 locale tag for reviewer context and UX copy.

- **`actor.type`** (SHOULD)  
  - `end_user | operator | system | ...` (adopter-defined but stable within your system).

- **`actor.role`** (MAY/SHOULD)  
  - Optional role in your product (e.g. `admin`, `moderator`, `guest`).

- **`scene.sensitivity[]`** (SHOULD)  
  - Tags for privacy/safety sensitivity (adopter-defined). Example tags: `personal_data`, `medical_info`, `minors`, `financial_advice`.

## 3) Traceability / future routing hints (MAY)

These fields are **semantic references only** (no promise that any Engine evaluates them).

- **`policy.spec_version`** (MAY): e.g. `HDGP-1.0`  
- **`policy.strategy_id`** (MAY): your internal strategy/profile id  
- **`policy.bundles[]`** (MAY): referenced bundle ids (for docs/traceability)  
- **`policy.override_flags[]`** (MAY): non-normative flags you record

## 3.5) Transparency & provenance (MAY, recommended for content/factual outputs)

These fields are **adopter-side** (Meta-only) additions that help reduce governance ambiguity in common “content site / compliance intake / documentation” scenarios.

- **`adopter.disclosure`** (MAY)  
  - Relationship/promo transparency (sponsorship, affiliate, ownership ties), plus data-boundary disclosures.
- **`adopter.sources`** (MAY)  
  - Provenance/source hints for factual claims (citations list, thresholds, timestamps, “no-sources” flags).

## 4) Adopter extension pattern (RECOMMENDED)

If you need additional fields, prefer one of:

- **Namespaced top-level key**: `adopter.<your_org>.<field>`  
- **Namespaced scene key**: `scene.<your_org>_<field>`  

Avoid inventing a completely different `scene`/`risk` hierarchy; keep the MUST set stable.

