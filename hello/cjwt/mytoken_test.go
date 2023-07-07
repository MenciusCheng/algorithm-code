package cjwt

import (
	"testing"
	"time"
)

func TestMyToken_GenerateToken(t *testing.T) {
	m := NewMyToken([]byte("aaabbb123"), 3*time.Second)
	username := "大猫"

	got, err := m.GenerateToken(username)
	if err != nil {
		t.Errorf("GenerateToken() error = %v", err)
		return
	}

	claims, err := m.ParseToken(got)
	if err != nil {
		t.Errorf("ParseToken() error = %v", err)
		return
	}

	t.Logf("claims: %+v", claims)
	if claims.Username != username {
		t.Errorf("claims.Username = %s", claims.Username)
		return
	}
}

func TestMyToken_GenerateToken_signature_invalid(t *testing.T) {
	m := NewMyToken([]byte("aaabbb123"), 1*time.Second)
	username := "大猫"

	got, err := m.GenerateToken(username)
	if err != nil {
		t.Errorf("GenerateToken() error = %v", err)
		return
	}

	m2 := NewMyToken([]byte("a23x23452"), 1*time.Second)
	claims, err := m2.ParseToken(got)
	if err != nil {
		t.Errorf("ParseToken() error = %v", err)
		return
	}

	t.Logf("claims: %+v", claims)
	if claims.Username != username {
		t.Errorf("claims.Username = %s", claims.Username)
		return
	}
}

func TestMyToken_GenerateToken_expired(t *testing.T) {
	m := NewMyToken([]byte("aaabbb123"), 500*time.Millisecond)
	username := "大猫"

	got, err := m.GenerateToken(username)
	if err != nil {
		t.Errorf("GenerateToken() error = %v", err)
		return
	}

	time.Sleep(time.Second)

	claims, err := m.ParseToken(got)
	if err != nil {
		t.Errorf("ParseToken() error = %v", err)
		return
	}

	t.Logf("claims: %+v", claims)
	if claims.Username != username {
		t.Errorf("claims.Username = %s", claims.Username)
		return
	}
}

func TestMyToken_GenerateToken1(t *testing.T) {
	type fields struct {
		jwtSecret     []byte
		validDuration time.Duration
	}
	type args struct {
		username string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		want            string
		wantErr         bool
		wantGotParseErr error
	}{
		{
			name: "",
			fields: fields{
				jwtSecret:     []byte("aaabbb123"),
				validDuration: 3 * time.Second,
			},
			args: args{
				username: "大猫",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMyToken(tt.fields.jwtSecret, tt.fields.validDuration)
			got, err := m.GenerateToken(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			claims, err := m.ParseToken(got)
			if err != tt.wantGotParseErr {
				t.Errorf("ParseToken() error = %v", err)
				return
			}

			t.Logf("claims: %+v", claims)
			if claims.Username != tt.args.username {
				t.Errorf("claims.Username = %s", claims.Username)
				return
			}
		})
	}
}
