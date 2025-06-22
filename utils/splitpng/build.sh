#!/bin/bash

echo "正在编译Windows版本..."

# 创建输出目录
mkdir -p dist

# 编译Windows 64位版本
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/image-splitter.exe main.go

# 编译macOS版本
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o dist/image-splitter main.go

echo "编译完成！文件在 dist/ 目录中"
