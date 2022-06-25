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

type CreateUser struct {
	Fullname    string `db:"fullname"`
	Address     string `db:"address"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Verified    string `db:"verified_date"`
	Role        int    `db:"role"`
	Photo       string `db:"picture_profile"`
}

type UpdateUser struct {
	ID          int64  `db:"id"`
	Fullname    string `db:"fullname"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
	Photo       string `db:"picture_profile"`
}

type UserRepository interface {
	FetchUserByID(id int64) (User, error)
	Login(email string, password string) (User, error)
	Create(fullname, email, password, address, phoneNumber string, role int) (User, error)
	CheckAccountEmail(email string) bool
	UpdateUserProfile(id int64, fullname, address, phoneNumber, photo string) (UpdateUser, error)
}

type AuthUsecase interface {
	Login(user User) (User, error)
	Register(user CreateUser) (User, error)
}

type UserUsecase interface {
	UpdateUserProfile(user UpdateUser, id int64) (UpdateUser, error)
	FetchUserByID(id int64) (User, error)
}
