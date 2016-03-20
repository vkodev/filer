package common

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
)

type ApiUser struct {
	ID        int     	`gorm:"primary_key"`
	Login     string
	Password  string
	CanWrite  bool
	CreatedAt time.Time
}

func (u *ApiUser) TableName() string {
	return "api_users"
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

func HashPass(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPass), err
}

type ApiUserGormRepository struct {
	DB gorm.DB
}

func MakeApiUserGormRepository(db *gorm.DB) ApiUserRepository {
	db.AutoMigrate(&ApiUser{})

	return &ApiUserGormRepository{DB:db}
}

func (r *ApiUserGormRepository) FindById(id int) (*ApiUser, error) {
	apiUser := &ApiUser{}
	if err := r.DB.First(id).Error; err != nil {
		return nil, err
	}

	return apiUser, nil
}

func (r *ApiUserGormRepository) FindByLogin(login string) (*ApiUser, error) {
	apiUser := &ApiUser{}
	if err := r.DB.Where("login = ?", login).First(&apiUser).Error; err != nil {
		return nil, err
	}

	return apiUser, nil
}

func (r *ApiUserGormRepository) FindAll() (apiUsers []*ApiUser, error) {
	if err := r.DB.Find(apiUsers).Error; err != nil {
		return nil, err
	}

	return apiUsers, nil
}

func (r *ApiUserGormRepository) Create(apiUser *ApiUser) error {
	tx := r.DB.Begin()

	if err := tx.Create(apiUser).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *ApiUserGormRepository) Update(apiUser *ApiUser) error {
	tx := r.DB.Begin()

	if err := tx.Update(apiUser).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *ApiUserGormRepository) Delete(id int) error {
	tx := r.DB.Begin()

	if err := tx.Delete(&ApiUser{ID:id}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}