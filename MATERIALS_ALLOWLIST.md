# HDGP-Core Materials Allowlist (file-level, maintained by Core)

> **Status**: Final; **open release `v1.0.0`** dated **2026-05-03** (Git tag `v1.0.0`). **Final mainline→Core Meta-only pick** completed **2026-05-03** (mainline anchor `e60732ea78a29ddb168a41be8792dff96af5ee59`). PRE-GATE draft reviewed into this file. Traceability & mainline acceptance cross-reference: `docs/CORE_EXTRACT_SNAPSHOT.md`. **Ongoing relationship**: repositories **fully isolated** except **quarterly ethics alignment** (`docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_YYYYQn.md` and CHIP/ethics process).  
> **Purpose**: define the minimal **Meta-only** file set allowed in `HDGP-Core`, to avoid unintentionally importing Judge/Audit/Ops commitments.  
> **Principles**: include only materials required for **Meta / ethics / mapping / governance** and the mandatory Core README statements. Engine, audit, release gates, and commercial deliverables are **out of scope** by default.  
> **Semantic boundary index**: see `spec/HDGP_META_VS_JUDGE_SCOPE.md`.

---

## 0. Copy rules (MUST follow)

- **Only include files listed in this allowlist**; do not bulk-sync directories.  
- **Preserve original attribution and license notices** (keep existing headers; do not rewrite meanings).  
- **Excerpt rule**: for documents that include Engine/Judge narratives, Core may **excerpt** Meta-related sections only, but must clearly mark them as excerpts and cite the source path.  
- **Prohibited items**: anything not listed here—such as `cmd/`, `internal/`, `scripts/`, `docs/operations/`, `docs/HDGP_G3_*`, `docs/audit-reports/`—must not enter the initial Core baseline.

---

## 1. Required for initial release (file-level)

| Path | Notes |
|------|------|
| `spec/HDGP_META_VS_JUDGE_SCOPE.md` | Semantic boundary between Meta and Judge/Audit |
| `spec/HDGP_ETHICS_BASELINE.md` | Ethics baseline |
| `spec/HDGP_CORE_MAPPING_SPEC.md` | Mapping from principles to executable constraints |
| `spec/HDGP_INTEGRATION_SPEC.md` | Integration semantics (**excerpt** Meta-related sections) |
| `GOVERNANCE.md` | Governance |
| `docs/index.html` | GitHub Pages landing page (static index linking to repo specs & governance) |
| `docs/.nojekyll` | Disables Jekyll processing for GitHub Pages (`docs/` publishing) |
| `docs/ADOPTION_BUNDLE.md` | Pack-and-go adoption entry (Meta-only boundary, attribution, contact; no Engine/Judge) |
| `docs/CORE_EXTRACT_SNAPSHOT.md` | Post–Gate G extract snapshot (mainline SHA + §6 copy-paste block) |
| `docs/CHIP_PROCESS.md` | CHIP process |
| `docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_YYYYQn.md` | Quarterly ethics alignment attestation (traceable link required; used by external review) |
| `CONTRIBUTING.md` | Contributing guide |
| `CODE_OF_CONDUCT.md` | Code of conduct |
| `LETTERS.md` | Open Letters narrative (致硅基 / 致碳基); non-normative; styled canonical HTML on main site |
| `HDGP_OPEN_FRAMEWORK.md` | Framework overview (**excerpt** Meta-related sections) |
| `README.md` | Mandatory statements and scope boundaries |
| `SECURITY.md` | Security disclosure policy (prefer GitHub private reporting; optional email published there — must not be fabricated) |

---

## 2. Excluded by default (mainline / reference implementations / operations)

| Path/pattern | Notes |
|----------|------|
| `cmd/`, `internal/`, `scripts/` | Reference implementations (Engine/tooling/release gates) |
| `docs/operations/`, `docs/HDGP_G3_*`, `docs/audit-reports/` | Audit and G3 operational evidence |
| Any runbooks/deliverables that commit to Judge/Audit/Ops | Would expand commitments beyond Meta-only (must not enter the initial Core baseline) |

---

## HDGP-Core 材料白名单（文件级，Core 自维护）(ZH-CN)

