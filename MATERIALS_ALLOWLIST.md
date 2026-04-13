# HDGP-Core Materials Allowlist (file-level, maintained by Core)

> **Status**: Final (initial baseline).  
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
| `docs/CHIP_PROCESS.md` | CHIP process |
| `CONTRIBUTING.md` | Contributing guide |
| `CODE_OF_CONDUCT.md` | Code of conduct |
| `HDGP_OPEN_FRAMEWORK.md` | Framework overview (**excerpt** Meta-related sections) |
| `README.md` | Mandatory statements and scope boundaries |
| `SECURITY.md` | Security disclosure policy (prefer GitHub private reporting; email/PGP may be TBD, must not be fabricated) |

---

## 2. Excluded by default (mainline / reference implementations / operations)

| Path/pattern | Notes |
|----------|------|
| `cmd/`, `internal/`, `scripts/` | Reference implementations (Engine/tooling/release gates) |
| `docs/operations/`, `docs/HDGP_G3_*`, `docs/audit-reports/` | Audit and G3 operational evidence |
| Any runbooks/deliverables that commit to Judge/Audit/Ops | Would expand commitments beyond Meta-only (must not enter the initial Core baseline) |

---

## HDGP-Core 材料白名单（文件级，Core 自维护）(ZH-CN)

> **状态**：定稿（首发基线）。  
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
| `docs/CHIP_PROCESS.md` | CHIP 流程 |
| `CONTRIBUTING.md` | 贡献指南 |
| `CODE_OF_CONDUCT.md` | 行为准则 |
| `HDGP_OPEN_FRAMEWORK.md` | 框架总览（Core 侧 **节选** Meta 相关章节） |
| `README.md` | 必备声明与边界（见仓库根目录 README） |
| `SECURITY.md` | 安全披露渠道（首选 GitHub Private vulnerability reporting；邮箱/PGP 可 TBD，不得编造） |

---

## 2. 默认不纳入（主系统 / 参考实现 / 运维）

| 路径/模式 | 说明 |
|----------|------|
| `cmd/`、`internal/`、`scripts/` | 参考实现（Engine/工具链/运维门禁） |
| `docs/operations/`、`docs/HDGP_G3_*`、`docs/audit-reports/` | 审计与 G3 运营证据 |
| 任何承诺 Judge/Audit/Ops 的运行手册与交付物 | 会把承诺面带入 Meta-only（禁止进入 Core 首发基线） |

