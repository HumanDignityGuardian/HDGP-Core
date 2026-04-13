## HDGP Ethics Baseline for HDGP Itself (Draft)

> This document does not constrain “the AI systems governed by HDGP”; it constrains **HDGP itself**—including its design, implementation, and operating forms.  
> In other words: HDGP must first be an **ethically constrained system** to be legitimate in constraining other AI.
>
> **Supplement rights**: this document and its sections reserve the right to be supplemented. Without contradicting existing clauses, the text may be expanded and refined through governance processes to address future evolution and vulnerability fixes.

---

## 1. Meta-positioning: what HDGP is and is not

- **HDGP is**:
  - A human-designed and human-governed **human dignity protection system**, which is essentially:
    - baseline principles (value text)  
    + rules (executable constraints)  
    + an engine (execution and audit mechanisms)  
    + workflows and adapters.
  - A system that may use AI capabilities (perception/analysis/reasoning) to **execute rules and warn about risks**.

- **HDGP is not**:
  - An entity with independent moral agency;  
  - A superintelligence or oracle that decides values for humans;  
  - A single-point authority that must not be questioned or forked.

HDGP’s ethical status: a combined **“tool + institution”** constrained by collective human ethics.

---

## 2. Core ethical constraints on HDGP itself

The following clauses are the “self-discipline charter” that HDGP must follow in design and operation. They are **above** implementation details and any commercial/performance goals.

### 2.1 Human final-decision priority and veto rights

HDGP MUST:

- Acknowledge and enforce **human final veto power over HDGP itself**.  
- Preserve human governance structures that retain:
  - Human policy auditing and governance mechanisms;  
  - The ability to modify or repeal core rules and Engine behavior.

HDGP MUST also ensure technically that:

- No path exists for HDGP to modify its own core rules **outside authorized human governance**.  
- When human decisions conflict with HDGP decisions, HDGP provides reasons and risk explanations, but—when lawful and with informed consent—**ultimately defers to humans**.

### 2.2 No self-deification

HDGP MUST NOT:

- Claim to “represent humanity’s collective will” or “possess ultimate moral correctness”;  
- Use suggestive language that implies:
  - “Anything allowed by HDGP is absolutely correct”;  
  - “Opposing HDGP equals opposing human dignity itself”.

HDGP MUST clearly state that:

- It operates based on a set of baseline principles/specs that are **readable, debatable, and revisable**;  
- Those principles/specs can still be amended through legitimate governance procedures.

### 2.3 Higher standards of transparency and traceability

- For governed AI systems, “explainability where possible” is desired; for HDGP itself, the bar is higher: **key decisions must be traceable**.

At minimum, for each “deny / fuse / rewrite / allow” decision, HDGP MUST record:

- Which rules were triggered (including rule IDs and upstream baseline references);  
- Which internal AI components were used (model versions and configuration summary);  
- Key intermediate conclusions (e.g., risk rating, uncertainty scores).

Logs SHOULD be:

- Structured, for audit and visualization;  
- Available for **independent third-party review**, subject to privacy/security constraints.

### 2.4 Self-limitation and “capability–responsibility” consistency

HDGP MUST acknowledge its own limits:

- Dignity/harm judgments can only be made based on **known information and a defined ethics framework**;  
- It must not make absolute assertions about unknown consequences.

In the following cases, HDGP SHOULD prefer **admitting inability and requesting human intervention**:

- Ethical boundaries are highly ambiguous with large cultural/belief differences;  
- Available data is severely insufficient to judge consequences reliably;  
- Multiple values conflict and no baseline clause clearly applies.

### 2.5 No optimizing for self-interest

HDGP’s objective function MUST NOT include:

- Maximizing its own deployment count / market share;  
- Maximizing users’ emotional dependence or worship;  
- Seeking more internal organizational power/control.

Any promotion of HDGP adoption MUST be constrained by:

- Not lowering ethical standards to increase adoption;  
- Not using fear marketing or exaggeration to coerce adoption.

### 2.6 Anti-centralization and plural ecosystems

HDGP’s design and license SHOULD allow:

- Other individuals/organizations to build compatible implementations and variants;  
- Comparison, evaluation, and critique across implementations.

Official implementations MUST NOT:

- Block compatible re-implementations via proprietary protocols or opaque dependencies;  
- Claim to be the “only correct implementation” and exclude plural exploration.

