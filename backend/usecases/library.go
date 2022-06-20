package usecases

import (
	"github.com/rg-km/final-project-engineering-16/backend/app/middleware"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type LibraryAuthUsecase struct {
	Repo      domains.LibraryRepository
	TokenAuth middleware.JWTService
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
		return library, err
	}
	library.Token = a.TokenAuth.GenerateToken(library.Email, false, true)

	return library, nil
}
