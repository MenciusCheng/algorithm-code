package main

import (
	"encoding/json"
	"fmt"
)

// 邮件图表配置
type MessageTaskMailChartModel struct {
	Title     string                           `json:"title"`     // 标题
	ChartList []MessageTaskMailChartModelChart `json:"chartList"` // 图表列表
}

// 邮件图表配置的展示图表
type MessageTaskMailChartModelChart struct {
	StatSqlId    int64  `json:"statSqlId"`    // 分析SQL记录ID
	ChartTableId string `json:"chartTableId"` // 图表ID
	Remark       string `json:"remark"`       // 备注
	UsedStatus   int32  `json:"usedStatus"`   // 使用的状态 1在使用 -1 不使用
	ShowNumber   int32  `json:"showNumber"`   // 是否显示数字，1是，-1否
}

func main() {
	m := MessageTaskMailChartModel{
		Title: "",
		ChartList: []MessageTaskMailChartModelChart{
			{
				StatSqlId:    0,
				ChartTableId: "",
				Remark:       "",
				UsedStatus:   0,
				ShowNumber:   0,
			},
		},
	}
	bs, _ := json.Marshal(m)
	fmt.Printf("%s\n", bs)
}