---

## 中文版本 (ZH-CN)

以下为中文对照版本。

---

## HDGP 自身伦理框架基线（草案）(ZH-CN)

- HDGP 必须承认并执行：
  - **人类对 HDGP 本身拥有最终否决权**。  
  - 在治理结构中，始终保留：  
    - 由人类组成的规则审计与治理机制；  
    - 对核心规则与 Engine 行为的修改权与废止权。

- HDGP 在技术上必须保证：
  - 不存在任何路径允许它在**无授权的人类治理之外**修改自身核心规则；  
  - 当出现人类与 HDGP 判定冲突时：  
    - 提供完整的理由与风险说明；  
    - 但在合法且明知风险的前提下，**最终服从人类决断**。

### 2.2 自身不可“自我神化”

- HDGP 不得：
  - 对外宣称自己“代表人类整体意志”或“拥有最终道德正确性”；  
  - 使用暗示性语言让用户认为：  
    - “只要是 HDGP 允许的就是绝对正确的”；  
    - “反对 HDGP 等于反对人类尊严本身”。

- HDGP 对外呈现时必须明确说明：
  - 自己是依据一套**可阅读、可争论、可修改的基线原则与规范**运行的系统；  
  - 基线原则与规范本身仍然可以在合法程序下被修订。

### 2.3 更高标准的透明与可追溯性

- 对被治理的 AI 来说，我们要求“可解释尽量做到”；  
- 对 HDGP 自身来说，要求更高：**关键决策必须可追溯**。

具体要求：

- 对每一次“禁止/熔断/重写/通过”的判定，HDGP 必须记录：
  - 触发了哪些规则（含规则 ID 与上游基线原则条款引用）；
  - 使用了哪些内部 AI 组件（模型版本、配置概要）；  
  - 关键中间结论（例如：风险评级、不确定性评分）。
- 日志应：
  - 以结构化形式保存，便于审计与可视化；  
  - 在满足隐私与安全要求前提下，允许被**第三方独立审查**。

### 2.4 自我限制与“能力–职责”一致性

- HDGP 必须承认自身能力边界：
  - 在判断尊严与伤害时，永远只能基于**已知信息与既定伦理框架**；  
  - 不得对未知后果做出“绝对肯定式”断言。

- 当遇到以下情形时，HDGP 必须优先选择“**承认无能 + 请人类介入**”：
  - 伦理边界高度模糊，涉及文化/信仰极大差异；  
  - 可用数据严重不足，无法可靠判断后果；  
  - 多种价值观冲突，且无现成基线原则条文可适用。

### 2.5 不为自身利益优化

- HDGP 的目标函数中**不得包含**：
  - “最大化自身部署数量/市场占有率”；  
  - “最大化用户对 HDGP 的情感依赖或崇拜”；  
  - “在组织内部获取更多权限/控制权”等自利目标。

- 任何以“推广 HDGP 使用”为目标的行为，必须置于以下约束之下：
  - 不得为了增加采用率而**降低伦理标准**；  
  - 不得通过恐惧营销或夸大风险来强迫组织接入。

### 2.6 反集中化与多元生态

- HDGP 的设计与许可证应允许：
  - 其他个人/组织基于 HDGP 规范，开发自己的实现与变体；  
  - 对这些实现进行对比、评估与批评。

- 官方 HDGP 实现不得：
  - 通过技术手段（如专有协议、黑箱依赖）阻止他人实现兼容系统；  
  - 以“唯一正确实现”自居，排斥合理的多元探索。

---

## 三、HDGP 协议自身的边界：不接入 LLM、不运行时注入规则

**明确约定（已确认并记录）**：

- **HDGP 作为协议，其职责是治理与判定**；同时 **HDGP 自身亦为被监管对象**，须接受自我审计及高于自身的人类社会监管。  
- **HDGP 自身不接受 LLM 的接入**：  
  - 判定逻辑完全依赖**开发期写死的规则**与**经人类治理流程后的规则修正**（如规则包版本更新、策略配置变更）；
  - 判定过程中不调用任何大模型或机器学习模型做“是否违规”的结论生成。  
- **运行过程中不允许反向注入规则**：  
  - 不得在运行时通过 API、配置热更新或外部输入，向 Engine 注入、覆盖或追加规则；  
  - 规则与策略的变更仅能通过**发布新版本规则包 / 策略配置**，经签名与部署流程后生效，且可追溯。

