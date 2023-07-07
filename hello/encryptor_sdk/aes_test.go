package encryptor_sdk

import "testing"

func TestAES_GenKey_Encrypt_Decrypt(t *testing.T) {
	type args struct {
		plainText string
	}
	tests := []struct {
		name           string
		args           args
		wantGenErr     bool
		wantEncryptErr bool
		wantDecryptErr bool
	}{
		{
			args: args{
				plainText: "123abcAbc@你好呀！",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aesKey, err := GenerateRandomString(AesKey256Len)
			if (err != nil) != tt.wantGenErr {
				t.Errorf("GenerateRandomString() error = %v, wantGenErr %v", err, tt.wantGenErr)
				return
			}

			cipherText, err := AESEncrypt(aesKey, tt.args.plainText)
			if (err != nil) != tt.wantEncryptErr {
				t.Errorf("AESEncrypt() error = %v, wantEncryptErr %v", err, tt.wantEncryptErr)
				return
			}
			t.Logf("aesKey = %s, plainText = %s, cipherText = %s", aesKey, tt.args.plainText, cipherText)

			plainText, err := AESDecrypt(aesKey, cipherText)
			if (err != nil) != tt.wantDecryptErr {
				t.Errorf("AESDecrypt() error = %v, wantDecryptErr %v", err, tt.wantDecryptErr)
				return
			}

			if plainText != tt.args.plainText {
				t.Errorf("AESDecrypt() failed, plainText = %s, tt.args.plainText = %s", plainText, tt.args.plainText)
				return
			}
		})
	}
}
