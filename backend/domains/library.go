package domains

type library struct {
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
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
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
