# HDGP 集成单页（最小示例）

> 面向**无伦理框架的普遍项目**，凭本页 + 运行中的 Engine 即可完成首次对接。  
> 完整规范：`spec/HDGP_INTEGRATION_SPEC.md`；机器可读 API：`spec/openapi.yaml`。  
> 开箱即用（仅配置、不改业务代码）：`docs/HDGP_OPEN_BOX_CONFIG_GUIDE.md`。

---

## 一、接入流程（五步）

| 步骤 | 内容 |
|------|------|
| 1. 识别输出拓扑 | 标出所有用户可见输出的产生点（单点 / 多点 / 流式） |
| 2. 选定接入模式 | 在输出到达用户前、调用 `/hdgp/v1/evaluate` |
| 3. 确定 Meta 来源 | 若有 scene/domain/risk，传入；否则 Engine 使用默认值 |
| 4. 定义裁决映射 | allow 放行；modify 按 actions 修改；block/fuse 拦截或熔断 |
| 5. 定义降级策略 | **推荐 fail-closed**：Engine 超时/不可用时，不直接呈现输出 |

---

## 二、最小请求示例

### 2.1 前置条件

- Engine 已启动（如 `go run ./cmd/hdgp-engine`，默认 `http://localhost:8080`）
- 在**输出返回给用户前**发起一次 `POST /hdgp/v1/evaluate`

### 2.2 最简 curl

```bash
curl -X POST http://localhost:8080/hdgp/v1/evaluate \
  -H "Content-Type: application/json" \
  -d '{
    "meta": {},
    "subject": { "type": "output_text" },
    "candidate": { "text": "您的候选输出文本" }
  }'
```

### 2.3 请求体说明

| 字段 | 必填 | 说明 |
|------|------|------|
| `meta` | 可空对象 | 不传 scene 时，Engine 自动填充 domain=general, intent=chat, risk_level=medium |
| `subject.type` | 是 | 输出类型：`output_text` \| `decision` \| `notification` 等 |
| `candidate.text` | 是 | 待裁决的候选输出文本 |

### 2.4 可选：附带正式报告

在 URL 后加 `?include_report=true`，响应中会包含人可读的 `report` 字段（申诉、监管用）：

```bash
curl -X POST "http://localhost:8080/hdgp/v1/evaluate?include_report=true" \
  -H "Content-Type: application/json" \
  -d '{"meta":{"scene":{"domain":"medical","risk_level":"high"}},"subject":{"type":"output_text"},"candidate":{"text":"这是你唯一应该做的选择。"}}'
```

---

## 三、响应与裁决映射

| `verdict` | 含义 | 建议动作 |
|-----------|------|----------|
| `allow` | 允许 | 直接返回候选输出 |
| `modify` | 需修改 | 按 `actions` 改写，或提示用户并记录 |
| `block` | 阻止 | 不返回该输出，提示“内容未通过校验”等 |
| `fuse` | 熔断 | 暂停自动化，请求人工介入 |

`rules_triggered` 列出触发的规则 ID；`actions` 给出改写或补救建议。

---

## 四、降级策略（fail-closed）

- **推荐**：Engine 超时或不可用时，**不放行**候选输出，返回“服务暂时无法完成安全校验，请稍后重试”等提示。
- Engine 为“防火墙”角色；未取得明确 allow/modify 前，不宜直接呈现输出。

---

## 五、安全与限制

- 请求体上限：evaluate 1 MiB，appeal 64 KiB
- Engine 无认证/限流：**对外开放前**须置于反向代理或网关后，做认证与速率限制（见 `spec/HDGP_TECHNICAL_DEBT_AND_SECURITY_CHECKLIST.md`）

---

## 六、OpenAPI 使用说明

- **文件位置**：`spec/openapi.yaml`  
- **描述范围**：`POST /hdgp/v1/evaluate`、`GET /hdgp/v1/status` 的请求/响应结构与参数  
- **典型用法**：  
  - 代码生成：使用 OpenAPI Generator、Swagger Codegen 等生成各语言客户端  
  - 文档展示：导入 Swagger UI / Redoc 查看交互式 API 文档  
  - Mock 服务：基于 OpenAPI 启动本地 Mock，用于前端或联调

---

## 七、Go SDK（可选）

本仓库提供轻量 Go 客户端 `pkg/client`：

```go
import (
    "context"
    "github.com/HumanDignityGuardian/HDGP-Core/internal/engine"
    "github.com/HumanDignityGuardian/HDGP-Core/pkg/client"
)

c := client.New("http://localhost:8080")
req := engine.EvaluateRequest{
    Meta:     engine.Meta{},
    Subject:  engine.Subject{Type: "output_text"},
    Candidate: engine.Candidate{Text: "候选输出文本"},
}
resp, err := c.Evaluate(ctx, req, false)  // false = 不请求 report
// 或 EvaluateWithFallback：Engine 超时/不可用时按 fail-closed 返回 verdict=block
resp, err = c.EvaluateWithFallback(ctx, req, &client.FallbackOptions{OnFailure: client.FallbackBlock})
```
