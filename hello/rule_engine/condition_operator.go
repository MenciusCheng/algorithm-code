package rule_engine

import (
	"reflect"
	"strconv"
	"time"
)

// EvaluateCondition 执行单个条件评估
func EvaluateCondition(condition *Condition, factValue interface{}) bool {
	operator := condition.Operator
	conditionValue := condition.FieldValue

	switch operator {
	case OperatorEqual:
		return reflect.DeepEqual(factValue, conditionValue)
	case OperatorNotEqual:
		return !reflect.DeepEqual(factValue, conditionValue)
	case OperatorGreaterEqual:
		return compareValues(factValue, conditionValue, ">=")
	case OperatorGreaterThan:
		return compareValues(factValue, conditionValue, ">")
	case OperatorLessThan:
		return compareValues(factValue, conditionValue, "<")
	case OperatorLessEqual:
		return compareValues(factValue, conditionValue, "<=")
	case OperatorIn:
		return containsValue(factValue, conditionValue)
	case OperatorNotIn:
		return !containsValue(factValue, conditionValue)
	case OperatorTrue:
		return evaluateOperatorTrue(factValue)
	case OperatorFalse:
		return evaluateOperatorFalse(factValue)
	default:
		return false // 不支持的操作符
	}
}

// compareValues 比较两个值的大小 (支持数值、字符串和日期比较)
func compareValues(factValue interface{}, conditionValue interface{}, operator string) bool {
	factValueFloat, okFactFloat := toFloat64(factValue)
	conditionValueFloat, okConditionFloat := toFloat64(conditionValue)

	if okFactFloat && okConditionFloat { // 数值比较
		switch operator {
		case ">=":
			return factValueFloat >= conditionValueFloat
		case ">":
			return factValueFloat > conditionValueFloat
		case "<":
			return factValueFloat < conditionValueFloat
		case "<=":
			return factValueFloat <= conditionValueFloat
		}
	}

	factValueString, okFactString := factValue.(string)
	conditionValueString, okConditionString := conditionValue.(string)
	if okFactString && okConditionString { // 字符串比较 (字典序)
		// 尝试解析为日期进行日期比较 (ISO 8601 格式)
		factTime, errFactTime := time.Parse(time.RFC3339, factValueString)
		conditionTime, errConditionTime := time.Parse(time.RFC3339, conditionValueString)
		if errFactTime == nil && errConditionTime == nil { // 日期比较
			switch operator {
			case ">=":
				return factTime.After(conditionTime) || factTime.Equal(conditionTime)
			case ">":
				return factTime.After(conditionTime)
			case "<":
				return factTime.Before(conditionTime)
			case "<=":
				return factTime.Before(conditionTime) || factTime.Equal(conditionTime)
			}
		} else { // 字符串字典序比较
			switch operator {
			case ">=":
				return factValueString >= conditionValueString
			case ">":
				return factValueString > conditionValueString
			case "<":
				return factValueString < conditionValueString
			case "<=":
				return factValueString <= conditionValueString
			}
		}
	}

	return false // 类型不匹配或不支持的比较
}

// containsValue 判断 factValue 是否在 conditionValue (数组) 中
func containsValue(factValue interface{}, conditionValue interface{}) bool {
	conditionValueSlice, ok := conditionValue.([]interface{}) // 假设 FieldValue 是 []interface{} 类型
	if !ok {
		return false // conditionValue 不是数组
	}
	for _, item := range conditionValueSlice {
		if reflect.DeepEqual(factValue, item) {
			return true
		}
	}
	return false
}

// toFloat64 尝试将 interface{} 转换为 float64
func toFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return v, true
	case string:
		floatVal, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return floatVal, true
		}
	default:
		return 0, false
	}
	return 0, false
}

// evaluateOperatorTrue 评估 OperatorType 为 true 的条件
func evaluateOperatorTrue(factValue interface{}) bool {
	boolValue, ok := factValue.(bool)
	if ok {
		return boolValue
	}
	return false // 默认 false，如果 Fact 值不是 bool 类型
}

// evaluateOperatorFalse 评估 OperatorType 为 false 的条件
func evaluateOperatorFalse(factValue interface{}) bool {
	return !evaluateOperatorTrue(factValue) // OperatorFalse 是 OperatorTrue 的反向逻辑
}
