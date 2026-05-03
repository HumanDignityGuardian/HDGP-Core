# Core extract snapshot (post–Gate G · final mainline pick)

> **When to use**: after mainline **Gate G** is closed; records **pickable** specification states for `HDGP-Core`. **Not** a code-sync credential.  
> **Final batch**: per mainline **`docs/HDGP_CORE_FINAL_PULL_PLAN.md`** §7–§9, Core performed the **last Meta-only file-level pick** from private `HDGP-Protocol` on **2026-05-03**. After that pick, the two repositories are **operationally isolated** except **quarterly ethics alignment** (see **Repository isolation** below).  
> **Mainline acceptance & freeze**: recorded in mainline **`docs/HDGP_CORE_FINAL_HANDOFF_ACCEPTANCE_2026-05-03.md`** (private `HDGP-Protocol`; not distributed in Core). Default: no further iteration of mainline→Core **batch pick** checklists; exceptions per that document (CHIP, security disclosure, written emergency alignment).  
> **Core v1.0.0**: open Meta-only baseline **released** on **2026-05-03** with Git tag **`v1.0.0`** (first stable publication state after acceptance).  
> **Core v1.0.1** (2026-05-04): patch tag **`v1.0.1`** — adoption ergonomics (JSON Schema, `snippets/`, `docs/ADOPTION_STATEMENT_TEMPLATE.md`, `docs/ADOPTION_BUNDLE.md`, `.gitignore`) plus release documentation alignment (`README.md`, `docs/index.html`, `MATERIALS_ALLOWLIST.md`, this file); **does not** alter this Gate G snapshot or the mainline pick.  
> **Related**: mainline policy `docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md`; allowlist **Final** in **`MATERIALS_ALLOWLIST.md`** (from PRE-GATE review of `docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md`).

---

## Snapshot metadata (current anchor — final Meta-only pull)

| Field | Value |
|--------|--------|
| **Final pull date** | 2026-05-03 |
| **Mainline repo** | Private Git `HDGP-Protocol` (local path on maintainer machine: `D:\HDGP\HDGP-Protocol`) |
| **Mainline Git SHA (anchor)** | `e60732ea78a29ddb168a41be8792dff96af5ee59` |
| **Branch** | `main` |
| **Gate G** | Closed; picks executed under finalized `MATERIALS_ALLOWLIST.md` only. |
| **Allowlist draft revision (mainline, at anchor)** | File `docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md` at commit `8a18ec946c83af03816c913ffbcb6218ddae1464` |
| **Core HEAD before this final pull commit** | `4c8a4608decf9a7c1532e02063379eb540fe4135` (pin the bundle that lands this pick with `git log -1 --format=%H` after merge of the final-pull commit) |
| **Executor** | Core maintainers (per `HDGP_CORE_FINAL_PULL_PLAN.md` §7) |

### Prior snapshot (stage C — superseded by anchor above)

| Field | Value |
|--------|--------|
| **Stage C record date** | 2026-04-20 |
| **Mainline Git SHA (then)** | `ef8e041e6c3e42657d719656e2cc120b3fb9a25d` |
| **Allowlist draft revision (then)** | `21275605b5999f3acab31696e05057328e18afae` |

---

## §7 checklist — Core-side execution (marked complete for final pull)

Per **`docs/HDGP_CORE_FINAL_PULL_PLAN.md`** §7:

1. **Anchor**: mainline `HEAD` at pick time recorded as **`e60732ea78a29ddb168a41be8792dff96af5ee59`** (`main`).  
2. **Files**: allowlisted paths copied **file-by-file** from mainline to Core (see **`MATERIALS_ALLOWLIST.md`** §1); excerpts marked at source where applicable.  
3. **This document**: updated with anchor SHA, date, executor, and cross-reference to **`MATERIALS_ALLOWLIST.md`**.  
4. **README**: relationship / capability boundaries reviewed; **governance snapshot footnote** added (mainline ↔ Core governance mechanics may differ).  
5. **Not copied**: `cmd/`, `internal/`, `scripts/`, ops/audit evidence trees — **absent** from Core (verified).  
6. **Optional `rulepacks/` index** (plan §6): **not** added in this batch (optional).  
7. **Self-check**: bilingual sections and links reviewed for the touched documents.

