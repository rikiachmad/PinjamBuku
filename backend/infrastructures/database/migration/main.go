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
			verified_date DATETIME,
			role_id INTEGER NOT NULL,
            phone_number VARCHAR(255) NOT NULL,
			picture_profile VARCHAR(255),
			no_ktp VARCHAR(255),
			picture_ktp VARCHAR(255),
			created_at DATETIME,
			updated_at DATETIME,
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

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS libraries (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			address VARCHAR(255) NOT NULL,
			phone_number VARCHAR(255) NOT NULL,
			picture_profile VARCHAR(255) NOT NULL,
			account_id INTEGER NOT NULL,
			created_at DATETIME,
			updated_at DATETIME,
			FOREIGN KEY(account_id) REFERENCES bank_accounts(id)
		);
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS bank_accounts (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(256) NOT NULL,
			number VARCHAR(256) NOT NULL,
			bank_name VARCHAR(256) NOT NULL
        );
    `)

	if err != nil {
		panic(err)
	}

	// i create for dummy data
	_, err = db.Exec(`
		INSERT INTO libraries (name, email, password, address, phone_number, picture_profile, account_id)
		VALUES("PERPUS KAB. TANGERANG", "perpuskabtang@gmail.com", "123", "Tangerang", "0210324234", "perpuskabtang.jpg", 1),
		("PERPUS KAB. BANDUNG", "perpuskabbandungng@gmail.com", "123", "Bandung", "0210324234", "perpuskabbandung.jpg", 1)
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			cover BLOB NOT NULL,
			page_number INTEGER NOT NULL,
			stock INTEGER NOT NULL,
			deposit INTEGER NOT NULL,
			category_id INTEGER NOT NULL,
			library_id INTEGER NOT NULL,
			is_publish BOOLEAN NOT NULL,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME,
		FOREIGN KEY (category_id) REFERENCES book_categories(id)
		FOREIGN KEY (library_id) REFERENCES libraries(id)
		)
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS book_categories (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL
		)
	`)

	if err != nil {
		panic(err)
	}

	// For Testing
	// _, err = db.Exec(`INSERT INTO book_categories (name) VALUES ("EDUCATION"), ("ROMANTIC"), ("FIKSI")`)

	// if err != nil {
	// 	panic(err)
	// }
	// For Testing
	// _, err = db.Exec(`
	// 	INSERT INTO books (title, author, description, cover, page_number, stock, deposit, category_id, library_id, is_publish)
	// 	VALUES ("Berdamai dengan Emosi", "Asti Musman", "Ragam Emosi, Gangguan Emosi, Emosi dan Ekspresi Wajah", "berdamai.jpg", 120, 40, 30000, 1, 1, true ),
	// 	("Paket komplit having fun 7in one", "Tim Penerbit", "Paket buku ini disusun agar kita sebagai orang tua.", "7in.jpg", 190, 2, 70000, 2, 2, true )
	// `)

	// if err != nil {
	// 	panic(err)
	// }
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

	sqlStmt = `DROP TABLE libraries;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE bank_accounts;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE books;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE book_categories;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
