package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

type Apple struct {
	Age int `json:"age,omitempty"`
	H   int `json:"h,omitempty"`
	D   Dog `json:"d,omitempty"`
}

type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RP struct {
	Name  string `json:"name"`
	BName string `json:"name"`
	Age   int    `json:"age"`
}

func main() {
	structToJsonToMap()
}

func testRP() {
	//rp := RP{
	//	Name: "aa",
	//	Age:  10,
	//}
	//bs, _ := json.Marshal(rp)
	//fmt.Printf("bs: %s\n", bs)

	bs := []byte(" {\"name\":\"heihei\",\"age\":10}")

	rp2 := RP{}
	_ = json.Unmarshal(bs, &rp2)
	fmt.Printf("rp2: %+v\n", rp2)
}

func testApple() {
	a := Apple{
		Age: 2,
	}
	bs, _ := json.Marshal(a)
	fmt.Printf("%s\n", bs)
}

func mapToJsonToStruct() {
	param := make(map[string]interface{})
	param["name"] = "Good Dog"
	param["age"] = 12

	bytes, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	dog := Dog{}
	err = json.Unmarshal(bytes, &dog)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dog: %+v\n", dog)
}

func structToJsonToMap() {
	dog := Dog{
		Name: "Good Dog",
		Age:  12,
	}

	bytes, err := json.Marshal(dog)
	if err != nil {
		panic(err)
	}

	param := make(map[string]interface{})
	err = json.Unmarshal(bytes, &param)
	if err != nil {
		panic(err)
	}
	fmt.Printf("param: %+v\n", param)
}

func UnmarshalPromotionV2Data(extension string) error {
	data := PromotionV2Data{}
	err := json.Unmarshal([]byte(extension), &data)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", data)
	return nil
}

func Unmarshal2PromotionV2Data(extension string) error {
	data := PromotionV2Data{}
	err := json2.Unmarshal([]byte(extension), &data)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", data)
	return nil
}

func Unmarshal2PromotionV2DataXX(extension string) error {
	data := PromotionV2Data{}
	err := json2.Unmarshal([]byte(extension), &data)
	if err != nil {
		return err
	}
	fmt.Printf("data: %+v\n", data)

	thirdData := &data
	extensionByte, err := json2.Marshal(thirdData)
	if err != nil {
		return err
	}
	fmt.Printf("extensionByte: %s\n", extensionByte)

	return nil
}

// 把 Map 序列化为 JSON，再反序列化为 Map ，再序列化为 JSON
// 测试序列化与反序列化过程中，哪些字段类型会被改变
func MapToJsonToMapToJson() {
	// https://developer.mozilla.org/zh-CN/docs/Glossary/JSON
	param := make(map[string]interface{})
	param["name"] = "cat"
	param["age"] = 12
	param["length"] = 15.78
	param["null"] = nil
	param["bool"] = true
	param["sub"] = map[string]interface{}{
		"subname": "wei",
		"age":     15,
	}
	param["ints"] = []int{1, 2}
	param["sp"] = "<div>cat&dog</div>" //  replaces "<", ">", "&", U+2028, and U+2029 are escaped to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029".

	bytes, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	unmarshalParam := make(map[string]interface{})
	err = json.Unmarshal(bytes, &unmarshalParam)
	if err != nil {
		panic(err)
	}

	unmarshalBytes, err := json.Marshal(unmarshalParam)
	if err != nil {
		panic(err)
	}

	fmt.Printf("param: %+v, age type: %s\n", param, reflect.TypeOf(param["age"]).String())
	fmt.Printf("bytes: %s\n", string(bytes))
	// json 反序列化后，整型变成了浮点数
	fmt.Printf("unmarshalParam: %+v, age type: %s\n", unmarshalParam, reflect.TypeOf(unmarshalParam["age"]).String())
	fmt.Printf("unmarshalBytes: %s\n", string(unmarshalBytes))
}

