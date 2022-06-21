package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/handler"
	"github.com/rg-km/final-project-engineering-16/backend/infrastructures/repository"
	"github.com/rg-km/final-project-engineering-16/backend/usecases"
)

func InitRoutesCart(db *sql.DB, route *gin.Engine) {
	cartRepository := repository.NewCartRepository(db)
	cartUsecase := usecases.NewCartUsecase(cartRepository)
	cartController := handler.NewCartController(cartUsecase)

	apiv1 := route.Group("/api/v1")
	{
		cart := apiv1.Group("/cart")
		{
			cart.GET("/", cartController.ShowCartByUserID)
			cart.POST("/", cartController.InsertToCart)
			cart.DELETE("/:id", cartController.DeleteCartByID)
		}

	}
}
