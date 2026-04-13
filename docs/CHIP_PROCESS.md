# Baseline Improvement Proposal (CHIP) Process

> **CHIP** = Baseline Improvement Proposal. It is the formal process for proposing changes to HDGP baseline articles or the ethics baseline.  
> This process aligns with the “multi-layer governance for ethics/baseline changes” in `GOVERNANCE.md` and §8 of `spec/HDGP_ETHICS_BASELINE.md`, so participants have a consistent procedure for baseline/ethics changes.

---

## 1. When CHIP is required

The following changes **MUST** go through CHIP (or an equivalent ethics change process) and must not be merged solely via a regular PR:

- Modifying **baseline articles (A)** in the HDGP technical whitepaper;  
- Modifying the ethics baseline in **`spec/HDGP_ETHICS_BASELINE.md`** (including anti-hijacking and governance sections);  
- Modifying **key principles (P)** or **core prohibitions** directly derived from baseline/ethics, in ways that affect the boundary of “human dignity / human final-decision priority”.

Changes limited to implementation details, wording polish, or spec clarifications that explicitly do not touch baseline/ethics may follow the normal PR flow (but opening an Issue first is still recommended).

---

## 2. CHIP overview (multi-layer governance)

1. **Proposal & system self-check (Layer 1)**  
   - File an Issue prefixed with “CHIP” (use the [CHIP template](.github/ISSUE_TEMPLATE/chip-proposal.md) if available).  
   - The proposal must describe: changes, rationale, impacted clauses/rules, and expected risks.  
   - **System self-check**: maintainers (or delegated roles) evaluate the proposal against the currently effective ethics baseline and baseline articles.  
   - If the self-check concludes “inconsistent with HDGP’s own ethics”, the proposal is **suspended** until revised or withdrawn; the reason must be recorded.

2. **Accountable owner & risk assessment (Layer 2)**  
   - For changes involving key rules (P/R/B), assign an **accountable owner** (a natural person who can bear legal consequences) and record it in the ethics changelog.  
   - Complete risk/vulnerability assessment. If major risks remain unassessed, mitigation must be done with the accountable owner before moving forward.

3. **Public notice & time window (Layer 3; baseline/ethics only)**  
   - Changes to **baseline articles or the ethics baseline** require an announced, pre-defined time window and public channels for global users to receive and respond.  
   - After the window, feedback may be summarized and a vote (or equivalent broad decision mechanism) may be used; **a simple majority is insufficient**—a higher threshold must be met per governance rules.

4. **Multi-sign and activation**  
   - Once the above layers are satisfied, changes require **at least two authorized natural-person sign-offs** before merge/release.  
   - All ethics-related changes must be recorded in the **Ethics Changelog** and the relevant spec version updated.

---

## 3. Emergency changes

In major security emergencies, temporary changes may be applied per `GOVERNANCE.md` §6.5, but must be formally reviewed within a predefined window (e.g., 7 days) and must **retroactively complete self-check + accountable owner/risk assessment**. If not approved, the change must be reverted, with full audit records preserved.

---

## 4. References

- Roles and decision types: `GOVERNANCE.md` §2–§3  
- Multi-layer governance for baseline/ethics changes: `GOVERNANCE.md` §6  
- Ethics baseline (self-check, accountability, notice & thresholds): `spec/HDGP_ETHICS_BASELINE.md` §8  
- Kernel checklist (governance checklist items): `spec/HDGP_KERNEL_CHECKLIST.md` §4

This document defines the process; concrete time windows and voting thresholds are defined by governance rules or later updates to GOVERNANCE.

---

# 基线原则修订提案（CHIP）流程说明 (ZH-CN)

> **CHIP** = Baseline Improvement Proposal，指对 HDGP 基线原则条款或伦理基线进行的正式修订提案。  
> 本流程与 `GOVERNANCE.md` 第六节「伦理与基线原则变更的多层治理」及 `spec/HDGP_ETHICS_BASELINE.md` 第八节一致，便于参与者在提出或审议基线/伦理变更时有所依循。

---

## 一、何时需要走 CHIP 流程

以下变更**必须**经过基线原则修订提案（CHIP）或等效的伦理变更流程，不得仅凭 PR 直接修改：

- 修改《厚德归朴（HDGP）技术白皮书 v1.0》中的**基线原则条款（A）**；  
- 修改 **`spec/HDGP_ETHICS_BASELINE.md`** 中的伦理基线（含 §7 防挟持、§8 治理与伦理变更）；  
- 修改由基线原则/伦理基线直接导出的**关键原则（P）或核心禁止类规则**，且影响“人类尊严/最终决策优先”边界的。

仅涉及实现细节、文案润色、或明确不触及基线/伦理的规范补充，可按普通 PR 流程处理（仍建议在 Issue 中先说明变更性质）。

---

## 二、CHIP 流程概览（多层治理）

1. **提案与系统自检（第一层）**  
   - 在 Issue 中以「CHIP」为前缀提交提案（建议使用 [CHIP 提案模板](.github/ISSUE_TEMPLATE/chip-proposal.md) 若已配置）。  
   - 提案须说明：拟议变更内容、理由、影响的条款与规则、预期风险。  
   - **系统自检**：由维护者或指定角色以当前已生效的伦理基线与基线原则为参照，评估拟议变更是否与之一致。  
   - 若自检结论为「拟议变更违背 HDGP 自身伦理」：提案**挂起**，由提案发起人在限定时间内修正或撤销；挂起原因记入伦理变更日志。

2. **责任方与风险评估（第二层）**  
   - 涉及关键规则（P/R/B）时，须指定**实际负责人**（能承担法律后果的自然人），并在伦理 Changelog 中记录。  
   - 完成风险与漏洞评估；若存在未评估的重大风险，须在责任方参与下修补或缓释后再进入后续步骤。

3. **公告与时间窗（第三层，仅基线原则/伦理基线修改）**  
   - 对**基线原则或伦理基线**的修改：须**公告至少一个事先约定的时间窗口**，并通过官方渠道使全球用户可收悉、可反馈。  
   - 时间窗结束后可组织反馈汇总，并可选择公投或等效广泛表决；**仅“投票过半”不足以保证通过**，须满足治理细则规定之更高门槛。

4. **多签与生效**  
   - 在满足上述各层前提下，变更须由**至少两名具相应权限的自然人共同签署**后方可合并/发布。  
   - 所有伦理相关变更须记入**伦理变更日志（Ethics Changelog）**，并更新相应规范版本号。

---

## 三、紧急变更

在重大安全紧急情况下，可依 `GOVERNANCE.md` §6.5 执行临时变更，但须在预设时间窗内（如 7 天）提交正式复核，并**补走自检与责任方/风险评估**；未通过则自动回退，保留完整审计记录。

---

## 四、参考文档

- 角色与决策类型：`GOVERNANCE.md` 第二、三节  
- 基线原则/伦理变更多层治理：`GOVERNANCE.md` 第六节  
- 伦理基线（自检、责任方、公告与高门槛）：`spec/HDGP_ETHICS_BASELINE.md` 第八节  
- 内核自查清单（治理流程勾选）：`spec/HDGP_KERNEL_CHECKLIST.md` 第四节

本文档为流程说明，具体时间窗、表决门槛等由治理细则或后续 GOVERNANCE 修订明确。

