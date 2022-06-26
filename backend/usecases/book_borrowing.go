package usecases

import (
	"fmt"

	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type BorrowingUsecase struct {
	BorrowingRepo domains.BorrowingRepository
	BookRepo      domains.BookRepository
	CartRepo 	  domains.CartRepository
}

func NewBorrowingUsecase(borrowingRepo domains.BorrowingRepository, bookRepo domains.BookRepository, cartRepo domains.CartRepository) BorrowingUsecase {
	return BorrowingUsecase{
		BorrowingRepo: borrowingRepo,
		BookRepo:      bookRepo,
		CartRepo: 	   cartRepo,
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

func (c BorrowingUsecase) InsertToBorrowing(userID int64, cartIDs []int64, totalCost int64) (domains.BorrowingWithBook, error) {
	if userID == 0 {
		return domains.BorrowingWithBook{}, exceptions.ErrBadRequest
	}

	bookIDs := []int64{}
	for _, cartID := range cartIDs {
		cart, err := c.CartRepo.FetchCartByID(cartID)
		if err != nil {
			return domains.BorrowingWithBook{}, err
		}
		bookIDs = append(bookIDs, cart.Book.ID)
	}
	fmt.Println("bookIDs", bookIDs)

	totalDeposit, err := c.BookRepo.GetTotalDeposit(bookIDs)
	if err != nil {
		fmt.Println("err", err)
		return domains.BorrowingWithBook{}, err
	}
	fmt.Printf("totalDeposit: %d\n", totalDeposit)

	borrowing, err := c.BorrowingRepo.InsertToBorrowing(userID, bookIDs, totalDeposit, totalCost)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return domains.BorrowingWithBook{}, err
	}

	borrowing, err = c.BorrowingRepo.FetchBorrowingByID(borrowing.ID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return domains.BorrowingWithBook{}, err
	}

	bookBorrowingList, err := c.BorrowingRepo.FetchBookListByBorrowingID(borrowing.ID)
	if err != nil {
		return domains.BorrowingWithBook{}, err
	}

	err = c.CartRepo.DeleteCartByUserIDAndBookIDs(userID, bookIDs)
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