---

## Repository isolation (post–final pull)

**English**

- **Final Meta-only pick**: the batch anchored at **`e60732ea78a29ddb168a41be8792dff96af5ee59`** completes the **last** mainline-directed bulk pick described in **`HDGP_CORE_FINAL_PULL_PLAN.md`**.  
- **No standing sync**: neither side has an ongoing obligation to mirror specs, keep SHAs aligned, or repeat directory-level pulls **unless** a **new explicit agreement** and a **new snapshot** say otherwise.  
- **Exception — quarterly ethics alignment**: Core continues to maintain **`docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_YYYYQn.md`** (and related allowlisted ethics materials) on a **quarterly** cadence per **`MATERIALS_ALLOWLIST.md`** and CHIP/ethics process. This is **not** general mainline↔Core code or Engine sync.  
- **Relationship**: unchanged policy wording — open Core and private mainline remain **independent**; mainline does **not** upstream patches from Core (see README / mainline policy documents).

**简体中文**

- **终局 Meta-only 拣选**：以主系统提交 **`e60732ea78a29ddb168a41be8792dff96af5ee59`** 为锚点的本批次，完成 **`HDGP_CORE_FINAL_PULL_PLAN.md`** 所述的由主系统主导的 **最后一次** 批量拣选。  
- **无持续同步义务**：除另行书面约定并发布新快照外，**不**承担规范镜像、SHA 对齐或目录级持续拉取的义务。  
- **例外——季度伦理对齐**：Core 仍按 **`MATERIALS_ALLOWLIST.md`** 与 CHIP/伦理流程，**按季度**维护 **`docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_YYYYQn.md`**（及白名单内伦理相关材料）；**不属于** Engine 或通用代码的双向同步。  
- **关系口径**：开源 Core 与闭源主系统 **相互独立**；主系统 **不** 将 Core 贡献作为上游合入（见 README / 主系统政策文档）。

---

## Mainline formal acceptance, freeze & Core v1.0.0 (cross-reference)

**English**

- The private mainline records formal **handoff acceptance** in **`docs/HDGP_CORE_FINAL_HANDOFF_ACCEPTANCE_2026-05-03.md`** (`HDGP-Protocol`; **not** copied into Core): acceptance checklist, hash verification, Core `HEAD` at acceptance, and **§3** — **freeze** of further **batch “mainline → Core pick” work items**. **`docs/HDGP_CORE_FINAL_PULL_PLAN.md`** and **`docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md`** on mainline cross-reference this acceptance file.  
- **Default thereafter (mainline)**: do **not** iterate standing “batch pull to Core” lists. **Exceptions**: CHIP-related coordination, **security disclosure**, or **written emergency alignment** with Core — as written in that acceptance document.  
- **HDGP-Core `v1.0.0`**: this repository’s Meta-only baseline is **released as version 1.0.0** on **2026-05-03** (annotated tag **`v1.0.0`** on `main`), marking the first stable public baseline **after** mainline acceptance.  
- **HDGP-Core `v1.0.1`** (2026-05-04): patch tag **`v1.0.1`** — adoption ergonomics and release-facing docs; **does not** change the pick or this snapshot.

**简体中文**

- 私有主系统在 **`docs/HDGP_CORE_FINAL_HANDOFF_ACCEPTANCE_2026-05-03.md`**（**未**纳入本 Core 仓）记录终局 **验收**：清单、哈希核验、验收时 Core `HEAD`、**§3** 对主系统→Core **批量拣选议题的冻结**。主系统 **`docs/HDGP_CORE_FINAL_PULL_PLAN.md`**、**`docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md`** 已交叉引用该验收文。  
- **此后默认（主系统）**：**不**再持续迭代「向 Core 批量拉取」类清单；**例外**：CHIP 相关协调、**安全披露**、或与 Core 的 **书面应急对齐**（见验收文）。  
- **HDGP-Core `v1.0.0`**：本仓库 Meta-only 基线于 **2026-05-03** 以 **`v1.0.0`**（在 `main` 上打附注标签 **`v1.0.0`**）**正式发布**，作为通过主系统验收后的首个稳定公开状态。  
- **HDGP-Core `v1.0.1`**（2026-05-04）：补丁标签 **`v1.0.1`** —— 采用侧配套与面向发布的文档对齐；**不**改变拣选与本快照。

