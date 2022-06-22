package repository

import (
	"database/sql"
	"time"

	domains "github.com/rg-km/final-project-engineering-16/backend/domains"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) domains.BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Add(title, author, description, cover string, pageNumber, stock, deposit, categoryId, libraryId int64) (domains.Book, error) {
	var book domains.Book
	sqlStmt := `
		INSERT INTO books(title, author, description, cover, page_number, stock, deposit, category_id, library_id, is_publish, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ? ,? ,? ,?, ?, ?)
	`

	_, err := b.db.Exec(sqlStmt, title, author, description, cover, pageNumber, stock, deposit, categoryId, libraryId, true, time.Now(), time.Now())

	if err != nil {
		return domains.Book{}, err
	}

	return book, nil
}

func (b *BookRepository) Update(title, author, description, cover string, pageNumber, stock, deposit, categoryId, id int64) (domains.Book, error) {
	var book domains.Book

	sqlStmt := `
		UPDATE FROM books
			SET title = ?,
			SET author = ?,
			SET description = ?,
			SET cover = ?,
			SET page_number = ?,
			SET stock = ?,
			SET deposit = ?,
			SET category_id = ?,
			SET updated_at = ?
		WHERE id = ?
	`

	err := b.db.QueryRow(sqlStmt, title, author, description, cover, pageNumber, stock, deposit, categoryId, time.Now(), id).Scan(
		&book.Title,
		&book.Author,
		&book.Description,
		&book.Cover,
		&book.PageNumber,
		&book.Stock,
		&book.Deposit,
		&book.CategoryName,
	)

	if err != nil {
		return domains.Book{}, err
	}

	return book, nil

}

func (b *BookRepository) GetById(id int64) (domains.Book, error) {
	var book domains.Book

	sqlStmt := `
		SELECT
		 b.id,
		 b.title,
		 b.author,
		 b.description,
		 b.cover,
		 b.page_number,
		 b.stock,
		 b.deposit,
		 bc.name as category_name,
		 l.name as library_name,
		 l.address as library_address
		FROM books b
		INNER JOIN book_categories bc ON b.category_id = bc.id
		INNER JOIN libraries l ON b.library_id = l.id
		WHERE b.id = ?
	`

	err := b.db.QueryRow(sqlStmt, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Description,
		&book.Cover,
		&book.PageNumber,
		&book.Stock,
		&book.Deposit,
		&book.CategoryName,
		&book.LibraryName,
		&book.LibraryAddress,
	)

	if err != nil {
		return domains.Book{}, err
	}

	return book, nil
}

func (b *BookRepository) GetAll() ([]domains.Book, error) {
	var books = []domains.Book{}

	sqlStmt := `
	SELECT 
	 b.id, 
	 b.title, 
	 b.author, 
	 b.description, 
	 b.cover, 
	 b.page_number, 
	 b.stock, 
	 b.deposit, 
	 bc.name as category_name, 
	 l.name as library_name
	FROM books b 
	INNER JOIN book_categories bc ON b.category_id = bc.id
	INNER JOIN libraries l ON b.library_id = l.id
	`

	rows, err := b.db.Query(sqlStmt)

	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book domains.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.Cover,
			&book.PageNumber,
			&book.Stock,
			&book.Deposit,
			&book.CategoryName,
			&book.LibraryName,
		); err != nil {
			return books, err
		}

		books = append(books, book)
	}

	if err != nil {
		return books, err
	}

	return books, nil
}

func (b *BookRepository) GetSearchByTitle(words string) ([]domains.Book, error) {
	var books []domains.Book

	sqlStmt := `
		SELECT 
		b.id, 
		b.title, 
		b.author, 
		b.description, 
		b.cover, 
		b.page_number, 
		b.stock, 
		b.deposit, 
		bc.name as category_name, 
		l.name as library_name
		FROM books b 
		INNER JOIN book_categories bc ON b.category_id = bc.id
		INNER JOIN libraries l ON b.library_id = l.id
		WHERE 
			b.title LIKE ? 
	`
	likeWords := `%` + words + `%`
	rows, err := b.db.Query(sqlStmt, likeWords)

	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book domains.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.Cover,
			&book.PageNumber,
			&book.Stock,
			&book.Deposit,
			&book.CategoryName,
			&book.LibraryName,
		); err != nil {
			return books, err
		}

		books = append(books, book)
	}

	if err != nil {
		return books, err
	}

	return books, nil

}

func (b *BookRepository) GetSort(key string) ([]domains.Book, error) {
	var books []domains.Book

	sqlStmt := `
		SELECT 
			b.id, 
			b.title,
			b.author,
			b.description, 
			b.cover,
			b.page_number,
			b.stock, 
			b.deposit, 
			bc.name as category_name, 
			l.name as library_name
		FROM books b 
		INNER JOIN book_categories bc ON b.category_id = bc.id
		INNER JOIN libraries l ON b.library_id = l.id
		ORDER BY b.title
	` + key
	rows, err := b.db.Query(sqlStmt)

	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book domains.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.Cover,
			&book.PageNumber,
			&book.Stock,
			&book.Deposit,
			&book.CategoryName,
			&book.LibraryName,
		); err != nil {
			return books, err
		}

		books = append(books, book)
	}

	if err != nil {
		return books, err
	}

	return books, nil
}
