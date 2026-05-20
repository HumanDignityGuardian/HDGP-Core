# Ethics Alignment Attestation 鈥?2026Q2 (HDGP-Core)

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
| Commit (optional) | `666a25e4603862338880cf3912add11c159aab4e` (post鈥搗1.0.2 attestation UTF-8 fix; re-verify if branch moved) |

---

## Anchor files & hashes (SHA-256) 鈥?Core repository

| File (in this repo) | SHA-256 (Core) | Verified |
|---------------------|----------------|----------|
| `spec/HDGP_ETHICS_BASELINE.md` | `af9b5f5a71941ae6b3aaede7223d0bd9d59bc1b76b4b93387e4f86cbec0284e2` | Yes 鈥?matches file bytes in Core |
| `spec/HDGP_META_VS_JUDGE_SCOPE.md` | `cd612df52814ce25f7caf6b6c811e1b53b26f706a1078da774a489aa5f0a86ee` | Yes 鈥?matches file bytes in Core |

> **Note**: `af9b5f5a鈥 applies only to the ethics baseline file; `cd612df5鈥 applies only to the Meta/Judge scope file (distinct anchors).  
> Cross-repo comparison vs private mainline (2026-Q2): mainline anchors are **`48eaa0aa鈥** (ethics baseline) and **`b0de2978鈥** (Meta/Judge scope) 鈥?**not byte-identical** to Core; see **Alignment status** below.

---

## Alignment status (quarterly: Core 鈫?private mainline)

- [ ] aligned
- [x] diverged

**Core internal (declaration 鈫?files in this repo):** consistent 鈥?anchor hashes above match Core files at sign-off.

