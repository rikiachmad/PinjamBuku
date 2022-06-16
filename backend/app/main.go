package main

import (
	"database/sql"
	"log"

	"github.com/rg-km/final-project-engineering-16/backend/app/handler"
	"github.com/rg-km/final-project-engineering-16/backend/app/middleware"
	"github.com/rg-km/final-project-engineering-16/backend/app/routes"
	"github.com/rg-km/final-project-engineering-16/backend/infrastructures/repository"
	"github.com/rg-km/final-project-engineering-16/backend/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_ = godotenv.Load()
	router := gin.Default()

	db, err := sql.Open("sqlite3", "pinjambuku.db")
	if err != nil {
		panic(err)
	}

	tokenAuthService := middleware.JWTAuthService()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecases.NewAuthUsecase(userRepository, tokenAuthService)
	authController := handler.NewAuthController(userUsecase)
	authRoutes := routes.NewAuthRoutes(routes.RequestHandler{Gin: router}, authController)

	routes := routes.NewRoutes(authRoutes)
	routes.Setup()

	err = router.Run(":" + "8000")
	if err != nil {
		log.Fatal(err)
	}
}
