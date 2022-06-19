package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/handler"
	"github.com/rg-km/final-project-engineering-16/backend/app/middleware"
	"github.com/rg-km/final-project-engineering-16/backend/infrastructures/repository"
	"github.com/rg-km/final-project-engineering-16/backend/usecases"
)

func InitRoutesAuth(db *sql.DB, route *gin.Engine) {
	tokenAuthService := middleware.JWTAuthService()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecases.NewAuthUsecase(userRepository, tokenAuthService)
	authController := handler.NewAuthController(userUsecase)

	apiv1 := route.Group("/api/v1")
	{
		auth := apiv1.Group("/auth/")
		{
			auth.POST("/login", authController.SignIn)
			auth.POST("/register", authController.Register)
		}

	}
}
