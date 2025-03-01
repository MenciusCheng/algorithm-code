package main

import (
	"context"
	"fmt"
	"github.com/MenciusCheng/algorithm-code/hello/rule_engine"
	"time"
)

func main() {
	// 示例规则
	rule1 := &rule_engine.Rule{
		Name:        "Rule for high temperature",
		Description: "This rule fires when the temperature is above 30 degrees Celsius.",
		Enabled:     true,
		Conditions: []*rule_engine.Condition{
			{FieldName: rule_engine.FieldNameTemperature, Operator: rule_engine.OperatorGreaterThan, FieldValue: float64(30)},
		},
		Actions: []*rule_engine.Action{
			{ActionType: rule_engine.ActionTypePrint, Parameters: map[string]interface{}{"message": "Temperature is high!"}},
		},
	}

	rule2 := &rule_engine.Rule{
		Name:        "Rule for admin user and tag 'vip'",
		Description: "This rule fires when the user role is admin and has 'vip' tag.",
		Enabled:     true,
		Conditions: []*rule_engine.Condition{
			{FieldName: rule_engine.FieldNameRole, Operator: rule_engine.OperatorEqual, FieldValue: "admin"},
			{FieldName: rule_engine.FieldNameTags, Operator: rule_engine.OperatorIn, FieldValue: []interface{}{"vip"}}, // FieldValue 使用 []interface{}
		},
		Actions: []*rule_engine.Action{
			{ActionType: rule_engine.ActionTypePrint, Parameters: map[string]interface{}{"message": "Admin VIP user detected!"}},
		},
	}

	rule3 := &rule_engine.Rule{
		Name:        "Rule for sending email to high temperature users",
		Description: "This rule sends an email to users when the temperature is high.",
		Enabled:     true,
		Conditions: []*rule_engine.Condition{
			{FieldName: rule_engine.FieldNameTemperature, Operator: rule_engine.OperatorGreaterEqual, FieldValue: float64(32)},
		},
		Actions: []*rule_engine.Action{
			{ActionType: rule_engine.ActionTypeSendEmail, Parameters: map[string]interface{}{
				"recipient": "{{.fact.email}}",
				"message":   "高温预警！请注意防暑降温。",
			}},
		},
	}

	rule4 := &rule_engine.Rule{
		Name:        "Rule with ConditionSet - Admin or VIP and Enabled",
		Description: "This rule fires if the user is admin OR has 'vip' tag AND is enabled.",
		Enabled:     true,
		Conditions: []*rule_engine.Condition{
			{
				ConditionSet: &rule_engine.ConditionSet{
					IsOr:    true,
					Enabled: true,
					Conditions: []*rule_engine.Condition{
						{FieldName: rule_engine.FieldNameRole, Operator: rule_engine.OperatorEqual, FieldValue: "admin"},
						{FieldName: rule_engine.FieldNameTags, Operator: rule_engine.OperatorIn, FieldValue: []interface{}{"vip"}},
					},
				},
			},
			{FieldName: rule_engine.FieldNameEnabled, Operator: rule_engine.OperatorTrue, FieldValue: true}, // 使用 OperatorTrue
		},
		Actions: []*rule_engine.Action{
			{ActionType: rule_engine.ActionTypePrint, Parameters: map[string]interface{}{"message": "Rule with ConditionSet fired!"}},
		},
	}

	rule5 := &rule_engine.Rule{
		Name:        "Rule for temperature less than or equal to 25 degrees Celsius",
		Description: "This rule fires when the temperature is less than or equal to 25 degrees Celsius.",
		Enabled:     true,
		Conditions: []*rule_engine.Condition{
			{FieldName: rule_engine.FieldNameTemperature, Operator: rule_engine.OperatorLessEqual, FieldValue: float64(25)},
		},
		Actions: []*rule_engine.Action{
			{ActionType: rule_engine.ActionTypePrint, Parameters: map[string]interface{}{"message": "Temperature is comfortable."}},
		},
	}

	rule6 := &rule_engine.Rule{
		Name:        "Rule for temperature on or after 2025-03-02",
		Description: "This rule fires when the temperature is checked on or after 2025-03-02.",
		Enabled:     true,
		Conditions: []*rule_engine.Condition{
			{FieldName: rule_engine.FieldNameTemperature, Operator: rule_engine.OperatorGreaterEqual, FieldValue: "2025-03-02T00:00:00Z"}, // 使用日期字符串 (ISO 8601)
		},
		Actions: []*rule_engine.Action{
			{ActionType: rule_engine.ActionTypePrint, Parameters: map[string]interface{}{"message": "Temperature check after 2025-03-02."}},
		},
	}

	rule7 := &rule_engine.Rule{
		Name:        "Rule for status is false",
		Description: "This rule fires when the status is false.",
		Enabled:     true,
		Conditions: []*rule_engine.Condition{
			{FieldName: rule_engine.FieldNameStatus, Operator: rule_engine.OperatorFalse, FieldValue: false}, // 使用 OperatorFalse
		},
		Actions: []*rule_engine.Action{
			{ActionType: rule_engine.ActionTypePrint, Parameters: map[string]interface{}{"message": "Status is false."}},
		},
	}

	engine := rule_engine.NewRuleEngine()
	engine.AddRule(rule1)
	engine.AddRule(rule2)
	engine.AddRule(rule3)
	engine.AddRule(rule4)
	engine.AddRule(rule5)
	engine.AddRule(rule6)
	engine.AddRule(rule7)

	// 示例 Fact
	fact1 := &UserFact{
		Temperature: 35,
		Role:        "guest",
		Tags:        []string{"normal"},
		Email:       "user1@example.com",
		Status:      true,
		Enabled:     true,
	}

	fact2 := &UserFact{
		Temperature: 25,
		Role:        "admin",
		Tags:        []string{"vip", "staff"},
		Email:       "admin@example.com",
		Status:      false,
		Enabled:     true,
	}

	fact3 := &UserFact{
		Temperature: 33,
		Role:        "normal",
		Tags:        []string{"normal"},
		Email:       "user3@example.com",
		Status:      true,
		Enabled:     false, // disabled for rule4
	}

	fact4 := &UserFact{
		Temperature: 20,
		Role:        "guest",
		Tags:        []string{"normal"},
		Email:       "user4@example.com",
		Status:      true,
		Enabled:     true,
	}

	fact5 := &UserFact{
		Temperature: 33,
		Role:        "vip_guest",
		Tags:        []string{"vip"},
		Email:       "user5@example.com",
		Status:      true,
		Enabled:     true,
	}

	fact6 := &UserFact{
		Temperature: 33,
		Role:        "normal",
		Tags:        []string{"normal"},
		Email:       "user6@example.com",
		Status:      true,
		Enabled:     true,
	}

	ctx := context.Background()
	// 评估 Fact
	fmt.Println("Actions for fact1:")
	actions1 := engine.EvaluateFact(ctx, fact1)
	for _, action := range actions1 {
		fmt.Printf("- Action Type: %s, Parameters: %+v\n", action.ActionType, action.Parameters)
		fact1.ExecuteAction(ctx, string(action.ActionType), action.Parameters) // 需要将 ActionType 转换为 string
	}

	fmt.Println("\nActions for fact2:")
	actions2 := engine.EvaluateFact(ctx, fact2)
	for _, action := range actions2 {
		fmt.Printf("- Action Type: %s, Parameters: %+v\n", action.ActionType, action.Parameters)
		fact2.ExecuteAction(ctx, string(action.ActionType), action.Parameters)
	}

	fmt.Println("\nActions for fact3:")
	actions3 := engine.EvaluateFact(ctx, fact3)
	for _, action := range actions3 {
		fmt.Printf("- Action Type: %s, Parameters: %+v\n", action.ActionType, action.Parameters)
		fact3.ExecuteAction(ctx, string(action.ActionType), action.Parameters)
	}

	fmt.Println("\nActions for fact4:")
	actions4 := engine.EvaluateFact(ctx, fact4)
	for _, action := range actions4 {
		fmt.Printf("- Action Type: %s, Parameters: %+v\n", action.ActionType, action.Parameters)
		fact4.ExecuteAction(ctx, string(action.ActionType), action.Parameters)
	}
	fmt.Println("\nActions for fact5:")
	actions5 := engine.EvaluateFact(ctx, fact5)
	for _, action := range actions5 {
		fmt.Printf("- Action Type: %s, Parameters: %+v\n", action.ActionType, action.Parameters)
		fact5.ExecuteAction(ctx, string(action.ActionType), action.Parameters)
	}
	fmt.Println("\nActions for fact6:")
	actions6 := engine.EvaluateFact(ctx, fact6)
	for _, action := range actions6 {
		fmt.Printf("- Action Type: %s, Parameters: %+v\n", action.ActionType, action.Parameters)
		fact6.ExecuteAction(ctx, string(action.ActionType), action.Parameters)
	}

	fmt.Println("\nFact1 Role after actions:", fact1.Role) // 演示 Action 修改 Fact 的效果
	fmt.Println("Fact2 Role after actions:", fact2.Role)
	fmt.Println("Fact3 Role after actions:", fact3.Role)
	fmt.Println("Fact4 Role after actions:", fact4.Role)
	fmt.Println("Fact5 Role after actions:", fact5.Role)
	fmt.Println("Fact6 Role after actions:", fact6.Role)

	fmt.Println("\nFact6 Temperature after actions:", fact6.GetValue(ctx, string(rule_engine.FieldNameTemperature))) // 演示 GetValue 取值
	fmt.Println("Fact6 Enabled after actions:", fact6.GetValue(ctx, string(rule_engine.FieldNameEnabled)))
	fmt.Println("Fact6 Status after actions:", fact6.GetValue(ctx, string(rule_engine.FieldNameStatus)))

	now := time.Now().Format(time.RFC3339)
	fmt.Println("\nCurrent Time:", now)
}
