package domains

type BorrowingStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type Borrowing struct {
	ID            int64           `db:"id"`
	UserID        int64           `db:"user_id"`
	StatusID      int64           `db:"status_id"`
	User          User            `db:"user"`
	Book          Book            `db:"book"`
	Library       Library         `db:"library"`
	TotalDeposit  int64           `db:"total_deposit"`
	TotalCost     int64           `db:"total_cost"`
	BorrowingDate string          `db:"borrowing_date"`
	DueDate       string          `db:"due_date"`
	FinishDate    string          `db:"finish_date"`
	Status        BorrowingStatus `db:"status"`
	CreatedAt     string          `db:"created_at"`
	UpdatedAt     string          `db:"updated_at"`
}

type BorrowingWithBook struct {
	Borrowing Borrowing `json:"borrowing"`
	Books     []Book    `json:"books"`
}

type BorrowingRepository interface {
	FetchBorrowingByID(id int64) (Borrowing, error)
	FetchBorrowingByUserID(userID int64) ([]Borrowing, error)
	FetchBookListByBorrowingID(borrowingID int64) ([]Book, error)
	InsertToBorrowing(userID int64, bookID []int64, totalDeposit int64, totalCost int64) (Borrowing, error)
	DeleteBorrowingByID(id int64) error
}

type BorrowingUsecase interface {
	ShowBorrowingByUserID(id int64) ([]Borrowing, error)
	InsertToBorrowing(userID int64, cartIDs []int64, totalCost int64) (BorrowingWithBook, error)
	DeleteBorrowingByID(id int64) error
}
