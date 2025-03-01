package main

import (
	"context"
	"fmt"
	"github.com/MenciusCheng/algorithm-code/hello/rule_engine"
	"strings"
)

// 示例 Fact 实现
type UserFact struct {
	Temperature float64
	Role        string
	Tags        []string
	Email       string
	Status      bool // 示例字段
	Enabled     bool // 示例字段
}

// 实现了 Fact 接口的 GetValue 方法
func (f *UserFact) GetValue(ctx context.Context, fieldName string) interface{} {
	switch rule_engine.FieldName(fieldName) { // 使用 FieldName 枚举
	case rule_engine.FieldNameTemperature:
		return f.Temperature
	case rule_engine.FieldNameRole:
		return f.Role
	case rule_engine.FieldNameTags:
		return f.Tags
	case rule_engine.FieldNameEmail:
		return f.Email
	case rule_engine.FieldNameStatus:
		return f.Status
	case rule_engine.FieldNameEnabled:
		return f.Enabled
	default:
		return nil
	}
}

// 实现了 Fact 接口的 ExecuteAction 方法
func (f *UserFact) ExecuteAction(ctx context.Context, actionTypeStr string, parameters map[string]interface{}) {
	actionType := rule_engine.ActionType(actionTypeStr) // 类型转换

	switch actionType {
	case rule_engine.ActionTypePrint:
		message, err := rule_engine.GetStringParam(parameters, rule_engine.ActionParameterMessage) // 使用 GetStringParam 获取 message
		if err != nil {
			fmt.Printf("ExecuteAction: %v\n", err)
			return
		}
		fmt.Println(message)

	case rule_engine.ActionTypeModifyFact:
		fieldNameStr, errField := rule_engine.GetStringParam(parameters, rule_engine.ActionParameterFieldName) // 获取 fieldName
		newValue, errValue := rule_engine.GetInterfaceParam(parameters, rule_engine.ActionParameterNewValue)   // 获取 newValue, 保留 interface{} 类型
		if errField != nil {
			fmt.Printf("ExecuteAction: %v\n", errField)
			return
		}
		if errValue != nil {
			fmt.Printf("ExecuteAction: %v\n", errValue)
			return
		}
		fieldName := rule_engine.FieldName(fieldNameStr) // 转换为 FieldName 类型
		f.ModifyFieldValue(string(fieldName), newValue)

	case rule_engine.ActionTypeSendEmail:
		recipientTemplate, errRecipient := rule_engine.GetStringParam(parameters, rule_engine.ActionParameterRecipient)
		messageTemplate, errMessage := rule_engine.GetStringParam(parameters, rule_engine.ActionParameterMessage)
		if errRecipient != nil {
			fmt.Printf("ExecuteAction: %v\n", errRecipient)
			return
		}
		if errMessage != nil {
			fmt.Printf("ExecuteAction: %v\n", errMessage)
			return
		}

		recipient := strings.ReplaceAll(recipientTemplate, "{{.fact.email}}", f.Email)
		message := strings.ReplaceAll(messageTemplate, "{{.fact.email}}", f.Email)
		fmt.Printf("Sending email to: %s, with message: %s\n", recipient, message)

	case rule_engine.ActionTypeModifyRole:
		newRole, errRole := rule_engine.GetStringParam(parameters, rule_engine.ActionParameterNewValue) // 获取 newRole
		if errRole != nil {
			fmt.Printf("ExecuteAction: %v\n", errRole)
			return
		}
		f.ModifyFieldValue(string(rule_engine.FieldNameRole), newRole) // 修改 Role 字段

	case rule_engine.ActionTypeCallFunction:
		fmt.Println("ActionTypeCallFunction is triggered but not implemented in example.") // 示例，可以根据需要实现

	default:
		fmt.Printf("Unknown ActionType: %s\n", actionType)
	}
}

// ModifyFieldValue 示例：修改 Fact 字段值的通用方法 (需要根据实际 Fact 结构进行调整)
func (f *UserFact) ModifyFieldValue(fieldName string, newValue interface{}) {
	switch rule_engine.FieldName(fieldName) { // 使用 FieldName 枚举
	case rule_engine.FieldNameRole:
		if roleStr, ok := newValue.(string); ok {
			f.Role = roleStr
		}
	case rule_engine.FieldNameStatus:
		if statusBool, ok := newValue.(bool); ok {
			f.Status = statusBool
		}
	case rule_engine.FieldNameEnabled:
		if enabledBool, ok := newValue.(bool); ok {
			f.Enabled = enabledBool
		}
	// ... 可以根据需要添加更多字段的修改逻辑
	default:
		fmt.Println("Unknown field name for modification:", fieldName)
	}
}

// CallFunction 示例：调用自定义函数 (需要根据实际需求实现)
func (f *UserFact) CallFunction(functionName string, arguments map[string]interface{}) {
	switch functionName {
	case "exampleFunction":
		// ... 执行 exampleFunction 的逻辑，可以使用 arguments 中的参数
		fmt.Println("Calling exampleFunction with arguments:", arguments)
	default:
		fmt.Println("Unknown function name:", functionName)
	}
}
