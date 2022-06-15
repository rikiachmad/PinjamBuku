package domains

type User struct {
	ID          int64  `db:"id"`
	Fullname    string `db:"fullname"`
	Address     string `db:"address"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Verified    string `db:"is_verified"`
	Role        string `db:"role"`
	Photo       string `db:"picture_profile"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