---

## §6 — Pull Guide (historical copy-paste block)

The following **verbatim** block is from mainline `docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md` §6 (for earlier batches). The **final** batch instruction is **`§9` below**.

```text
【来源】主系统 HDGP-Protocol（私有/非公开工作区），拣选须严格按 allowlist，不建立长期自动同步。

【请先读】docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md（全文）、docs/HDGP_CORE_README_REQUIREMENTS.md、
docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md（§〇、§五：无代码流通与 README 区隔）。

【快照 SHA】以主系统书面告知为准（见 docs/CORE_EXTRACT_SNAPSHOT.md 或 Release/标签说明）。

【Core 侧动作】
1) 在 Core 仓库创建 MATERIALS_ALLOWLIST.md（由 docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md 定稿复制）。
2) 仅按 allowlist 逐文件拣选；补齐 SECURITY.md 与 README（按 HDGP_CORE_README_REQUIREMENTS）。
3) 不复制 cmd/、internal/、scripts/run-release-gate、docs/operations/ 等（见 Pull Guide §2.3）。
4) 若无法直接访问主系统 Git，请向主系统索取「allowlist 内文件」打包或约定只读交接方式。

【关系口径】开源发布后两仓独立；主系统不从 Core 合入代码；详见主系统文档 §5。
```

If mainline has **not** finished stage C snapshot and finalized allowlist, complete checklist **§C** before issuing the Pull Guide §6 block (same caveat as in the Pull Guide).

---

## §9 — Mainline owner → Core maintainers (final pull · filled block)

```text
【批次类型】HDGP-Core 最后一次 Meta-only 拣选拉取（无 Engine、无门禁脚本、无运维证据默认纳入）。

【锚点】主系统 Git SHA：e60732ea78a29ddb168a41be8792dff96af5ee59；分支 main。

【范围】仅复制 MATERIALS_ALLOWLIST 内路径；并按 docs/HDGP_CORE_FINAL_PULL_PLAN.md §3–§4 执行。

【显式禁止】不得复制 cmd/、internal/、scripts/、docs/operations/、docs/audit-reports/、docs/HDGP_G3_*；conformance-tests 默认不纳入。

【Core 侧完成判据】
1) docs/CORE_EXTRACT_SNAPSHOT.md 已更新锚点 SHA 与日期；
2) README / SECURITY / MATERIALS_ALLOWLIST 满足 Core 必备结构；
3) 仓库内无 Engine 二进制与 Go 执行核源码路径。

【关系口径】开源发布后两仓独立；主系统不从 Core 合入代码；本次之后无长期自动同步义务。
```

---

## Notes

- Packaged handoff (no Git access): **`docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md`** §2.  
- Pin the Core commit that contains this snapshot and the finalized allowlist:  
  `git log -1 --format=%H -- docs/CORE_EXTRACT_SNAPSHOT.md` on `main` after pull.

---

# Core 提取快照（门槛 G 之后 · 终局拣选）

> **用途**：记录 Gate G 闭合后的可拣选状态；并于 **`HDGP_CORE_FINAL_PULL_PLAN.md` §7–§9** 指导下完成 **最后一次** 主系统 → Core 的 **Meta-only 文件级拣选**（2026-05-03）。此后 **双仓完全隔离运作**，仅保留 **季度伦理对齐** 例外（见上文 **Repository isolation**）。  
> **主系统验收与冻结**：见主系统 **`docs/HDGP_CORE_FINAL_HANDOFF_ACCEPTANCE_2026-05-03.md`**（**未**纳入 Core）；默认不再迭代「向 Core 批量拉取」清单，例外见该文（CHIP / 安全披露 / 书面应急对齐）。  
> **Core v1.0.0**：**2026-05-03** 以 Git 标签 **`v1.0.0`** 正式发布首个稳定 Meta-only 公开基线。  
> **Core v1.0.1**（2026-05-04）：补丁标签 **`v1.0.1`** —— 采用侧配套与面向发布的文档对齐；**不**改变 Gate G 快照与主系统拣选。  
> **关联**：主系统策略见 `docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md`；白名单见根目录 **`MATERIALS_ALLOWLIST.md`**。

