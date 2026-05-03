# Baseline Improvement Proposal (CHIP) Process

> **CHIP** = Baseline Improvement Proposal: a formal proposal to amend baseline principle clauses or the ethics baseline.  
> This process aligns with **§6 “Multi-layer governance for ethics and baseline changes”** in `GOVERNANCE.md` and **§8** in `spec/HDGP_ETHICS_BASELINE.md`, so participants have a single reference when proposing or reviewing baseline/ethics changes.

---

## 1. When CHIP is required

The following changes **must** go through CHIP (or an equivalent ethics-change process) and **must not** be merged by a routine PR alone:

- Amending **baseline principle clauses (A)** in *HDGP Technical White Paper v1.0*;  
- Amending the ethics baseline in **`spec/HDGP_ETHICS_BASELINE.md`** (including §7 anti‑capture and §8 governance & ethics changes);  
- Amending **key principles (P)** or **core prohibition-class rules** derived directly from the baseline principles/ethics baseline when the change affects the “human dignity / human final decision priority” boundary.

Pure implementation detail, wording polish, or normative additions that clearly do not touch baseline/ethics may follow the normal PR flow (still describe the nature of the change in an Issue first).

---

## 2. CHIP overview (multi-layer governance)

1. **Proposal & system self-check (layer 1)**  
   - File a proposal in an Issue prefixed with `CHIP` (use [.github/ISSUE_TEMPLATE/chip-proposal.md](.github/ISSUE_TEMPLATE/chip-proposal.md) if configured).  
   - The proposal must state: intended change, rationale, affected clauses/rules, and expected risks.  
   - **System self-check**: maintainers or a designated role assess consistency against the currently effective ethics baseline and baseline principles.  
   - If the self-check concludes the proposal **contradicts HDGP’s own ethics**: the proposal is **suspended**; the author must correct or withdraw within a set period; reasons are recorded in the ethics change log.

2. **Accountability & risk review (layer 2)**  
   - For key rules (P/R/B), name an **accountable party** (a natural person who can bear legal consequence) and record it in the ethics Changelog.  
   - Complete risk and gap analysis; unresolved major risks must be fixed or mitigated with the accountable party before proceeding.

3. **Notice & time window (layer 3 — baseline principle / ethics baseline only)**  
   - Changes to **baseline principles or the ethics baseline** require **at least one pre-announced time window** and official channels so global users can see and respond.  
   - After the window, aggregate feedback; optional referendum or equivalent broad vote may apply; **a simple majority alone is insufficient** — higher thresholds in governance rules apply.

4. **Multi-party sign-off & effectiveness**  
   - **Target state**: after the layers above, changes are merged/released only after **at least two authorized natural persons** co-sign.  
   - **Genesis / single-maintainer exception**: if only one accountable person exists objectively, follow the Single Maintainer mode in `GOVERNANCE.md` with compensating controls (owner, delay window, evidence chain, rollback) and public disclosure.  
   - All ethics-related changes go to the **Ethics Changelog** and bump spec versions as required.

---

## 3. Emergency changes

Under major security emergencies, temporary changes may follow `GOVERNANCE.md` §6.5, but **formal review must be filed within a preset window** (e.g. 7 days) with **self-check and accountability/risk review** completed; failure rolls back automatically with a full audit trail.

---

## 4. References

- Roles & decision types: `GOVERNANCE.md` §§2–3  
- Multi-layer governance for baseline/ethics: `GOVERNANCE.md` §6  
- Ethics baseline (self-check, accountability, notice, high bar): `spec/HDGP_ETHICS_BASELINE.md` §8  
- Kernel checklist (governance ticks): `spec/HDGP_KERNEL_CHECKLIST.md` §4 (mainline reference if not in Core)

This document describes process only; notice windows and voting thresholds are defined in governance rules or future `GOVERNANCE` revisions.

---

## 中文版本 (ZH-CN)

> 以下中文与上文英文对应；社区阅读顺序以英文为先。

# 基线原则修订提案（CHIP）流程说明

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
   - **目标态**：在满足上述各层前提下，变更应由**至少两名具相应权限的自然人共同签署**后方可合并/发布。  
   - **Genesis / Single Maintainer 例外**：若项目客观上仅有单一责任人，可按 `GOVERNANCE.md` 中的 Single Maintainer 模式执行，但必须满足其补偿控制（责任人、延迟窗口、证据链、撤销回滚）并进行公开披露。  
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
