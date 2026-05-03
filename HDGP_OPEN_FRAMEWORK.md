## HDGP Community Baseline Framework Overview (Meta + Skill + Workflow + Engine)

> **Scope**: This document describes the **full Open Framework narrative** (including Engine). In **`HDGP-Core` (Meta-only)**, adopt only chapters related to **`meta`** / ethics semantics; Judge/Audit/signing gates ship with the **private mainline** reference implementation and **have no obligation** to merge code into Core (see `docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md`). **Semantic layering index**: `spec/HDGP_META_VS_JUDGE_SCOPE.md`.

This document summarizes HDGP’s technical baseline stack and answers:

- Does HDGP need an independent ethics framework?  
- Is “Meta + Skill + Workflow” enough? Do we still need an Engine?  
- Under “most comprehensive & safest” goals, what overall architecture do we target?

---

## 1. Does HDGP need an independent ethics framework?

**Engineering conclusion**: **Yes** — and it must be **explicit, auditable, versioned**.

Reasons:

- HDGP optimizes for **human dignity & human final decision priority**, not raw efficiency — that is a strong value choice;
- Without an explicit ethics baseline, implementations drift under UX / conversion / business KPI pressure;
- A versioned ethics layer can:
  - Separate value assumptions from code for debate/comparison;
  - Localize to cultures/scenarios;
  - Anchor conformance, certification, and audit criteria.

Therefore specs & implementations should:

- Maintain an ethics baseline under `spec/` (examples: human consciousness is non‑quantifiable; human final decision priority; diversity beats single‑metric efficiency; uncertainty prefers halt vs guessing; no exploiting psychological weaknesses for platform gain);  
- Require every policy to declare its ethical premises and bump versions on change.

In short: **HDGP cannot pretend to be “just code”.** We write the ethics explicitly and treat it as part of the **read‑only kernel + signed rules**.

---

## 2. Is Meta + Skill + Workflow enough? Why still an Engine?

### 2.1 Strengths of Meta + Skill + Workflow

Natural fit for HDGP:

- **Meta**: shared context — user type, scenario (medical/finance/education/…), risk tier, preference profiles, active HDGP spec versions — shaping downstream decisions.

- **Skills**: capabilities as governed units — LLM calls, retrieval, codegen, image gen, external APIs — each with description, risk labels, allowed/forbidden contexts.

- **Workflows**: compose Meta + Skills per scenario — medical advice / tutoring / creative flows — inserting uncertainty checks, circuit breaks, human hand‑offs, audit hooks.

Many projects stop here without a separate Engine.

### 2.2 Why HDGP still adds an Engine layer

Given HDGP’s goals we still add a unified **HDGP Engine** because:

1. **Safety & consistency** — ad‑hoc workflows may interpret rules differently or omit critical checks.  
2. **Read‑only kernel & signed rules** — kernel ethics/rules must load as signed bundles, not arbitrary workflow scripts.  
3. **Unified circuit breaking & audit** — halt logic and structured audit collection must be cross‑workflow.  
4. **Prevent bypass** — without a mandatory Engine, a new workflow could skip every guard while claiming “uses HDGP components”.

Recommended safest architecture:

- **Meta + Skill + Workflow** express intent;  
- **Engine + Policy** decide whether intent/output complies and may halt.

**Workflows propose; Engine adjudicates.**

For weaker threat models with fully trusted stacks, you might omit a standalone Engine — HDGP targets **civilizational‑scale** risk and therefore keeps a **small, hard Engine** as the final gate.

---

## 3. Safest overall structure (summary)

### 3.1 Logical layers

- **Ethics & spec layer** (`docs/`, `spec/`) — baseline principles, ethics baseline, interfaces & behavioral norms; docs + machine‑readable rules.

- **Engine layer** — load read‑only kernel & signed bundles; expose uniform evaluation (`evaluate(context, action, result)`); implement uncertainty halts, ethics risk scoring, human escalation; central audit logs; mandatory checkpoint for Meta/Skill/Workflow stacks.

- **Meta layer** — users/scenarios/risk/preferences/spec versions fed into Engine.

- **Skills** — wrap models/tools; calls pass through Engine authorization.

- **Workflows** — compose flows; explicitly mark human‑confirm nodes, risk branches, recovery after faults.

### 3.2 Safety mechanisms

- **Read‑only kernel & signed rules** — verify signatures at load; reject unsigned/invalid bundles.

- **Engine as mandatory gate** — external logic only proposes intents/candidates; Engine returns allow/modify/block/fuse; bypass paths must be impossible or explicitly “non‑HDGP mode”.

- **Unified audit & observability** — structured logs, anonymization policies, external audit APIs.