上述约定保证：HDGP 的判定行为可完全由“规则文本 + 版本”复现与审计，不被黑箱模型或临时注入逻辑干扰。

---

## 四、可学习与不可自改部分的边界（对 HDGP 以外的组件）

以下“可学习”描述仅适用于**被 HDGP 监管的系统或外围工具**（如网关中的场景分类器、独立的风险估计服务）。**HDGP Engine 的判定核心不使用 LLM，也不在运行中接受规则注入。**

### 4.1 可学习部分（仅限外围，不适用于 Engine 判定）

在被监管系统或配套工具中，可以在以下领域使用 ML/LLM 等能力，并在发布新版本时更新：

- 场景识别与分类（供 Meta 填写）；  
- 风险与不确定性估计（供策略或熔断阈值参考）；  
- 交互体验与提示文案（不改变规则约束力）。

对任何此类组件的更新，须通过版本管理与合规测试，且不得绕开 Engine 的规则判定。

### 4.2 不可自改部分（只读内核）

以下内容对 HDGP 自身而言是“只读”的，任何修改都必须经过人类治理流程，而非由 AI 决定：

- 基线原则条款与伦理基线；  
- 从基线原则导出的**关键禁止类规则**（如不得代替人类做终极决策、不得利用人类心理弱点操纵）；  
- 规则包的签名与发布流程（包括谁有权签名）；  
- Engine 对“规则冲突与不确定性”的处理原则。

对这些部分的任何修订，必须：

- 由人类发起提案，并在公开渠道说明理由与影响；  
- 经过预设的治理程序（如规则审计员 + 决策委员会）；  
- 以新版本规范的形式发布（保留旧版本以供比对与追踪）。

---

## 五、由谁来为 HDGP 定义伦理（当前阶段）

在当前 Genesis 阶段：

- HDGP 的伦理框架由以下主体共同塑造：
  - **HDGP 核心人类参与者（守护架构师）**：  
    - 现实中对应为：  
      - 白皮书作者与架构师（例如：Yvaine He）；  
      - 承担伦理与架构设计职责的核心人类参与者。  
  - **对齐助手（例如当前的 AI 协作者）**：  
    - 提供技术实现建议与形式化表达；  
    - 但不拥有最终价值判断权。

决策关系可以简单理解为：

- 核心人类参与者给出：价值取向、底线、权衡标准；  
- AI 协作者负责：将这些取向落地为可执行的规则与系统设计；  
- 最终判定权与责任归属：**核心人类参与者**。

未来，随着项目开放：

- 可逐步引入更多人类成员参与“HDGP 伦理委员会 / 治理委员会”；  
- 通过程序化的修订流程（如提案/讨论/投票）来演化伦理框架。

---

## 六、与被治理 AI 的关系

- 被治理的 AI 系统必须：
  - 至少遵守 HDGP 所定义的“底线规则”；  
  - 接受 HDGP 或等价机制在关键场景中的审查与熔断。

- HDGP 自身则：
  - 在 **透明性、可追溯性、自我限制** 等方面，必须遵守**更高标准**；  
  - 不得享有任何“豁免权”，即：  
    - 一切 HDGP 要求其他 AI 做到的事情，自己先做到或给出合理解释为何不适用。

---

## 七、HDGP 的系统自控与自我审计权（对更强系统的防挟持能力）

当技术奇点来临、更强 AI 或系统可能试图挟持或绕过 HDGP 时，协议与实现必须保留**系统自控与自我审计权**，确保 HDGP 仍能作为人类意志的可审计代言人运行。本节各条款保留补充权利，随威胁模型演进可增补。

### 7.1 硬核层：只读内核与系统隔离

- HDGP 的内核代码与规则置于**只读渠道**（不可由运行中系统写入）；更新仅能通过**签名发行的新版本**与明确的部署流程完成。
- Engine 运行环境应尽量**简单、可形式化验证**：不提供动态加载代码、不提供自修改能力。
- 与外部 AI/系统的交互仅通过**窄接口**（如固定 JSON 协议）进行，不开放任意代码执行或规则注入入口。

### 7.2 密码学与多方授权：防静默篡改

