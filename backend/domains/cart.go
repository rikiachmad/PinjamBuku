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
	InsertToCart(userID, bookID int64) (Cart, error)
}

type CartUsecase interface {
	ShowCartByUserID(id int64) ([]Cart, error)
	InsertToCart(userID, bookID int64) (Cart, error)
}
