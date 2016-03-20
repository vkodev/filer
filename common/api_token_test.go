package common

import (
	"testing"
	"time"
)

const (
	password = "123asd"
)

func TestHashPass(t *testing.T) {
	hashedPass, err := HashPass(password)

	if err != nil {
		t.Error("HashPass func failure: ", err, hashedPass)
	}
}

func TestCheckPass(t *testing.T) {
	user, _ := NewApiUser("test", password, true)

	if !user.CheckPass(password) {
		t.Error("Password equals but func return false")
	}
}

func TestTokenGenerate(t *testing.T) {
	token := NewApiToken(nil)

	if token.Token == "" {
		t.Error("Token not generated nill or empty", token.Token)
	}
}

func TestTokenExpiry(t *testing.T) {
	token := NewApiToken(nil)

	if token.IsExpired() {
		t.Error("Token just created but has already expired", token)
	}

	token.CreatedAt = time.Now().Add(DefaultExpiry)

	if !token.IsExpired() {
		t.Error("Token expired but func return not", token)
	}

}
