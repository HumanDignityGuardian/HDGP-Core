# Meta guidance layer vs Judge / Audit responsibilities

> **Purpose**: Fix **semantic layering** so “ethics guidance fields” are not conflated with “hard rule enforcement / audit persistence”. This copy in **HDGP-Core** is the Meta-only normative anchor for adopters; HTTP and field details for a full Engine live in the private mainline’s `spec/HDGP_ENGINE_API_SPEC.md` (not shipped in Core).  
> **Related**: Open Framework overview: **`HDGP_OPEN_FRAMEWORK.md`**.

---

## 1. Definitions

| Concept | Responsibility | Typical carrier (private mainline) |
|--------|----------------|-------------------------------------|
| **Meta guidance layer** | Semantics & contracts: scenario, risk, accountable party, escalation & human review hints; **not** automatic enforcement | `meta` field family; “input side” text in ethics baseline & mapping specs |
| **Judge (real-time rule evaluation)** | **Hard decisions** on candidate outputs: allow / modify / block / fuse | Engine `/hdgp/v1/evaluate`, `/hdgp/v1/chat`, rule bundles |
| **Audit** | Structured records, hash chains, retention, forensics **ops semantics** | `/hdgp/v1/audit`, persistence, G3 evidence flows |

**Boundary sentence**: Meta **describes risk and context**; Judge **executes policy**; Audit **records and enables verification**. Adopters may implement **Meta only** (weaving guidance fields) without deploying the mainline Engine; if Judge/Audit are enabled, follow the corresponding `spec/` and runbooks in the mainline.

---

## 2. Spec ownership (read-only index)

| Topic | Canonical spec | Meta-related | Judge/Audit-related |
|-------|----------------|--------------|---------------------|
| Ethics & principles | `spec/HDGP_ETHICS_BASELINE.md` | Baseline clauses, disclaimers | Kernel anti-capture, change governance |
| Principles → executable mapping | `spec/HDGP_CORE_MAPPING_SPEC.md` | Mapping semantics | Rule tiers, implementation constraints |
| Integration & gateway | `spec/HDGP_INTEGRATION_SPEC.md` | `meta`, scenario, routing semantics | Engine surface, errors, circuit-breaking |
| Engine API | `spec/HDGP_ENGINE_API_SPEC.md` (mainline) | `meta` structure & recommended fields | Verdict, audit fields, integrity |
| Kernel checklist | `spec/HDGP_KERNEL_CHECKLIST.md` (mainline) | Governance & traceability | Release & audit gates |

**Pick guidance**: For Core **Meta-only** pulls, prefer **excerpts** from the “Meta-related” column; do not wholesale-sync specs dominated by Judge narratives unless explicitly allowlisted.

---

## 3. Relationship to HDGP-Core

- The **private mainline** maintains the canonical index above; there is **no** standing obligation of bidirectional code sync with Core (see mainline `docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md` §5).  
- After **Gate G** in `docs/HDGP_MAINLINE_BASELINE_GATE_CHECKLIST.md`, finalize the allowlist and extraction snapshot (**stage C** in that document).

---

## 4. Revisions

Semantic changes follow **`docs/CHIP_PROCESS.md`**; typo/link fixes may use a normal docs PR with a short rationale.

---

## 中文版本 (ZH-CN)

> 以下中文与上文英文对应；社区阅读顺序以英文为先。

# Meta 指导层与 Judge / Audit 职责边界（主系统规范锚点）

> **目的**：在 `HDGP-Protocol` 内固定 **语义分层**，避免将「伦理指导字段」与「硬规则执行 / 审计落盘」混为一谈；并为日后 **`HDGP-Core`（Meta-only）** 拣选材料提供 **规范层单一事实源**。  
> **关系**：Open Framework 总览见根目录 **`HDGP_OPEN_FRAMEWORK.md`**；HTTP 与字段细节见 **`spec/HDGP_ENGINE_API_SPEC.md`**。

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
