## HDGP Governance Draft

This document describes HDGP governance: **roles, decision processes, and change-management principles**.  
This governance model may evolve as the project matures; the current version is a **Genesis draft**.

---

## 1. Vision and scope

- HDGP aims to build an **open, transparent, auditable, and evolvable** public technical baseline.  
- The primary goals of governance are:
  - To **prevent silent dilution or tampering** of HDGP’s core values and rules;  
  - While still enabling **full discussion and experimentation** under transparency.

Governance does not directly dictate how downstream applications use HDGP, but is responsible for:

- Version management of official specs (Spec);  
- Direction and quality of official reference implementations (Implementation);  
- Granting and revoking “HDGP certification / marks”.

---

## 2. Roles and responsibilities

> Roles may be held by individuals or teams, and may expand to multi-party participation over time.

### 2.1 Project Founder / Founding Guardian

- In the Genesis phase:
  - Sets the initial vision and value boundaries;  
  - Appoints the first maintainers and policy auditors;  
  - Serves as **one of the necessary consent parties** for major matters such as changes to core charter principles, brand/trademark policies, and license changes (not a unilateral veto; must work with the multi-layer process in §6).
- As the project matures, founder power may be delegated or constrained by community consensus; see §6 and `docs/CHIP_PROCESS.md`.

### 2.2 Core Maintainers

- Responsibilities:
  - Maintain the repository and CI/CD;  
  - Review and merge routine PRs;  
  - Manage Issues and plan the roadmap;  
  - Evolve architecture and implementation details without touching core clauses.
- Permissions:
  - Merge PRs to the main branch after CI and at least one maintainer review;  
  - Initiate minor releases (e.g., `1.1.0` → `1.2.0`).

### 2.3 Policy Auditors

- Responsibilities:
  - Review changes involving `policies/`, `spec/`, and ethics framework content;  
  - Evaluate conformance test coverage and quality;  
  - Review policy bundle releases and signing requests.
- Permissions:
  - Veto changes that “lower safety thresholds / relax constraints”;  
  - Require additional tests/audits for specific changes.

### 2.4 Certification Committee (planned)

- Formed during the ecosystem phase; members may include technical/ethics experts and academic/industry representatives.
- Responsibilities:
  - Maintain the baseline and upgrade strategy for conformance testing;  
  - Decide grant/revocation of “HDGP Certified / Compatible” marks;  
  - Handle disputes related to HDGP naming/branding.

### 2.5 Contributors and users

- Anyone may:
  - File Issues and PRs;  
  - Discuss specs and implementations;  
  - Fork and derive works under the license.
- Proposals to change specs or core rules are handled via public discussion and the formal processes below.

---

## 3. Decision types and processes

### 3.1 Routine technical changes (implementation layer)

Examples: refactoring, performance optimizations, new adapters, bug fixes.

- Process:
  1. File an Issue / PR;  
  2. Pass automated and conformance tests (if applicable);  
  3. Approved by at least one core maintainer;  
  4. Merge to main or feature branches.

- These changes do **not** require policy auditors or the certification committee.

### 3.2 Policy / spec changes

Examples: adjusting circuit-break thresholds, adding/removing allowed/forbidden behaviors, changing ethics priorities.

- Process:
  1. Describe motivation and impact as a proposal in an Issue;  
  2. Label PR as `policy-change` or `spec-change`;  
  3. Pass automated + conformance tests;  
  4. Reviewed by at least one policy auditor, confirming:
    - Human dignity / human final-decision priority is not weakened; or  
    - Any weakening is explicitly documented and sufficiently discussed;  
  5. For major changes, extend public discussion before merging.

### 3.3 Major directional changes (values / core clauses)

Examples: changing the formulation of “human consciousness is unquantifiable”; changing the priority of “humans are always above systems”; license changes; fundamental brand/trademark policy changes.

- Recommended process:
  - Mark as a “Baseline Improvement Proposal (CHIP)”;  
  - Follow **`docs/CHIP_PROCESS.md`** (self-check, accountability, public notice + high thresholds);  
  - Set a public feedback period of at least 30 days;  
  - Requires explicit consent from **all** of:
    - The founding guardian (or designated successor);  
    - Multiple policy auditors;  
    - Certification committee representatives (once established);  
    - And all requirements of §6 (multi-layer governance);
  - Approved changes produce a new spec version (e.g., `HDGP-1.1`); old versions remain but may be marked “not recommended”.

