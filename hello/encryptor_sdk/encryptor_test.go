package encryptor_sdk

import (
	"testing"
	"time"
)

func TestNewAesJwtEncryptor(t *testing.T) {
	type args struct {
		jwtSecret   []byte
		jwtDuration time.Duration
		plainText   string
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
				jwtSecret:   []byte("xingyu"),
				jwtDuration: 12 * time.Hour,
				plainText:   "{\"text\":\"123abcAbc@你好呀！\"}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := NewAesJwtEncryptor(tt.args.jwtSecret, tt.args.jwtDuration)

			key, err := entity.GenKey()
			if (err != nil) != tt.wantGenErr {
				t.Errorf("GenKey() error = %v, wantGenErr %v", err, tt.wantGenErr)
				return
			}

			cipherText, err := entity.Encrypt(key, tt.args.plainText)
			if (err != nil) != tt.wantEncryptErr {
				t.Errorf("Encrypt() error = %v, wantEncryptErr %v", err, tt.wantEncryptErr)
				return
			}
			t.Logf("key = %s, plainText = %s, cipherText = %s", key, tt.args.plainText, cipherText)

			plainText, err := entity.Decrypt(key, cipherText)
			if (err != nil) != tt.wantDecryptErr {
				t.Errorf("Decrypt() error = %v, wantDecryptErr %v", err, tt.wantDecryptErr)
				return
			}

			if plainText != tt.args.plainText {
				t.Errorf("Decrypt() failed, plainText = %s, tt.args.plainText = %s", plainText, tt.args.plainText)
				return
			}
		})
	}
}
