# Adopter Checklist (HDGP-Core · Meta-only)

This checklist is for teams adopting **HDGP-Core as Meta-only baseline**. It focuses on “low change, high clarity” items that reduce misuse and ambiguity.

## A. Boundary & positioning

- [ ] Public docs explicitly state **Meta-only** and **no Judge/Engine shipped** (copy the boundary sentence from `docs/ADOPTION_BUNDLE.md`).
- [ ] Your product/site does **not** claim runtime enforcement, certification, audit hosting, or SLAs “from HDGP-Core”.
- [ ] You have a short **Do/Don’t** block (can reuse `README.md` Do/Don’t wording from Core).

## B. Meta field weaving (engineering)

- [ ] You include at least `scene.domain`, `scene.intent`, `scene.risk_level` for every governed output path.
- [ ] Defaults are defined when missing (`general` / `chat` / `medium`) and documented.
- [ ] `request_id` exists for traceability (or you document why not).
- [ ] You have a minimal “Meta mapping table” (where each field is sourced from: request, route, UI, config).

## C. Governance / review readiness

- [ ] You define what triggers **human review** (e.g. `risk_level >= high`, or specific sensitivity tags).
- [ ] You define **disclosure rules** (relationships, sponsorship, affiliate, data boundary).
- [ ] You define **source-quality thresholds** (if your system presents factual claims, summaries, or “authoritative” statements).

## D. Attribution & licensing

- [ ] You comply with `LICENSE` and include a **link-back** to `HDGP-Core` in a stable place (e.g. `NOTICE`, `ADOPTION.md`, website footer).
- [ ] If you copied any Core text, you keep **verbatim attribution** (see `docs/adopter-pack/ATTRIBUTION_AND_CITATION.md`).
- [ ] If you only adopted the semantics (no text copy), you still include a **semantic attribution** statement.

## E. Optional module (content sites / governance portals)

- [ ] If you are a content site, you decide whether to enable the optional **anti-promo / anti-diversion** module (`MODULE_CONTENT_SITE_ANTI_PROMO.md`).
- [ ] If enabled, you publish: anti-link policy + disclosure policy + enforcement scope (what you do and don’t do).

