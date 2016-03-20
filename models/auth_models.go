package models

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

const (
	FastExpiry    = 17 * time.Minute
	DefaultExpiry = 24 * time.Hour
	LongExpiry    = 24 * time.Hour * 7
)

type ApiToken struct {
	Token     string
	ApiUser   *ApiUser
	CreatedAt time.Time
	ExpiryAt  time.Time
}

type ApiUser struct {
	ID        int
	Login     string
	Password  string
	CanWrite  bool
	CreatedAt time.Time
}

func NewApiUser(login, password string, canWrite bool) (*ApiUser, error) {
	hashedPass, err := HashPass(password)
	if err != nil {
		return nil, err
	}
	user := &ApiUser{
		Login:     login,
		Password:  hashedPass,
		CanWrite:  canWrite,
		CreatedAt: time.Now(),
	}
	return user, err
}

func (u *ApiUser) CheckPass(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// TODO: need add choice token expiry from config or from REST after success auth user
func NewApiToken(user *ApiUser) *ApiToken {
	now := time.Now()
	return &ApiToken{
		Token:     GenToken(),
		ApiUser:   user,
		CreatedAt: now,
		ExpiryAt:  now.Add(DefaultExpiry),
	}
}

func (t *ApiToken) IsExpired() bool {
	return t.CreatedAt.Unix() >= t.ExpiryAt.Unix()
}

func HashPass(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPass), err
}

func GenToken() string {
	uuid := uuid.NewV4().String()
	token := strings.Replace(uuid, "-", "", -1)

	return token
}
