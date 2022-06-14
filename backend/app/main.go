package main

import (
	"database/sql"

	repository "github.com/rg-km/final-project-engineering-16/backend/infrastructures/repository"
)


func main(){
	db, err := sql.Open("sqlite3", "backend/infrastructures/database/pinjambuku.db")
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)
}
