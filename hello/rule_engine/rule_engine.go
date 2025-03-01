package rule_engine

import "context"

func NewRuleEngine() *RuleEngine {
	return &RuleEngine{
		Rules: []*Rule{},
	}
}

// RuleEngine 规则引擎
type RuleEngine struct {
	Rules []*Rule `json:"rules"` // 规则列表
}

// Rule 表示一个规则，包含条件集和动作集。
type Rule struct {
	Name        string       `json:"name"`        // 规则名称
	Description string       `json:"description"` // 规则描述
	Conditions  []*Condition `json:"conditions"`  // 条件集
	Actions     []*Action    `json:"actions"`     // 动作集
	Enabled     bool         `json:"enabled"`     // 规则是否启用
}

// Condition 表示一个条件，包含字段名、操作符和值。
type Condition struct {
	FieldName    FieldName     `json:"fieldName"`              // 字段名
	Operator     OperatorType  `json:"operator"`               // 操作符
	FieldValue   interface{}   `json:"fieldValue"`             // 值
	ConditionSet *ConditionSet `json:"conditionSet,omitempty"` // 嵌套条件集合, 用于实现嵌套条件，支持 AND/OR 逻辑
}

// ConditionSet 表示一个嵌套条件集合，用于实现嵌套条件，支持 AND/OR 逻辑。
type ConditionSet struct {
	IsOr       bool         `json:"isOr"`       // 是则使用 OR 逻辑，否则使用 AND 逻辑
	Conditions []*Condition `json:"conditions"` // 条件集合
	Enabled    bool         `json:"enabled"`    // 是否启用
}

// Action 表示一个动作，包含动作类型和参数。
type Action struct {
	ActionType ActionType             `json:"actionType"` // 动作类型
	Parameters map[string]interface{} `json:"parameters"` // 动作参数
}

// Fact 接口，定义 Fact 需要实现的方法
type Fact interface {
	GetValue(ctx context.Context, fieldName string) interface{}                              // 根据字段名获取 Fact 中的值
	ExecuteAction(ctx context.Context, actionType string, parameters map[string]interface{}) // 执行动作
}

// AddRule 添加规则
func (re *RuleEngine) AddRule(rule *Rule) {
	re.Rules = append(re.Rules, rule)
}

// EvaluateFact 返回 Fact 中的 Action
func (re *RuleEngine) EvaluateFact(ctx context.Context, fact Fact) []*Action {
	var firedActions []*Action
	for _, rule := range re.Rules {
		if rule.Enabled && re.evaluateRule(ctx, rule, fact) { // 评估单个规则
			firedActions = append(firedActions, rule.Actions...)
		}
	}
	return firedActions
}

// ExecuteFact 执行 Fact 中的 Action
func (re *RuleEngine) ExecuteFact(ctx context.Context, fact Fact) {
	actions := re.EvaluateFact(ctx, fact) // 调用 EvaluateFact 获取 Actions
	for _, action := range actions {
		fact.ExecuteAction(ctx, string(action.ActionType), action.Parameters) //  执行 Action
	}
}

// evaluateRule 评估单个规则是否满足 Fact
func (re *RuleEngine) evaluateRule(ctx context.Context, rule *Rule, fact Fact) bool {
	return re.evaluateConditions(ctx, rule.Conditions, fact)
}

// evaluateConditions 评估多个条件是否满足，AND 逻辑
func (re *RuleEngine) evaluateConditions(ctx context.Context, conditions []*Condition, fact Fact) bool {
	if len(conditions) == 0 {
		return true // 没有条件则默认满足
	}
	for _, condition := range conditions {
		if condition.ConditionSet != nil && condition.ConditionSet.Enabled { // 优先处理 ConditionSet
			if !re.evaluateConditionSet(ctx, condition.ConditionSet, fact) { // 评估 ConditionSet
				return false
			}
		} else { // 处理单个 Condition
			fieldValueFromFact := fact.GetValue(ctx, string(condition.FieldName)) // 使用 string(FieldName) 获取字段名
			if fieldValueFromFact == nil {
				return false // Fact 中不存在 Condition 中指定的字段
			}
			if !EvaluateCondition(condition, fieldValueFromFact) { // 调用 condition_operator.go 中的 EvaluateCondition
				return false
			}
		}
	}
	return true // 所有条件都满足 (AND 逻辑)
}

// evaluateConditionSet 评估 ConditionSet (支持 AND/OR 逻辑)
func (re *RuleEngine) evaluateConditionSet(ctx context.Context, conditionSet *ConditionSet, fact Fact) bool {
	if conditionSet.IsOr { // OR 逻辑
		for _, condition := range conditionSet.Conditions {
			if condition.ConditionSet != nil && condition.ConditionSet.Enabled { // 优先处理 ConditionSet
				if re.evaluateConditionSet(ctx, condition.ConditionSet, fact) { // 递归评估 ConditionSet
					return true // 嵌套 ConditionSet 满足，整个 ConditionSet 满足 OR 条件
				}
			} else {
				fieldValueFromFact := fact.GetValue(ctx, string(condition.FieldName))
				if fieldValueFromFact != nil && EvaluateCondition(condition, fieldValueFromFact) {
					return true // 任意一个 Condition 满足，整个 ConditionSet 满足 OR 条件
				}
			}
		}
		return false
	} else { // AND 逻辑
		for _, condition := range conditionSet.Conditions {
			if condition.ConditionSet != nil && condition.ConditionSet.Enabled { // 优先处理 ConditionSet
				if !re.evaluateConditionSet(ctx, condition.ConditionSet, fact) { // 递归评估 ConditionSet
					return false // 嵌套 ConditionSet 不满足，整个 ConditionSet 不满足 AND 条件
				}
			} else {
				fieldValueFromFact := fact.GetValue(ctx, string(condition.FieldName))
				if fieldValueFromFact == nil || !EvaluateCondition(condition, fieldValueFromFact) {
					return false // 任意一个 Condition 不满足，整个 ConditionSet 不满足 AND 条件
				}
			}
		}
		return true
	}
}
