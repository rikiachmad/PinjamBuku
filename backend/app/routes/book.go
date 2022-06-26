package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/handler"
	"github.com/rg-km/final-project-engineering-16/backend/infrastructures/repository"
	"github.com/rg-km/final-project-engineering-16/backend/usecases"
)

func InitRoutesBook(db *sql.DB, route *gin.Engine) {
	// tokenAuthService := middleware.JWTAuthService()
	bookRepository := repository.NewBookRepository(db)
	bookUsecase := usecases.NewBookUsecase(bookRepository)
	bookController := handler.NewBookController(bookUsecase)

	apiv1 := route.Group("/api/v1")
	{
		auth := apiv1.Group("/book")
		{
			auth.GET("/", bookController.GetAllBook)
		}
		{
			auth.GET("/:id", bookController.GetBookById)
		}
		{
			auth.GET("/search", bookController.GetSearchBook)
		}
	}
}
