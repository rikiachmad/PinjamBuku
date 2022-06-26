package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-16/backend/app/middleware"
	"github.com/rg-km/final-project-engineering-16/backend/app/routes"
)

func main() {
	_ = godotenv.Load()
	router := gin.Default()

	db, err := sql.Open("sqlite3", "backend/infrastructures/database/migration/pinjambuku.db")

	if err != nil {
		panic(err)
	}

	router.Use(middleware.CORSMiddleware())
	routes.InitRoutesAuth(db, router)
	routes.InitRoutesBook(db, router)
	routes.InitRoutesCart(db, router)
	routes.InitRoutesLibrary(db, router)
	routes.InitRoutesUser(db, router)

	err = router.Run(":" + "8080")
	if err != nil {
		log.Fatal(err)
	}
}
