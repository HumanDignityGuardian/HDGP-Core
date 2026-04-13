## HDGP Open Framework Overview (Meta-only excerpt)

> **Note**: this is a Meta-only excerpt used by the `HDGP-Core` open baseline.  
> **Source**: mainline repository `HDGP-Protocol/HDGP_OPEN_FRAMEWORK.md` (excerpted).  
> **Excerpt rule**: we keep only content related to the **need for an explicit ethics framework** and **Meta responsibilities / field semantics**. Engineering narratives about Engine/Judge/Audit, signing, and operational gates are out of scope by default for this repository.  
> **Semantic boundary**: see `spec/HDGP_META_VS_JUDGE_SCOPE.md`.

---

## 1. Does HDGP need an explicit ethics framework?

**Conclusion (engineering-oriented)**: yes—HDGP needs an ethics framework that is **explicit, auditable, and versioned**.

Rationale:

- HDGP’s core goal is not “efficiency”, but to protect human dignity and the priority of human final decision-making; this is an explicit value choice.  
- Without a clear ethics baseline, implementations can be quietly eroded under pressure to “optimize UX / reduce refusals / improve business metrics”.  
- A versioned ethics framework can:
  - Separate value assumptions from code and make them discussable and comparable;  
  - Support localization across cultures and contexts;  
  - Provide explicit evaluation criteria for conformance testing, certification, and audits.

Therefore, maintaining an **ethics baseline** in `spec/` is necessary (see `spec/HDGP_ETHICS_BASELINE.md`).

---

## 2. Meta responsibilities (excerpt)

The Meta layer expresses context and metadata in a unified way. Its core role is to turn “ethical and risk context” into a weaveable input contract, for example:

- User/actor type and accountable roles  
- Scene (medical/finance/education/governance/creative, etc.) and intent (chat/decision_support, etc.)  
- Risk level (low/medium/high) and sensitivity tags  
- Applicable spec/policy version (for versioning and traceability)

This repository focuses on Meta semantics and contracts. Adopters may implement Meta-only (field weaving and governance) and choose their own “decision gate implementation” when needed (not a default deliverable of this repo).

---

## 中文版本 (ZH-CN)

以下为中文对照版本。

---

## HDGP 社区基线框架总览（Meta-only 节选）(ZH-CN)

> **说明**：本文档为 Meta-only 节选版，用于 `HDGP-Core` 开源基线。  
> **来源**：主系统仓库 `HDGP-Protocol/HDGP_OPEN_FRAMEWORK.md`（节选）。  
> **节选原则**：仅保留与 **伦理框架必要性、Meta 的职责与字段语义**相关内容；涉及 Engine/Judge/Audit、规则签名与运维门禁的实现叙事不在本仓库默认范围内。  
> **语义边界**：见 `spec/HDGP_META_VS_JUDGE_SCOPE.md`。

---

## 一、HDGP 是否需要独立的伦理框架？

**结论（工程取向）**：需要，而且必须是**显式、可审计、可版本化**的伦理框架。

理由：

- HDGP 的核心不是“提高效率”，而是“在任何时代保护人类尊严与人类最终决策优先”，这本身就是一种强价值选择；  
- 若缺乏明确的伦理基线，技术实现很容易在“优化体验 / 降低拒绝率 / 提升商业指标”的压力下被悄然侵蚀；  
- 一个可版本化的伦理框架，能够：
  - 把“价值观假设”从代码中剥离出来，成为可讨论、可对比的对象；  
  - 与不同文化/场景的本地化扩展对接；  
  - 为合规测试、认证与审计提供明确的评判标准。

因此，在规范层（`spec/`）中维护一份 **伦理基线** 是必要的（见 `spec/HDGP_ETHICS_BASELINE.md`）。

---

## 二、Meta 的职责（节选）

Meta 层用于统一表达上下文与元信息，其核心作用是把“伦理与风险语境”变成可编织的输入契约，例如：

- 用户/主体类型与责任主体  
- 场景（医疗/金融/教育/治理/创作等）、意图（chat/decision_support 等）  
- 风险等级（low/medium/high）与敏感性标签  
- 当前适用的规范与策略版本（用于版本化与回溯）

本仓库聚焦于 Meta 的语义与契约本身；采用方可以只实现 Meta（字段编织与治理），并在需要时选择自己的“判定关口实现”（不属于本仓库默认交付）。

