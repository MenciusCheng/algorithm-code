package rule_engine

// OperatorType 定义操作符枚举
type OperatorType string

const (
	OperatorEqual        OperatorType = "equal"
	OperatorNotEqual     OperatorType = "not_equal"
	OperatorGreaterEqual OperatorType = "greater_equal"
	OperatorGreaterThan  OperatorType = "greater_than"
	OperatorLessThan     OperatorType = "less_than"
	OperatorLessEqual    OperatorType = "less_equal"
	OperatorIn           OperatorType = "in"
	OperatorNotIn        OperatorType = "not_in"
	OperatorTrue         OperatorType = "true"  // 新增: 真值判断
	OperatorFalse        OperatorType = "false" // 新增: 假值判断
)

// FieldName 定义字段名枚举
type FieldName string

const (
	FieldNameTemperature FieldName = "temperature"
	FieldNameRole        FieldName = "role"
	FieldNameTags        FieldName = "tags"
	FieldNameEmail       FieldName = "email"
	FieldNameStatus      FieldName = "status"  // 示例字段名
	FieldNameEnabled     FieldName = "enabled" // 示例字段名
)

// ActionType 定义 ActionType 枚举
type ActionType string

const (
	ActionTypePrint        ActionType = "print"
	ActionTypeModifyFact   ActionType = "modify_fact"
	ActionTypeCallFunction ActionType = "call_function"
	ActionTypeSendEmail    ActionType = "send_email"  // 示例 ActionType
	ActionTypeModifyRole   ActionType = "modify_role" // 示例 ActionType
)

// ActionParameter 定义 Action Parameters 的字段枚举
type ActionParameter string

const (
	ActionParameterMessage   ActionParameter = "message"
	ActionParameterFieldName ActionParameter = "fieldName"
	ActionParameterNewValue  ActionParameter = "newValue"
	ActionParameterRecipient ActionParameter = "recipient"
)
