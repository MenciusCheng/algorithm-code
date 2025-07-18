package main

import (
	"fmt"
	"github.com/Eyevinn/mp4ff/mp4"
	"io"
	"os"
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go <MP4文件路径>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// 打开MP4文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件错误: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// 解析MP4文件
	parsedMp4 := GetMp4UdtaBox(file)
	if parsedMp4 == nil {
		fmt.Println("解析MP4文件失败")
		os.Exit(1)
	}

	// 打印MP4文件信息
	printMp4Info(parsedMp4)
}

func GetMp4UdtaBox(r io.Reader) *mp4.File {
	// 解析 MP4 文件
	parsedMp4, err := mp4.DecodeFile(r)
	if err != nil {
		fmt.Printf("解析 MP4 文件错误: %v\n", err)
		return nil
	}
	return parsedMp4
}

func printMp4Info(mp4File *mp4.File) {
	fmt.Println("=== MP4 文件信息 ===")

	// 打印基本信息
	fmt.Printf("文件类型: %s\n", mp4File.Ftyp.MajorBrand)
	fmt.Printf("兼容品牌: %v\n", mp4File.Ftyp.CompatibleBrands)

	// 打印moov box信息
	if mp4File.Moov != nil {
		fmt.Printf("时长: %d 时间单位\n", mp4File.Moov.Mvhd.Duration)
		fmt.Printf("时间刻度: %d\n", mp4File.Moov.Mvhd.Timescale)
		if mp4File.Moov.Mvhd.Timescale > 0 {
			durationSeconds := float64(mp4File.Moov.Mvhd.Duration) / float64(mp4File.Moov.Mvhd.Timescale)
			fmt.Printf("时长: %.2f 秒\n", durationSeconds)
		}
		fmt.Printf("轨道数量: %d\n", len(mp4File.Moov.Traks))

		// 打印每个轨道的信息
		for i, trak := range mp4File.Moov.Traks {
			fmt.Printf("\n--- 轨道 %d ---\n", i+1)
			if trak.Tkhd != nil {
				fmt.Printf("  轨道ID: %d\n", trak.Tkhd.TrackID)
				fmt.Printf("  轨道时长: %d\n", trak.Tkhd.Duration)
				fmt.Printf("  轨道宽度: %d\n", trak.Tkhd.Width>>16)
				fmt.Printf("  轨道高度: %d\n", trak.Tkhd.Height>>16)
			}

			if trak.Mdia != nil && trak.Mdia.Hdlr != nil {
				fmt.Printf("  媒体类型: %s\n", trak.Mdia.Hdlr.HandlerType)
				fmt.Printf("  媒体名称: %s\n", trak.Mdia.Hdlr.Name)
			}

			if trak.Mdia != nil && trak.Mdia.Minf != nil && trak.Mdia.Minf.Stbl != nil {
				if trak.Mdia.Minf.Stbl.Stsd != nil {
					fmt.Printf("  样本数量: %d\n", len(trak.Mdia.Minf.Stbl.Stsd.Children))
				}
			}
		}
	}

	// 打印文件大小信息
	fmt.Println("\n=== 文件结构信息 ===")
	if mp4File.Ftyp != nil {
		fmt.Printf("FTYP Box 大小: %d 字节\n", mp4File.Ftyp.Size())
	}
	if mp4File.Moov != nil {
		fmt.Printf("MOOV Box 大小: %d 字节\n", mp4File.Moov.Size())
	}
}
