package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/handler"
	"github.com/rg-km/final-project-engineering-16/backend/app/middleware"
	"github.com/rg-km/final-project-engineering-16/backend/infrastructures/repository"
	"github.com/rg-km/final-project-engineering-16/backend/usecases"
)

func InitRoutesUser(db *sql.DB, route *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := handler.NewUserController(&userUsecase)

	apiv1 := route.Group("/api/v1")
	{
		apiv1.Use(middleware.AuthorizeJWT(), middleware.AuthMiddleware("user"))
		user := apiv1.Group("/user")
		{
			user.GET("/:id", middleware.ValidateIDMiddleware(), userController.GetUserDetail)
		}
		{
			user.PUT("/:id", middleware.ValidateIDMiddleware(), userController.UpdateUserProfile)
		}
	}

}
