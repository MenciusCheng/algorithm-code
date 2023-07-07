package encryptor_sdk

import "time"

type Encryptor interface {
	// GenKey 生成密钥
	GenKey() (string, error)
	// Encrypt 加密，返回密文字符串
	Encrypt(key string, plainText string) (string, error)
	// Decrypt 解密，返回明文字符串
	Decrypt(key string, cipherText string) (string, error)
}

func NewAesJwtEncryptor(jwtSecret []byte, jwtDuration time.Duration) Encryptor {
	return &AesJwtEncryptor{
		jwtSecret:   jwtSecret,
		jwtDuration: jwtDuration,
	}
}

type AesJwtEncryptor struct {
	jwtSecret   []byte        // jwt 密钥
	jwtDuration time.Duration // jwt token 有效期
}

// GenKey 生成 jwt token，包装 aes 密钥
func (a *AesJwtEncryptor) GenKey() (string, error) {
	// 生成 aesKey
	aesKey, err := GenerateRandomString(AesKey256Len)
	if err != nil {
		return "", err
	}

	// 生成 jwt token，包装密钥
	token, err := generateJwtToken(aesKey, a.jwtSecret, a.jwtDuration)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Encrypt 加密
func (a *AesJwtEncryptor) Encrypt(key string, plainText string) (string, error) {
	aesKey, err := a.parseJwtAesKey(key)
	if err != nil {
		return "", err
	}

	encrypt, err := AESEncrypt(aesKey, plainText)
	if err != nil {
		return "", err
	}

	return encrypt, nil
}

// Decrypt 解密
func (a *AesJwtEncryptor) Decrypt(key string, cipherText string) (string, error) {
	aesKey, err := a.parseJwtAesKey(key)
	if err != nil {
		return "", err
	}

	encrypt, err := AESDecrypt(aesKey, cipherText)
	if err != nil {
		return "", err
	}

	return encrypt, nil
}

// parseJwtAesKey 从 jwtToken 解析出 aesKey
func (a *AesJwtEncryptor) parseJwtAesKey(jwtToken string) (string, error) {
	claims, err := parseJwtToken(jwtToken, a.jwtSecret)
	if err != nil {
		return "", err
	}

	return claims.SecretKey, nil
}
