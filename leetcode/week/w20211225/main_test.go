package w20211225

import (
	"reflect"
	"testing"
)

func Test_findAllRecipes(t *testing.T) {
	type args struct {
		recipes     []string
		ingredients [][]string
		supplies    []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				recipes:     []string{"bread", "sandwich"},
				ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}},
				supplies:    []string{"yeast", "flour", "meat"},
			},
			want: []string{"bread", "sandwich"},
		},
		{
			args: args{
				recipes:     []string{"ju", "fzjnm", "x", "e", "zpmcz", "h", "q"},
				ingredients: [][]string{{"d"}, {"hveml", "f", "cpivl"}, {"cpivl", "zpmcz", "h", "e", "fzjnm", "ju"}, {"cpivl", "hveml", "zpmcz", "ju", "h"}, {"h", "fzjnm", "e", "q", "x"}, {"d", "hveml", "cpivl", "q", "zpmcz", "ju", "e", "x"}, {"f", "hveml", "cpivl"}},
				supplies:    []string{"f", "hveml", "cpivl", "d"},
			},
			want: []string{"bread", "sandwich"}, // 需要自己计算
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllRecipes(tt.args.recipes, tt.args.ingredients, tt.args.supplies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllRecipes() = %v, want %v", got, tt.want)
			}
		})
	}
}