- **Composable Meta/Skill/Workflow** — customize Meta schemas, add Skills, author Workflows **only** if Engine cannot be bypassed.

---

## 4. Conclusion & next steps

We adopt:

- **Explicit ethics baseline specs**;  
- **Meta + Skill + Workflow** orchestration with **Engine** as final gate;  
- **Signed bundles + unified audit** to prevent silent erosion.

Next steps (reference mainline deliverables where absent in Core):

- Expand `spec/` with conflict‑resolution detail & minimal Engine API surfaces;  
- Draft gateway/reference structures for Meta/Skill/Workflow/Engine interactions;  
- Ship minimal demos validating halt & audit chains.

This overview evolves with architecture.


---

## 中文版本 (ZH-CN)

> 以下中文与上文英文对应；社区阅读顺序以英文为先。

## HDGP 社区基线框架总览（Meta + Skill + Workflow + Engine）

> **文档范围**：本文件覆盖 **Open Framework 全栈叙事**（含 Engine）。对外开源仓库 **`HDGP-Core`** 以 **Meta-only** 为主时，仅拣选与 `meta`/伦理语义相关的章节即可；Judge/Audit/签名门禁等工程能力以**本仓库主系统**为参考实现载体，且**与 Core 无代码合入义务**（见 `docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md`）。**语义分层索引**见 **`spec/HDGP_META_VS_JUDGE_SCOPE.md`**。

本文件用于总结 HDGP 在社区基线实现层面的总体技术框架，重点回答：

- HDGP 是否需要独立的伦理框架支撑？  
- “Meta + Skill + Workflow” 架构是否足够？是否还需要引擎？  
- 在“最全面、最安全”的设计目标下，本项目拟采用的整体方案。

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

因此，本项目在规范与实现上，会采用以下做法：

- 在 `spec/` 中维护一份 **伦理基线文档**（例如：  
  - 人类意识不可量化；  
  - 人类最终决策优先；  
  - 多样性与反熵增优先于单一效率；  
  - 不确定性优先熔断而非“拍脑袋给答案”；  
  - 不利用人类心理弱点实现系统利益最大化，等等）。
- 在实现层（网关 / 引擎），所有策略与规则都必须显式声明其所依赖的伦理前提，并在变更时更新版本与说明。

简单说：  
**HDGP 自身就是一个带有强烈伦理取向的协议，它不能假装“只是代码”。  
我们会把这套伦理取向显式写出来，并作为“只读内核 + 可签名规则”的一部分。**

---

## 二、Meta + Skill + Workflow 架构是否足够？是否还需要「引擎」？

### 1. Meta + Skill + Workflow 的优势

这种架构天然适合 HDGP：

- **Meta 层**：  
  - 统一管理上下文与元信息：  
    - 用户类型、场景（医疗/金融/教育/创作/治理…）、风险等级；  
    - 用户价值偏好（保守/激进、重安全/重效率…）；  
    - 当前使用的 HDGP 规范与规则版本。  
  - 这些信息直接影响后续规则判断与工作流分支。

- **Skill 层**：  
  - 把各种“能力”抽象为可管控的技能：  
    - LLM 对话、检索、代码生成、图像生成、外部 API 调用等；  
  - 每个技能都有：  
    - 能力描述；  
    - 可见风险标签；  
    - 允许/禁止的调用场景约束。

- **Workflow 层**：  
  - 不同场景下，组合 Meta 与 Skills：  
    - 医疗咨询工作流、理财建议工作流、教育辅导工作流、创作辅助工作流…  
  - 工作流内自然可插入：  
    - 不确定性评估节点；  
    - 熔断与人工介入节点；  
    - 审计与日志节点。

在很多普通项目中，仅靠 Meta + Skill + Workflow 就足以描述绝大多数逻辑，无需额外写“引擎”代码。

### 2. 为什么 HDGP 仍然需要「引擎层」？

考虑到 HDGP 的特殊目标，我们仍然需要在上述三层之下/之上增加一层**统一的“HDGP 引擎（Engine）”**，原因包括：

1. **安全与一贯性**  
   - 如果完全依赖工作流编排器，容易出现不同工作流对同一规则理解不一致的情况；  
   - 一旦某个工作流在实现上“忘记插入关键规则检查节点”，就可能绕开 HDGP 保护。

2. **只读内核与规则签名**  
   - HDGP 的核心规则和伦理框架需要以“只读 + 签名”的形式加载与执行；  
   - 这类规则不能作为“普通 Workflow 脚本”随意编辑，否则失去“不可篡改”的效力。

