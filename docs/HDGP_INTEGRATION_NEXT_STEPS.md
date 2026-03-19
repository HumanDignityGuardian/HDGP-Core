# HDGP 集成技术落地与下一步行动

> 在集成规范（`spec/HDGP_INTEGRATION_SPEC.md`）以「无伦理框架的普遍项目」为主、流程框架已固化为 MD 的前提下，本文档说明：架构上 HDGP 是否仅靠「只读」即可工作、当前已有技术实现、以及可落地的下一步技术项与优先级。

---

## 一、架构上 HDGP 是否「只读即可工作」

**结论：是。**

- 集成方只需**读取**协议与 API 说明（MD 或未来的 OpenAPI），向 Engine **发送**一次 `POST /hdgp/v1/evaluate`（JSON），再**读取**响应中的 `verdict`、`rules_triggered`、`actions` 等，据此决定放行 / 修改 / 拦截。
- Engine **不向业务系统写入**任何数据，不要求数据库、配置注入或回调；单请求无状态。
- 因此从架构上，**HDGP 依赖「只读协议 + 单次请求/响应」即可工作**，无需在集成方侧部署 HDGP 自有组件（除非选用 Gateway 模式）。

---

## 二、当前已有技术实现（非仅 MD）

| 组件 | 说明 |
|------|------|
| **HDGP Engine（Go）** | `cmd/hdgp-engine`：`/hdgp/v1/evaluate`、`/status`、`/audit`、`/appeal`、`/chat`；4 条规则、模板化正式报告、`include_report` 参数 |
| **Meta 默认值** | `applyMetaDefaults`：缺失 scene 时填充 domain=general、intent=chat、risk_level=medium |
| **安全加固** | 请求体限制（evaluate/chat 1 MiB、appeal 64 KiB）、Server 超时与 MaxHeaderBytes |
| **合规测试** | `cmd/hdgp-conftest` + `conformance-tests/cases/*.json`，7 个用例 |
| **OpenAPI** | `spec/openapi.yaml`：evaluate、status 的机器可读描述 |
| **集成单页** | `docs/HDGP_INTEGRATION_ONEPAGER.md` |
| **Go SDK** | `pkg/client`：`Evaluate`、`EvaluateWithFallback`（fail-closed）|

尚未实现：**独立 Gateway 二进制**。

---

## 三、可落地的技术实现项（按优先级）

### P0：Engine 端 Meta 默认值（低成本、高收益）

- **目的**：对接「无伦理框架」且无显式 scene/domain/risk 的调用方，降低接入门槛。
- **做法**：在 `Evaluate` 入口或 handler 中，若 `meta.Scene` 为空或未设，则填充：`domain=general`、`intent=chat`、`risk_level=medium`。
- **产出**：调用方可不传 scene，Engine 行为与集成规范 §2.4 一致。

### P1：集成单页 / 最小示例

- **目的**：让「只读 MD + 一次请求」即可完成首次集成。
- **做法**：新增一页（如 `docs/HDGP_INTEGRATION_ONEPAGER.md` 或站内单页）：五步流程摘要 + 最小 `curl`/请求体示例 + 默认 Meta 说明 + 裁决含义与 fail-closed 建议。
- **产出**：集成方可仅凭此页 + 运行中的 Engine 完成对接。

### P2：机器可读 API 描述（OpenAPI 3.x） ✅ 已完成

- **目的**：从「仅 MD」升级为可被工具与 SDK 使用的契约。
- **做法**：`spec/openapi.yaml` 描述 `/hdgp/v1/evaluate`、`/hdgp/v1/status`。
- **产出**：支持代码生成、Mock 服务、文档站点与后续 SDK 生成。

### P3：轻量 SDK（单语言优先） ✅ 已完成

- **目的**：封装「构建请求 → POST → 解析响应」，减少集成方样板代码。
- **产出**：`pkg/client` Go 客户端，提供 `Evaluate`、`EvaluateWithFallback`（fail-closed）；见集成单页 §7。

### P4：Fail-closed 辅助 ✅ 已并入 P3

- **产出**：`EvaluateWithFallback` 在 Engine 不可用时返回 verdict=block，满足规范第八节建议。

### P5：独立 Gateway 二进制（后续）

- **目的**：对不愿改业务代码的集成方，提供「流量经 Gateway 再转上游」的部署形态。
- **做法**：独立服务：接收请求 → 转发上游（如 LLM）→ 对响应调用 Engine → 按 verdict 返回或阻断。
- **产出**：可选组件，与当前「只读 + 单次调用」的架构兼容。

---

## 四、建议的下一步行动顺序

1. **立即**：实现 **P0 Meta 默认值**（在 Engine 内，约一处逻辑）。
2. **短期**：编写 **P1 集成单页**（单 MD 或站内页），与 P0 一起保证「只读 + 最小请求」即可接入。
3. **中期**：产出 **P2 OpenAPI**，再视需要做 **P3 单语言 SDK**（建议 TypeScript 或 Go）。
4. **按需**：P4 fail-closed 辅助（可并入 P3 SDK）；P5 Gateway 作为独立迭代。

上述顺序保证：在**不改变「只读即可工作」架构**的前提下，从「仅 MD」逐步增加可执行、可工具化的技术实现，且优先服务「无伦理框架」项目的低门槛接入。
