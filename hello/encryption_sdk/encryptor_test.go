package encryption_sdk

import (
	"github.com/MenciusCheng/algorithm-code/hello/crypto"
	"testing"
)

func Test_encryptor_Encrypt(t *testing.T) {

	field := "18666666666"

	// 生成aeskey
	aesKey, err := crypto.GenerateRandomString(32)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("aesKey: %s", aesKey)

	cp := NewEncryptor()

	encryptField, err := cp.Encrypt(aesKey, field)
	if err != nil {
		t.Error(err)
		return
	}

	decryptField, err := cp.Decrypt(aesKey, encryptField)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("aesKey: %s, encryptField: %s, decryptField: %s", aesKey, encryptField, decryptField)
	if decryptField != field {
		t.Errorf("AesDecrypt() = %v, want %v", decryptField, field)
		return
	}
}
