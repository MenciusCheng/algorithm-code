package encryptor_sdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

const (
	// Symbols 随机字符串可用字符集
	Symbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// AesKey256Len aes 密钥长度
	AesKey256Len = 32
)

// GenerateRandomString 生成一个指定长度的随机字符串（只包含大小写字母与数字）
func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(Symbols))
	for i, b := range randomBytes {
		randomBytes[i] = Symbols[b%symbolsByteLength]
	}
	return string(randomBytes), nil
}

// AESEncrypt aes 加密，返回的密文由base64编码
func AESEncrypt(aesKey, plainText string) (string, error) {
	// 转成字节数组
	key := []byte(aesKey)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData := PKCS7Padding([]byte(plainText), blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// 创建数组
	cipherText := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cipherText, origData)
	ciphertextBase64 := base64.StdEncoding.EncodeToString(cipherText)
	return ciphertextBase64, nil
}

// AESDecrypt aes 解密，返回明文字符串
func AESDecrypt(aesKey, cipherText string) (string, error) {
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("cipher too short")
	}
	key := []byte(aesKey)

	// 转成字节数组
	cipherByte, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	// 分组秘钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	// 创建数组
	plainText := make([]byte, len(cipherByte))
	// 解密
	blockMode.CryptBlocks(plainText, cipherByte)
	// 去补全码
	plainText = PKCS7UnPadding(plainText)
	return string(plainText), err
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
