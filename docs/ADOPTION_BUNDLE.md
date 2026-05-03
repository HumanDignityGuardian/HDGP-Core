# HDGP-Core Adoption Bundle (Meta-only) — v1.0.x

This document is a **pack-and-go entry** for adopting **HDGP-Core** as a **Meta-only** ethics layer during development — without bundling a runnable **Engine/Judge**.

> **Non-negotiable boundary (first impression)**: **HDGP-Core is Meta-only.**  
> It provides **semantics + contracts + governance**. It does **not** ship a Judge/Engine, release gates, or audit/evidence hosting.  
> If you extend Core materials into your own Engine/Judge, that extension is **your implementation**, under your accountability.

---

## What you can “take and weave”

- **Meta contracts** (domain / intent / risk level and related semantics) for routing, logging, and review hooks  
- **Ethics baseline** constraints for HDGP itself: `spec/HDGP_ETHICS_BASELINE.md`  
- **Boundary index** (Meta vs Judge/Audit): `spec/HDGP_META_VS_JUDGE_SCOPE.md`  
- **Integration semantics** (Meta-focused): `spec/HDGP_INTEGRATION_SPEC.md`  
- **Governance & change process**: `GOVERNANCE.md`, `docs/CHIP_PROCESS.md`
- **Machine-readable Meta shape (Draft-07 JSON Schema)**: `schemas/hdgp-core-meta.schema.json` — optional validator for weaving `scene` / `policy` hints; **not** an Engine/Judge API or verdict contract (see schema `$id` and descriptions).
- **Starter snippets (TypeScript / Python)**: `snippets/` — copy-paste types/helpers; **not** a published SDK.
- **Adoption statement template**: `docs/ADOPTION_STATEMENT_TEMPLATE.md` — optional wording for your project docs (attribution + Meta-only boundary).

---

## Machine-readable contract (optional)

For tooling and CI, you may validate or generate structs from **`schemas/hdgp-core-meta.schema.json`** (JSON Schema Draft-07). It describes a **Meta-only** object: required `scene` (with `domain`, `intent`, `risk_level` patterns aligned to `spec/HDGP_INTEGRATION_SPEC.md` §2.4 defaults and documented enumerations). It deliberately **excludes** verdict, audit chain, enforcement, or runtime Judge fields.

- **In-repo path**: `schemas/hdgp-core-meta.schema.json`  
- **Stable raw URL** (for `$ref` / remote validators): `https://raw.githubusercontent.com/HumanDignityGuardian/HDGP-Core/main/schemas/hdgp-core-meta.schema.json`

---

## Starter snippets (optional)

The **`snippets/`** directory contains minimal **TypeScript** and **Python** shapes aligned with the JSON Schema, plus default `scene` values from `spec/HDGP_INTEGRATION_SPEC.md` §2.4. Copy into your codebase as needed; there is **no** npm/PyPI package from this repository.

---

## Adoption statement template (optional)

If you want a short, human-readable declaration next to your attribution block, start from **`docs/ADOPTION_STATEMENT_TEMPLATE.md`** (English + Chinese). It is **not** a substitute for the license or legal advice.

---

## What you must NOT assume this repo provides

- Any **runnable** Judge/Engine implementation  
- Any **runtime** verdict enforcement, circuit-breaking, or “default safe behavior” guarantees  
- Any audit storage, evidence chain hosting, compliance certification, or operations gates  
- Any commercial SLAs or managed services

If you need those, they belong to a separate mainline/commercial track — not this Meta-only baseline.

---

## If you build an Engine/Judge from these materials

You may implement your own runtime layer on top of Meta contracts — but:

- **Scope**: your Engine/Judge is **not** “provided by HDGP-Core” unless you explicitly ship and maintain it.  
- **Accountability**: behavior, safety, compliance, and incidents are **your** responsibility.  
- **No implied warranty**: HDGP-Core does not warrant your implementation outcomes.

If issues arise, you may **contact the maintainer** for discussion and coordination, but this is **not a support / warranty / after-sales service** commitment.

Contact: Yvaine He · `xyan8921@gmail.com` (see `README.md` / `SECURITY.md`).

---

## Attribution & open-source rules (required)

When you adopt or excerpt HDGP-Core materials:

- **Follow the repository license** (`LICENSE`).  
- Keep **attribution** and include a **link back** to this repository in your project docs, e.g.:  
  - `https://github.com/HumanDignityGuardian/HDGP-Core`  
  - or the GitHub Pages index: `https://humandignityguardian.github.io/HDGP-Core/`

This is not an “endorsement”; it is **basic attribution and traceability**.

---

## 中文版本 (ZH-CN)

