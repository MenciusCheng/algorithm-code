package urlf

import (
	"net/http"
	"testing"
)

func TestUrlParse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//{
		//	args: args{path: "https://www.baidu.com/ttt"},
		//},
		//{
		//	args: args{path: "nei/wang/img2img_662_1699946020229_01.jpg"},
		//},
		{
			args: args{path: "https%3A%2F%2Ftrack.sigmob.cn%2Ftrack%3Fc%3DCiRmMmIzMzYxOS05ZjZmLTExZWMtYTg2Zi0wMDE2M2UzMjk0N2YSJGYyYjMzNjE5LTlmNmYtMTFlYy1hODZmLTAwMTYzZTMyOTQ3ZioLZTdjYzI5NDk1ZTk6cQoCNjQSBTU0MTk3GgYxMzA4MzAgAyiIJzCIJ0oeaHR0cDovL2Euc2lnbW9iLmNvbS8jL2luZGV4MTExUgUzNjI2NloEMTA1OWkotnCu-IzCP3FkI3QPsdGCP3gBgAHaFYgBqpsCmAH3pAWgAWSoAYgnSAJQAVqJAhIQMzQxZmU0MjY2MWI2ODM3ZCIPODY3Nzg3MDMzMjk2ODI3KhAzNDFmZTQyNjYxYjY4MzdkMgd1bmtub3duQiQzMmY5Y2YyOC1lZWVhLTQ2MDItOGJiOC0wN2RmMGI1YjUxNjhKEDA4NTI4NTI0MzZjM2FmZTJSBlhpYW9taVoPODY3Nzg3MDMzMjk2ODI3Yg84Njc3ODcwMzMyOTY4MzWSASA2YjkxYTJlYTIyODVmNjEzMmIxMTYzMjg4YjljZmJkNpoBIDE3N2NhMDUxMzVjYjI3MzA4M2VhMTZhOGI1MGM3OTM1ogEgYjgwOGQ2NmM1NDg4ODhjOTRlYzk4NjRmMTRiZDRlZDdiBDEyODJqDjEyMC4xMzMuNDIuMTM1kAEHmAECoAEBwAEE0AEB%26e%3Dactive%26p%3DOML9xjIE2GpXa5Auktzzng"},
		},
		//{
		//	//args:    args{path: "netdisc/prod/userTmp/user/103/我的图片/1699934030980_`%O{TC8W~L}8EY$H5_%7H]9.png"},
		//	args:    args{path: "netdisc%2Fprod%2FuserTmp%2Fuser%2F103%2F%E6%88%91%E7%9A%84%E5%9B%BE%E7%89%87%2F1699934030980_%60%25O%7BTC8W~L%7D8EY%24H5_%257H%5D9.png"},
		//	//wantErr: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UrlParse(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("UrlParse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIpParse(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{ipStr: "2406:3003:2003:1661:d8d2:a583:cb03:469a"},
		},
		{
			args: args{ipStr: "F3E1:8214:53C1:B86A:CCDC:905A:5434:9000"},
		},
		{
			args: args{ipStr: "192.168.1.1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IpParse(tt.args.ipStr); (err != nil) != tt.wantErr {
				t.Errorf("IpParse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSplitHostPort(t *testing.T) {
	type args struct {
		hostport string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				hostport: "192.168.1.1:80",
			},
		},
		{
			args: args{
				hostport: "F3E1:8214:53C1:B86A:CCDC:905A:5434:9000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SplitHostPort(tt.args.hostport); (err != nil) != tt.wantErr {
				t.Errorf("SplitHostPort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestForwarded(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				req: &http.Request{
					Header: map[string][]string{
						"X-Forwarded-For": {"2001:db8:85a3:8d3:1319:8a2e:370:7348"},
					},
				},
			},
		},
		{
			args: args{
				req: &http.Request{
					Header: map[string][]string{
						"X-Forwarded-For": {"203.0.113.195"},
					},
				},
			},
		},
		{
			args: args{
				req: &http.Request{
					Header: map[string][]string{
						"X-Forwarded-For": {"203.0.113.195, 2001:db8:85a3:8d3:1319:8a2e:370:7348"},
					},
				},
			},
		},
		{
			args: args{
				req: &http.Request{
					Header: map[string][]string{
						"X-Forwarded-For": {"203.0.113.195,2001:db8:85a3:8d3:1319:8a2e:370:7348,198.51.100.178"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Forwarded(tt.args.req); got != tt.want {
				t.Errorf("Forwarded() = %v, want %v", got, tt.want)
			}
		})
	}
}