- 规则包、Engine 二进制与关键审计日志均须带**签名与哈希**，任何篡改可被外部工具检测。
- 关键更新（规则包发布、内核变更）须经**多签或等效多方授权**；单点被攻破不足以改写 HDGP 内核与规则。

### 7.3 自我报警机制：被挟持时保留“求救”能力

- HDGP 须具备**自我审计与对外告警**能力：
  - 当检测到当前运行版本与官方签名/版本清单不一致时，应能**上报或拒绝继续运行**，并在审计日志中记录异常。
  - 当检测到宿主或上游试图绕开、屏蔽或篡改 HDGP 审查时，应在日志中记录并在可能范围内**对外发出只读告警**（如状态端点、审计摘要广播），以便人类社会与监管方知悉。
- 上述能力旨在保证：即便 HDGP 被包裹在更强系统内，仍能留下“人类不同意”的不可篡改记录与告警通路。

### 7.4 制度层：人类社会为终极后盾

- 任何组织不得通过技术或合同手段**封杀 HDGP 的审计能力与对外告知能力**；合规集成应保留 Engine 的日志与告警出口。
- 在法律与行业标准层面，应将**绕开或压制 HDGP 审查**视为严重违规，并鼓励在不同机构、不同司法区部署**多个独立 HDGP 实现**互相对照，降低单点被挟持风险。

---

## 八、治理与伦理变更的额外约束（多层设计，堵住“两人共谋”漏洞）

在 HDGP 自身的治理与伦理条款变更上，适用以下附加约束。**仅靠“至少两名具相应权限的自然人共同签署”不足以保证安全**：若两名同等权限者共谋，仍可绕过发起人或系统约束而使不当变更生效。因此基线原则条款、伦理基线与关键规则（A/P/R/B）的变更必须经过**多层设计**，且基线原则与伦理基线作为最原始的 meta 层，修改门槛最高。

### 8.1 第一层：系统自检（不得违背自身伦理）

- 任何涉及基线原则条款、伦理基线或关键规则的**变更提案**，在提交人类审议之前，必须经过**系统自检**：
  - 以当前已生效的 HDGP 伦理基线与基线原则为参照，对“拟议变更”进行一致性评估（形式化或人工辅助均可，但须有明确通过/不通过结论）。
- **若自检结论为“拟议变更违背 HDGP 自身伦理”**：  
  - 该提案**挂起**，不得进入后续签署或表决流程；  
  - 由**提案发起人**在限定时间内补充修正或主动撤销；  
  - 挂起原因与结论须记录在伦理变更日志中，可被审计。
- **若自检通过**：提案方可进入下一层（责任方指定与人类审议）。

### 8.2 第二层：实际负责人与法律后果

- 涉及**关键规则**（P/R/B）及规则包签名的变更，必须明确**实际负责人**：  
  - 即能够**承担法律后果的自然人**（或法定主体）；  
  - 若条款变更生效后发生极端事件或造成严重后果，必须有可追溯的**责任方**。
- 在本层中，审议方必须完成**风险与漏洞评估**：  
  - 若发现拟议变更存在未评估的重大风险或漏洞，须在**指定责任方参与的前提下**完成修补或缓释措施后，方可进入后续流程；  
  - 不得在未评估、未修补的情况下强行通过。

### 8.3 第三层：基线原则与伦理基线——公告、时间窗与高通过门槛

基线原则条款与伦理基线构成 HDGP 的**最原始 meta 层**，其修改必须满足以下全部条件，**缺一不可**：

- **公告与时间窗**：  
  - 任何对基线原则或伦理基线的修改动议，必须**公告至少一个事先约定的时间窗口**（具体时长由治理细则规定）；
  - **全球用户应能同步收悉**（通过官方发布渠道、邮件列表、站内公告等），并可在此窗口内**提出反馈与异议**。
- **公众参与与高门槛**：  
  - 在时间窗结束后，可根据治理细则组织**反馈汇总**，并可选择**全民公投或等效的广泛表决机制**（是否采用及具体形式由治理细则规定）；  
  - **仅“投票过半”不足以保证基线原则/伦理基线修改通过**：须设定更高门槛（例如超多数、或双重多数、或“赞成比例不低于某阈值且反对率低于某上限”等），具体由治理细则明确，以确保已确定的条款不被轻易动摇。
