package leetcode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Subject struct {
	// 原始内容
	Desc string // 题目描述
	Url  string // 题目链接
	Ans  string // 解答函数

	// 解析内容
	QuestionNum   int            // 题目编号
	AnsFuncName   string         // 解答函数名
	AnsParams     []SubjectParam // 解答函数参数列表
	AnsReturnType string         // 解答函数返回类型
	SubjectTests  []SubjectTest  // 示例列表
}

type SubjectParam struct {
	Name string
	Type string
}

type SubjectTest struct {
	Params []SubjectTestParam // 示例参数列表
}

type SubjectTestParam struct {
	Name  string
	Value string
}

func NewSubject(desc string, url string, ans string) (*Subject, error) {
	subject := &Subject{
		Desc: strings.TrimSpace(desc),
		Url:  strings.TrimSpace(url),
		Ans:  strings.TrimSpace(ans),
	}

	// 解析参数
	if err := subject.parseAns(); err != nil {
		return nil, err
	}

	// 解析题目编号
	questionNumReg := regexp.MustCompile(`^\s*(\d+)`)
	questionNumSubmatch := questionNumReg.FindStringSubmatch(subject.Desc)
	if len(questionNumSubmatch) != 2 {
		return nil, fmt.Errorf("questionNumSubmatch length is not equal to 2: %+v", questionNumSubmatch)
	}
	subject.QuestionNum, _ = strconv.Atoi(questionNumSubmatch[1])

	// 解析示例，失败不影响生成文件
	if err := subject.parseTest(); err != nil {
		fmt.Printf("subject.parseTest err: %v\n", err)
	}

	return subject, nil
}

// 解析参数
func (s *Subject) parseAns() error {
	ansArr := strings.Split(strings.TrimSpace(s.Ans), "\n")

	if len(ansArr) == 0 {
		return fmt.Errorf("ans is empty")
	}

	reg := regexp.MustCompile(`func\s+(\w+)\s*\(([A-Za-z0-9\[\], *]+)\)\s*([A-Za-z0-9\[\]*]*)`)
	submatch := reg.FindStringSubmatch(ansArr[0])
	if len(submatch) != 4 {
		return fmt.Errorf("submatch length is not equal to 4: %+v", submatch)
	}

	var paramStr string
	s.AnsFuncName, paramStr, s.AnsReturnType = submatch[1], submatch[2], submatch[3]

	paramReg := regexp.MustCompile(`(\w+)\s+([A-Za-z0-9\[\]*]+)`)
	for _, item := range strings.Split(paramStr, ",") {
		paramMatchs := paramReg.FindStringSubmatch(item)
		if len(paramMatchs) != 3 {
			return fmt.Errorf("paramMatchs length is not equal to 3: %+v", paramMatchs)
		}

		s.AnsParams = append(s.AnsParams, SubjectParam{
			Name: paramMatchs[1],
			Type: paramMatchs[2],
		})
	}

	if len(ansArr) >= 3 && strings.TrimSpace(ansArr[len(ansArr)-1]) == "}" {
		var body string
		switch s.AnsReturnType {
		case "int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8":
			body = "return 0"
		case "string":
			body = "return \"\""
		case "[]int":
			body = "return []int{}"
		case "[][]int":
			body = "return [][]int{}"
		case "[]string":
			body = "return []string{}"
		case "bool":
			body = "return false"
		}
		if body != "" {
			ansArr[len(ansArr)-2] = body
			s.Ans = strings.Join(ansArr, "\n")
		}
	}

	return nil
}

// 解析示例
func (s *Subject) parseTest() error {
	descArr := strings.Split(strings.TrimSpace(s.Desc), "\n")
	if len(descArr) == 0 {
		return fmt.Errorf("desc is empty")
	}

	testDemos := make([]SubjectTest, 0)
	testParams := make([]SubjectTestParam, 0)

	for _, descLine := range descArr {
		line := strings.TrimSpace(descLine)
		if strings.HasPrefix(line, "输入") {
			// 解析每个参数的示例
			for _, param := range s.AnsParams {
				subjectTestParam, err := CalInputSubjectTestParam(param, line)
				if err != nil {
					return err
				}
				testParams = append(testParams, subjectTestParam)
			}
		} else if strings.HasPrefix(line, "输出") {
			// 解析输出参数示例
			subjectTestParam, err := CalOutputSubjectTestParam(s.AnsReturnType, line)
			if err != nil {
				return err
			}
			testParams = append(testParams, subjectTestParam)

			// 添加到示例
			testDemos = append(testDemos, SubjectTest{
				Params: testParams,
			})
			testParams = make([]SubjectTestParam, 0)
		}
	}

	s.SubjectTests = testDemos

	return nil
}

// 解析输入参数用例
func CalInputSubjectTestParam(param SubjectParam, line string) (SubjectTestParam, error) {
	res := SubjectTestParam{}

	paramRegStr := ""
	switch param.Type {
	case "int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8":
		paramRegStr = fmt.Sprintf(`%s\s*=\s*([0-9\-]+)`, param.Name)
	case "string":
		paramRegStr = fmt.Sprintf(`%s\s*=\s*("[\w\[\], \-]*")`, param.Name)
	case "[]int", "[][]int", "[]string":
		paramRegStr = fmt.Sprintf(`%s\s*=\s*\[([\w\[\], \-]*)\]`, param.Name)
	case "bool":
		paramRegStr = fmt.Sprintf(`%s\s*=\s*(true|false)`, param.Name)
	}
	paramReg := regexp.MustCompile(paramRegStr)
	paramMatchs := paramReg.FindStringSubmatch(line)
	if len(paramMatchs) != 2 {
		return res, fmt.Errorf("paramMatchs length is not equal to 2: %+v, line: %s", paramMatchs, line)
	}

	res.Name = param.Name
	switch param.Type {
	case "[]int":
		res.Value = fmt.Sprintf("[]int{%s}", paramMatchs[1])
	case "[][]int":
		value := paramMatchs[1]
		value = strings.ReplaceAll(value, "[", "{")
		value = strings.ReplaceAll(value, "]", "}")
		res.Value = fmt.Sprintf("[][]int{%s}", value)
	case "[]string":
		value := paramMatchs[1]
		value = strings.ReplaceAll(value, "[", "{")
		value = strings.ReplaceAll(value, "]", "}")
		res.Value = fmt.Sprintf("[][]string{%s}", value)
	default:
		res.Value = paramMatchs[1]
	}

	return res, nil
}

// 解析输出参数用例
func CalOutputSubjectTestParam(ansReturnType string, line string) (SubjectTestParam, error) {
	res := SubjectTestParam{}

	var paramReg *regexp.Regexp
	switch ansReturnType {
	case "int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8":
		paramReg = regexp.MustCompile(`([0-9\-]+)`)
	case "string":
		paramReg = regexp.MustCompile(`("[\w\[\], \-]*")`)
	case "[]int", "[][]int", "[]string":
		paramReg = regexp.MustCompile(`\[([\w\[\], \-]*)\]`)
	case "bool":
		paramReg = regexp.MustCompile(`(true|false)`)
	}
	paramMatchs := paramReg.FindStringSubmatch(line)
	if len(paramMatchs) != 2 {
		return res, fmt.Errorf("paramMatchs length is not equal to 2: %+v, line: %s", paramMatchs, line)
	}

	res.Name = "want"
	res.Value = paramMatchs[1]
	return res, nil
}
