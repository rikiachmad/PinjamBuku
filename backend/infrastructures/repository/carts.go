package repository

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	domains "github.com/rg-km/final-project-engineering-16/backend/domains"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) domains.CartRepository {
	return &CartRepository{db: db}
}

func (u *CartRepository) FetchCartByID(id int64) (domains.Cart, error) {
	sqlStmt := `SELECT c.id, u.id, u.fullname, u.address, u.email, u.phone_number, u.verified_date, 
	b.id, b.title, b.author, b.page_number, b.deposit, b.cover
	FROM carts c 
	INNER JOIN users u ON c.user_id = u.id
	INNER JOIN books b ON c.book_id = b.id
	WHERE c.id = ?`

	cart := domains.Cart{}

	row := u.db.QueryRow(sqlStmt, id)
	err := row.Scan(
		&cart.ID,
		&cart.User.ID,
		&cart.User.Fullname,
		&cart.User.Address,
		&cart.User.Email,
		&cart.User.PhoneNumber,
		&cart.User.Verified,
		&cart.Book.ID,
		&cart.Book.Title,
		&cart.Book.Author,
		&cart.Book.PageNumber,
		&cart.Book.Deposit,
		&cart.Book.Cover,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domains.Cart{}, exceptions.ErrCartNotFound
		}
		return domains.Cart{}, err
	}

	return cart, nil
}

func (u *CartRepository) CheckCartByUserIDAndBookID(userID int64, bookID int64) (domains.Cart, error) {
	sqlStmt := `SELECT id, user_id, book_id FROM carts WHERE user_id = ? AND book_id = ?`

	cart := domains.Cart{}

	row := u.db.QueryRow(sqlStmt, userID, bookID)
	err := row.Scan(
		&cart.ID,
		&cart.UserID,
		&cart.BookID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domains.Cart{}, exceptions.ErrCartNotFound
		}
		return domains.Cart{}, err
	}

	return cart, nil
}

func (u *CartRepository) FetchCartByUserID(userID int64) ([]domains.Cart, error) {
	sqlStmt := `SELECT c.id, u.id, u.fullname, u.address, u.email, u.phone_number, u.verified_date, 
	b.id, b.title, b.author, b.page_number, b.deposit, b.cover
	FROM carts c 
	INNER JOIN users u ON c.user_id = u.id
	INNER JOIN books b ON c.book_id = b.id
	WHERE c.user_id = ?`

	carts := []domains.Cart{}

	rows, err := u.db.Query(sqlStmt, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cart := domains.Cart{}

		err := rows.Scan(
			&cart.ID,
			&cart.User.ID,
			&cart.User.Fullname,
			&cart.User.Address,
			&cart.User.Email,
			&cart.User.PhoneNumber,
			&cart.User.Verified,
			&cart.Book.ID,
			&cart.Book.Title,
			&cart.Book.Author,
			&cart.Book.PageNumber,
			&cart.Book.Deposit,
			&cart.Book.Cover,
		)

		if err != nil {
			return nil, err
		}

		carts = append(carts, cart)
	}

	return carts, nil
}

func (u *CartRepository) InsertToCart(userID int64, bookID int64) (domains.Cart, error) {
	sqlStmt := `INSERT INTO carts (user_id, book_id) VALUES (?, ?) RETURNING id, user_id, book_id`

	cart := domains.Cart{}

	err := u.db.QueryRow(sqlStmt, userID, bookID).Scan(
		&cart.ID,
		&cart.UserID,
		&cart.BookID,
	)
	if err != nil {
		return domains.Cart{}, err
	}

	return cart, nil
}

func (u *CartRepository) DeleteCartByID(id int64) error {
	sqlStmt := `DELETE FROM carts WHERE id = ?`

	_, err := u.db.Exec(sqlStmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *CartRepository) DeleteCartByUserID(userID int64) error {
	sqlStmt := `DELETE FROM carts WHERE user_id = ?`

	_, err := u.db.Exec(sqlStmt, userID)
	if err != nil {
		return err
	}

	return nil
}
