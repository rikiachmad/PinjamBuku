package domains

type Book struct {
	ID             int64  `db:"id"`
	KatalogId      string `db:"katalog_id"`
	Title          string `db:"title"`
	Author         string `db:"author"`
	Description    string `db:"description"`
	Cover          string `db:"cover"`
	PageNumber     int64  `db:"page_number"`
	Stock          int64  `db:"stock"`
	Deposit        int64  `db:"deposit"`
	CategoryId     int64  `db:"category_id"`
	CategoryName   string `db:"category_name"`
	LibraryId      int64  `db:"library_id"`
	LibraryName    string `db:"library_name"`
	LibraryAddress string `db:"library_address"`
	IsPublish      bool   `db:"is_publish"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
}

type CreateBook struct {
	KatalogId   string `db:"katalog_id"`
	Title       string `db:"title"`
	Author      string `db:"author"`
	Description string `db:"description"`
	Cover       string `db:"cover"`
	PageNumber  int64  `db:"page_number"`
	Stock       int64  `db:"stock"`
	Deposit     int64  `db:"deposit"`
	CategoryId  int64  `db:"category_id"`
	LibraryId   int64  `db:"library_id"`
	IsPublish   bool   `db:"is_publish"`
}

type BookRepository interface {
	GetAll() ([]Book, error)
	GetById(id int64) (Book, error)
	GetSearchByTitle(words string) ([]Book, error)
	GetSort(key string) ([]Book, error)
}

type BookUsecase interface {
	FetchAll() ([]Book, error)
	FetchById(book Book) (Book, error)
	FetchSearchBook(by, words string) ([]Book, error)
	FetchSort(key string) ([]Book, error)
}
