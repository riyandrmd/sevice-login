package usecase

import (
	"net/http"
	"users/apperror"
	"users/authenticate"

	jwtinternal "users/util/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	authRepo authenticate.AuthRepository
}

func NewAuthUsecase(authRepo authenticate.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		authRepo,
	}
}

func (auth *AuthUseCase) Login(c *gin.Context) (*authenticate.AuthRespone, error) {
	var loginRequest *authenticate.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		return nil, apperror.New(http.StatusBadRequest, "Request Login Error")
	}

	data, err := auth.authRepo.GetDataLogin(loginRequest.Username)
	if err != nil {
		return nil, err
	}

	matching := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(loginRequest.Password))
	if matching != nil {
		return nil, apperror.New(http.StatusUnprocessableEntity, "Invalid Credential")
	}

	token, err := jwtinternal.CreateToken(int(data.ID))
	if err != nil {
		return nil, apperror.New(http.StatusBadRequest, "Can't Access JWT")
	}

	if err != nil {
		return nil, apperror.New(http.StatusUnprocessableEntity, "Generated Token Failed")
	}

	LoginResp := &authenticate.AuthRespone{
		Token:        token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return LoginResp, nil

}

func (auth *AuthUseCase) Register(c *gin.Context) error {
	var result authenticate.Users
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = auth.authRepo.CreateUser(&result)
	if err != nil {
		return err
	}

	return nil

}
