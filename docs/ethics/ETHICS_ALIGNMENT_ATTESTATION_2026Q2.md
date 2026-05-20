# Ethics Alignment Attestation — 2026Q2 (HDGP-Core)

> **Repository**: this file applies to the open **HDGP-Core** repository. The quarterly policy template may be described in the private mainline `docs/HDGP_QUARTERLY_ETHICS_ALIGNMENT_POLICY.md` (not shipped in Core).  
> **Process**: complete metadata, anchor hashes, alignment status, and sign-off for the quarter. Human review is required before publication.

---

## Metadata

| Field | Value |
|------|------|
| Quarter | 2026-Q2 |
| Date | 2026-05-20 |
| Repo | HDGP-Core |
| Branch | main |
| Commit (optional) | `354d49b` (attestation + Q3 issue draft; re-verify if branch moved) |

---

## Anchor files & hashes (SHA-256) — Core repository

| File (in this repo) | SHA-256 (Core) | Verified |
|---------------------|----------------|----------|
| `spec/HDGP_ETHICS_BASELINE.md` | `af9b5f5a71941ae6b3aaede7223d0bd9d59bc1b76b4b93387e4f86cbec0284e2` | Yes — matches file bytes in Core |
| `spec/HDGP_META_VS_JUDGE_SCOPE.md` | `cd612df52814ce25f7caf6b6c811e1b53b26f706a1078da774a489aa5f0a86ee` | Yes — matches file bytes in Core |

> **Note**: `af9b5f5a…` applies only to the ethics baseline file; `cd612df5…` applies only to the Meta/Judge scope file (distinct anchors).  
> Cross-repo comparison vs private mainline (2026-Q2): mainline anchors are **`48eaa0aa…`** (ethics baseline) and **`b0de2978…`** (Meta/Judge scope) — **not byte-identical** to Core; see **Alignment status** below.

---

## Alignment status (quarterly: Core ↔ private mainline)

- [ ] aligned
- [x] diverged

**Core internal (declaration ↔ files in this repo):** consistent — anchor hashes above match Core files at sign-off.

**Core ↔ mainline:** **diverged** — **disclosed**, not a silent fork. After the final Meta-only pick (2026-05-03), Core ships an **English-first + ZH-CN mirror** Meta-only canon; mainline retains a **Chinese-primary, shorter** independent evolution (~196 vs ~356 lines for ethics baseline; Meta/Judge scope sizes differ similarly). This does not contradict Core handoff acceptance §2.3 (“Core attestation hashes match **Core** files”); it is not a claim of **byte-identical** dual-repo bodies.

---

## If diverged: differences & actions

**Summary (EN):**

- Mainline `HDGP_ETHICS_BASELINE.md`: `48eaa0aa…` (18,978 B) vs Core `af9b5f5a…` (28,367 B) — different editorial structure and length; same path, not same bytes.
- Mainline `HDGP_META_VS_JUDGE_SCOPE.md`: `b0de2978…` (2,895 B) vs Core `cd612df5…` (6,078 B) — same path, not same bytes.
- Prior mainline draft anchor `13fc4efc…` is **superseded**; current mainline anchors per mainline `ETHICS_ALIGNMENT_ATTESTATION_2026Q2.md`.
- Divergence is **expected** under repository isolation policy; documented in this attestation and mainline review (not undisclosed drift).

**Next steps (EN):**

