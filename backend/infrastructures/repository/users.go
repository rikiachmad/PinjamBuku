package repository

import (
	"database/sql"

	domains "github.com/rg-km/final-project-engineering-16/backend/domains"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domains.UserRepository{
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (domains.User, error) {
	sqlStmt := `SELECT 
	u.id,
	u.fullname,
	u.address,
	u.email,
	u.password,
	u.phone_number,
	u.verified_date,
	ur.name
	FROM users u INNER JOIN user_roles ur ON u.role_id = ur.id WHERE u.id = ?`

	user := domains.User{}

	row := u.db.QueryRow(sqlStmt, id)
	err := row.Scan(
		&user.ID,
		&user.Fullname,
		&user.Address,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Verified,
		&user.Role,
	)

	if err != nil {
		return domains.User{}, err
	}

	return user, nil
}

func (u *UserRepository) Login(email string, password string) (domains.User, error) {
	sqlStmt := `SELECT 
	u.id,
	u.fullname,
	u.address,
	u.email,
	u.password,
	u.phone_number,
	u.verified_date,
	ur.name
	FROM users u INNER JOIN user_roles ur ON u.role_id = ur.id WHERE u.email = ? AND u.password = ?`

	user := domains.User{}

	row := u.db.QueryRow(sqlStmt, email, password)
	err := row.Scan(
		&user.ID,
		&user.Fullname,
		&user.Address,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Verified,
		&user.Role,
	)

	if err != nil {
		return domains.User{}, err
	}

	return user, nil
}
