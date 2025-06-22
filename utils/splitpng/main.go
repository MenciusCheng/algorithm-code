package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ImageSplitter 图片切割器
type ImageSplitter struct {
	img            image.Image
	segmentHeight  int
	whiteThreshold uint8
	whiteBuffer    int // 前后空白行缓冲区大小
}

// NewImageSplitter 创建新的图片切割器
func NewImageSplitter(img image.Image, segmentHeight int) *ImageSplitter {
	return &ImageSplitter{
		img:            img,
		segmentHeight:  segmentHeight,
		whiteThreshold: 240, // 白色阈值，可以调整
		whiteBuffer:    5,   // 前后需要的空白行数
	}
}

// isWhiteLine 检查某一行是否主要是白色
func (is *ImageSplitter) isWhiteLine(y int) bool {
	bounds := is.img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	whitePixels := 0

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		r, g, b, _ := is.img.At(x, y).RGBA()
		// 转换为8位值
		r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

		// 检查是否接近白色
		if r8 >= is.whiteThreshold && g8 >= is.whiteThreshold && b8 >= is.whiteThreshold {
			whitePixels++
		}
	}

	// 如果80%以上的像素都是白色，认为这是白色行
	return float64(whitePixels)/float64(width) >= 0.8
}

// hasWhiteBuffer 检查指定位置前后是否都有足够的空白行
func (is *ImageSplitter) hasWhiteBuffer(y int, startY int) bool {
	bounds := is.img.Bounds()

	// 检查向上的空白行
	upwardWhiteLines := 0
	for i := 1; i <= is.whiteBuffer && y-i >= startY && y-i >= bounds.Min.Y; i++ {
		if is.isWhiteLine(y - i) {
			upwardWhiteLines++
		} else {
			break // 遇到非空白行就停止
		}
	}

	// 检查向下的空白行
	downwardWhiteLines := 0
	for i := 1; i <= is.whiteBuffer && y+i < bounds.Max.Y; i++ {
		if is.isWhiteLine(y + i) {
			downwardWhiteLines++
		} else {
			break // 遇到非空白行就停止
		}
	}

	// 检查当前行是否也是空白行
	currentLineIsWhite := is.isWhiteLine(y)

	// 必须满足：当前行是空白行，且前后都有足够的空白行
	return currentLineIsWhite && upwardWhiteLines >= is.whiteBuffer && downwardWhiteLines >= is.whiteBuffer
}

// findWhiteRegionWithBuffer 在指定范围内寻找有足够缓冲区的白色区域
func (is *ImageSplitter) findWhiteRegionWithBuffer(startY, endY int) int {
	// 从目标位置向上下搜索白色行
	targetY := startY + is.segmentHeight
	bounds := is.img.Bounds()

	if targetY >= bounds.Max.Y {
		return bounds.Max.Y - 1
	}

	// 搜索范围
	searchRange := 1000 // 扩大搜索范围，因为条件更严格

	// 首先检查目标位置
	if is.hasWhiteBuffer(targetY, startY) {
		fmt.Printf("在目标位置 %d 找到合适的切割点（有%d像素缓冲区）\n", targetY, is.whiteBuffer)
		return targetY
	}

	// 向下搜索
	for i := 1; i <= searchRange && targetY+i < bounds.Max.Y; i++ {
		if is.hasWhiteBuffer(targetY+i, startY) {
			fmt.Printf("在位置 %d 找到合适的切割点（有%d像素缓冲区）\n", targetY+i, is.whiteBuffer)
			return targetY + i
		}
	}

	// 向上搜索
	for i := 1; i <= searchRange && targetY-i > startY+is.whiteBuffer*2; i++ {
		if is.hasWhiteBuffer(targetY-i, startY) {
			fmt.Printf("在位置 %d 找到合适的切割点（有%d像素缓冲区）\n", targetY-i, is.whiteBuffer)
			return targetY - i
		}
	}

	// 如果没找到有缓冲区的位置，尝试寻找普通的白色行
	fmt.Printf("警告：在位置 %d 附近未找到有%d像素缓冲区的切割点，使用普通白色行\n", targetY, is.whiteBuffer)

	// 向下搜索普通白色行
	for i := 1; i <= searchRange && targetY+i < bounds.Max.Y; i++ {
		if is.isWhiteLine(targetY + i) {
			return targetY + i
		}
	}

	// 向上搜索普通白色行
	for i := 1; i <= searchRange && targetY-i > startY; i++ {
		if is.isWhiteLine(targetY - i) {
			return targetY - i
		}
	}

	// 如果都没找到，返回目标位置
	fmt.Printf("警告：在位置 %d 附近未找到任何合适的切割点，使用目标位置\n", targetY)
	return targetY
}

