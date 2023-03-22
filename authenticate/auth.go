package authenticate

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var AuthSecret = "rumah-sakit"
var RefreshAuthSecret = "rumah-sakit-refresh"
var From = "http://localhost:8081"

type Users struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenDetail struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
}

type AuthRespone struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTClaims struct {
	jwt.StandardClaims
	UID   int      `json:"uid"`
	Group []string `json:"group"`
}

type AuthRepository interface {
	GetDataLogin(string) (*Users, error)
	CreateUser(*Users) error
}

type AuthUseCase interface {
	Login(*gin.Context) (*AuthRespone, error)
	Register(*gin.Context) error
}