---

## 快照元数据（当前锚点 — 终局拣选）

| 字段 | 值 |
|------|-----|
| **终局拣选日期** | 2026-05-03 |
| **主系统仓库** | 私有 Git `HDGP-Protocol`（本地路径示例：`D:\HDGP\HDGP-Protocol`） |
| **主系统 Git SHA（锚点）** | `e60732ea78a29ddb168a41be8792dff96af5ee59` |
| **分支** | `main` |
| **门槛 G** | 已闭合；拣选仅依据定稿 **`MATERIALS_ALLOWLIST.md`**。 |
| **白名单草案版本（锚点处主系统）** | `docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md` @ `8a18ec946c83af03816c913ffbcb6218ddae1464` |
| **终局拣选提交前 Core HEAD** | `4c8a4608decf9a7c1532e02063379eb540fe4135`（落地终局拣选提交后，用 `git log -1 --format=%H` 钉住包含本快照的提交） |
| **执行人** | Core 维护者（按 `HDGP_CORE_FINAL_PULL_PLAN.md` §7） |

### 早前快照（阶段 C — 已由上方锚点取代）

| 字段 | 值 |
|------|-----|
| **阶段 C 记录日期** | 2026-04-20 |
| **主系统 Git SHA（当时）** | `ef8e041e6c3e42657d719656e2cc120b3fb9a25d` |
| **白名单草案版本（当时）** | `21275605b5999f3acab31696e05057328e18afae` |

---

## §7 执行清单（终局拣选 — 已完成勾选）

依据 **`docs/HDGP_CORE_FINAL_PULL_PLAN.md`** §7：

1. **锚点**：拣选当日主系统 `HEAD` 记为 **`e60732ea78a29ddb168a41be8792dff96af5ee59`**（`main`）。  
2. **文件**：按 **`MATERIALS_ALLOWLIST.md`** §1 **逐文件**自主系统复制；节选处按 Meta-only 规则标注来源。  
3. **本文档**：写入锚点 SHA、日期、执行人，并与 **`MATERIALS_ALLOWLIST.md`** 互引。  
4. **README**：复核闭源主系统/商业区隔与能力边界；增加 **治理快照脚注**（主系统与 Core 治理机制可能不同）。  
5. **禁止项**：未纳入 `cmd/`、`internal/`、`scripts/`、运维证据目录（仓库内已核查）。  
6. **可选 `rulepacks/` 索引**（方案 §6）：本批次 **未** 添加。  
7. **合并前自检**：涉及更新的文档链接与双语段落已复核。

---

## 双仓隔离（终局拣选之后）

见上文 **Repository isolation** 中英对照。

---

## 主系统验收、冻结与 Core v1.0.0（交叉引用）

见上文 **Mainline formal acceptance, freeze & Core v1.0.0** 中英对照（同一内容）。

---

## §6 — Pull Guide（历史可复制块）

上文英文 **§6 — Pull Guide** 中 fenced `text` 块为 Pull Guide §6 **原文**；**终局批次**以英文 **§9** 块为准。

---

## §9 — 主系统负责人 → Core 维护者（终局拉取 · 已填写块）

见上文英文 **§9 — Mainline owner → Core maintainers** 中的 ```text``` 块（内容已填写锚点与日期判据）。

---

## 说明

- 无 Git 权限时：**`docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md`** §2 打包交接。  
- 钉住包含本快照与定稿白名单的 Core 提交：`git log -1 --format=%H -- docs/CORE_EXTRACT_SNAPSHOT.md`（在拉取后的 `main` 上执行）。
