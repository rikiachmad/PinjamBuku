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

func (l *LibraryRepository) CreateBook(katalogId, title, author, description, cover string, pageNumber, stock, deposit, categoryId, libraryId int64) error {
	// var book domains.Book
	sqlStmt := `
		INSERT INTO books(katalog_id, title, author, description, cover, page_number, stock, deposit, category_id, library_id, is_publish, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ? ,? ,? ,?, ?, ?, ?)
	`

	_, err := l.db.Exec(sqlStmt, katalogId, title, author, description, cover, pageNumber, stock, deposit, categoryId, libraryId, true, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (l *LibraryRepository) UpdateBook(katalogId, title, author, description, cover string, pageNumber, stock, deposit, categoryId, id, libraryId int64) error {
	// var res FMResultSet
	sqlStmt := `
		UPDATE books
			SET katalog_id = ?,
				 title = ?,
				 author = ?,
				 description = ?,
				 cover = ?,
				 page_number = ?,
				 stock = ?,
				 deposit = ?,
				 category_id = ?,
				 updated_at = ?
		WHERE id = ? AND library_id = ?
	`
	_, err := l.db.Exec(sqlStmt, katalogId, title, author, description, cover, pageNumber, stock, deposit, categoryId, time.Now(), id, libraryId)

	if err != nil {
		return err
	}
	return nil

}

func (b *BookRepository) GetById(id int64) (domains.Book, error) {
	var book domains.Book

	sqlStmt := `
		SELECT
		 b.id,
		 b.katalog_id,
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
		&book.KatalogId,
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
	 b.katalog_id, 
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
			&book.KatalogId,
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
		b.katalog_id,
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
			&book.KatalogId,
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
			b.katalog_id,
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
			&book.KatalogId,
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

func (l *LibraryRepository) GetAllBookById(id int64) ([]domains.Book, error) {
	var books []domains.Book

	sqlStmt := `
		SELECT
		 b.id,
		 b.title,
		 b.katalog_id,
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
		WHERE l.id = ? 
	`

	rows, err := l.db.Query(sqlStmt, id)
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var book domains.Book
		if err := rows.Scan(
			&book.ID,
			&book.KatalogId,
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
		); err != nil {
			return books, err
		}

		books = append(books, book)
	}

	// if err != nil {
	// 	return books, err
	// }
	// log.Printf("query test => %v %v", books, book)
	return books, nil

}

func (l LibraryRepository) CheckBook(id int64) bool {
	var res int64
	sqlSmt := `SELECT id FROM books WHERE id = ?`

	err := l.db.QueryRow(sqlSmt, id).Scan(&res)
	if err != nil {
		if err != sql.ErrNoRows {
			return false
		}
	}
	return res == id
}
