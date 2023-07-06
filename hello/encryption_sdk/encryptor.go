package encryption_sdk

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/MenciusCheng/algorithm-code/hello/crypto"
)

type Encryptor interface {
	// Encrypt 加密，返回密文字符串
	Encrypt(key string, plaintext string) (string, error)

	// Decrypt 解密，返回明文字符串
	Decrypt(key string, ciphertext string) (string, error)
}

func NewEncryptor() Encryptor {
	return &encryptor{}
}

type encryptor struct {
}

func (e *encryptor) Encrypt(key, plaintext string) (string, error) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData := crypto.PKCS7Padding([]byte(plaintext), blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, []byte(key)[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	return base64.StdEncoding.EncodeToString(cryted), nil
}

func (e *encryptor) Decrypt(key, ciphertext string) (string, error) {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(ciphertext)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = crypto.PKCS7UnPadding(orig)
	return string(orig), nil
}
