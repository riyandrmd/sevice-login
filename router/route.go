package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	authHandler "users/authenticate/handler"
	authRepo "users/authenticate/repository"
	authUC "users/authenticate/usecase"

	"users/middleware"
)

type Handlers struct {
	Ctx   context.Context
	DB    *gorm.DB
	R     *gin.Engine
	Redis *redis.Client
}

func (h *Handlers) Routes() {
	// Repository
	authRepo := authRepo.NewAuthRepository(h.DB)

	// Use Case
	AuthUseCase := authUC.NewAuthUsecase(authRepo)

	middleware.Add(h.R, middleware.CORSMiddleware())

	v1 := h.R.Group("rumahsakit")
	authHandler.AuthRoute(AuthUseCase, v1)
}
