package leetcode

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

func Gen(desc, url, cal, month string) error {
	subject, err := NewSubject(desc, url, cal)
	if err != nil {
		panic(err)
	}

	question := fmt.Sprintf("q%d", subject.QuestionNum)

	tmpl, err := template.New("CheckInSubject").Parse(CheckInSubject)
	if err != nil {
		panic(err)
	}

	fileName := "main.go"
	directory := fmt.Sprintf("/Users/chengmengwei/goProject/algorithm-code/leetcode/%s/%s", month, question)
	filePath := fmt.Sprintf("%s%s%s", directory, string(os.PathSeparator), fileName)

	if err := BuildDir(filePath); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, subject)
	if err != nil {
		return err
	}

	cmd, err := exec.Command("gofmt", "-l", "-w", filePath).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(cmd))

	return nil
}

// BuildDir 创建目录
func BuildDir(absDir string) error {
	return os.MkdirAll(path.Dir(absDir), os.ModePerm) //生成多级目录
}

func ArrStr(str string) string {
	s2 := strings.ReplaceAll(str, "[", "{")
	s2 = strings.ReplaceAll(s2, "]", "}")
	return s2
}
