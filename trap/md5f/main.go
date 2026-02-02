package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	//fmt.Printf("%s\n", GetMD5Hash([]byte("zzzssss")))

	idfa := "EAE643DA-1122-47FB-A610-5044DC2240D4"
	fmt.Println(Md5IDFA(idfa))
}

func GetMD5Hash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func Md5IDFA(idfa string) string {
	upper := strings.ToUpper(idfa)     // 保持大写
	hash := md5.Sum([]byte(upper))     // 计算 MD5
	return hex.EncodeToString(hash[:]) // 转小写 32 位
}