---

## 4. Policy bundles and release governance

> Implementation details will be documented later; this section describes governance constraints.

### 4.1 Candidate build

- CI/CD builds from a specified branch:
  - Guard/gateway binaries or images;  
  - Policy bundle archives;  
  - Conformance test reports and change summaries.

### 4.2 Audit and signing

- Auditors check:
  - Behavior differences vs the previous release;  
  - Whether conformance tests passed, or accepted reasons for failures;  
  - Whether risk increases lack corresponding protections.

- For releases that pass audit:
  - Use controlled keys for signing (**multi-sign required**; no single-person signing);  
  - Produce publicly verifiable signatures and fingerprints.

### 4.3 Public notice and revocation

- Each official release:
  - Is announced in Releases/docs;  
  - States the applicable HDGP spec version;  
  - Provides signature verification instructions.

- If severe issues are found later:
  - Publish a revocation notice;  
  - Mark the version “unsafe / not recommended”, and add to blacklists if necessary.

---

## 5. Conflicts and appeals

- For major disputes about specs, implementations, or certification decisions:
  1. Describe the issue publicly (Issue/discussions);  
  2. Request written responses from relevant roles;  
  3. If unresolved, propose an open meeting/workgroup;  
  4. Record the final decision publicly.

- The project discourages “private agreements that change rules without public records”, as this undermines HDGP credibility.

---

## 6. Multi-layer governance for ethics and baseline changes (closing the “two-person collusion” loophole)

To prevent abuse by any individual or small group (including collusion by two peers), **“two authorized sign-offs” is not the only threshold**. Changes involving baseline articles (A), ethics baseline, key rules (P/R/B), and policy bundle signing must follow a **multi-layer design**. Baseline and ethics are the most fundamental Meta layer and therefore have the highest change threshold. See §8 of `spec/HDGP_ETHICS_BASELINE.md`.

### 6.1 Layer 1: system self-check (must not violate HDGP’s own ethics)

- Before human deliberation, proposals must pass a self-check against the currently effective ethics baseline and baseline articles.  
- If the result is “violates HDGP’s own ethics”: the proposal is **suspended**; the proposer must revise or withdraw; reasons must be recorded in the ethics changelog.  
- If passed, proceed to the next layer.

### 6.2 Layer 2: accountable owner and legal consequences

- Changes involving key rules (P/R/B) or signing must specify an **accountable owner**—a natural person (or legal entity) able to bear legal consequences.  
- Risk/vulnerability assessment is required. Unassessed major risks must be mitigated with the accountable owner before continuing.

### 6.3 Layer 3: public notice, time window, and high thresholds (baseline/ethics)

For changes to baseline articles or the ethics baseline, all of the following are required:

- **Public notice and time window**: announce a predefined time window; ensure global users can receive and respond via official channels.  
- **Public participation and high thresholds**: after the window, feedback may be summarized and a vote (or equivalent broad decision mechanism) may be used; **a simple majority is insufficient**—higher thresholds (supermajority, double majority, etc.) must be defined to stabilize the baseline and reduce disruption.

### 6.4 Multi-sign (necessary but not sufficient)

- After 6.1–6.3 are satisfied, changes still require **at least two authorized natural-person sign-offs**. Multi-sign is necessary, but does not replace self-check, accountability, or notice/high-threshold mechanisms.

### 6.5 Emergency changes and post-review

- In major security emergencies, temporary changes are permitted, but must be formally reviewed within a predefined window (e.g., 7 days) and must retroactively complete self-check + accountability/risk assessment. If not approved, changes are reverted with full audit records preserved.

### 6.6 Ethics changelog and operator audit logs

- Ethics change logs must record: change content, rationale, impact, proposer, **self-check result, accountable owner**, deliberation and voting records, effective time and version.  
- Privileged operator actions (proposal, comments, votes, signing) should be auditable and distinguished from automation.

### 6.7 Non-monopoly, but no internal governance backdoors

- HDGP does not claim to satisfy everyone; non-adoption and alternative protocols are legitimate.  
- Internally, governance must not leave loopholes for collusion or single-point abuse; the multi-layer design ensures baseline/ethics evolution is auditable, accountable, and socially supervised.

---

## 7. Violations and accountability

To protect HDGP credibility and stakeholders, participants with governance power who engage in the following should bear responsibility:

