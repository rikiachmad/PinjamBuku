package usecases

import (
	"fmt"

	"github.com/rg-km/final-project-engineering-16/backend/app/middleware"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type AuthUsecase struct {
	Repo domains.UserRepository
	TokenAuth middleware.JWTService
}

func NewAuthUsecase(repo domains.UserRepository, tokenAuth middleware.JWTService) AuthUsecase {
	return AuthUsecase{
		Repo: repo,
		TokenAuth: tokenAuth,
	}
}

func (a AuthUsecase) Login(user domains.User) (domains.User, error) {
	fmt.Printf("user:\n%+v\n", user)
	user, err := a.Repo.Login(user.Email, user.Password)
	if err != nil {
		return user, err
	}
	if user.Role == "admin" {
		user.Token = a.TokenAuth.GenerateToken(user.Email, true, false)
	} else {
		user.Token = a.TokenAuth.GenerateToken(user.Email, false, false)
	}
	return user, nil
}

func (a AuthUsecase) Register(user domains.CreateUser) (domains.User, error) {
	isEmailExist := a.Repo.CheckAccountEmail(user.Email)
	if isEmailExist {
		fmt.Printf("%s\nEmail %+v is already exist\n", user.Email, isEmailExist)
		return domains.User{}, exceptions.ErrUserAlreadyExists
	}

	createdUser, err := a.Repo.Create(user.Fullname, user.Email, user.Password, user.Address, user.PhoneNumber, 1)
	if err != nil {
		return createdUser, err
	}
	return createdUser, nil
}