> 以下中文与上文英文对应；社区阅读顺序以英文为先。

# HDGP-Core 采用入口包（Meta-only）— v1.0.x

本文档是采用 **HDGP-Core** 的“**打包带走入口**”：用于在开发期把 **Meta-only** 伦理层以**编织式**方式接入你的系统，而不引入可运行的 **Engine/Judge**。

> **不可动摇的边界（第一印象）**：**HDGP-Core 仅为 Meta-only。**  
> 它提供 **语义 + 契约 + 治理**，**不**提供 Judge/Engine、发布门禁、审计/证据链托管。  
> 若你基于 Core 材料扩展并实现自己的 Engine/Judge，那是**你的实现**，由你承担责任。

---

## 你可以“带走并编织”的内容

- **Meta 契约**（domain / intent / risk_level 等语义）用于路由、日志、复核钩子  
- HDGP 自身伦理基线：`spec/HDGP_ETHICS_BASELINE.md`  
- 语义边界索引（Meta vs Judge/Audit）：`spec/HDGP_META_VS_JUDGE_SCOPE.md`  
- 集成语义（以 Meta 为中心）：`spec/HDGP_INTEGRATION_SPEC.md`  
- 治理与变更流程：`GOVERNANCE.md`、`docs/CHIP_PROCESS.md`
- **机器可读的 Meta 形态（Draft-07 JSON Schema）**：`schemas/hdgp-core-meta.schema.json` — 可选校验/代码生成；**不是** Engine/Judge API 或 verdict 契约（见 schema 内 `$id` 与说明）。
- **起步片段（TypeScript / Python）**：`snippets/` — 可复制类型与默认值辅助；**不是**已发布的 SDK。
- **采用声明模板**：`docs/ADOPTION_STATEMENT_TEMPLATE.md` — 可选文案（致谢 + Meta-only 边界）。

---

## 机器可读契约（可选）

若需在工具链或 CI 中校验/生成结构体，可使用 **`schemas/hdgp-core-meta.schema.json`**（JSON Schema Draft-07）。它描述 **Meta-only** 对象：必填 `scene`（`domain`、`intent`、`risk_level` 与 `spec/HDGP_INTEGRATION_SPEC.md` §2.4 默认值及文档枚举对齐）。**不包含** verdict、审计链、执法或运行时 Judge 字段。

- **仓内路径**：`schemas/hdgp-core-meta.schema.json`  
- **稳定 raw URL**（供 `$ref` 或远程校验器）：`https://raw.githubusercontent.com/HumanDignityGuardian/HDGP-Core/main/schemas/hdgp-core-meta.schema.json`

---

## 起步片段（可选）

**`snippets/`** 提供与 JSON Schema 对齐的 **TypeScript** 与 **Python** 最小结构，以及 `spec/HDGP_INTEGRATION_SPEC.md` §2.4 的默认 `scene`。按需复制到工程即可；本仓库**不**发布 npm/PyPI 包。

---

## 采用声明模板（可选）

若希望在致谢旁附一段可读声明，可从 **`docs/ADOPTION_STATEMENT_TEMPLATE.md`**（英 / 中）起稿。**不能**替代许可证或法务意见。

---

## 你不得默认认为本仓提供的能力

- 任何**可运行**的 Judge/Engine 实现  
- 任何运行时 verdict 执法、熔断或“默认安全”保证  
- 任何审计落盘、证据链托管、合规认证或运维门禁  
- 任何商业 SLA 或托管服务

如需上述能力，应属于独立的主系统/商业轨道，而非本 Meta-only 基线。

---

## 如果你据此构建 Engine/Judge

你可以在 Meta 契约之上实现自己的运行时层，但需明确：

- **范围**：除非你明确交付并维护，否则你的 Engine/Judge **不**属于 “HDGP-Core 自带能力”。  
- **责任**：行为、安全、合规与事故由**集成方/实现方**承担。  
- **无默示担保**：HDGP-Core 不对你的实现结果作出担保。

后续如出现问题，你可以**联系维护者**进行讨论与协调，但这**不构成售后/保修/支持承诺**。

联系人：Yvaine He · `xyan8921@gmail.com`（见 `README.md` / `SECURITY.md`）。

---

## 致谢与开源规则（必须遵守）

采用或节选 HDGP-Core 材料时：

- **遵守仓库许可证**（`LICENSE`）。  
- 保留**署名与致谢**，并在你的项目文档中提供**回链**，例如：  
  - `https://github.com/HumanDignityGuardian/HDGP-Core`  
  - 或 GitHub Pages：`https://humandignityguardian.github.io/HDGP-Core/`

这不是“背书”，而是**基本的署名与可追溯性要求**。

