package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

const (
	// NonceSymbols 随机字符串可用字符集
	NonceSymbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// NonceLength 随机字符串的长度
	NonceLength = 32
)

// 平台证书公钥
const PublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqz21WOq4989db66Yvt3f
U7p/rRdc/YFXZMLhuAcPRRjPK2x2aJyKHNIb0il7p1VB0A3mZoncOwxqS6GwsbBZ
5cb4aOnJex/FoIaizzfk7aloqoiieaXF+BWzeZifeTBFZO+1JR66f74QW/GtefdS
YmlENomJRAXz4GhpUoyyhJCqIG2ANWQeLi0QZmzzFo90oIJonFhSI9yB285Gq5/D
Dg6HOdWcF5wsFpyXmWftkpeS6OnRUO+7EIP0TbBwhxTHTJAYcN2tbs1rZhobNer6
0MDedzayZV3VTVOh2+EpD8cZJvEDf5ZHII+VXz9qGJP0Efnx2MwOxDRemORBhzQm
PQIDAQAB
-----END PUBLIC KEY-----
`

// 商户私钥
const PrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCrPbVY6rj3z11v
rpi+3d9Tun+tF1z9gVdkwuG4Bw9FGM8rbHZonIoc0hvSKXunVUHQDeZmidw7DGpL
obCxsFnlxvho6cl7H8WghqLPN+TtqWiqiKJ5pcX4FbN5mJ95MEVk77UlHrp/vhBb
8a1591JiaUQ2iYlEBfPgaGlSjLKEkKogbYA1ZB4uLRBmbPMWj3SggmicWFIj3IHb
zkarn8MODoc51ZwXnCwWnJeZZ+2Sl5Lo6dFQ77sQg/RNsHCHFMdMkBhw3a1uzWtm
Ghs16vrQwN53NrJlXdVNU6Hb4SkPxxkm8QN/lkcgj5VfP2oYk/QR+fHYzA7ENF6Y
5EGHNCY9AgMBAAECggEBAKm4aJPvWFKsoSP+/X83MpCu6LpqzIypdC/3A3a88IKA
e1zDjHxUooK8MDzVRqxr0OXf0PrSKogP3w6E/Daw3fjigvSgukQhEmPu7VDKSkjP
vmA/OiGtjxA6mf6rqWuYzz6iYnye/yN7AoXOfCzFV08YnDYdjQHxkGxbr//UWV8O
z1W4LYOtI8GzsfSs5a+3jIw9d9VXODHxqHVU1w2zO+f2ah6MJpJjWkT0OeCfi2cA
yrAkcTSO4fT+hzd9XSc8xpHHUiuR4rqmSY4zj9EEiw+8iXBGaMKnb2qYrHufOnHI
gTwiBov12y8NH2BLs3GsugRsTUFEt6ZEhtfeYk+PrPUCgYEA2nwD0R9LZtXvhMK+
bnS5Kz16ysYaUy6UDq7OSXvZ4RJGWduUtFRRToyiak400bUiqevRRlyky5FQPO6I
GgB4/hZIc/3VUIDI4hjYjhksIHKwEKoT87Dcr2XxjRr9oEuCuHQh2Sv5Voxt0IF9
tZA7cqjtBUBkGEdHKyU04Zub2BsCgYEAyKUAgidEBzWnma0rMjs2I6mf0S/5mC7e
SdYwrEqjkX/swk9Rvdqj7gQQw7nCm4Ocz9sfTvYjCGvkaCpg+WuzyuOH83NwOlsD
5KMOKHS5KV8WhpzCGlZnY17aISCtN2ylkkOpp4O+xjkY2zYNA7WqMBlcwBXRnlvu
I3gD+b7CkIcCgYB5P62Wlnlv3nYIBVNNVTWVy46jjYD6fLTp2RTeLv9hKpUkNPm3
gbuFjgJsYG5rmsxb5GTFWZCS8FfJGM8rLuv2hkM7K7j/7hiNcSBs64dTkpInDVv5
N0Ohiz+cAUiTdpRa7QgDz6WV5GTk+5fZ9Oso3Jp8+ZkS6CuUxBcuxP6d1QKBgHYU
rSTryNxZVkZZpEdOk4EPbY+lpUVLp+RaWpC66GJTn9mG5rtNthX4bIlThFUTcrDc
6yMENf/ZxzUKY8HGAayQlUzQDic5syPWVhm2/9V6MX6NOKpZWUUg6EyAt5Abr6FR
ksGUBi0QIzG3MJeTTRI7its8u/1vasmS8vwCqOx3AoGAFqKJVZjZqEZis7S5qazl
X/qaYHBSyciGj5EluUjaiZ0SbaBXD1uhd+pmH6ws9gT2lIP7AhioZ8D4rFD0puPx
IYOxfyaOqVoRXNbFPWXTCBKyMCWyd6idtO4FmIiHzsYPHDpxyjx5VaRKttbbwliq
5ACSA+xbNBNcKD1eiZVdTx4=
-----END PRIVATE KEY-----
`

