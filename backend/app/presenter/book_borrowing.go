package presenter

import (
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type InsertBorrowing struct {
	ID            int64      `json:"id"`
	UserID        int64      `json:"userId"`
	User          CreateUser `json:"user"`
	TotalDeposit  int64      `json:"totalDeposit"`
	TotalCost     int64      `json:"totalCost"`
	BorrowingDate string     `json:"borrowingDate"`
	DueDate       string     `json:"dueDate"`
	FinishDate    string     `json:"finishDate"`
	Status        string     `json:"status"`
	CreatedAt     string     `json:"createdAt,omitempty"`
	UpdatedAt     string     `json:"updatedAt,omitempty"`
}

func InsertBorrowingFromDomain(c domains.Borrowing) InsertBorrowing {
	return InsertBorrowing{
		ID:            c.ID,
		UserID:        c.UserID,
		User:          CreateUserFromDomain(c.User),
		TotalDeposit:  c.TotalDeposit,
		TotalCost:     c.TotalCost,
		BorrowingDate: c.BorrowingDate,
		DueDate:       c.DueDate,
		FinishDate:    c.FinishDate,
		Status:        StatusToString(c.Status),
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}
}

type BorrowingWithBook struct {
	Borrowing InsertBorrowing `json:"borrowing"`
	Books     []Book          `json:"books"`
}

type Borrowing struct {
	ID            int64      `json:"id"`
	BookID        int64      `json:"bookId,omitempty"`
	UserID        int64      `json:"userId,omitempty"`
	User          CreateUser `json:"user,omitempty"`
	Book          Book       `json:"book,omitempty"`
	Library       Library    `json:"library,omitempty"`
	TotalDeposit  int64      `json:"totalDeposit"`
	TotalCost     int64      `json:"totalCost"`
	BorrowingDate string     `json:"borrowingDate"`
	DueDate       string     `json:"dueDate"`
	FinishDate    string     `json:"finishDate"`
	Status        string     `json:"status"`
	CreatedAt     string     `json:"createdAt,omitempty"`
	UpdatedAt     string     `json:"updatedAt,omitempty"`
}

func BorrowingFromDomain(b domains.Borrowing) Borrowing {
	return Borrowing{
		ID:            b.ID,
		UserID:        b.UserID,
		User:          CreateUserFromDomain(b.User),
		Book:          FetchBookDefault(b.Book),
		Library:       FetchLibraryDefault(b.Library),
		TotalDeposit:  b.TotalDeposit,
		TotalCost:     b.TotalCost,
		BorrowingDate: b.BorrowingDate,
		DueDate:       b.DueDate,
		FinishDate:    b.FinishDate,
		Status:        StatusToString(b.Status),
		CreatedAt:     b.CreatedAt,
		UpdatedAt:     b.UpdatedAt,
	}
}

func BorrowingListFromDomain(b []domains.Borrowing) []Borrowing {
	var carts []Borrowing
	for _, v := range b {
		carts = append(carts, BorrowingFromDomain(v))
	}
	return carts
}

func BorrowingWithBookFromDomain(b domains.BorrowingWithBook) BorrowingWithBook {
	books := []Book{}
	for _, v := range b.Books {
		books = append(books, FetchBookDefault(v))
	}
	return BorrowingWithBook{
		Borrowing: InsertBorrowingFromDomain(b.Borrowing),
		Books:     books,
	}
}

func StatusToString(status domains.BorrowingStatus) string {
	return status.Status
}
