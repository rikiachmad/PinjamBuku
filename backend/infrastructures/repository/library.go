package repository

import (
	"database/sql"

	exceptions "github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type LibraryRepository struct {
	db *sql.DB
}

func NewLibraryRepository(db *sql.DB) domains.LibraryRepository {
	return &LibraryRepository{db: db}
}

func (l *LibraryRepository) Login(email string, password string) (domains.Library, error) {
	sqlstmt := `SELECT id, name, email, password FROM libraries WHERE email = ?`

	library := domains.Library{}

	row := l.db.QueryRow(sqlstmt, email)
	err := row.Scan(
		&library.ID,
		&library.Name,
		&library.Email,
		&library.Password,
	)

	if err != nil {
		return domains.Library{}, err
	}

	// if !helpers.IsMatched(library.Password, password) {
	// 	return domains.Library{}, exceptions.ErrInvalidCredentials
	// }

	if library.Password != password {
		return domains.Library{}, exceptions.ErrInvalidCredentials
	}

	// empty the password in sake of security
	library.Password = ""

	return library, nil
}

func (l *LibraryRepository) GetAllLibrary() ([]domains.Library, error) {
	var libraries = []domains.Library{}

	sqlStmt := `
	SELECT l.id, 
	l.name, 
	l.email, 
	l.address, 
	l.phone_number, 
	l.picture_profile, 
	ba.number, 
	ba.name, 
	ba.bank_name 
	FROM libraries l
	INNER JOIN bank_accounts ba ON l.account_id = ba.id
	`

	rows, err := l.db.Query(sqlStmt)

	if err != nil {
		return libraries, err
	}

	defer rows.Close()

	for rows.Next() {
		var library domains.Library

		if err := rows.Scan(
			&library.ID,
			&library.Name,
			&library.Email,
			&library.Address,
			&library.PhoneNumber,
			&library.Photo,
			&library.AccountNumber,
			&library.AccountName,
			&library.BankName,
		); err != nil {
			return libraries, err
		}

		libraries = append(libraries, library)
	}

	if err != nil {
		return libraries, err
	}

	return libraries, nil
}

func (l *LibraryRepository) GetLibraryByID(id int64) (domains.Library, error) {
	sqlstmt := `SELECT l.id, l.name, l.email, l.address, l.phone_number, l.picture_profile, ba.number, ba.name, ba.bank_name FROM libraries l
				INNER JOIN bank_accounts ba ON l.account_id = ba.id WHERE l.id = ?`

	library := domains.Library{}
	row := l.db.QueryRow(sqlstmt, id)
	err := row.Scan(
		&library.ID,
		&library.Name,
		&library.Email,
		&library.Address,
		&library.PhoneNumber,
		&library.Photo,
		&library.AccountNumber,
		&library.AccountName,
		&library.BankName,
	)

	if err != nil {
		return domains.Library{}, err
	}

	return library, nil
}

func (l *LibraryRepository) UpdateLibraryProfileByID(id int64, name, address, phoneNumber, photo string) (domains.UpdateLibrary, error) {
	sqlstmt := `UPDATE libraries
					SET name = ?,
						address = ?,
						phone_number = ?,
						picture_profile = ?
					WHERE id = ?`
	_, err := l.db.Exec(sqlstmt, name, address, phoneNumber, photo, id)

	if err != nil {
		return domains.UpdateLibrary{}, err
	}

	library := domains.UpdateLibrary{}
	library.ID = id
	library.Name = name
	library.Address = address
	library.PhoneNumber = phoneNumber
	library.Photo = photo

	return library, nil
}

func (l *LibraryRepository) CheckExistLibrary(id int64) bool {
	var res int64
	sqlStmt := `SELECT id FROM libraries WHERE id = ?`

	err := l.db.QueryRow(sqlStmt, id).Scan(&res)

	if err != nil {
		if err != sql.ErrNoRows {
			return false
		}
	}
	return res == id
}
