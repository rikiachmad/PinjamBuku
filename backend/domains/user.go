package domains

type User struct {
	ID          int64  `db:"id"`
	Fullname    string `db:"fullname"`
	Address     string `db:"address"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Verified    string `db:"verified_date"`
	Role        string `db:"role"`
	Token       string `db:"token"`
	Photo       string `db:"picture_profile"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

type UserRepository interface {
	FetchUserByID(id int64) (User, error)
	Login(email string, password string) (User, error)
}

type UserUsecase interface {
	Login(user User) (User, error)
}
