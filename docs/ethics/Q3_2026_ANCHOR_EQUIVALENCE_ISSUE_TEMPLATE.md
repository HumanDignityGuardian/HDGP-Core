# Issue draft — 2026-Q3 anchor equivalence index (copy into GitHub Issue)

> **Suggested title:** `[Governance] 2026-Q3: cross-repo anchor equivalence index (section mapping; bytes may differ)`  
> **Labels (optional):** `governance`, `ethics`, `documentation`  
> **Repos:** open in **HDGP-Core**; link private mainline counterpart in comments.

---

## Summary (EN)

Following **2026-Q2** attestation (**diverged**, disclosed) in `docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_2026Q2.md`, propose a **2026-Q3** path that does **not** assume byte-identical files across Core and private mainline, but **does** give external reviewers a **stable, published mapping** between anchor sections.

**Chosen track (recommended):** **Section-level equivalence index** — maintain separate canonical bodies per repo; publish a quarterly **equivalence table** (section ID / clause ref ↔ intent summary ↔ both sides “satisfied Y/N”) so `aligned` in Q3 can mean **semantic equivalence under mapping**, not same SHA-256.

---

## Background

| Item | Core | Mainline (private) |
|------|------|------------------|
| Ethics baseline hash | `af9b5f5a…` | `48eaa0aa…` |
| Meta/Judge scope hash | `cd612df5…` | `b0de2978…` |
| Editorial model | EN-first + ZH mirror (Meta-only pick) | ZH-primary, shorter independent evolution |
| Q2 status | Core internal hashes OK; **Core ↔ mainline diverged (disclosed)** | per mainline Q2 attestation |

---

## Proposal

1. Add allowlisted doc (name TBD), e.g. `docs/ethics/ANCHOR_EQUIVALENCE_INDEX_2026Q3.md`, maintained **quarterly**.
2. Table columns (minimum):
   - Anchor ID (stable across quarters where possible)
   - Core path + section heading
   - Mainline path + section heading (reference only; file may live in private repo)
   - Equivalence intent (1–2 sentences)
   - Reviewer sign-off column (Y/N/NA)
3. **2026-Q3 attestation** may cite this index; target status: **aligned under published mapping** OR **diverged with explicit gaps listed**.
4. No standing obligation to merge or sync full file bytes between repos.

---

## Out of scope

- Engine/Judge code sync
- Forcing Core and mainline to a single language/edition canon in Q3
- Replacing CHIP for actual ethics baseline **clause** changes

---

## Acceptance criteria

- [ ] Draft equivalence index reviewed by Core maintainer + mainline representative
- [ ] Linked from `docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_2026Q3.md` (when created)
- [ ] `MATERIALS_ALLOWLIST.md` updated if new file is promoted to allowlisted material
- [ ] Q3 attestation records outcome: mapped-aligned vs residual diverged gaps

---

## 摘要（中文）

在 **2026-Q2** 已披露 **diverged** 的前提下，Q3 不追求双仓同字节，而发布**章节级等价索引表**，使外部复核可判断「语义对齐」而非仅比对 SHA。

**推荐路径：** 维护各自正本 + 季度 **`ANCHOR_EQUIVALENCE_INDEX_YYYYQn.md`** 映射表；Q3 声明可写「在已发布映射下 aligned」或「映射后仍存在的差异」。

---

## References

- Core: `docs/ethics/ETHICS_ALIGNMENT_ATTESTATION_2026Q2.md`
- Core: `GOVERNANCE.md` §8 (quarterly ethics alignment)
- Core: `docs/CORE_EXTRACT_SNAPSHOT.md` (repository isolation)
