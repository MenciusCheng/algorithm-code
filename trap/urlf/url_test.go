package urlf

import "testing"

func TestUrlParse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{path: "https://www.baidu.com/ttt"},
		},
		{
			args: args{path: "nei/wang/pc.html"},
		},
		{
			args:    args{path: "netdisc/prod/userTmp/user/103/我的图片/1699934030980_`%O{TC8W~L}8EY$H5_%7H]9.png"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UrlParse(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("UrlParse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