- Bypassing established processes to change ethics clauses or rules;  
- Hiding, deleting, or tampering with audit records related to ethics/rules decisions;  
- Bribery, coercion, or threats to influence governance;  
- Using privileges to obtain improper exemptions or relaxations for specific parties.

**Handling principles**:

- Immediately suspend governance/admin privileges upon discovery;  
- Depending on severity, take one or more of:
  - Independent investigation and public explanation;  
  - Revocation of titles/roles;  
  - Referral to relevant judicial/regulatory authorities;  
- **No exceptions** due to past contributions or prior status.

---

## 8. Future evolution

This governance draft will be tested through practice and time. Possible evolutions include:

- Establishing an independent HDGP foundation / non-profit to assume parts of governance;  
- Bringing in more international members for auditing and certification;  
- Exploring on-chain or multi-sign records for key governance decisions;  
- **Open-source vs commercialization separation**: the project may later establish independent commercial entities (certification, hosting, enterprise services) with clear legal/governance separation. The open-source track remains community-governed; commercialization may be pursued via services and trademark policies, with transparency and auditability.

Suggestions are welcome via Issues/PRs.

---

## HDGP 项目治理草案（GOVERNANCE）(ZH-CN)

本文件描述 HDGP 治理体系的**角色分工、决策流程与变更管理原则**。  
本治理模型将随项目演进而调整，当前为 **Genesis 草案**。

---

## 一、项目愿景与边界

- HDGP 项目致力于构建一个 **公开透明、可审计、可演化** 的厚德归朴（HDGP）公开技术基线实现。  
- 项目治理的首要目标是：  
  - **确保 HDGP 的核心价值与规则不被悄然稀释或篡改**；  
  - **同时允许在公开透明的前提下进行充分讨论与实验**。

治理层不直接干预各个下游应用怎么使用 HDGP，但会对以下方面负责：

- 官方规范（Spec）的版本管理；  
- 官方参考实现（Implementation）的方向与质量；  
- “HDGP 认证 / 标识”的授予与撤销。

---

## 二、角色与职责

> 角色可由具体个人或团队承担，未来可扩展为多方组织共同参与。

### 1. 项目发起人与创始守护者（Project Founder / Founding Guardian）

- 在 Genesis 阶段承担以下职责：  
  - 设定 HDGP 的初始愿景与价值边界；  
  - 指定首批维护者与规则审计员；  
  - 在涉及**元章程核心原则修改、品牌/商标授权政策、项目许可证变更**等重大事项时，作为**必要同意方之一**参与决策（非单独否决权，须与 §6 多层治理流程配合）。
- 随项目成熟，发起人的权限可以通过社区共识逐步下放或限制；具体表述见第六节与 `docs/CHIP_PROCESS.md`。

### 2. 核心维护者（Core Maintainers）

- 职责：
  - 维护代码仓库与 CI/CD 流程；  
  - 对日常 PR 进行 Review 与合并；  
  - 管理 Issue、规划 Roadmap；  
  - 在不触及核心条款的前提下，演进实现细节与架构。
- 权限：
  - 可在通过 CI 与至少 1 名维护者 Review 的前提下，合并 PR 至主分支；  
  - 可发起 Minor 版本发布（如 `1.1.0` → `1.2.0`）。

### 3. 规则审计员（Policy Auditors）

- 职责：
  - 专门审查涉及 `policies/`、`spec/`、伦理框架的变更；  
  - 评估合规测试套件覆盖度与质量；  
  - 审核“规则包（Policy Bundle）”的发布与签名申请。
- 权限：
  - 对任何“降低安全阈值 / 放宽限制”的变更拥有否决权；  
  - 可要求为特定变更增加额外测试与审计。

### 4. 认证委员会（Certification Committee，规划中）

- 在生态阶段成立，成员可来自：  
  - 不同机构的技术/伦理专家；  
  - 学界/产业界代表。
- 职责：
  - 维护 HDGP 合规测试的基线与升级策略；  
  - 审议“HDGP Certified / Compatible”标识的授予与撤销；  
  - 处理与 HDGP 名称/品牌相关的争议。

### 5. 普通贡献者与用户（Contributors / Users）

- 任何人都可以：
  - 提 Issue 与 PR；  
  - 讨论规范与实现；  
  - 在遵守许可证的前提下 Fork 与衍生。
- 对于规范与核心规则的修改建议，将通过公开讨论与上述角色的正式流程处理。

