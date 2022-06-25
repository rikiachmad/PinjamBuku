package repository

import (
	"database/sql"
	"time"

	exceptions "github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	helpers "github.com/rg-km/final-project-engineering-16/backend/commons/helpers"
	domains "github.com/rg-km/final-project-engineering-16/backend/domains"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domains.UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (domains.User, error) {
	sqlStmt := `SELECT 
	u.id,
	u.fullname,
	u.address,
	u.email,
	u.phone_number,
	u.picture_profile 
	FROM users u WHERE u.id = ?`

	user := domains.User{}

	row := u.db.QueryRow(sqlStmt, id)
	err := row.Scan(
		&user.ID,
		&user.Fullname,
		&user.Address,
		&user.Email,
		&user.PhoneNumber,
		&user.Photo,
	)

	if err != nil {
		return domains.User{}, err
	}

	return user, nil
}

func (u *UserRepository) Login(email string, password string) (domains.User, error) {
	sqlStmt := `SELECT id, fullname, email, password FROM users WHERE email = ?`

	user := domains.User{}

	row := u.db.QueryRow(sqlStmt, email, password)
	err := row.Scan(
		&user.ID,
		&user.Fullname,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domains.User{}, exceptions.ErrInvalidCredentials
		}
		return domains.User{}, err
	}

	if !helpers.IsMatched(user.Password, password) {
		return domains.User{}, exceptions.ErrInvalidCredentials
	}

	return user, nil
}

func (u *UserRepository) Create(fullname, email, password, address, phoneNumber string, role int) (domains.User, error) {
	var user domains.User
	passEncrypt, err := helpers.HashPassword(password)

	if err != nil {
		return domains.User{}, err
	}

	sqlStmt := `
		INSERT INTO users(fullname, address, email, password, verified_date, role_id, phone_number, picture_profile, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) 
		RETURNING id, fullname, address, email, phone_number, verified_date, role_id, picture_profile, created_at, updated_at
	`

	err = u.db.QueryRow(sqlStmt, fullname, address, email, passEncrypt, "", role, phoneNumber, "", time.Now(), time.Now()).Scan(
		&user.ID,
		&user.Fullname,
		&user.Address,
		&user.Email,
		&user.Verified,
		&user.Role,
		&user.PhoneNumber,
		&user.Photo,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return domains.User{}, err
	}

	return user, nil
}

func (u *UserRepository) CheckAccountEmail(email string) bool {
	var res string
	sqlSmt := `SELECT email FROM users WHERE email = ?`

	err := u.db.QueryRow(sqlSmt, email).Scan(&res)
	if err != nil {
		if err != sql.ErrNoRows {
			return false
		}
	}
	return res == email
}

func (u *UserRepository) UpdateUserProfile(id int64, fullname, address, phoneNumber, photo string) (domains.UpdateUser, error) {
	sqlstmt := `UPDATE users
					SET fullname = ?,
						address = ?,
						phone_number = ?,
						picture_profile = ?
					WHERE id = ?`
	_, err := u.db.Exec(sqlstmt, fullname, address, phoneNumber, photo, id)

	if err != nil {
		return domains.UpdateUser{}, err
	}

	user := domains.UpdateUser{}
	user.ID = id
	user.Fullname = fullname
	user.Address = address
	user.PhoneNumber = phoneNumber
	user.Photo = photo

	return user, nil
}
