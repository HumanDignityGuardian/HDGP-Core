# HDGP 开箱即用配置说明

> 本文档说明如何通过**仅改配置、不改业务代码**的方式接入 HDGP，适用于已建成项目（如 X-TSOS、Intdone）与无伦理框架系统。  
> 前置条件：应用具备**可配置的出口**（如模型 API 地址、审核服务 URL、代理地址等）。

---

## 一、开箱即用含义

- **不修改业务逻辑**：不在应用源码中增加或改写与 HDGP 相关的业务代码。  
- **不重构源码**：不因接入 HDGP 而调整项目结构或核心流程。  
- **仅配置与部署**：通过环境变量、配置文件或部署拓扑，将「候选输出」指向 HDGP Engine（evaluate）。

若应用当前**没有**可配置的 HDGP Engine evaluate 地址，则需在应用侧增加**最少一处**配置项（例如「内容决策/裁决服务地址」），仍不属于业务逻辑重构。

---

## 二、典型接入方式（仅配置）

### 2.1 应用已支持「候选输出决策/审核」URL 可配置

- 将该 URL 指向 HDGP Engine 的 evaluate 地址（如 `http://localhost:8080/hdgp/v1/evaluate`）。  
- 应用在展示前把“候选输出”发送给 Engine；Engine 返回 verdict，应用按结果放行/拦截/改写。  
- 无需改业务逻辑，仅改配置。

### 2.2 应用具备降级/超时策略可配置

- 当 Engine 不可用或超时时，把回退策略配置为 fail-closed（例如返回 block/deny）。  
- 无需改业务逻辑，只需调整降级策略配置与 Engine 超时参数。

---

## 三、配置项速查（仅 Engine evaluate）

| 配置项 | 说明 | 示例 |
|--------|------|------|
| Engine 监听地址 | Engine 监听端口/地址 | `HDGP_ENGINE_ADDR=:8080` |
| 审计日志落盘 | 审计记录追加写入文件（可选） | `HDGP_AUDIT_LOG_PATH=./audit.jsonl` |

具体以各组件文档或 Release 说明为准。

---

## 四、回退与旁路

- 若需恢复「不经 HDGP」的行为：将 evaluate 地址恢复为原决策/审核逻辑，或在配置中停用对 HDGP Engine evaluate 的调用即可。  
- 无需改代码，仅改配置即可撤回 HDGP，保证可逆。

---

## 五、相关文档

- 集成流程与最小示例：`docs/HDGP_INTEGRATION_ONEPAGER.md`  
- 集成规范：`spec/HDGP_INTEGRATION_SPEC.md`  
- 开源与商业区隔与宪章保障：`HDGP_COMMERCIAL_PLAN.md` 第七节、`docs/HDGP_OPEN_SOURCE_AND_COMMERCIAL.md`