func encryptDemo() {
	field := "18666666666"

	// 生成aeskey
	aesKey, err := GenerateRandomString(32)
	if err != nil {
		fmt.Println("generate aes key failed", err)
		return
	}
	// 生成随机串
	nonce, err := GenerateRandomString(NonceLength)
	// 指定附加数据
	additionalData := []byte("")
	if err != nil {
		fmt.Println("generate nonce failed", err)
		return
	}
	// 对原始字段aes-gcm加密
	encryptField, err := AESGCMEncrypt(aesKey, field, []byte(nonce), additionalData)
	if err != nil {
		fmt.Println("aes-gcm encrypt failed", err)
		return
	}
	// 对aesKey进行RSA加密
	encryptKey, err := RSAEncrypt(PublicKey, aesKey)
	if err != nil {
		fmt.Println("rsa encrypt failed", err)
		return
	}
	fmt.Println("aesKey:", aesKey)
	fmt.Println("nonce:", nonce)
	fmt.Println("additionalData:", additionalData)
	fmt.Println("-------")

	fmt.Println("encryptKey:", encryptKey)
	fmt.Println("encryptField:", encryptField)
	// 请求 附带加密key、加密字段、nonce、附加数据、平台证书序列号
	// ....
	// 响应
	fmt.Println("-------")
	// 根据返回的SerialNo获取证书对应的私钥
	// 使用RSA对加密key解密
	plainAesKey, err := RSADecrypt(PrivateKey, encryptKey)
	if err != nil {
		fmt.Println("rsa decrypt failed", err)
		return
	}
	// 使用解密后的key，解密密文字段
	plainField, err := AESGCMDecrypt(plainAesKey, encryptField, []byte(nonce), additionalData)
	if err != nil {
		fmt.Println("aes-gcm decrypt failed", err)
		return
	}
	fmt.Println("plainField", plainField)
	fmt.Println("plainAesKey:", plainAesKey)
}

// AESGCMEncrypt aes-gcm 加密，返回的密文由base64编码
// GCM的nonce长度用的是标准的12个字符，gcmStandardNonceSize = 12, 建议采用标准长度
func AESGCMEncrypt(aesKey, plainText string, nonce, additionalData []byte) (string, error) {
	key := []byte(aesKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return "", err
	}
	ciphertext := aesGcm.Seal(nil, nonce, []byte(plainText), additionalData)
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)
	return ciphertextBase64, nil
}

// AESGCMDecrypt aes-gcm解密，返回明文字符串
func AESGCMDecrypt(aesKey, ciphertext string, nonce, additionalData []byte) (string, error) {
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("cipher too short")
	}
	key := []byte(aesKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return "", err
	}
	cipherData, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := aesGcm.Open(nil, []byte(nonce), cipherData, []byte(additionalData))
	if err != nil {
		return "", err
	}

	return string(plaintext), err
}

// GenerateRandomString 生成一个指定长度的随机字符串（只包含大小写字母与数字）
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(NonceSymbols))
	for i, b := range bytes {
		bytes[i] = NonceSymbols[b%symbolsByteLength]
	}
	return string(bytes), nil
}

// RSAEncrypt 公钥加密, 返回base64编码的密文
func RSAEncrypt(publicKey, plainText string) (string, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("invalid publick key")
	}
	// 加密
	ciphertext, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pubKey, []byte(plainText), nil)
	if err != nil {
		return "", err
	}
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)
	return ciphertextBase64, nil
}

// 私钥解密, 返回明文字符串
func RSADecrypt(privateKey, ciphertext string) (string, error) {
	// 获取私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}
	cipherData, err := base64.StdEncoding.DecodeString(ciphertext)
	//解析PKCS1格式的私钥
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		private, ok := privInterface.(*rsa.PrivateKey)
		if ok {
			privKey = private
		}
	}

	// 解密
	data, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privKey, cipherData, nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