3. **统一的熔断与审计机制**  
   - 熔断逻辑应当是跨工作流统一的：  
     - 不确定性评估；  
     - 伦理风险评估；  
     - 人类介入与报告生成。  
   - 审计日志也需要统一收集，以便合规测试与外部审查。

4. **防止滥用 Meta/Skill/Workflow 绕过 HDGP**  
   - 如果没有一个“不可绕过的引擎”，则完全可以构造一个新的 Workflow，  
     - 直接绕过所有规则检查与审计节点，  
     - 让系统在名义上“使用 HDGP 组件”，但实际上不尊重 HDGP 协议。

因此，本项目建议的最安全架构为：

- **Meta + Skill + Workflow** 负责“表达系统想做什么”；  
- **Engine + Policy** 负责“判断这些想做的事符不符合 HDGP”，并有权触发保护性熔断或拒绝。

换句话说：  
**Workflow 可以是“建议”，Engine 必须是“判定关口”。**  
在你的另一个项目中，如果 meta+skill+workflow 都在可信控制下，且威胁模型较弱，可以不写独立引擎；  
但 HDGP 面对的是“文明级风险”，**我们应当额外引入一个小而硬的 Engine 层，作为最终守门人**。

---

## 三、最全面最安全的整体设计（简要结构）

结合上述讨论，HDGP 的社区基线框架建议如下：

### 1. 逻辑分层

- **伦理与规范层（Ethics & Spec）**  
  - 位于 `docs/` 与 `spec/` 中：  
    - 基线原则与核心条款；  
    - 伦理基线与冲突处理原则；  
    - 接口与行为规范。  
  - 以文档 + 机器可读规则的形式存在。

- **引擎层（Engine）**  
  - 负责：
    - 加载只读内核与签名规则包；  
    - 提供统一的规则评估接口（例如 `evaluate(context, action, result)`）；  
    - 实现不确定性熔断、伦理风险评估、人工介入触发；  
    - 统一记录审计日志。  
  - 对 Meta + Skill + Workflow 暴露为“必经关口”。

- **Meta 层（Meta）**  
  - 描述：用户/场景/风险/偏好/当前规范版本等；  
  - 引擎以 Meta 为输入之一进行评估。

- **Skill 层（Skills）**  
  - 封装各种模型与工具调用；  
  - 每次调用前后，必须经过 Engine 的评估与授权。

- **Workflow 层（Workflows）**  
  - 按场景组合 Meta 与 Skills；  
  - Workflow 中显式标示出：  
    - 需要额外人工确认的节点；  
    - 低/中/高风险分支；  
    - 异常与熔断后的恢复流程。

### 2. 安全机制要点

- **只读内核与签名规则**  
  - 核心规则以只读方式打包，并由受控密钥签名；  
  - Engine 在启动/加载时验证签名；  
  - 未签名或签名校验失败的规则一律拒绝加载。

- **Engine 作为硬规则评估的唯一必经关口**  
  - 所有外部可变逻辑（Meta/Skill/Workflow）只能：
    - 提出“意图”和“候选行动”；  
    - 接受 Engine 的“允许 / 拒绝 / 需要人工介入”判定。  
  - 任何试图绕过 Engine 直接调用底层能力的路径，应在架构与实现上彻底禁止或严格标记为“非 HDGP 模式”。

- **统一的审计与可观测性**  
  - Engine 为每个关键决策生成结构化审计日志；  
  - 可配置匿名化与脱敏策略，平衡隐私与可追溯性；  
  - 为外部审计工具与合规测试提供标准化查询接口。

- **Meta + Skill + Workflow 的可替换性与可组合性**  
  - 在保证 Engine 不被绕过的前提下，允许不同项目：  
    - 自定义 Meta 结构（增加行业特定字段）；  
    - 引入新的 Skill 类型（新模型、新工具）；  
    - 编写新的 Workflow（新的业务场景）。  
  - 这样既保证底层伦理与规则的一致性，又保留足够的工程灵活性。

---

## 四、结论与后续工作

综合考虑“全面性”与“安全性”，本项目将采用：

- **有明确伦理基线的规范层**；  
- **以 Engine 为最终判定关口的 Meta+Skill+Workflow 架构**；  
- **通过签名规则包与统一审计，防止绕过与悄然稀释**。

后续工作建议：

- 在 `spec/` 中补充：  
  - 伦理基线条目与冲突处理规则；  
  - Engine 对外暴露的最小接口定义（输入/输出结构）。  
- 在实现层（例如 `gateway/`）起草：  
  - `Meta`、`Skill`、`Workflow`、`Engine` 的数据结构与交互顺序；  
  - 一个最小可运行的 Demo，用于验证熔断与审计链路。

本文件会随架构演进持续更新。


