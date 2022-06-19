package domains

type Book struct {
	ID          int64  `db:"id"`
	CategoryID  int64  `db:"category_id"`
	LibraryID   int64  `db:"library_id"`
	KatalogID   string `db:"katalog_id"`
	Title       string `db:"title"`
	Author      string `db:"author"`
	PageNumber  int64  `db:"page_number"`
	Stock       int64  `db:"stock"`
	Description string `db:"description"`
	Deposit     int64  `db:"deposit"`
	CoverPict   string `db:"cover"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

type BookRepository interface {
	FetchBookByID(id int64) (Book, error)
	FetchBookByLibraryID(userID int64) ([]Book, error)
	InsertBook(userID, bookID int64) (Book, error)
}

type BookUsecase interface {
	FetchBookByID(id int64) (Book, error)
	FetchBookByLibraryID(userID int64) ([]Book, error)
	InsertBook(userID, bookID int64) (Book, error)
}
