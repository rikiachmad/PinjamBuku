package usecases

import (
	"log"

	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type BookUsecase struct {
	Repo domains.BookRepository
}

func NewBookUsecase(repo domains.BookRepository) BookUsecase {
	return BookUsecase{
		Repo: repo,
	}
}

func (b BookUsecase) FetchById(book domains.Book) (domains.Book, error) {
	books, err := b.Repo.GetById(book.ID)

	if err != nil {
		log.Printf("error usecases-book FetchById %s", err)
		return domains.Book{}, exceptions.ErrBadRequest
	}

	return books, nil
}

func (b BookUsecase) FetchAll() ([]domains.Book, error) {
	books, err := b.Repo.GetAll()

	if err != nil {
		log.Printf("error usecases-book FetchAll %s", err)
		return []domains.Book{}, exceptions.ErrInternalServerError
	}

	return books, nil
}

func (b BookUsecase) Insert(book domains.CreateBook) (domains.Book, error) {
	books, err := b.Repo.Add(book.Title, book.Author, book.Description, book.Cover, book.PageNumber, book.Stock, book.Deposit, book.CategoryId, book.LibraryId)

	if err != nil {
		log.Printf("error usecases-book Insert %s", err)
		return domains.Book{}, exceptions.ErrInternalServerError
	}

	return books, nil
}

func (b BookUsecase) FetchSearchBook(by, words string) ([]domains.Book, error) {

	books, err := b.Repo.GetSearchByTitle(words)

	if err != nil {
		log.Printf("error usecases-book FetchSearchBook %s", err)
		return []domains.Book{}, err
	}

	return books, nil

}
