# Semantic Boundary: Meta vs Judge/Audit (normative anchor)

> **Purpose**: fix the semantic layering in the mainline (`HDGP-Protocol`) to avoid conflating “Meta guidance fields” with “hard rule execution / audit persistence”, and to provide a single normative anchor for **`HDGP-Core` (Meta-only)** materials selection.  
> **Related**: the Open Framework overview is in `HDGP_OPEN_FRAMEWORK.md`; HTTP/field details are in `HDGP_ENGINE_API_SPEC.md` (mainline).

---

## 1. Definitions

| Term | Responsibility | Typical artifact (mainline) |
|------|------|----------------------|
| **Meta (guidance layer)** | Semantics and contracts such as scene, risk, accountable roles, escalation and human-review suggestions; **not** equivalent to automated enforcement | `meta` field family; “input-side” descriptions in ethics baseline and mapping specs |
| **Judge (real-time rule evaluation)** | Executes hard verdicts on candidate outputs, such as **allow / modify / block / fuse** | Engine `/hdgp/v1/evaluate`, `/hdgp/v1/chat`, policy bundles |
| **Audit (evidence)** | Operational semantics for structured records, hash chains, retention policies, and forensics | `/hdgp/v1/audit`, audit persistence, G3 evidence streams |

**Boundary sentence**: Meta **describes risk and context**; Judge **executes policy**; Audit **preserves traceability and verifiability**. Adopters may implement **Meta only** (field weaving) without deploying any Engine from this repository; if Judge/Audit is enabled, corresponding specs/runbooks apply.

---

## 2. Spec document ownership (read-only index)

| Topic | Primary spec | Meta-related | Judge/Audit-related |
|------|------------|-----------|------------------|
| Ethics & principles | `spec/HDGP_ETHICS_BASELINE.md` | Baseline clauses, disclaimers | Anti-hijacking and change-governance sections |
| Principles → executable mapping | `spec/HDGP_CORE_MAPPING_SPEC.md` | Mapping relations and semantics | Rule tiers and implementation constraints |
| Integration semantics | `spec/HDGP_INTEGRATION_SPEC.md` | Meta, scenes, routing semantics | Engine-facing API surface, errors, circuit breaking |
| Engine API | `spec/HDGP_ENGINE_API_SPEC.md` | `meta` structure and recommended fields | Verdicts, audit fields, integrity |
| Kernel checklist | `spec/HDGP_KERNEL_CHECKLIST.md` | Governance and test traceability | Release and audit gates |

**Selection note**: for **Meta-only** Core extraction, prefer **excerpts** from the “Meta-related” column; do not bulk-sync full Judge-heavy specs by default.

---

## 3. Relationship to HDGP-Core open-source

- The **mainline** maintains this file and the index above; the existence of the open-source Core repository does **not** create an automatic bi-directional code sync obligation.  
- After gate **G** is satisfied, finalize the allowlist and extraction snapshot per the mainline process.

---

## 4. Revisions

Semantic changes should follow **`docs/CHIP_PROCESS.md`**. For typos or link fixes, a normal documentation PR is sufficient.

---

# Meta 指导层与 Judge / Audit 职责边界（主系统规范锚点）(ZH-CN)

> **目的**：在 `HDGP-Protocol` 内固定 **语义分层**，避免将「伦理指导字段」与「硬规则执行 / 审计落盘」混为一谈；并为日后 **`HDGP-Core`（Meta-only）** 拣选材料提供 **规范层单一事实源**。  
> **关系**：Open Framework 总览见根目录 **`HDGP_OPEN_FRAMEWORK.md`**；HTTP 与字段细节见 **`HDGP_ENGINE_API_SPEC.md`**。

---

## 1. 定义

| 概念 | 职责 | 典型载体（主系统） |
|------|------|----------------------|
| **Meta 指导层** | 场景、风险、责任主体、升级与人工复核建议等 **语义与契约**；**不**等同于自动执法 | `meta` 字段族、伦理基线与映射规范中的「输入侧」描述 |
| **Judge（实时规则评估）** | 在候选输出上执行 **allow / modify / block / fuse** 等 **硬判定** | Engine `/hdgp/v1/evaluate`、`/hdgp/v1/chat` 与规则包 |
| **Audit（审计证据）** | 结构化记录、哈希链、留存策略与取证 **运维语义** | `/hdgp/v1/audit`、审计落盘、G3 证据流 |

**边界句**：Meta **描述风险与语境**；Judge **执行策略**；Audit **留痕与可验证性**。采用方可以 **只实现 Meta**（指导字段编织）而不部署本仓库 Engine；若启用 Judge/Audit，须遵循对应 `spec/` 与 Runbook。

---

## 2. 规范文档归属（只读索引）

| 主题 | 主规范文件 | Meta 相关 | Judge/Audit 相关 |
|------|------------|-----------|------------------|
| 伦理与原则 | `spec/HDGP_ETHICS_BASELINE.md` | 基线条款、免责声明 | 与内核防挟持、变更治理相关章节 |
| 原则 → 可执行映射 | `spec/HDGP_CORE_MAPPING_SPEC.md` | 映射关系与语义 | 与规则分级、实现约束相关节 |
| 集成与网关 | `spec/HDGP_INTEGRATION_SPEC.md` | meta、场景、路由语义 | Engine 调用面、错误与熔断 |
| Engine API | `spec/HDGP_ENGINE_API_SPEC.md` | `meta` 结构与推荐字段 | 请求/响应 verdict、审计字段、完整性 |
| 内核清单 | `spec/HDGP_KERNEL_CHECKLIST.md` | 治理与测试可追溯 | 发布与审计门禁 |

**拣选提示**：面向 Core 的 **Meta-only** 拉取，优先从表中「Meta 相关」列对应文件 **节选**；全文含 Judge 的规范勿整体默认同步（见 `docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md`）。

---

## 3. 与对外开源（HDGP-Core）的关系

- **主系统**维护本文件与上表索引；**不**因开源仓存在而自动双向同步代码（见 `docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md` §五）。
- 达到 **`docs/HDGP_MAINLINE_BASELINE_GATE_CHECKLIST.md`** 门槛 **G** 后，再 **定稿** 拣选白名单与提取快照（见该文档阶段 **C**）。

---

## 4. 修订

语义变更应走 **`docs/CHIP_PROCESS.md`**；若仅修正笔误或链接，可通过文档 PR 说明。

