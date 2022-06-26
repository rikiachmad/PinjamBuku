package repository

import (
	"database/sql"
	"time"

	domains "github.com/rg-km/final-project-engineering-16/backend/domains"
)

type BorrowingRepository struct {
	db *sql.DB
}

func NewBorrowingRepository(db *sql.DB) domains.BorrowingRepository {
	return &BorrowingRepository{db: db}
}

func (u *BorrowingRepository) FetchBorrowingByID(id int64) (domains.Borrowing, error) {
	sqlStmt := `SELECT 
		bb.id, bb.user_id, bb.status_id, bb.total_deposit, bb.total_cost, bb.borrowing_date, bb.due_date, bb.finish_date, 
		bs.id, bs.status,
		u.id, u.email
	FROM book_borrowing bb
	INNER JOIN borrowing_status bs ON bb.status_id = bs.id
	INNER JOIN users u ON bb.user_id = u.id
	WHERE bb.id = ?`

	borrowing := domains.Borrowing{}

	row := u.db.QueryRow(sqlStmt, id)
	err := row.Scan(
		&borrowing.ID,
		&borrowing.UserID,
		&borrowing.StatusID,
		&borrowing.TotalDeposit,
		&borrowing.TotalCost,
		&borrowing.BorrowingDate,
		&borrowing.DueDate,
		&borrowing.FinishDate,
		&borrowing.Status.ID,
		&borrowing.Status.Status,
		&borrowing.User.ID,
		&borrowing.User.Email,
	)

	if err != nil {
		return domains.Borrowing{}, err
	}

	return borrowing, nil
}

func (u *BorrowingRepository) FetchBookListByBorrowingID(borrowingID int64) ([]domains.Book, error) {
	sqlStmt := `SELECT 
		b.id, b.title, b.author, b.description, b.cover, b.page_number, b.stock, b.deposit,
		l.name, l.address
	FROM book_borrowing_list bbl
	INNER JOIN books b ON bbl.book_id = b.id
	INNER JOIN libraries l ON b.library_id = l.id
	WHERE borrowing_id = ?`

	book_borrowing := []domains.Book{}

	rows, err := u.db.Query(sqlStmt, borrowingID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := domains.Book{}

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.Cover,
			&book.PageNumber,
			&book.Stock,
			&book.Deposit,
			&book.LibraryName,
			&book.LibraryAddress,
		)

		if err != nil {
			return nil, err
		}

		book_borrowing = append(book_borrowing, book)
	}

	return book_borrowing, nil
}

func (u *BorrowingRepository) FetchBorrowingByUserID(userID int64) ([]domains.Borrowing, error) {
	sqlStmt := `SELECT 
		bb.id, bb.user_id, bb.status_id, bb.total_cost, bb.borrowing_date, bb.due_date, bb.finish_date, 
		bs.status,
		b.id, b.title, b.stock, b.deposit,
		l.id, l.name, l.address
	FROM book_borrowing bb
	INNER JOIN book_borrowing_list bbl ON bb.id = bbl.borrowing_id
	INNER JOIN borrowing_status bs ON bb.status_id = bs.id
	INNER JOIN books b ON bbl.book_id = b.id
	INNER JOIN libraries l ON b.library_id = l.id
	WHERE bb.user_id = ?`

	book_borrowing := []domains.Borrowing{}

	rows, err := u.db.Query(sqlStmt, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		borrowing := domains.Borrowing{}

		err := rows.Scan(
			&borrowing.ID,
			&borrowing.UserID,
			&borrowing.StatusID,
			&borrowing.TotalCost,
			&borrowing.BorrowingDate,
			&borrowing.DueDate,
			&borrowing.FinishDate,
			&borrowing.Status,
			&borrowing.Book.ID,
			&borrowing.Book.Title,
			&borrowing.Book.Stock,
			&borrowing.Book.Deposit,
			&borrowing.Library.ID,
			&borrowing.Library.Name,
			&borrowing.Library.Address,
		)

		if err != nil {
			return nil, err
		}

		book_borrowing = append(book_borrowing, borrowing)
	}

	return book_borrowing, nil
}

func (u *BorrowingRepository) InsertToBorrowing(userID int64, bookID []int64, totalDeposit int64, totalCost int64) (domains.Borrowing, error) {
	sqlStmt := `INSERT INTO book_borrowing
	(user_id, status_id, total_deposit, total_cost, borrowing_date, due_date, finish_date, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	RETURNING id, user_id, status_id, total_deposit, total_cost, borrowing_date, due_date, finish_date, created_at, updated_at`

	borrowing := domains.Borrowing{}

	err := u.db.QueryRow(sqlStmt, userID, 1, totalDeposit, totalCost, time.Now(), time.Now().AddDate(0, 0, 21), time.Time{}, time.Now(), time.Now()).Scan(
		&borrowing.ID,
		&borrowing.UserID,
		&borrowing.StatusID,
		&borrowing.TotalDeposit,
		&borrowing.TotalCost,
		&borrowing.BorrowingDate,
		&borrowing.DueDate,
		&borrowing.FinishDate,
		&borrowing.CreatedAt,
		&borrowing.UpdatedAt,
	)

	if err != nil {
		return domains.Borrowing{}, err
	}

	sqlStmt = `INSERT INTO book_borrowing_list
	(borrowing_id, book_id, created_at)
	VALUES (?, ?, ?)`

	for _, bookID := range bookID {
		_, err := u.db.Exec(sqlStmt, borrowing.ID, bookID, time.Now())
		if err != nil {
			return domains.Borrowing{}, err
		}
	}

	return borrowing, nil
}

func (u *BorrowingRepository) DeleteBorrowingByID(id int64) error {
	sqlStmt := `DELETE FROM book_borrowing_list WHERE borrowing_id = ?`

	_, err := u.db.Exec(sqlStmt, id)
	if err != nil {
		return err
	}

	sqlStmt = `DELETE FROM book_borrowing WHERE id = ?`

	_, err = u.db.Exec(sqlStmt, id)
	if err != nil {
		return err
	}

	return nil
}
