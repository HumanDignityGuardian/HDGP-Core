## HDGP Integration Specification (Meta-only excerpt)

> **Note**: this is a Meta-only excerpt used by the `HDGP-Core` open baseline.  
> **Source**: mainline repository `HDGP-Protocol/spec/HDGP_INTEGRATION_SPEC.md` (excerpted).  
> **Excerpt rule**: we keep only content related to **Meta semantics**, **Meta sourcing**, and **where to place decision gates**. Implementation/commitments around Gateway/SDK/Engine/Judge/Audit are out of scope by default for this repository.  
> **Semantic boundary**: see `spec/HDGP_META_VS_JUDGE_SCOPE.md`.

---

## 1. Scope of applicability

HDGP’s “governed targets” include (but are not limited to):

- **大模型与 ML 系统**：  
  - 通用 LLM（API 服务、本地部署实例等）；  
  - 特定领域模型（推荐、评分、画像、风控等）；  
  - 自研 AI 模型或含 ML 模块的子系统。

- **传统应用与服务**：  
  - Web / App / 桌面应用（含或不含 AI 功能）；  
  - 各类聊天软件、IM 系统、社交平台；  
  - 内部业务系统（CRM、ERP、客服中台等）。

- **自动化与脚本**：  
  - Agent 与自动化工作流（RPA、任务编排系统、自动运营脚本）；  
  - 各类批量通知、自动回复、机器人。

**Core criterion**:  
If a system generates an output from its inputs/state and that output is seen by humans and relied upon for action, then it **can—and should—** be considered a candidate for HDGP integration.

### 1.1 Target audience

This specification primarily targets projects without an explicit ethics layer (e.g., baseline principles, shields, review pipelines): LLM integrations, recommender systems, notification systems, agents, etc. For such systems, HDGP can serve as the **first** ethics/safety checkpoint. For systems that already have a full ethics framework, use **only the applicable steps** and apply them to output paths that need HDGP alignment.

---

## 2. Integration workflow (generic five-step method)

This workflow is architecture-agnostic and applies to any system that produces human-visible outputs. It is primarily for projects without an ethics framework; projects that already have one may excerpt applicable steps.

### 2.1 Five steps

| 步骤 | 内容 | 产出 |
|------|------|------|
| 1. 识别输出拓扑 | 标出所有用户可见输出的产生点及调用链 | 输出清单、拓扑类型 |
| 2. 选定接入模式 | 按拓扑选择单点/门面/多点/流式 | 接入点位置、调用时机 |
| 3. 确定 Meta 来源 | scene、domain、risk_level、intent 可从何处获取 | Meta 映射表 |
| 4. 定义判定映射 | allow/modify/block/fuse 在本系统的业务含义 | 判定到动作映射 |
| 5. 定义降级策略 | 当“规则判定关口”不可用时的行为 | fail-closed/fail-open 策略 |

> Note: this repository does not provide a Judge/Engine reference implementation. This section keeps only the **semantic requirement** that adopters must define decision-to-action mapping and fallback strategy in advance.

### 2.2 Output topology and integration points (semantics)

| 拓扑类型 | 特征 | 接入策略（语义） |
|----------|------|------------------|
| 单点输出 | 一个主输出路径（Pipeline 末端、单一 API 响应） | 在该点前增加“判定关口” |
| 多点输出 | 多个独立输出路径 | 统一门面接入，或每条路径单独接入 |
| 流式输出 | 持续产生消息或 chunk | 对每条消息/每批次在发送前完成判定 |

### 2.3 Common architectures and key integration points (semantics)

Pipeline 末段输出前；Request-Response 响应返回前；Event-Driven 各 Handler 出口或统一出口；Microservices 服务边界或 API Gateway；Streaming 发送前对消息或 batch 评估。

### 2.4 Minimal Meta set and defaults (recommended)

Systems without an ethics framework often lack explicit Meta. Recommended minimum:

- `meta.scene.domain`（例如：medical/finance/education/general）  
- `meta.scene.intent`（例如：chat/decision_support/notification）  
- `meta.scene.risk_level`（例如：low/medium/high）

If missing, use conservative defaults (e.g., `domain=general`, `intent=chat`, `risk_level=medium`) and record that defaults were used, so the mapping can be improved later.

---

## 中文版本 (ZH-CN)

以下为中文对照版本。

---

## 《HDGP 对模型 / 系统集成方式规范》（Meta-only 节选）(ZH-CN)

> **说明**：本文档为 Meta-only 节选版，用于 `HDGP-Core` 开源基线。  
> **来源**：主系统仓库 `HDGP-Protocol/spec/HDGP_INTEGRATION_SPEC.md`（节选）。  
> **节选原则**：仅保留与 **Meta 语义、Meta 来源、接入点选择**相关内容；涉及 Gateway/SDK/Engine/Judge/Audit 的实现与承诺面不在本仓库默认范围内。  
> **语义边界**：见 `spec/HDGP_META_VS_JUDGE_SCOPE.md`。

