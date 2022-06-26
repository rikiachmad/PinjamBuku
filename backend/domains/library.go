package domains

type Library struct {
	ID            int64  `db:"id"`
	Name          string `db:"name"`
	Email         string `db:"email"`
	Password      string `db:"password"`
	Address       string `db:"address"`
	PhoneNumber   string `db:"phone_number"`
	Photo         string `db:"picture_profile"`
	AccountNumber string `db:"account_number"`
	AccountName   string `db:"account_name"`
	BankName      string `db:"bank_name"`
	Token         string `db:"token"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}

type UpdateLibrary struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
	Photo       string `db:"picture_profile"`
}

type CreateLibrary struct {
	Name          string `db:"name"`
	Email         string `db:"email"`
	Password      string `db:"password"`
	Address       string `db:"address"`
	PhoneNumber   string `db:"phone_number"`
	Photo         string `db:"picture_profile"`
	AccountNumber string `db:"account_number"`
	AccountName   string `db:"account_name"`
	BankName      string `db:"bank_name"`
}

type LibraryRepository interface {
	Login(email string, password string) (Library, error)
	GetAllLibrary() ([]Library, error)
	GetLibraryByID(id int64) (Library, error)
	UpdateLibraryProfileByID(id int64, name, address, phoneNumber, photo string) (UpdateLibrary, error)
	CheckExistLibrary(id int64) bool
	GetAllBookById(id int64) ([]Book, error)
	CreateBook(katalogId, title, author, description, cover string, pageNumber, stock, deposit, categoryId, libraryId int64) error
	UpdateBook(katalogId, title, author, description, cover string, pageNumber, stock, deposit, categoryId, id, libraryId int64) error
	CheckBook(id int64) bool
}

type LibraryAuthUsecase interface {
	Login(library Library) (Library, error)
}

type LibraryUsecase interface {
	GetAllLibrary() ([]Library, error)
	GetLibraryByID(id int64) (Library, error)
	UpdateLibraryProfileByID(library UpdateLibrary, id int64) (UpdateLibrary, error)
	GetAllBookById(id int64) ([]Book, error)
	CreateBook(book Book, id int64) error
	UpdateBook(book Book, id int64) error
}