func printJson() {
	a := JsonHuntTaskConf{
		Condition: &JsonHuntTaskCondition{
			GiftIds: []int32{},
		},
		Levels: []*JsonHuntTaskLevelData{
			{
				TaskId:          1,
				TaskName:        "任务1：抓住马杀鸡",
				MinGold:         20000,
				UpgradeRelateId: 101,
				UpgradeNoticeId: 102,
				TargetEffectId:  0,
				TargetPools: []JsonHuntTaskPool{
					{
						PoolId:   1,
						NoticeId: 103,
					},
				},
			},
			{
				TaskId:          2,
				TaskName:        "任务2：抓住三黄鸡",
				MinGold:         50000,
				UpgradeRelateId: 201,
				UpgradeNoticeId: 202,
				TargetEffectId:  0,
				TargetPools: []JsonHuntTaskPool{
					{
						PoolId:   1,
						NoticeId: 203,
					},
				},
			},
		},
		Broadcast: &JsonHuntTaskBroadcast{
			AwardIds:       []int64{1, 2},
			AwardNotice:    "",
			UpgradeTaskIds: []int64{5},
			UpgradeNotice:  "",
		},
		RankConfId: 0,
	}

	v, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(v))
}

type JsonHuntTaskConf struct {
	Condition  *JsonHuntTaskCondition   `json:"condition"`    // 玩法完成条件
	Levels     []*JsonHuntTaskLevelData `json:"levels"`       // 任务等级配置
	Broadcast  *JsonHuntTaskBroadcast   `json:"broadcast"`    // 中奖广播条件
	RankConfId int32                    `json:"rank_conf_id"` // 捕猎榜单配置ID
}

type JsonHuntTaskCondition struct {
	GiftIds []int32 `json:"gift_ids"` // 送礼礼物ID列表
}

type JsonHuntTaskLevelData struct {
	TaskId          int64              `json:"task_id"`           // 任务ID
	TaskName        string             `json:"task_name"`         // 任务名称
	MinGold         int64              `json:"min_gold"`          // 升级解锁金币
	UpgradeRelateId int64              `json:"upgrade_relate_id"` // 升级奖励关联id
	UpgradeNoticeId int64              `json:"upgrade_notice_id"` // 升级通知关联id
	TargetEffectId  int64              `json:"target_effect_id"`  // 捕猎目标特效id
	TargetPools     []JsonHuntTaskPool `json:"target_pools"`      // 捕猎奖池id
}

type JsonHuntTaskPool struct {
	PoolId   int64 `json:"pool_id"`   // 捕猎奖池id
	NoticeId int64 `json:"notice_id"` // 捕猎通知id
}

type JsonHuntTaskBroadcast struct {
	AwardIds       []int64 `json:"award_ids"`        // 新奖励ID列表
	AwardNotice    string  `json:"award_notice"`     // 新奖励广播文本
	UpgradeTaskIds []int64 `json:"upgrade_task_ids"` // 升级任务ID列表
	UpgradeNotice  string  `json:"upgrade_notice"`   // 升级广播文本
}

// 将结构体转换为 JSON 字符串，并为每个字段填充默认值
func StructToJSONWithDefaults(data interface{}) (string, error) {
	defaultStruct := populateDefaults(reflect.TypeOf(data))
	jsonBytes, err := json.MarshalIndent(defaultStruct, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// 根据类型生成默认值
func populateDefaults(t reflect.Type) interface{} {
	if t.Kind() == reflect.Ptr {
		// 如果是指针类型，递归获取指针所指向的默认值
		value := populateDefaults(t.Elem())
		ptr := reflect.New(t.Elem())
		if value != nil {
			ptr.Elem().Set(reflect.ValueOf(value))
		}
		return ptr.Interface()
	}

	switch t.Kind() {
	case reflect.Struct:
		// 创建结构体的每个字段的默认值
		result := reflect.New(t).Elem()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" { // 忽略未导出的字段
				continue
			}
			defaultValue := populateDefaults(field.Type)
			if defaultValue != nil {
				result.Field(i).Set(reflect.ValueOf(defaultValue))
			}
		}
		return result.Interface()
	case reflect.Slice:
		// 初始化一个包含单个默认值的切片
		elemType := t.Elem()
		slice := reflect.MakeSlice(t, 1, 1)
		defaultValue := populateDefaults(elemType)
		if defaultValue != nil {
			slice.Index(0).Set(reflect.ValueOf(defaultValue))
		}
		return slice.Interface()
	case reflect.Array:
		// 初始化一个包含单个默认值的数组
		elemType := t.Elem()
		array := reflect.New(t).Elem()
		if t.Len() > 0 {
			defaultValue := populateDefaults(elemType)
			if defaultValue != nil {
				array.Index(0).Set(reflect.ValueOf(defaultValue))
			}
		}
		return array.Interface()
	case reflect.Map:
		// 返回空 Map
		return reflect.MakeMap(t).Interface()
	case reflect.String:
		return ""
	case reflect.Bool:
		return false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.Zero(t).Interface()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.Zero(t).Interface()
	case reflect.Float32, reflect.Float64:
		return reflect.Zero(t).Interface()
	default:
		// 默认返回 nil
		return nil
	}
}

