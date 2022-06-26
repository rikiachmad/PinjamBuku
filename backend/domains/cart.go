package domains

type Cart struct {
	ID        int64  `db:"id"`
	UserID    int64  `db:"user_id"`
	BookID    int64  `db:"book_id"`
	User      User   `db:"user"`
	Book      Book   `db:"book"`
	CreatedAt string `db:"created_at"`
	DeletedAt string `db:"deleted_at"`
}

type CartRepository interface {
	FetchCartByID(id int64) (Cart, error)
	FetchCartByUserID(userID int64) ([]Cart, error)
	CheckCartByUserIDAndBookID(userID int64, bookID int64) (Cart, error)
	InsertToCart(userID, bookID int64) (Cart, error)
	DeleteCartByID(id int64) error
	DeleteCartByUserID(userID int64) error
	DeleteCartByUserIDAndBookIDs(userID int64, bookIDs []int64) error
}

type CartUsecase interface {
	GetCartByID(id int64) (Cart, error)
	ShowCartByUserID(id int64) ([]Cart, error)
	InsertToCart(userID, bookID int64) (Cart, error)
	DeleteCartByID(id int64) error
}