---

## 三、决策类型与流程

### 1. 日常技术变更（实现层）

示例：代码重构、性能优化、新增适配器、修复 Bug 等。

- 流程：
  1. 提交 Issue / PR；  
  2. 通过自动化测试与合规测试（如适用）；  
  3. 至少 1 名核心维护者 Review 通过；  
  4. 合并至主分支或特性分支。

- 这类变更**不需要**规则审计员或认证委员会介入。

### 2. 规则与策略变更（Policy / Spec 变更）

示例：调整熔断阈值、增删某些允许/禁止行为、修改伦理优先级等。

- 流程：
  1. 在 Issue 中以“提案（Proposal）”形式阐述变更动机与影响；  
  2. 标记 PR 为 `policy-change` 或 `spec-change`；  
  3. 通过自动化测试 + 合规测试；  
  4. 至少 1 名规则审计员 Review 并确认：  
    - 未削弱人类尊严/最终决策优先的保护边界；  
     - 或者削弱之处已在规范层明确记录并经充分讨论；  
  5. 对于重大变更，可能需要延长公开讨论期后再合并。

### 3. 重大方向调整（价值观 / 核心条款）

示例：修改“人类意识不可量化”的表述；改变“人类永远高于系统”的优先级；项目许可证变更；品牌/商标授权政策的根本调整等。

- 流程建议：
  - 标记为“基线原则修订提案（CHIP）”；  
  - 详细流程见 **`docs/CHIP_PROCESS.md`**（含自检、责任方、公告与高门槛）；  
  - 设定不少于 30 天的意见征集期；  
  - 至少需要以下**全部**方的明确同意（创始守护者为必要同意方之一，非单独决定）：
    - **创始守护者**（项目发起人或其指定接班人）；  
    - 多名规则审计员；  
    - 认证委员会（成立后）代表；  
    - 且满足第六节多层治理（系统自检、责任方、公告与高门槛）的全部要求；
  - 对通过的变更生成**新版本规范**（如 `HDGP-1.1`），旧版本保留但标注为“不再推荐”。

---

## 四、规则包与发行版的发布流程

> 具体细节会在实现阶段写入相应技术文档，这里描述治理侧约束。

### 1. 候选版本生成

- 通过 CI/CD 从指定分支构建：
  - 守门人服务/网关的二进制或镜像；  
  - 规则包（Policy Bundle）的压缩包；  
  - 合规测试报告与变更摘要。

### 2. 审计与签名

- 审计员检查：
  - 与上一版本相比，规则/行为的差异；  
  - 合规测试是否全部通过，若未通过，原因是否被接受；  
  - 是否存在明显风险增加而缺乏对等防护。

- 对通过审计的规则包与发行版：
  - 使用受控密钥进行签名（**须为多签**，不得由单一自然人独立完成）；  
  - 生成公开可验证的签名与指纹。

### 3. 公告与撤销

- 每个正式版本会：
  - 在 Release 页面与文档中公告；  
  - 标明适用的 HDGP 规范版本；  
  - 提供签名验证方式。

- 若后续发现严重问题：
  - 可发布“撤销声明（Revocation Notice）”；  
  - 将版本标记为“不安全 / 不再推荐”，必要时添加到黑名单列表。

---

## 五、冲突与上诉机制

- 对规范、实现或认证决定有重大异议时，建议流程：
  1. 在公开 Issue 或讨论区陈述具体问题；  
  2. 请求相关角色（维护者 / 审计员 / 认证委员会）给出书面回应；  
  3. 如分歧无法解决，可提议召开线上公开会议或工作组讨论；  
  4. 最终结论与决议需记录在公开文档中。

- 项目不鼓励“私下协定更改规则而不公开记录”的行为，这会破坏 HDGP 的可信度。

---

## 六、伦理与基线原则变更的多层治理（堵住“两人共谋”漏洞）

为避免任何个人或少数人（包括两名同等权限者共谋）滥用对 HDGP 伦理的影响力，**仅靠“至少两名具相应权限的自然人共同签署”不作为唯一门槛**。涉及基线原则条款（A）、伦理基线、关键规则（P/R/B）及规则包签名的变更，须经以下**多层设计**，且基线原则与伦理基线作为最原始 meta 层，修改门槛最高。详见 `spec/HDGP_ETHICS_BASELINE.md` 第八节。

