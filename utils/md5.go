package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(input string) string {
	m := md5.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}

func Md5Byte(inputByte []byte) string {
	m := md5.New()
	m.Write(inputByte)
	return hex.EncodeToString(m.Sum(nil))
}
