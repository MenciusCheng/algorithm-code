package timef

import (
	"time"
)

func PlayerBirthdayChangeUnix(month, day int) int64 {
	now := time.Now()
	birthdayInt := time.Date(now.Year(), time.Month(month), day, 0, 0, 0, 0, time.Local).Unix()
	return birthdayInt
}

var OffsetTime int64 = 0

// 获取当前时间
func GetNow() time.Time {
	now := time.Now().Add(time.Duration(OffsetTime) * time.Second)
	return now
}

func GetDayDot() int64 {
	timeStr := GetNow().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	_, offset := GetNow().Zone()
	timeNumber := t.Unix() - int64(offset)
	return timeNumber
}

// 将时间字符串转换为时间戳
func ConvertGroupStrToGroup(groupStr string) int64 {
	if groupStr == "" {
		return 0
	}
	// 尝试解析时间格式
	t, err := time.ParseInLocation("2006-01-02", groupStr, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}