**Core 鈫?mainline:** **diverged** 鈥?**disclosed**, not a silent fork. After the final Meta-only pick (2026-05-03), Core ships an **English-first + ZH-CN mirror** Meta-only canon; mainline retains a **Chinese-primary, shorter** independent evolution (~196 vs ~356 lines for ethics baseline; Meta/Judge scope sizes differ similarly). This does not contradict Core handoff acceptance 搂2.3 (鈥淐ore attestation hashes match **Core** files鈥?; it is not a claim of **byte-identical** dual-repo bodies.

---

## If diverged: differences & actions (fill by humans)

**Summary (EN):**

- Mainline `HDGP_ETHICS_BASELINE.md`: `48eaa0aa鈥 (18,978 B) vs Core `af9b5f5a鈥 (28,367 B) 鈥?different editorial structure and length; same path, not same bytes.
- Mainline `HDGP_META_VS_JUDGE_SCOPE.md`: `b0de2978鈥 (2,895 B) vs Core `cd612df5鈥 (6,078 B) 鈥?same path, not same bytes.
- Prior mainline draft anchor `13fc4efc鈥 is **superseded**; current mainline anchors per mainline `ETHICS_ALIGNMENT_ATTESTATION_2026Q2.md`.
- Divergence is **expected** under repository isolation policy; documented in this attestation and mainline review (not undisclosed drift).

**Next steps (EN):**

- **2026-Q3 onward**: pursue **aligned** only via explicit CHIP / written equivalence mapping (unified anchor edition or section-level mapping) 鈥?**do not** assume same path implies same bytes.
- Track in Core governance (Issue/CHIP as needed); no standing obligation for continuous code or document sync between repos.

**鎽樿锛堜腑鏂囷級锛?*

- 涓荤郴缁熶笌 Core 鍚屽悕閿氱偣鏂囦欢**闈炲瓧鑺備竴鑷?*锛汣ore 涓虹粓灞€鎷ｉ€夊悗鐨?Meta-only 鑻辨枃鍦ㄥ墠缁堢 + 涓枃瀵圭収锛屼富绯荤粺涓轰腑鏂囦富绋跨嫭绔嬫紨杩涳紝绡囧箙涓庡搱甯屽潎涓嶅悓銆?- 鍒嗘涓?*宸叉姭闇?*鐨勯鏈熷唴宸紓锛岄潪闈欓粯鍒嗗弶锛涗富绯荤粺鏃ч敋鐐?`13fc4efc鈥 宸茶繃鏃讹紝褰撳涓荤郴缁熼敋鐐逛互 `48eaa0aa鈥 / `b0de2978鈥 涓哄噯锛堣涓荤郴缁熷綋瀛ｅ０鏄庯級銆?- 鑻?2026-Q3 鍙婁互鍚庨渶杈惧埌 **aligned**锛岄』缁?CHIP/涔﹂潰銆岀瓑浠风珷鑺傘€嶆垨缁熶竴浣撲緥鏄犲皠锛屼笉鑳藉亣瀹氳矾寰勫悓鍚嶅嵆鍚屽瓧鑺傘€?
**涓嬩竴姝ワ紙涓枃锛夛細**

- 2026-Q3 璧锋寜闇€璧?CHIP/涔﹂潰鏄犲皠锛涙棤鎸佺画鍙屼粨鏂囨。鍚屾涔夊姟銆?
---

## Sign-off (fill by humans)

| Role | Name/Handle | Date |
|------|-------------|------|
| Maintainer / authorized representative | Yvaine He | 2026-05-20 |

---

## 涓枃鐗堟湰 (ZH-CN)

> 浠ヤ笅涓枃涓庝笂鏂囪嫳鏂囧搴旓紱绀惧尯闃呰椤哄簭浠ヨ嫳鏂囦负鍏堛€?
# 浼︾悊瀵归綈瀛ｅ害澹版槑 鈥?2026Q2锛圚DGP-Core锛?
> **浠撳簱**锛氭湰鏂囦欢閫傜敤浜庡紑婧?**`HDGP-Core`**銆傚搴︽斂绛栨ā鏉垮彲瑙佺鏈変富绯荤粺 `docs/HDGP_QUARTERLY_ETHICS_ALIGNMENT_POLICY.md`锛?*鏈?*闅?Core 鍏ㄦ枃鍙戝竷鏃讹紝浠呬綔鍙ｅ緞寮曠敤锛夈€? 
> **娴佺▼**锛氬～鍐欏厓鏁版嵁銆侀敋鐐瑰搱甯屻€佸榻愮粨璁轰笌绛剧讲锛涘彂甯冨墠椤荤粡浜哄伐澶嶆牳涓庣缃层€?
---

## 鍏冩暟鎹?
| 瀛楁 | 鍊?|
|------|-----|
| 瀛ｅ害 | 2026-Q2 |
| 鏃ユ湡 | 2026-05-20 |
| 浠撳簱 | HDGP-Core |
| 鍒嗘敮 | main |
| 鎻愪氦锛堝彲閫夛級 | `666a25e4603862338880cf3912add11c159aab4e`锛坴1.0.2 鍚?attestation 淇锛涜嫢鍒嗘敮宸插墠杩涘彂甯冨墠璇峰鏍革級 |

---

## 閿氱偣鏂囦欢涓庡搱甯岋紙SHA-256锛夆€?Core 浠撳簱

| 鏂囦欢锛堟湰浠撳簱鍐咃級 | SHA-256锛圕ore锛?| 宸叉牳瀵?|
|------------------|-----------------|--------|
| `spec/HDGP_ETHICS_BASELINE.md` | `af9b5f5a71941ae6b3aaede7223d0bd9d59bc1b76b4b93387e4f86cbec0284e2` | 鏄?鈥?涓?Core 鍐呮枃浠跺瓧鑺備竴鑷?|
| `spec/HDGP_META_VS_JUDGE_SCOPE.md` | `cd612df52814ce25f7caf6b6c811e1b53b26f706a1078da774a489aa5f0a86ee` | 鏄?鈥?涓?Core 鍐呮枃浠跺瓧鑺備竴鑷?|

> **璇存槑**锛歚af9b5f5a鈥 浠呮寚浼︾悊鍩虹嚎鏂囦欢锛沗cd612df5鈥 浠呮寚 Meta/Judge 杈圭晫鏂囦欢锛堜袱涓敋鐐逛笉鍙贩鐢級銆? 
> 涓庣鏈変富绯荤粺浜ゅ弶姣斿锛?026-Q2锛夛細涓荤郴缁熼敋鐐逛负 **`48eaa0aa鈥**锛堜鸡鐞嗗熀绾匡級銆?*`b0de2978鈥**锛圡eta/Judge 杈圭晫锛夆€?涓?Core **闈炲瓧鑺備竴鑷?*锛涜涓嬫枃 **瀵归綈鐘舵€?*銆?
---

## 瀵归綈鐘舵€侊紙瀛ｅ害锛欳ore 鈫?绉佹湁涓荤郴缁燂級

- [ ] aligned锛堜竴鑷达級
- [x] diverged锛堜笉涓€鑷达級

**Core 鍐呴儴锛堟湰浠撳０鏄?鈫?鏈粨鏂囦欢锛夛細** 涓€鑷?鈥?涓婃枃閿氱偣鍝堝笇涓庣缃叉椂 Core 鏂囦欢鐩哥銆?
**Core 鈫?涓荤郴缁燂細** **涓嶄竴鑷达紙diverged锛?* 鈥?灞?*宸叉姭闇?*鍒嗘锛岄潪闈欓粯鍒嗗弶銆傜粓灞€鎷ｉ€夛紙2026-05-03锛夊悗锛孋ore 涓?**鑻辨枃鍦ㄥ墠 + 涓枃瀵圭収** 鐨?Meta-only 缁堢锛涗富绯荤粺淇濈暀 **涓枃涓荤銆佹洿鐭?* 鐨勭嫭绔嬫紨杩涚増鏈€傝繖涓嶄笌 Core 楠屾敹鏉愭枡 搂2.3銆屽０鏄庡搱甯屼笌 **Core** 鏂囦欢涓€鑷淬€嶇煕鐩撅紱**涓?*琛ㄧず鍙屼粨姝ｆ枃鍚屽瓧鑺傘€?
---

## 鑻ヤ笉涓€鑷达細宸紓涓庡缃?
**鎽樿锛?* 瑙佽嫳鏂囪妭 **Summary锛圗N锛?* 涓?**鎽樿锛堜腑鏂囷級**锛堜富绯荤粺鍝堝笇 `48eaa0aa鈥 / `b0de2978鈥锛汣ore 涓?`af9b5f5a鈥 / `cd612df5鈥锛涗富绯荤粺鏃х `13fc4efc鈥 宸茶繃鏃讹級銆?
**涓嬩竴姝ワ細** 瑙佽嫳鏂?**Next steps** 涓庝腑鏂?**涓嬩竴姝?*锛?026-Q3 鑻ラ渶 aligned 椤?CHIP/涔﹂潰鏄犲皠锛夈€?
---

## 绛剧讲

| 瑙掕壊 | 濮撳悕/Handle | 鏃ユ湡 |
|------|-------------|------|
| 缁存姢鑰?/ 鎺堟潈浠ｈ〃 | Yvaine He | 2026-05-20 |
