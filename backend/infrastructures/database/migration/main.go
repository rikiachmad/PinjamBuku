package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "backend/infrastructures/database/migration/pinjambuku.db")
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
		CREATE TABLE IF NOT EXISTS book_categories (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL,
			created_at DATETIME,
			updated_at DATETIME
		);
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			katalog_id VARCHAR(10) NOT NULL,
			library_id INTEGER NOT NULL,
			category_id INTEGER NOT NULL,
			title VARCHAR(512) NOT NULL,
			author VARCHAR(255) NOT NULL,
			page_number INT NOT NULL,
			stock INT NOT NULL,
			description TEXT NOT NULL,
			deposit BIGINT NOT NULL,
			cover VARCHAR(1024) NOT NULL,
			is_publish BOOLEAN NOT NULL,
			created_at DATETIME,
			updated_at DATETIME,
			FOREIGN KEY(library_id) REFERENCES libraries(id),
			FOREIGN KEY(category_id) REFERENCES book_categories(id)
		);
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS carts (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			book_id INTEGER NOT NULL,
			created_at DATETIME,
			deleted_at DATETIME,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(book_id) REFERENCES books(id)
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

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS book_categories (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL
		)
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		INSERT INTO book_categories (name) 
		VALUES 
		("Umum"), ("Filsafat dan Psikologi"), ("Agama"), ("Sosial"), ("Bahasa"), ("Sains dan Matematika"), ("Teknologi"), ("Seni dan Rekreasi"), ("Literatur dan Sastra"), ("Sejarah dan Geografi")
	`)

	if err != nil {
		panic(err)
	}

	//for testing
	_, err = db.Exec(`
	INSERT INTO bank_accounts (name, number, bank_name)
	VALUES("PERPUS SBY", "144410101", "MANDIRI"), ("PERPUS SDA", "1444101012", "BCA")
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	INSERT INTO libraries (name, email, password, address, phone_number, picture_profile, account_id, created_at, updated_at)
	VALUES 
	("Perpus SBY", "perpussby@gmail.com", "123", "Surabaya", "123123", "sby.jpg", 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	("Perpus SDA", "perpussda@gmail.com", "123", "Sidoarjo", "123123", "sda.jpg", 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	INSERT INTO books (katalog_id, library_id, category_id, title, author, page_number, stock, description, deposit, cover, is_publish, created_at, updated_at) VALUES 
	('AB001', '1', '1', 'Ayat-Ayat Cinta', 'Habiburrahman', '241', '10', 'ini deskripsi ayat ayat cinta', '30000', 'ayat.jpg', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('AB002', '2', '2', 'Laskar Pelangi', 'Budi', '150', '30', 'ini deksripsi laskar', '45000', 'laskar.jpg', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('AB003', '1', '1', 'Heartbreak Motel', 'Wawan', '250', '40', 'ini deskripsi heartbreak', '120000', 'heart.jpg', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('AB004', '2', '2', 'Sagaras', 'Tere Liye', '285', '89', 'ini deskripsi sagaras', '150000', 'sagaras.jpg', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('AB005', '1', '1', 'Nebula', 'Wati', '201', '30', 'ini deskripsi nebula', '98000', 'nebula.jpg', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`)

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

	sqlStmt = `DROP TABLE book_categories;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE books;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE carts;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
