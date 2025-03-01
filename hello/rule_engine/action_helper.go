package rule_engine

import "fmt"

// GetStringParam 从参数 map 中获取字符串类型的参数值
func GetStringParam(params map[string]interface{}, key ActionParameter) (string, error) {
	value, ok := params[string(key)].(string)
	if !ok {
		return "", fmt.Errorf("parameter '%s' is not a string", key)
	}
	return value, nil
}

// GetIntParam 从参数 map 中获取 int 类型的参数值
func GetIntParam(params map[string]interface{}, key ActionParameter) (int, error) {
	value, ok := params[string(key)].(int)
	if !ok {
		return 0, fmt.Errorf("parameter '%s' is not an integer", key)
	}
	return value, nil
}

// GetFloat64Param 从参数 map 中获取 float64 类型的参数值
func GetFloat64Param(params map[string]interface{}, key ActionParameter) (float64, error) {
	value, ok := params[string(key)].(float64)
	if !ok {
		return 0.0, fmt.Errorf("parameter '%s' is not a float64", key)
	}
	return value, nil
}

// GetBoolParam 从参数 map 中获取 bool 类型的参数值
func GetBoolParam(params map[string]interface{}, key ActionParameter) (bool, error) {
	value, ok := params[string(key)].(bool)
	if !ok {
		return false, fmt.Errorf("parameter '%s' is not a boolean", key)
	}
	return value, nil
}

// GetStringSliceParam 从参数 map 中获取 string slice 类型的参数值
func GetStringSliceParam(params map[string]interface{}, key ActionParameter) ([]string, error) {
	value, ok := params[string(key)].([]string) // 直接断言为 []string
	if !ok {
		interfaceSlice, okInterfaceSlice := params[string(key)].([]interface{}) // 尝试断言为 []interface{}
		if !okInterfaceSlice {
			return nil, fmt.Errorf("parameter '%s' is not a string slice", key)
		}
		stringSlice := make([]string, len(interfaceSlice))
		for i, v := range interfaceSlice {
			strVal, okStr := v.(string)
			if !okStr {
				return nil, fmt.Errorf("element in parameter '%s' is not a string", key)
			}
			stringSlice[i] = strVal
		}
		return stringSlice, nil
	}
	return value, nil
}

// GetInterfaceParam 从参数 map 中获取 interface{} 类型的参数值 (如果需要原始 interface{} 值)
func GetInterfaceParam(params map[string]interface{}, key ActionParameter) (interface{}, error) {
	value, ok := params[string(key)]
	if !ok {
		return nil, fmt.Errorf("parameter '%s' not found", key)
	}
	return value, nil
}
