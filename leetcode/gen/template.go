package leetcode

const CheckInSubject = `package main

import (
	"fmt"
	"reflect"
)

/*
{{.Url}}

{{.Desc}}
*/
func main() {
	var tests = []struct {
		{{- range .AnsParams}}
		{{.Name}} {{.Type}}
		{{- end}}
		want {{.AnsReturnType}}
	}{
		{{- if not .SubjectTests}}
		{},
		{{- end}}
		{{- range .SubjectTests}}
		{
			{{- range .Params}}
			{{ .Name }}: {{ .Value }},
			{{- end}}
		},
		{{- end}}
	}

	for _, item := range tests {
		if ans := {{.AnsFuncName}}({{range $index, $item := .AnsParams}}{{if $index}}, {{end}}item.{{$item.Name}}{{end}}); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

{{.Ans}}
`
