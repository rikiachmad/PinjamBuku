package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "backend/infrastructures/database/pinjambuku.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id integer not null primary key AUTOINCREMENT,
            fullname varchar(255) not null,
            address varchar(255) not null,
            email varchar(255) not null,
            password varchar(255) not null,
            phone_number varchar(255) not null,
            verified datetime not null
        );
    `)

	if err != nil {
		panic(err)
	}
}
