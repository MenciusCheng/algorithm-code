package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", GetMD5Hash([]byte("zzzssss")))
}

func GetMD5Hash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
