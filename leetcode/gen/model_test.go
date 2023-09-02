package leetcode

import (
	"reflect"
	"testing"
)

func TestCalInputSubjectTestParam(t *testing.T) {
	type args struct {
		param SubjectParam
		line  string
	}
	tests := []struct {
		name    string
		args    args
		want    SubjectTestParam
		wantErr bool
	}{
		{
			args: args{
				param: SubjectParam{
					Name: "forts",
					Type: "[]int",
				},
				line: "输入：forts = [1,0,0,-1,0,0,0,0,1]",
			},
			want: SubjectTestParam{
				Name:  "forts",
				Value: "[]int{1,0,0,-1,0,0,0,0,1}",
			},
		},
		{
			args: args{
				param: SubjectParam{
					Name: "target",
					Type: "int",
				},
				line: "输入：nums = [-1,2,1,-4], target = -1",
			},
			want: SubjectTestParam{
				Name:  "target",
				Value: "-1",
			},
		},
		{
			args: args{
				param: SubjectParam{
					Name: "jewels",
					Type: "string",
				},
				line: "jewels = \"aA\", stones = \"aAAbbbb\"",
			},
			want: SubjectTestParam{
				Name:  "jewels",
				Value: "\"aA\"",
			},
		},
		{
			args: args{
				param: SubjectParam{
					Name: "forts",
					Type: "bool",
				},
				line: "输入：forts = true",
			},
			want: SubjectTestParam{
				Name:  "forts",
				Value: "true",
			},
		},
		{
			args: args{
				param: SubjectParam{
					Name: "grid",
					Type: "[][]int",
				},
				line: "输入：grid = [[1,2,4],[3,3,1]]",
			},
			want: SubjectTestParam{
				Name:  "grid",
				Value: "[][]int{{1,2,4},{3,3,1}}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalInputSubjectTestParam(tt.args.param, tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalInputSubjectTestParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalInputSubjectTestParam() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalOutputSubjectTestParam(t *testing.T) {
	type args struct {
		ansReturnType string
		line          string
	}
	tests := []struct {
		name    string
		args    args
		want    SubjectTestParam
		wantErr bool
	}{
		{
			args: args{
				ansReturnType: "int",
				line:          "输出：4",
			},
			want: SubjectTestParam{
				Name:  "want",
				Value: "4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalOutputSubjectTestParam(tt.args.ansReturnType, tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalOutputSubjectTestParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalOutputSubjectTestParam() got = %v, want %v", got, tt.want)
			}
		})
	}
}