> **状态**：定稿；**公开发布 `v1.0.0`**（**2026-05-03**，Git 标签 `v1.0.0`）。**终局 Meta-only 拣选**已于 **2026-05-03** 完成（主系统锚点 `e60732ea78a29ddb168a41be8792dff96af5ee59`）。PRE-GATE 草案已审阅定稿为本文件。可追溯及主系统验收交叉引用：`docs/CORE_EXTRACT_SNAPSHOT.md`。**持续关系**：双仓 **完全隔离运作**，仅保留 **季度伦理对齐**（`docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_YYYYQn.md` 与 CHIP/伦理流程）。  
> **目的**：定义 `HDGP-Core` 可纳入的 **最小文件集合（Meta-only）**，避免无意引入 Judge/Audit/Ops 承诺面。  
> **原则**：仅列入 **Meta / 伦理 / 映射 / 治理流程** 与 Core README 必备声明所需材料；Engine、审计、门禁、商业交付 **不** 属于默认同步范围。  
> **语义边界索引**：见 `spec/HDGP_META_VS_JUDGE_SCOPE.md`。

---

## 0. 复制规则（必须遵守）

- **只纳入本白名单中的文件**；不要整体同步目录。  
- **保留原始作者署名与许可证信息**（若文件头部已有声明，保留；不要重写含义）。  
- **节选规则**：对含 Engine/Judge 叙事的文件，Core 侧允许“**节选** Meta 相关章节”，但必须在 Core 文档中标注为节选并给出来源文件路径。  
- **禁止项**：本清单未列出的 `cmd/`、`internal/`、`scripts/`、`docs/operations/`、`docs/HDGP_G3_*`、`docs/audit-reports/` 等，不得进入 Core 首发基线。

---

## 1. 首发必选（文件级）

| 路径 | 说明 |
|------|------|
| `spec/HDGP_META_VS_JUDGE_SCOPE.md` | Meta 与 Judge/Audit 边界（拣选时的语义索引） |
| `spec/HDGP_ETHICS_BASELINE.md` | 伦理基线 |
| `spec/HDGP_CORE_MAPPING_SPEC.md` | 原则到可执行映射 |
| `spec/HDGP_INTEGRATION_SPEC.md` | 集成语义（**节选** Meta 相关章节） |
| `GOVERNANCE.md` | 治理 |
| `docs/index.html` | GitHub Pages 首页（静态索引，链至仓库内规范与治理文档） |
| `docs/.nojekyll` | GitHub Pages 禁用 Jekyll（从 `docs/` 发布时） |
| `docs/ADOPTION_BUNDLE.md` | 打包带走入口（Meta-only 边界、致谢/回链、联系方式；不含 Engine/Judge） |
| `docs/CORE_EXTRACT_SNAPSHOT.md` | 门槛 G 之后提取快照（主系统 SHA + §6 可复制块） |
| `docs/CHIP_PROCESS.md` | CHIP 流程 |
| `docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_YYYYQn.md` | 伦理对齐季度声明（需可追溯链接；供外部复核） |
| `CONTRIBUTING.md` | 贡献指南 |
| `CODE_OF_CONDUCT.md` | 行为准则 |
| `LETTERS.md` | 公开信叙事（致硅基 / 致碳基）；非规范性；主站提供排版定稿 HTML |
| `HDGP_OPEN_FRAMEWORK.md` | 框架总览（Core 侧 **节选** Meta 相关章节） |
| `README.md` | 必备声明与边界（见仓库根目录 README） |
| `SECURITY.md` | 安全披露渠道（首选 GitHub Private vulnerability reporting；备选邮箱见该文件 — 不得编造） |

---

## 2. 默认不纳入（主系统 / 参考实现 / 运维）

| 路径/模式 | 说明 |
|----------|------|
| `cmd/`、`internal/`、`scripts/` | 参考实现（Engine/工具链/运维门禁） |
| `docs/operations/`、`docs/HDGP_G3_*`、`docs/audit-reports/` | 审计与 G3 运营证据 |
| 任何承诺 Judge/Audit/Ops 的运行手册与交付物 | 会把承诺面带入 Meta-only（禁止进入 Core 首发基线） |

