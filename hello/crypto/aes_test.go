package crypto

import "testing"

func TestAesEncrypt(t *testing.T) {
	field := "18666666666"

	// 生成aeskey
	aesKey, err := GenerateRandomString(32)
	if err != nil {
		t.Error(err)
		return
	}

	encryptField := AesEncrypt(field, aesKey)
	decryptField := AesDecrypt(encryptField, aesKey)
	if decryptField != field {
		t.Errorf("AesDecrypt() = %v, want %v", decryptField, field)
		return
	}
}
