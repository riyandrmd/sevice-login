package handler

import (
	"net/http"
	"users/authenticate"
	"users/response"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authUc authenticate.AuthUseCase
}

func AuthRoute(authUc authenticate.AuthUseCase, v1 *gin.RouterGroup) {
	uc := authHandler{
		authUc,
	}

	v2 := v1.Group("users")
	v2.POST("/login", uc.Login)
	v2.POST("/register", uc.Register)
}

func (authHandler *authHandler) Login(c *gin.Context) {
	token, err := authHandler.authUc.Login(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Status:  "success",
		Message: "success",
		Meta:    token,
	})
}

func (authHandler *authHandler) Register(c *gin.Context) {
	err := authHandler.authUc.Register(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: "success ",
	})
}
