package timef

import (
	"testing"
	"time"
)

func TestPlayerBirthdayChangeUnix(t *testing.T) {
	type args struct {
		month int
		day   int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				month: 2,
				day:   29,
			},
			want: time.Date(2025, 2, 29, 0, 0, 0, 0, time.Local).Unix(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PlayerBirthdayChangeUnix(tt.args.month, tt.args.day)
			t.Logf("PlayerBirthdayChangeUnix() got=%v", time.Unix(got, 0).Format("2006-01-02 15:04:05"))
			if got != tt.want {
				t.Errorf("PlayerBirthdayChangeUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}
