package usecases

import (
	"log"
	"strings"

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

func (b BookUsecase) FetchSearchBook(by, words string) ([]domains.Book, error) {

	books, err := b.Repo.GetSearchByTitle(words)

	if err != nil {
		log.Printf("error usecases-book FetchSearchBook %s", err)
		return []domains.Book{}, err
	}

	return books, nil

}

func (b BookUsecase) FetchSort(key string) ([]domains.Book, error) {

	if strings.EqualFold("asc", key) || strings.EqualFold("desc", key) {
		books, err := b.Repo.GetSort(key)

		if err != nil {
			log.Printf("error usecases-book FetchSearchBook %s", err)
			return []domains.Book{}, err
		}

		return books, nil
	}

	return []domains.Book{}, exceptions.ErrBadRequest
}