### 6.1 第一层：系统自检（不得违背自身伦理）

- 变更提案在提交人类审议**之前**，须经**系统自检**：以当前已生效的 HDGP 伦理基线与基线原则为参照，评估拟议变更是否与之一致。
- **若自检结论为“拟议变更违背 HDGP 自身伦理”**：提案**挂起**，不得进入后续签署或表决；由**提案发起人**在限定时间内补充修正或撤销；挂起原因须记入伦理变更日志。
- 自检通过后，提案方可进入下一层。

### 6.2 第二层：实际负责人与法律后果

- 涉及**关键规则**（P/R/B）及规则包签名的变更，必须指定**实际负责人**——能**承担法律后果的自然人**（或法定主体）；若变更生效后发生极端事件或严重后果，须有可追溯的责任方。
- 审议中须完成**风险与漏洞评估**；若存在未评估的重大风险或漏洞，须在责任方参与下**修补或缓释**后方可进入后续流程。

### 6.3 第三层：基线原则与伦理基线——公告、时间窗与高通过门槛

对**基线原则条款与伦理基线**的修改，须同时满足：

- **公告与时间窗**：须**公告至少一个事先约定的时间窗口**；**全球用户可同步收悉**（官方渠道、邮件列表等），并可在窗口内**提出反馈与异议**。
- **公众参与与高门槛**：时间窗结束后可组织反馈汇总，并可选择**全民公投或等效广泛表决**；**仅“投票过半”不足以保证通过**，须设定更高门槛（如超多数、双重多数等，由治理细则规定），以稳固已确定条款并规避突发添乱或刻意干扰。

### 6.4 多签（必要但不充分）

- 在满足 6.1–6.3 的前提下，上述变更及规则包签名仍须**至少两名具相应权限的自然人共同签署**，任何单人无权单独生效。多签为必要条件，但不替代自检、责任方与公告/高门槛机制。

### 6.5 紧急变更与事后复核

- 重大安全紧急情况下可执行临时变更，但须在预设时间窗内（如 7 天）提交正式复核，并**补走自检与责任方/风险评估**；未通过则自动回退，保留完整审计记录。

### 6.6 伦理 Changelog 与个人操作日志

- 伦理变更日志须记录：变更内容、理由、影响、提案人、**自检结论、责任方**、审议与表决记录、生效时间与版本号。  
- 具治理权限者的操作（提案、评论、投票、签名）须可审计记录，与自动化行为区分。

### 6.7 不垄断、内部不留祸根

- HDGP 不声称符合所有人意愿；任何人有权不采纳、不合作；未来会有更多协议形式，我们不垄断。  
- **内部治理**中不得留下可被少数人共谋或单点滥用的漏洞；多层设计旨在使基线原则与伦理基线的演进可审计、可问责、可受人类社会监督。

---

## 七、违规行为与问责机制

为保护 HDGP 的信誉与被治理者的权益，任何参与治理的成员如有以下情形，应承担相应责任：

- 故意绕开既定流程私自修改伦理条款或规则；  
- 隐匿、删除或篡改与伦理/规则决策相关的审计记录；  
- 接受贿赂、施加或接受威胁，以影响 HDGP 的治理决策；  
- 利用其权限为特定组织/个人获取不当的规则豁免或放宽。

**处理原则**：

- 一经发现，立即暂停其在 HDGP 项目中的治理与管理权限；  
- 视情节轻重，采取以下一项或多项措施：  
  - 内部独立调查与公开说明；  
  - 撤销其在项目中的头衔与角色；  
  - 将相关证据移交有管辖权的司法或监管机构；  
- **无人例外**，不因其历史贡献或原有职务而豁免。

---

## 八、未来演进

本治理草案本身也会经过社区实践与时间检验。  
在未来可能出现的演进方向包括：

- 建立独立的 HDGP 基金会 / 非营利组织，承担部分治理职责；  
- 引入更多元的国际成员参与规则审计与认证；  
- 探索使用链上/多签等技术记录重要治理决策；  
- **开源与商业化过渡**：项目未来可能建立独立的商业实体（如认证、托管、企业服务），与开源项目在法律与治理上**明确区隔**——开源项目（HDGP-Core）保持 MIT 与社区治理，商业化通过认证、商标授权与服务实现；具体安排将随项目演进另行公告，确保透明与可审计。

任何对本治理文档的修改建议，欢迎通过 Issue / PR 提出。

