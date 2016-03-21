package common

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

const (
	FastExpiry    = 17 * time.Minute
	DefaultExpiry = 24 * time.Hour
	LongExpiry    = 24 * time.Hour * 7
)

type ApiToken struct {
	Token     string `gorm:"primary_key"`
	ApiUser   *ApiUser
	CreatedAt time.Time
	ExpiryAt  time.Time
}

func (t *ApiToken) TableName() string {
	return "api_tokens"
}

func (t *ApiToken) IsExpired() bool {
	return t.CreatedAt.Unix() >= t.ExpiryAt.Unix()
}

// TODO: need add choice token expiry from config or from REST with user auth
func NewApiToken(user *ApiUser) *ApiToken {
	now := time.Now()
	return &ApiToken{
		Token:     GenToken(),
		ApiUser:   user,
		CreatedAt: now,
		ExpiryAt:  now.Add(DefaultExpiry),
	}
}

func GenToken() string {
	uuid := uuid.NewV4().String()
	token := strings.Replace(uuid, "-", "", -1)

	return token
}

type ApiTokenGormRepository struct {
	DB *gorm.DB
}

func MakeApiTokenRepository(db *gorm.DB) ApiTokenRepository {
	db.AutoMigrate(&ApiToken{})

	return &ApiTokenGormRepository{DB: db}
}

func (r *ApiTokenGormRepository) FindByToken(token string) (*ApiToken, error) {
	apiToken := &ApiToken{}
	r.DB.First(apiToken, token)

	if r.DB.RecordNotFound() {
		return nil, r.DB.Error
	}
	return apiToken, r.DB.Error
}

func (r *ApiTokenGormRepository) FindOneByUser(userId int) (*ApiToken, error) {
	apiToken := &ApiToken{}
	apiUser := &ApiUser{ID: userId}
	r.DB.Model(apiToken).Related(apiUser)
	if r.DB.RecordNotFound() {
		return nil, r.DB.Error
	}
	return apiToken, r.DB.Error
}

func (r *ApiTokenGormRepository) Create(apiToken *ApiToken) error {
	tx := r.DB.Begin()

	if err := tx.Create(apiToken).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *ApiTokenGormRepository) UpdateExpiry(token string, expiry time.Time) error {
	tx := r.DB.Begin()

	if err := tx.Model(&ApiToken{Token: token}).Update("ExpiryAt", expiry).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *ApiTokenGormRepository) Delete(token string) error {
	tx := r.DB.Begin()

	if err := tx.Delete(&ApiToken{Token: token}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
