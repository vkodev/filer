package rest_api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/vkodev/filer/common"
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

func HandleAuth(tokenRepository common.ApiTokenRepository, userRepository common.ApiUserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		login := c.Form("login")
		password := c.Form("password")

		apiUser, err := userRepository.FindByLogin(login)

		var authResponse AuthResponse
		var status int

		// TODO: need elegant solution for error handling
		if err != nil {
			authResponse.SetError("user not found")
			status = http.StatusNotFound
		} else if !apiUser.CheckPass(password) {
			authResponse.SetError("password not compare")
			status = http.StatusBadRequest
		} else {
			apiToken := common.NewApiToken(apiUser)
			err := tokenRepository.Create(apiToken)

			if err != nil {
				authResponse.SetError("token not created")
				status = http.StatusBadRequest
			} else {
				authResponse.SetToken(apiToken.Token)
				status = http.StatusOK
			}
		}

		return c.JSON(status, authResponse)
	}
}
