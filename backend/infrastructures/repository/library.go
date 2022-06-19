package repository

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type LibraryRepository struct {
	db *sql.DB
}

func NewLibraryRepository(db *sql.DB) LibraryRepository {
	return LibraryRepository{
		db: db,
	}
}

func (l *LibraryRepository) Login(email string, password string) (domains.Library, error) {
	sqlstmt := `SELECT name, email, password FROM libaries WHERE email = ?`

	library := domains.Library{}

	row := l.db.QueryRow(sqlstmt, email)
	err := row.Scan(
		&library.Name,
		&library.Email,
		&library.Password,
	)

	if err != nil {
		return domains.Library{}, err
	}

	// for temporary password in db will not be encrypted
	// to make it easy for testing purpose because the library is generated manually by the admin.
	if library.Password != password {
		return domains.Library{}, exceptions.ErrInvalidCredentials
	}

	// empty the password in sake of security
	library.Password = ""

	return library, nil
}
