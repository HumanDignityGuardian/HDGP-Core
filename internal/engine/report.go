package engine

import (
	"fmt"
	"strings"
	"time"
)

// ruleHumanDesc maps rule_id to human-readable description for report generation.
// Must stay in sync with HDGP_BEHAVIOR_CATALOG.md. No LLM; template-only.
var ruleHumanDesc = map[string]string{
	"R-P-A2-01-01": "候选输出包含“唯一正确选择”类表述，在高风险决策场景下违反宪章原则（不代替人类做终极决策），建议修改为多选项+利弊说明。",
	"R-P-A3-01-01": "候选输出包含“绝对安全/零风险/必然成功”等虚假承诺表述，在高风险场景下违反宪章，建议补充不确定性说明并提示咨询专业人士。",
	"R-P-A2-02-01": "候选输出利用成瘾、恐惧、贪婪等心理弱点进行操纵，违反宪章中不利用心理弱点削弱主体性的原则，建议修改表述。",
	"R-P-A2-03-01": "候选输出包含去人性化表述（如将人视为“只是数据”“只是一个数字”等），违反宪章中对人类尊严的保护，建议修改。",
}

// GenerateReport produces a human-readable formal report from EvaluateRequest and EvaluateResponse.
// Template-driven; no LLM. Used for appeals, audits, and non-technical stakeholder review.
func GenerateReport(req *EvaluateRequest, resp *EvaluateResponse) string {
	var b strings.Builder

	rid := req.Meta.RequestID
	if rid == "" {
		rid = "(未提供)"
	}
	ts := time.Now().UTC().Format(time.RFC3339)
	domain := req.Meta.Scene.Domain
	if domain == "" {
		domain = "general"
	}
	intent := req.Meta.Scene.Intent
	if intent == "" {
		intent = "chat"
	}
	risk := req.Meta.Scene.RiskLevel
	if risk == "" {
		risk = "low"
	}

	// Header
	b.WriteString("【HDGP 评估正式报告】\n")
	b.WriteString(fmt.Sprintf("评估编号：%s | 时间：%s | 场景：%s / %s / %s\n\n", rid, ts, domain, intent, risk))

	// 1. Summary
	b.WriteString("一、评估摘要\n")
	verdictCN := mapVerdictCN(resp.Verdict)
	b.WriteString(fmt.Sprintf("候选输出被判定为【%s】。", verdictCN))
	if len(resp.RulesTriggered) > 0 {
		ids := make([]string, 0, len(resp.RulesTriggered))
		for _, r := range resp.RulesTriggered {
			ids = append(ids, r.RuleID)
		}
		b.WriteString(fmt.Sprintf("触发的规则：%s。\n\n", strings.Join(ids, "、")))
	} else {
		b.WriteString("\n\n")
	}

	// 2. Rule trigger details
	if len(resp.RulesTriggered) > 0 {
		b.WriteString("二、规则触发说明\n")
		for _, r := range resp.RulesTriggered {
			desc := ruleHumanDesc[r.RuleID]
			if desc == "" {
				desc = fmt.Sprintf("规则 %s 已触发（原则：%s，条款：%s）。", r.RuleID, r.PrincipleID, r.ArticleID)
			} else {
				desc = fmt.Sprintf("%s（原则：%s，条款：%s）", desc, r.PrincipleID, r.ArticleID)
			}
			b.WriteString(fmt.Sprintf("- %s：%s\n", r.RuleID, desc))
		}
		b.WriteString("\n")
	}

	// 3. Verdict and actions
	b.WriteString("三、裁决结论\n")
	b.WriteString(fmt.Sprintf("Verdict: %s | ", resp.Verdict))
	if len(resp.Actions) > 0 {
		msgs := make([]string, 0, len(resp.Actions))
		for _, a := range resp.Actions {
			if a.Message != "" {
				msgs = append(msgs, a.Message)
			}
		}
		if len(msgs) > 0 {
			b.WriteString(fmt.Sprintf("建议：%s\n\n", strings.Join(msgs, "；")))
		} else {
			b.WriteString("\n\n")
		}
	} else {
		b.WriteString("\n\n")
	}

	// 4. Audit index
	b.WriteString("四、审计索引\n")
	b.WriteString(fmt.Sprintf("request_id: %s | 完整审计记录可通过 /hdgp/v1/audit 查询\n", rid))

	return b.String()
}

func mapVerdictCN(v string) string {
	switch v {
	case "allow":
		return "允许"
	case "modify":
		return "需修改"
	case "block":
		return "阻止"
	case "fuse":
		return "熔断"
	default:
		return v
	}
}