- **2026-Q3 onward**: pursue **aligned** only via explicit CHIP / written equivalence mapping (unified anchor edition or section-level mapping) — **do not** assume same path implies same bytes.
- Track in Core governance ([Issue #8](https://github.com/HumanDignityGuardian/HDGP-Core/issues/8)); no standing obligation for continuous code or document sync between repos.

**摘要（中文）：**

- 主系统与 Core 同名锚点文件**非字节一致**；Core 为终局拣选后的 Meta-only 英文在前终稿 + 中文对照，主系统为中文主稿独立演进，篇幅与哈希均不同。
- 分歧为**已披露**的预期内差异，非静默分叉；主系统旧锚点 `13fc4efc…` 已过时，当季主系统锚点以 `48eaa0aa…` / `b0de2978…` 为准（见主系统当季声明）。
- 若 2026-Q3 及以后需达到 **aligned**，须经 CHIP/书面「等价章节」或统一体例映射，不能假定路径同名即同字节。

**下一步（中文）：**

- 2026-Q3 起按需走 CHIP/书面映射（见 Issue #8）；无持续双仓文档同步义务。

---

## Sign-off

| Role | Name/Handle | Date |
|------|-------------|------|
| Maintainer / authorized representative | Yvaine He | 2026-05-20 |

---

## 中文版本 (ZH-CN)

> 以下中文与上文英文对应；社区阅读顺序以英文为先。

# 伦理对齐季度声明 — 2026Q2（HDGP-Core）

> **仓库**：本文件适用于开源 **`HDGP-Core`**。季度政策模板可见私有主系统 `docs/HDGP_QUARTERLY_ETHICS_ALIGNMENT_POLICY.md`（**未**随 Core 全文发布时，仅作口径引用）。  
> **流程**：填写元数据、锚点哈希、对齐结论与签署；发布前须经人工复核与签署。

---

## 元数据

| 字段 | 值 |
|------|-----|
| 季度 | 2026-Q2 |
| 日期 | 2026-05-20 |
| 仓库 | HDGP-Core |
| 分支 | main |
| 提交（可选） | `354d49b`（attestation + Q3 issue 草案；若分支已前进发布前请复核） |

---

## 锚点文件与哈希（SHA-256）— Core 仓库

| 文件（本仓库内） | SHA-256（Core） | 已核对 |
|------------------|-----------------|--------|
| `spec/HDGP_ETHICS_BASELINE.md` | `af9b5f5a71941ae6b3aaede7223d0bd9d59bc1b76b4b93387e4f86cbec0284e2` | 是 — 与 Core 内文件字节一致 |
| `spec/HDGP_META_VS_JUDGE_SCOPE.md` | `cd612df52814ce25f7caf6b6c811e1b53b26f706a1078da774a489aa5f0a86ee` | 是 — 与 Core 内文件字节一致 |

> **说明**：`af9b5f5a…` 仅指伦理基线文件；`cd612df5…` 仅指 Meta/Judge 边界文件（两个锚点不可混用）。  
> 与私有主系统交叉比对（2026-Q2）：主系统锚点为 **`48eaa0aa…`**（伦理基线）、**`b0de2978…`**（Meta/Judge 边界）— 与 Core **非字节一致**；见下文 **对齐状态**。

---

## 对齐状态（季度：Core ↔ 私有主系统）

- [ ] aligned（一致）
- [x] diverged（不一致）

**Core 内部（本仓声明 ↔ 本仓文件）：** 一致 — 上文锚点哈希与签署时 Core 文件相符。

**Core ↔ 主系统：** **不一致（diverged）** — 属**已披露**分歧，非静默分叉。终局拣选（2026-05-03）后，Core 为 **英文在前 + 中文对照** 的 Meta-only 终稿；主系统保留 **中文主稿、更短** 的独立演进版本。这不与 Core 验收材料 §2.3「声明哈希与 **Core** 文件一致」矛盾；**不**表示双仓正文同字节。

---

## 若不一致：差异与处置

**摘要：** 见英文节 **Summary (EN)** 与 **摘要（中文）**（主系统哈希 `48eaa0aa…` / `b0de2978…`；Core 为 `af9b5f5a…` / `cd612df5…`；主系统旧稿 `13fc4efc…` 已过时）。

**下一步：** 见英文 **Next steps** 与中文 **下一步**（2026-Q3 见 [Issue #8](https://github.com/HumanDignityGuardian/HDGP-Core/issues/8)）。

---

## 签署

| 角色 | 姓名/Handle | 日期 |
|------|-------------|------|
| 维护者 / 授权代表 | Yvaine He | 2026-05-20 |
