# Core extract snapshot (post–Gate G)

> **When to use**: after mainline **Gate G** is closed; records one **pickable** specification state for `HDGP-Core` alignment. **Not** a code-sync credential.  
> **Related**: mainline policy `docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md` (in private `HDGP-Protocol`); allowlist finalized from PRE-GATE review of `docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md` → **`MATERIALS_ALLOWLIST.md`** in this repo.

---

## Snapshot metadata

| Field | Value |
|--------|--------|
| **Date** | 2026-04-20 |
| **Mainline repo** | Private Git `HDGP-Protocol` (local path on maintainer machine: `D:\HDGP\HDGP-Protocol`) |
| **Mainline Git SHA** | `ef8e041e6c3e42657d719656e2cc120b3fb9a25d` |
| **Branch** | `main` (as recorded at snapshot time) |
| **Gate G** | Closed; Core extraction proceeds under finalized allowlist only. |
| **Allowlist draft revision (mainline)** | File `docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md` at commit `21275605b5999f3acab31696e05057328e18afae` |
| **Core repo SHA at snapshot** | `c564a9c293dd91d755de80426f72652a84cb302b` (HEAD on `main` immediately before stage C landed; see Notes for pinning the stage C commit) |
| **Executor** | Core maintainers (per mainline instruction) |

## Notes

- This snapshot does **not** grant Core automatic ongoing sync; later changes need a new snapshot or a deliberate Core-side pick.  
- If mainline Git access is unavailable, use **packaged handoff** per `docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md` §2.  
- To pin the Core revision that includes this snapshot file and the finalized allowlist: `git log -1 --format=%H -- docs/CORE_EXTRACT_SNAPSHOT.md` (on `main` after pull).

---

## §6 — Mainline owner → Core maintainers (copy-paste block)

The following is the **verbatim** §6 block from mainline `docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md` (fenced `text`), for Core-side execution without opening the private repo:

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

If mainline has **not** finished stage C snapshot and finalized allowlist, complete checklist **§C** before issuing the instruction block above (same caveat as in the Pull Guide §6).

---

# Core 提取快照（门槛 G 之后）

> **用途**：主系统 **`HDGP_MAINLINE_BASELINE_GATE_CHECKLIST.md`** 门槛 **G** 已闭合后，记录一次「可拣选」规范状态，供 `HDGP-Core` 对齐；**非**代码自动同步凭证。  
> **关联**：主系统策略见私有仓内 `docs/HDGP_MAINLINE_BASELINE_FOR_CORE_EXTRACTION.md`；白名单由 PRE-GATE 对 `docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md` 的审阅定稿为本仓 **`MATERIALS_ALLOWLIST.md`**。

---

## 快照元数据

| 字段 | 值 |
|------|-----|
| **日期** | 2026-04-20 |
| **主系统仓库** | 私有 Git `HDGP-Protocol`（维护机本地路径示例：`D:\HDGP\HDGP-Protocol`） |
| **主系统 Git SHA** | `ef8e041e6c3e42657d719656e2cc120b3fb9a25d` |
| **分支** | 快照记录时 `main` |
| **门槛 G** | 已闭合；Core 仅按定稿白名单拣选。 |
| **白名单草案版本（主系统）** | 文件 `docs/HDGP_CORE_MATERIALS_ALLOWLIST_DRAFT.md` 于提交 `21275605b5999f3acab31696e05057328e18afae` |
| **本 Core 仓库 SHA（快照时）** | `c564a9c293dd91d755de80426f72652a84cb302b`（阶段 C 落地前 `main` 的 HEAD；钉住阶段 C 提交见上文说明） |
| **执行人** | Core 维护者（按主系统指令） |

## 说明

- 本快照**不**授予 Core 自动更新权；后续变更需新快照或 Core 侧自行择日拣选。  
- 若无法访问主系统 Git，按 `docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md` **§2** 走打包交接。  
- 钉住包含本快照与定稿白名单的 Core 提交：`git log -1 --format=%H -- docs/CORE_EXTRACT_SNAPSHOT.md`（在拉取后的 `main` 上执行）。

---

## §6 — 主系统负责人 → Core 维护者（可复制指令）

上文「§6 — Mainline owner → Core maintainers」中的 ```text``` 块为 **§6 全文**，与主系统 `docs/HDGP_CORE_MATERIALS_PULL_GUIDE.md` 保持一致，可直接复制为正式指令。

若主系统 **尚未** 完成阶段 C 快照与 allowlist 定稿，请先完成清单 **§C** 再发出上述指令块（与 Pull Guide §6 文末一致）。