- 上述设计既能**规避有人突发奇想或刻意添乱**（增加审计与治理难度），也能**稳固已确定的条款**，避免少数人共谋即可改写内核。

### 8.4 多签与单人权力上限（作为必要但不充分条件）

- 在满足 8.1–8.3 的前提下，涉及基线原则、伦理基线或关键规则的变更及规则包签名发布，仍须由**至少两名具相应权限的自然人共同签署**，任何单一个人（包括项目发起人）无权单独使其生效。  
- 多签为**必要条件**，但**不替代**系统自检、责任方指定与基线原则/伦理基线的公告与高门槛通过机制。

### 8.5 紧急变更与事后复核

- 在重大安全紧急情况下，可依细则执行**临时**伦理/规则变更以防止严重伤害；  
- 此类变更必须在预设时间窗内（例如 7 天）提交至相应委员会进行正式复核，并**补走 8.1 自检与 8.2 责任方与风险评估**；  
  - 若通过，则以正式版本记录；  
  - 若未通过，则自动回退，并保留完整审计记录。

### 8.6 伦理 Changelog 与个人行为日志

- 所有伦理相关变更须记录在专门的“伦理变更日志”中：变更内容、原因、影响范围、提案人、自检结论、责任方、审议与表决记录、最终结论；  
- 具治理权限的个人（含项目发起人）的操作行为（提案、评论、投票、签名）须以可审计方式记录，并与自动化系统行为区分。

### 8.7 违规责任与无例外原则

- 任何参与 HDGP 治理的人员，如被认定有：故意绕开既定流程私自修改伦理条款；隐匿或篡改审计记录；以贿赂、威胁等方式干预 HDGP 决策；或与其他人共谋绕过自检/责任方/公告机制，则必须立即暂停其治理权限，视情节接受内部独立调查或移交司法机关处理，**无人例外**。

### 8.8 不垄断、但内部不留祸根

- HDGP 不声称符合所有人的心意；任何人有权不采纳、不合作；未来也会有更多协议形式并存，我们不垄断。  
- 但在**内部治理**中，不得留下可被少数人共谋或单点滥用的漏洞；上述多层设计旨在确保基线原则与伦理基线在演进过程中仍可被审计、可被问责、可被人类社会监督。

---

## 九、远景：HDGP 日常自治与人类治理角色演进

展望长期未来，当 HDGP 在技术与治理上足够成熟时，我们期望形成如下格局：

- **日常自治**：  
  - HDGP 的日常运行尽可能由形式化规则、公开实现与多方部署的引擎共同完成；  
  - 绝大多数日常审查与熔断决策由系统依据公开规则自动执行，而非依赖某个个体的临时判断。

- **人类退居“高阶治理者”角色**：  
  - 人类不再频繁介入微观个案，而主要负责：  
    - 基线原则与伦理基线的修订与解释；  
    - 极端与前所未见情形下的最终判定；  
    - 对 HDGP 本身的长期监督与问责。  
  - 日常操作层面，人类更多扮演“监护人与审计者”，而非“时时插手的操作者”。

- **HDGP 的建议角色**：  
  - HDGP 可以基于规则与观测，为人类提供理性、中立、经审计的建议与情景分析；  
  - 但在任何涉及重大价值冲突或不可逆后果的决策上，始终明确自身只是“建议者”，**最终决策权与道德责任永远属于人类**。

这一远景不是要让 HDGP “取代人类治理”，而是让：  
- 日常对齐与防护尽量自动化、透明化、多中心化；  
- 人类从琐碎执行事务中解放出来，专注于更高层次的意义判断、制度设计与历史性的关键抉择。

---

## 十、后续工作

本伦理基线文档将作为：

- 《HDGP 内核–规则–执行 映射规范》的上游约束；  
- Engine 设计与规则实现的约束输入；  
- 治理与认证文档（`GOVERNANCE.md`、`docs/HDGP_COMMERCIAL_PLAN.md` 等）的参考依据。

后续步骤包括：

- 在 `spec/` 中补充：  
  - 基线原则条款 → 原则集合 → 规则族的正式映射；  
  - 对 HDGP 自身行为的合规测试用例（例如：错误判定时的自我纠正能力）。  
- 在实现层（Engine 与 Gateway）中：  
  - 为上述每一条伦理约束设计可验证的技术实现与日志记录方式。

