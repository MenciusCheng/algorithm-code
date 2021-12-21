package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
https://leetcode-cn.com/problems/day-of-the-year/

1154. 一年中的第几天
给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。请你计算并返回该日期是当年的第几天。
通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。每个月的天数与现行公元纪年法（格里高利历）一致。

示例 1：

输入：date = "2019-01-09"
输出：9
示例 2：

输入：date = "2019-02-10"
输出：41
示例 3：

输入：date = "2003-03-01"
输出：60
示例 4：

输入：date = "2004-03-01"
输出：61

提示：

date.length == 10
date[4] == date[7] == '-'，其他的 date[i] 都是数字
date 表示的范围从 1900 年 1 月 1 日至 2019 年 12 月 31 日
*/
func main() {
	var tests = []struct {
		date string
		want int
	}{
		{
			date: "2019-01-09",
			want: 9,
		},
		{
			date: "2019-02-10",
			want: 41,
		},
		{
			date: "2003-03-01",
			want: 60,
		},
		{
			date: "2004-03-01",
			want: 61,
		},
	}

	for _, item := range tests {
		if ans := dayOfYear(item.date); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	var days = day
	for i := 1; i < month; i++ {
		days += getMothDays(year, i)
	}
	return days
}

func getMothDays(year int, month int) int {
	var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	days := monthDays[month-1]
	if month == 2 && isLeap(year) {
		days++
	}
	return days
}

func isLeap(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
