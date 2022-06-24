package usecases

import (
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type BorrowingUsecase struct {
	BorrowingRepo domains.BorrowingRepository
	BookRepo      domains.BookRepository
}

func NewBorrowingUsecase(borrowingRepo domains.BorrowingRepository, bookRepo domains.BookRepository) BorrowingUsecase {
	return BorrowingUsecase{
		BorrowingRepo: borrowingRepo,
		BookRepo:      bookRepo,
	}
}

func (c BorrowingUsecase) ShowBorrowingByUserID(userID int64) ([]domains.Borrowing, error) {
	if userID == 0 {
		return nil, exceptions.ErrUnauthorized
	}
	carts, err := c.BorrowingRepo.FetchBorrowingByUserID(userID)
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (c BorrowingUsecase) InsertToBorrowing(userID int64, bookIDs []int64, totalCost int64) (domains.BorrowingWithBook, error) {
	if userID == 0 {
		return domains.BorrowingWithBook{}, exceptions.ErrBadRequest
	}

	totalDeposit := int64(0)

	for _, id := range bookIDs {
		book, err := c.BookRepo.GetById(id)
		if id == 0 || err != nil || book.Stock == 0 {
			return domains.BorrowingWithBook{}, exceptions.ErrBadRequest
		}
		totalDeposit += book.Deposit
	}

	borrowing, err := c.BorrowingRepo.InsertToBorrowing(userID, bookIDs, totalDeposit, totalCost)
	if err != nil {
		return domains.BorrowingWithBook{}, err
	}

	bookBorrowingList, err := c.BorrowingRepo.FetchBookListByBorrowingID(borrowing.ID)
	if err != nil {
		return domains.BorrowingWithBook{}, err
	}

	return domains.BorrowingWithBook{
		Borrowing: borrowing,
		Books:     bookBorrowingList,
	}, nil
}

func (c BorrowingUsecase) DeleteBorrowingByID(id int64) error {
	if id == 0 {
		return exceptions.ErrBadRequest
	}
	err := c.BorrowingRepo.DeleteBorrowingByID(id)
	if err != nil {
		return err
	}
	return nil
}