// analyzeWhiteRegion 分析白色区域的连续性
func (is *ImageSplitter) analyzeWhiteRegion(centerY int) (upLines, downLines int) {
	bounds := is.img.Bounds()

	// 向上计算连续空白行
	for i := 1; centerY-i >= bounds.Min.Y; i++ {
		if is.isWhiteLine(centerY - i) {
			upLines++
		} else {
			break
		}
	}

	// 向下计算连续空白行
	for i := 1; centerY+i < bounds.Max.Y; i++ {
		if is.isWhiteLine(centerY + i) {
			downLines++
		} else {
			break
		}
	}

	return
}

// Split 执行图片切割
func (is *ImageSplitter) Split() []image.Image {
	bounds := is.img.Bounds()
	var segments []image.Image

	currentY := bounds.Min.Y
	segmentIndex := 0

	fmt.Printf("切割参数：目标高度=%d像素，空白缓冲区=%d像素\n", is.segmentHeight, is.whiteBuffer)

	for currentY < bounds.Max.Y {
		// 寻找下一个切割点
		nextY := is.findWhiteRegionWithBuffer(currentY, bounds.Max.Y)

		// 确保不会超出图片边界
		if nextY >= bounds.Max.Y {
			nextY = bounds.Max.Y - 1
		}

		// 分析切割点的空白区域情况
		upLines, downLines := is.analyzeWhiteRegion(nextY)

		// 创建子图片
		segmentBounds := image.Rect(bounds.Min.X, currentY, bounds.Max.X, nextY+1)
		segment := &subImage{
			img:    is.img,
			bounds: segmentBounds,
		}

		segments = append(segments, segment)

		fmt.Printf("切割片段 %d: Y轴 %d 到 %d (高度: %d) - 切割点空白区域：上%d行，下%d行\n",
			segmentIndex, currentY, nextY, nextY-currentY+1, upLines, downLines)

		currentY = nextY + 1
		segmentIndex++
	}

	return segments
}

// SetWhiteBuffer 设置空白缓冲区大小
func (is *ImageSplitter) SetWhiteBuffer(buffer int) {
	is.whiteBuffer = buffer
}

// SetWhiteThreshold 设置白色阈值
func (is *ImageSplitter) SetWhiteThreshold(threshold uint8) {
	is.whiteThreshold = threshold
}

// subImage 实现image.Image接口的子图片
type subImage struct {
	img    image.Image
	bounds image.Rectangle
}

func (s *subImage) ColorModel() color.Model {
	return s.img.ColorModel()
}

func (s *subImage) Bounds() image.Rectangle {
	return s.bounds
}

func (s *subImage) At(x, y int) color.Color {
	if !image.Pt(x, y).In(s.bounds) {
		return color.RGBA{}
	}
	return s.img.At(x, y)
}

// saveImage 保存图片
func saveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
	case ".png":
		return png.Encode(file, img)
	default:
		return png.Encode(file, img)
	}
}

// loadImage 加载图片
func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

func main() {
	// 使用示例
	if len(os.Args) < 4 {
		fmt.Println("使用方法: go run main.go <输入图片路径> <输出目录> <切割高度> [空白缓冲区大小]")
		fmt.Println("示例: go run main.go input.jpg output/ 800")
		fmt.Println("示例: go run main.go input.jpg output/ 800 5")
		fmt.Println("")
		fmt.Println("空白缓冲区大小：切割点前后需要的空白行数（默认5像素）")
		return
	}

	inputPath := os.Args[1]
	outputDir := os.Args[2]
	var segmentHeight int
	fmt.Sscanf(os.Args[3], "%d", &segmentHeight)

	// 可选的空白缓冲区参数
	whiteBuffer := 5 // 默认值
	if len(os.Args) >= 5 {
		fmt.Sscanf(os.Args[4], "%d", &whiteBuffer)
	}

	// 创建输出目录
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatal("创建输出目录失败:", err)
	}

	// 加载图片
	fmt.Println("正在加载图片:", inputPath)
	img, err := loadImage(inputPath)
	if err != nil {
		log.Fatal("加载图片失败:", err)
	}

	// 创建切割器
	splitter := NewImageSplitter(img, segmentHeight)
	splitter.SetWhiteBuffer(whiteBuffer)

	// 执行切割
	fmt.Printf("开始切割图片，目标高度: %d 像素，空白缓冲区: %d 像素\n", segmentHeight, whiteBuffer)
	segments := splitter.Split()

	// 保存切割后的图片
	fmt.Printf("正在保存 %d 个切割片段...\n", len(segments))
	for i, segment := range segments {
		filename := fmt.Sprintf("%s/segment_%03d.png", outputDir, i+1)
		err := saveImage(segment, filename)
		if err != nil {
			log.Printf("保存片段 %d 失败: %v", i+1, err)
			continue
		}
		fmt.Printf("保存片段 %d: %s\n", i+1, filename)
	}

	fmt.Println("切割完成！")
}
