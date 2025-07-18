package main

import (
	"fmt"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go <MP4文件路径>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("文件不存在: %s\n", filePath)
		os.Exit(1)
	}

	fmt.Printf("分析文件: %s\n", filePath)

	// 获取视频信息
	videoInfo, err := getVideoInfo(filePath)
	if err != nil {
		fmt.Printf("获取视频信息失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(videoInfo)
}

func getVideoInfo(filePath string) (string, error) {
	// 使用ffprobe获取视频信息
	probe, err := ffmpeg.Probe(filePath)
	if err != nil {
		return "", err
	}
	return probe, nil
}
