## HDGP-Core（Meta-only）：Human Dignity Guardian Protocol · Open Meta Baseline

This repository (`HDGP-Core`) provides an **open-source Meta ethics & safety baseline**: semantics, contracts, and community governance materials (Meta / ethics / mapping / governance). It helps adopters **weave Meta guidance fields** into their systems and establish an “ethical input contract” that is discussable and auditable.

> **Release `v1.0.0` · 2026-05-03** — first stable Meta-only baseline for this repository (**Git tag [`v1.0.0`](https://github.com/HumanDignityGuardian/HDGP-Core/releases/tag/v1.0.0)**). Formal mainline handoff acceptance & freeze are summarized in **`docs/CORE_EXTRACT_SNAPSHOT.md`** (cross-reference to mainline `docs/HDGP_CORE_FINAL_HANDOFF_ACCEPTANCE_2026-05-03.md`, not shipped in Core).

### Website & contact

- **GitHub Pages**: `https://humandignityguardian.github.io/HDGP-Core/` — lightweight index for this repository (links into specs & governance on GitHub).  
- **Maintainer**: Yvaine He · `xyan8921@gmail.com` (general repository inquiries; use GitHub **Private vulnerability reporting** first for security issues — see `SECURITY.md`).  
- **Product / mainline narrative** (Engine demos, PFaaS context — not part of Core materials): [hdgp-protocol.com](https://www.hdgp-protocol.com/).

### Who Core is for · 面向谁（工程师与 Agent 基座，同一套规范）

- **Product / application engineers** — Weave **Meta** fields at development time (domain, intent, risk, etc.); start with `spec/HDGP_INTEGRATION_SPEC.md` and `spec/HDGP_META_VS_JUDGE_SCOPE.md`.  
- **Agent-base / platform builders** — Use the same contracts as **schema and routing hints** for tool and policy layers. Core does **not** ship a Judge/Engine; staying aligned is **semantic** (traceable, discussable) rather than claiming runtime enforcement.

You do **not** need two separate documentation trees in this repository: **one Meta-only baseline**, **two integration personas** — same specs, different glue code.

### Pack-and-go adoption entry (Meta-only)

- Start here: **`docs/ADOPTION_BUNDLE.md`** — single-page “take-and-weave” entry with **hard boundary** (no Engine/Judge), attribution rules, and contact.

### Open Letters & White Paper (readable editions on main site)

Canonical **styled** HTML (permanent record on the product site; not duplicated as HTML in Core):

- **Open Letters** — [hdgp-protocol.com · letters](https://www.hdgp-protocol.com/docs/readable/letters.html) · Markdown mirror in repo: [`LETTERS.md`](LETTERS.md)  
- **Technical White Paper v1.0** — [hdgp-protocol.com · whitepaper](https://www.hdgp-protocol.com/docs/readable/whitepaper.html)

The white paper describes the **global** HDGP framework (including tracks beyond this repo’s Meta-only scope). For **what Core commits to**, use [`MATERIALS_ALLOWLIST.md`](MATERIALS_ALLOWLIST.md) and [`spec/HDGP_META_VS_JUDGE_SCOPE.md`](spec/HDGP_META_VS_JUDGE_SCOPE.md).

### 与闭源主系统及商业交付的关系（必须读）

This repository (HDGP-Core) is an **open-source Meta baseline**, governed and evolved independently by the community. A corresponding **private mainline** (R&D) and **commercial deliveries** are maintained in separate non-public repositories and delivery tracks. **After open-source publication, there is no standing obligation to keep code synchronized between the two tracks.**

**Mainline policy (for boundary clarity)**: the mainline **does not merge or transplant** contributor code from this open-source repository. This policy constrains **code integration only**; it does **not** claim exclusivity over common public ideas, industry practices, or convergent independent implementations.

**Not in scope by default** (typically provided in the mainline/commercial track): full policy execution cores (Judge/Engine) and release/ops gates, compliance/certification services, audit evidence-chain hosting and managed operations, and industry/enterprise integrations. If you need these, consult the mainline or official commercial channels—do not assume they are included here.

**Security disclosure**: this repository provides `SECURITY.md` for private security reporting (prefer GitHub Private vulnerability reporting; alternative: `xyan8921@gmail.com`; PGP fingerprint may remain TBD if not published).

**Governance snapshot note**: the **final** Meta-only document pick from the private mainline was performed at mainline commit **`e60732ea78a29ddb168a41be8792dff96af5ee59`** (**2026-05-03**), per `docs/CORE_EXTRACT_SNAPSHOT.md` and the mainline **HDGP_CORE_FINAL_PULL_PLAN** (in the private `HDGP-Protocol` repository; not distributed inside Core). Governance wording in this repo reflects that snapshot; **community governance mechanics here may differ** from the private mainline. **No ongoing spec/repo sync** is implied except **quarterly ethics alignment** attestation materials allowed by `MATERIALS_ALLOWLIST.md`.

---

## 仓库范围与能力边界

- **In scope (Meta-only)**:
  - Semantic boundary between Meta and Judge/Audit (see `spec/HDGP_META_VS_JUDGE_SCOPE.md`)
  - Open Letters narrative (`LETTERS.md`; styled HTML on [main site](https://www.hdgp-protocol.com/docs/readable/letters.html) — non-normative)
  - Ethics baseline (see `spec/HDGP_ETHICS_BASELINE.md`)
  - Mapping from principles to executable constraints (see `spec/HDGP_CORE_MAPPING_SPEC.md`)
  - Integration semantics (Meta-related excerpts only; see `spec/HDGP_INTEGRATION_SPEC.md`)
  - Governance and proposal processes (`GOVERNANCE.md`, `docs/CHIP_PROCESS.md`)

- **Not committed by default**:
  - Any runnable Engine/Judge reference implementation or operational release gates
  - Audit storage / evidence-chain hosting / compliance attestations
  - Production runbooks, G3 operations, or audit-report collections

---

## 材料白名单（Core 自维护）

This repository uses a file-level allowlist to avoid importing Judge/Audit/Ops commitments into the Meta-only baseline. See `MATERIALS_ALLOWLIST.md`.

---

## 安全报告

Please read `SECURITY.md`.

---

## 参与贡献

Please read `CONTRIBUTING.md` and `CODE_OF_CONDUCT.md`. For changes touching baseline principles or the ethics baseline, follow the multi-layer governance process in `docs/CHIP_PROCESS.md` and `GOVERNANCE.md`.

---

## 中文版本 (ZH-CN)

本文件英文为社区首选版本；以下为中文对照版本。

---

## HDGP-Core（Meta-only）：Human Dignity Guardian Protocol · Open Meta Baseline (ZH-CN)

本仓库（`HDGP-Core`）提供 **开源的 Meta 伦理安全基线**：语义、契约与社区治理材料（Meta / 伦理 / 映射 / 治理流程）。它用于帮助采用方在自己的系统里**编织 Meta 指导字段**、形成可讨论与可审计的“伦理输入契约”。

> **正式发布 `v1.0.0` · 2026-05-03** — 本仓库首个稳定的 Meta-only 公开基线（**Git 标签 [`v1.0.0`](https://github.com/HumanDignityGuardian/HDGP-Core/releases/tag/v1.0.0)**）。主系统终局验收与冻结口径摘要见 **`docs/CORE_EXTRACT_SNAPSHOT.md`**（交叉引用主系统 `docs/HDGP_CORE_FINAL_HANDOFF_ACCEPTANCE_2026-05-03.md`，**未**纳入 Core）。

### 站点与联系

- **GitHub Pages**：`https://humandignityguardian.github.io/HDGP-Core/` — 本仓库的轻量索引页（链向 GitHub 上的规范与治理文档）。  
- **维护联系人**：Yvaine He · `xyan8921@gmail.com`（一般仓库咨询；**安全问题**请优先使用 GitHub **私密漏洞报告**，见 `SECURITY.md`）。  
- **产品与主系统叙事**（Engine 演示、PFaaS 语境 — **不属于** Core 材料正文）：[hdgp-protocol.com](https://www.hdgp-protocol.com/)

### 面向谁（工程师与 Agent 基座 · 同一套规范）

- **产品 / 应用工程师**：在开发期编织 **Meta** 字段；入门见 `spec/HDGP_INTEGRATION_SPEC.md`、`spec/HDGP_META_VS_JUDGE_SCOPE.md`。  
- **Agent 基座 / 平台构建者**：将同一套契约当作 **字段模式与路由语义**；Core **不**提供 Judge/Engine，对齐停留在 **语义与可追溯**，而非宣称运行时执法。

本仓库**不需要**两套并列文档树：**一套 Meta-only 基线**，**两种接入角色** —— 规范相同，集成胶水不同。

### 公开信与技术白皮书（主站可读版）

主站排版定稿（永久记录；Core **不**重复发布 HTML）：

- **公开信** — [主站 letters 页面](https://www.hdgp-protocol.com/docs/readable/letters.html) · 仓库内 Markdown 镜像：[`LETTERS.md`](LETTERS.md)  
- **技术白皮书 v1.0** — [主站 whitepaper 页面](https://www.hdgp-protocol.com/docs/readable/whitepaper.html)

白皮书表述的是 **全域** HDGP 框架（含超出本仓 Meta-only 的方向）。**Core 承诺面**以 [`MATERIALS_ALLOWLIST.md`](MATERIALS_ALLOWLIST.md) 与 [`spec/HDGP_META_VS_JUDGE_SCOPE.md`](spec/HDGP_META_VS_JUDGE_SCOPE.md) 为准。

### 与闭源主系统及商业交付的关系（必须读）

本仓库（HDGP-Core）是 **开源的 Meta 伦理安全基线**，由社区独立治理与演进。与之对应的 **主系统研发与商业交付** 在独立的闭源/非公开仓库与商业路径中维护，**双方仓库在开源发布后不承担代码层面的相互同步义务**。

**主系统政策（供采用方理解边界）**：主系统 **不合并、不移植** 本开源仓库中的贡献者代码；该政策 **仅约束代码合入**，**不** 声称对公开领域中的通用思路、行业惯例或独立实现的相似性作出互斥或独占。

**不在本开源仓库默认范围内的能力**（通常在主系统或商业侧提供）：完整策略执行核（Judge/Engine）与运维发布门禁、合规与认证背书服务、审计证据链与托管级运维、行业与企业扩展集成。若你需要上述能力，请通过主系统或官方商业渠道了解，而非默认期待本仓库发布物包含同等范围。

**安全披露**：本仓库提供 `SECURITY.md` 用于私密报告安全问题（优先 GitHub Private vulnerability reporting；备选邮箱 `xyan8921@gmail.com`；PGP 指纹若未公布可仍为 TBD）。

**治理快照脚注**：自私有主系统完成的 **最后一次** Meta-only 文件拣选锚点为 **`e60732ea78a29ddb168a41be8792dff96af5ee59`**（**2026-05-03**），依据 `docs/CORE_EXTRACT_SNAPSHOT.md` 与主系统私有仓内 **`docs/HDGP_CORE_FINAL_PULL_PLAN.md`**（**未**作为文件纳入 Core，仅作来源说明）。本仓库治理表述对应该快照时点；**社区治理机制可与闭源主系统不一致**。除 **`MATERIALS_ALLOWLIST.md`** 允许的 **季度伦理对齐** 材料外，**不暗示**与主系统持续规范/仓库同步。

---

## 仓库范围与能力边界

- **本仓库聚焦（Meta-only）**：
  - Meta 与 Judge/Audit 的语义边界（见 `spec/HDGP_META_VS_JUDGE_SCOPE.md`）
  - 公开信叙事（`LETTERS.md`；主站排版见 [letters.html](https://www.hdgp-protocol.com/docs/readable/letters.html)，非规范性）
  - 伦理基线（见 `spec/HDGP_ETHICS_BASELINE.md`）
  - 原则到可执行映射（见 `spec/HDGP_CORE_MAPPING_SPEC.md`）
  - 集成语义（仅 Meta 相关节选，见 `spec/HDGP_INTEGRATION_SPEC.md`）
  - 治理与提案流程（`GOVERNANCE.md`、`docs/CHIP_PROCESS.md`）

- **本仓库不默认承诺**：
  - 任何可运行的 Engine/Judge 参考实现或运维门禁
  - 审计落盘/证据链托管与认证背书服务
  - 生产运行手册、G3 运营与审计报告集

---

## 材料白名单（Core 自维护）

本仓库采用文件级材料白名单，避免把 Judge/Audit/Ops 的承诺面带入 Meta-only 发行物。详见根目录 `MATERIALS_ALLOWLIST.md`。

---

## 安全报告

请阅读 `SECURITY.md`。

---

## 参与贡献

请阅读 `CONTRIBUTING.md` 与 `CODE_OF_CONDUCT.md`。涉及伦理基线与基线原则的变更，请遵循 `docs/CHIP_PROCESS.md` 与 `GOVERNANCE.md` 中的多层治理流程。

