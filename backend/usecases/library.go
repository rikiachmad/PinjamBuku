package usecases

import (
	"github.com/rg-km/final-project-engineering-16/backend/app/middleware"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type LibraryAuthUsecase struct {
	Repo      domains.LibraryRepository
	TokenAuth middleware.JWTService
}

type LibraryUsecase struct {
	Repo domains.LibraryRepository
}

func NewLibraryAuthUsecase(repo domains.LibraryRepository, tokenAuth middleware.JWTService) LibraryAuthUsecase {
	return LibraryAuthUsecase{
		Repo:      repo,
		TokenAuth: tokenAuth,
	}
}

func (a *LibraryAuthUsecase) Login(library domains.Library) (domains.Library, error) {
	library, err := a.Repo.Login(library.Email, library.Password)
	if err != nil {
		return domains.Library{}, err
	}
	library.Token = a.TokenAuth.GenerateToken(library.ID, library.Email, false, true)

	return library, nil
}

func NewLibraryUsecase(repo domains.LibraryRepository) LibraryUsecase {
	return LibraryUsecase{
		Repo: repo,
	}
}

func (l *LibraryUsecase) GetAllLibrary() ([]domains.Library, error) {
	libraries, err := l.Repo.GetAllLibrary()

	if err != nil {
		return []domains.Library{}, err
	}

	return libraries, nil
}

func (l *LibraryUsecase) GetLibraryByID(id int64) (domains.Library, error) {
	library, err := l.Repo.GetLibraryByID(id)
	if err != nil {
		return domains.Library{}, err
	}

	return library, nil
}

func (l *LibraryUsecase) UpdateLibraryProfileByID(library domains.UpdateLibrary, id int64) (domains.UpdateLibrary, error) {
	library, err := l.Repo.UpdateLibraryProfileByID(id, library.Name, library.Address, library.PhoneNumber, library.Photo)
	if err != nil {
		return domains.UpdateLibrary{}, nil
	}

	return library, nil
}

func (l *LibraryUsecase) GetAllBookById(id int64) ([]domains.Book, error) {
	ok := l.Repo.CheckExistLibrary(id)

	if ok {
		books, err := l.Repo.GetAllBookById(id)
		if err != nil {
			return []domains.Book{}, exceptions.ErrBadRequest
		}
		return books, nil
	}

	return []domains.Book{}, exceptions.ErrBadRequest
}

func (l *LibraryUsecase) CreateBook(b domains.Book, idLibrary int64) error {
	err := l.Repo.CreateBook(b.KatalogId, b.Title, b.Author, b.Description, b.Cover, b.PageNumber, b.Stock, b.Deposit, b.CategoryId, idLibrary)

	if err != nil {
		return exceptions.ErrBadRequest
	}

	return nil
}

func (l *LibraryUsecase) UpdateBook(b domains.Book, idLibrary int64) error {
	ok := l.Repo.CheckBook(b.ID)

	if ok {
		err := l.Repo.UpdateBook(b.KatalogId, b.Title, b.Author, b.Description, b.Cover, b.PageNumber, b.Stock, b.Deposit, b.CategoryId, b.ID, idLibrary)

		if err != nil {
			return exceptions.ErrBadRequest
		}
		return nil
	}

	return exceptions.ErrBadRequest
}
