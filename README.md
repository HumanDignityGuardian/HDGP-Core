## HDGP-Core（Meta-only）：Human Dignity Guardian Protocol · Open Meta Baseline

This repository (`HDGP-Core`) provides an **open-source Meta ethics & safety baseline**: semantics, contracts, and community governance materials (Meta / ethics / mapping / governance). It helps adopters **weave Meta guidance fields** into their systems and establish an “ethical input contract” that is discussable and auditable.

### 与闭源主系统及商业交付的关系（必须读）

This repository (HDGP-Core) is an **open-source Meta baseline**, governed and evolved independently by the community. A corresponding **private mainline** (R&D) and **commercial deliveries** are maintained in separate non-public repositories and delivery tracks. **After open-source publication, there is no standing obligation to keep code synchronized between the two tracks.**

**Mainline policy (for boundary clarity)**: the mainline **does not merge or transplant** contributor code from this open-source repository. This policy constrains **code integration only**; it does **not** claim exclusivity over common public ideas, industry practices, or convergent independent implementations.

**Not in scope by default** (typically provided in the mainline/commercial track): full policy execution cores (Judge/Engine) and release/ops gates, compliance/certification services, audit evidence-chain hosting and managed operations, and industry/enterprise integrations. If you need these, consult the mainline or official commercial channels—do not assume they are included here.

**Security disclosure**: this repository provides `SECURITY.md` for private security reporting (prefer GitHub Private vulnerability reporting; security email/PGP may remain TBD if not published).

---

## 仓库范围与能力边界

- **In scope (Meta-only)**:
  - Semantic boundary between Meta and Judge/Audit (see `spec/HDGP_META_VS_JUDGE_SCOPE.md`)
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

### 与闭源主系统及商业交付的关系（必须读）

本仓库（HDGP-Core）是 **开源的 Meta 伦理安全基线**，由社区独立治理与演进。与之对应的 **主系统研发与商业交付** 在独立的闭源/非公开仓库与商业路径中维护，**双方仓库在开源发布后不承担代码层面的相互同步义务**。

**主系统政策（供采用方理解边界）**：主系统 **不合并、不移植** 本开源仓库中的贡献者代码；该政策 **仅约束代码合入**，**不** 声称对公开领域中的通用思路、行业惯例或独立实现的相似性作出互斥或独占。

**不在本开源仓库默认范围内的能力**（通常在主系统或商业侧提供）：完整策略执行核（Judge/Engine）与运维发布门禁、合规与认证背书服务、审计证据链与托管级运维、行业与企业扩展集成。若你需要上述能力，请通过主系统或官方商业渠道了解，而非默认期待本仓库发布物包含同等范围。

**安全披露**：本仓库提供 `SECURITY.md` 用于私密报告安全问题（优先 GitHub Private vulnerability reporting；安全邮箱/PGP 若未公布则为 TBD）。

---

## 仓库范围与能力边界

- **本仓库聚焦（Meta-only）**：
  - Meta 与 Judge/Audit 的语义边界（见 `spec/HDGP_META_VS_JUDGE_SCOPE.md`）
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