type JsonYearFireworkCondition struct {
	GiftIds []int32 `json:"gift_ids"` // 送礼礼物ID列表
	MinGold int64   `json:"min_gold"` // 燃放烟花金币
}

type JsonYearFireworkRank struct {
	RankConfId   int32 `json:"rank_conf_id"`   // 榜单配置ID
	RankPoint    int64 `json:"rank_point"`     // 榜单心动值
	RankNoticeId int64 `json:"rank_notice_id"` // 榜单心动值通知关联id
}

type JsonBirthdayWeekConf struct {
	WeekStartNoticeId   int32 `json:"week_start_notice_id"`   // 生日周开始通知id
	WeekStartNoticeHour int   `json:"week_start_notice_hour"` // 生日周开始通知小时
	WeekEndNoticeId     int32 `json:"week_end_notice_id"`     // 生日周结束通知id
	WeekEndNoticeHour   int   `json:"week_end_notice_hour"`   // 生日周结束通知小时
	WeekRoomNoticeId    int32 `json:"week_room_notice_id"`    // 生日周首次进房通知id
}

// 将结构体定义的字符串解析为 JSON 默认值
func StructDefinitionToJSON(structDef string) (string, error) {
	// 解析字符串为 AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", "package main\n"+structDef, parser.AllErrors)
	if err != nil {
		return "", fmt.Errorf("failed to parse struct definition: %w", err)
	}

	// 查找第一个结构体类型
	var structType *ast.StructType
	ast.Inspect(node, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if ok {
			if st, ok := ts.Type.(*ast.StructType); ok {
				structType = st
				return false
			}
		}
		return true
	})

	if structType == nil {
		return "", fmt.Errorf("no struct type found in definition")
	}

	// 构造默认值
	defaultStruct := constructDefaultStruct(structType, node)

	// 转换为 JSON
	jsonBytes, err := json.MarshalIndent(defaultStruct, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return string(jsonBytes), nil
}

// 根据 AST 构造结构体的默认值
func constructDefaultStruct(st *ast.StructType, node *ast.File) map[string]interface{} {
	result := make(map[string]interface{})
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			continue // 忽略匿名字段
		}

		// 获取 JSON 标签名称
		jsonName := getJSONTag(field)
		if jsonName == "" {
			continue // 如果没有 JSON 标签，跳过该字段
		}

		defaultValue := getDefaultValueForField(field.Type, node)
		result[jsonName] = defaultValue
	}
	return result
}

// 获取字段的 JSON 标签名称
func getJSONTag(field *ast.Field) string {
	if field.Tag == nil {
		return ""
	}

	tagValue := strings.Trim(field.Tag.Value, "`")
	for _, tag := range strings.Split(tagValue, " ") {
		if strings.HasPrefix(tag, "json:") {
			parts := strings.SplitN(tag, ":", 2)
			if len(parts) == 2 {
				jsonTag := strings.Trim(parts[1], "\"")
				return strings.Split(jsonTag, ",")[0] // 只取标签的第一个部分
			}
		}
	}
	return ""
}

// 获取字段的默认值
func getDefaultValueForField(expr ast.Expr, node *ast.File) interface{} {
	switch ft := expr.(type) {
	case *ast.Ident:
		// 基础类型
		return getDefaultValueForType(ft.Name, node)
	case *ast.StarExpr:
		// 指针类型，递归处理指针指向的类型
		return getDefaultValueForField(ft.X, node)
	case *ast.ArrayType:
		// 数组类型，初始化一个包含单个默认值的切片
		return []interface{}{getDefaultValueForField(ft.Elt, node)}
	case *ast.StructType:
		// 内联结构体
		return constructDefaultStruct(ft, node)
	default:
		// 未处理的类型
		return nil
	}
}

// 获取基础类型的默认值
func getDefaultValueForType(typeName string, node *ast.File) interface{} {
	switch typeName {
	case "int", "int32", "int64":
		return 0
	case "float32", "float64":
		return 0.0
	case "string":
		return ""
	case "bool":
		return false
	default:
		// 如果是其他复杂类型，尝试匹配结构体定义
		for _, decl := range node.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if typeSpec.Name.Name == typeName {
					if st, ok := typeSpec.Type.(*ast.StructType); ok {
						return constructDefaultStruct(st, node)
					}
				}
			}
		}
	}
	return nil
}
