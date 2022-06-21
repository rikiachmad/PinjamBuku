package repository

import (
	"database/sql"

	exceptions "github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	helpers "github.com/rg-km/final-project-engineering-16/backend/commons/helpers"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type LibraryRepository struct {
	db *sql.DB
}

func NewLibraryRepository(db *sql.DB) domains.LibraryRepository {
	return &LibraryRepository{db: db}
}

func (l *LibraryRepository) Login(email string, password string) (domains.Library, error) {
	sqlstmt := `SELECT name, email, password FROM libraries WHERE email = ?`

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

	if !helpers.IsMatched(library.Password, password) {
		return domains.Library{}, exceptions.ErrInvalidCredentials
	}

	// if library.Password != password {
	// 	return domains.Library{}, exceptions.ErrInvalidCredentials
	// }

	// empty the password in sake of security
	library.Password = ""

	return library, nil
}
