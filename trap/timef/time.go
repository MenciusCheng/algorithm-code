package timef

import "time"

func PlayerBirthdayChangeUnix(month, day int) int64 {
	now := time.Now()
	birthdayInt := time.Date(now.Year(), time.Month(month), day, 0, 0, 0, 0, time.Local).Unix()
	return birthdayInt
}
