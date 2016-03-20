package rest_api

import (
	"github.com/vkodev/filer/models"
	"github.com/labstack/echo"
	"net/http"
)

const (
	Success = "success"
	Error   = "error"
)

type AuthResponse struct {
	Status string
	Token  string
	Error  string
}

func (a *AuthResponse) SetToken(token string) {
	a.Status = Success
	a.Token = token
}

func (a *AuthResponse) SetError(err string) {
	a.Status = Error
	a.Error = err
}

func HandleAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		login := c.Form("login")
		password := c.Form("password")

		// TODO: need replace with real UserRepository
		userRepo := NewTestMockRepository()

		apiUser := userRepo.FindByLogin(login)

		var authResponse AuthResponse
		var status int

		if apiUser == nil {
			authResponse.SetError("user not found")
			status = http.StatusNotFound
		} else if !apiUser.CheckPass(password) {
			authResponse.SetError("password not compare")
			status = http.StatusBadRequest
		} else {
			// TODO: need save this token in BD
			apiToken := models.NewApiToken(apiUser)
			authResponse.SetToken(apiToken.Token)
			status = http.StatusOK
		}

		return c.JSON(status, authResponse)
	}
}

// This repository only for test auth
type MockUserRepository struct {
	ApiUser *models.ApiUser
}

func NewTestMockRepository() *MockUserRepository {
	testUser, _ := models.NewApiUser("test", "123rfv", true)
	return &MockUserRepository{ApiUser: testUser}
}

func (r *MockUserRepository) FindByLogin(login string) *models.ApiUser {
	var user *models.ApiUser = nil

	if r.ApiUser.Login == login {
		user = r.ApiUser
	}

	return user
}
