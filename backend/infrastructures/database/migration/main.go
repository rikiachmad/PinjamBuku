package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "./pinjambuku.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            fullname VARCHAR(255) NOT NULL,
            address VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL,
            password VARCHAR(255) NOT NULL,
			is_verified DATETIME,
			role_id INTEGER,
            phone_number VARCHAR(255) NOT NULL,
			FOREIGN KEY(role_id) REFERENCES user_roles(id)
        );
    `)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS user_roles (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(10) NOT NULL
        );
    `)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO user_roles (name) VALUES ("ADMIN"), ("USER");`)

	if err != nil {
		panic(err)
	}
}

func Rollback(db *sql.DB) {
	sqlStmt := `DROP TABLE users;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE user_roles;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
