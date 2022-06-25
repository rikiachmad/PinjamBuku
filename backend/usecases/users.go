package usecases

import (
	"github.com/rg-km/final-project-engineering-16/backend/app/middleware"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type AuthUsecase struct {
	Repo      domains.UserRepository
	TokenAuth middleware.JWTService
}

type UserUsecase struct {
	Repo domains.UserRepository
}

func NewAuthUsecase(repo domains.UserRepository, tokenAuth middleware.JWTService) AuthUsecase {
	return AuthUsecase{
		Repo:      repo,
		TokenAuth: tokenAuth,
	}
}

func NewUserUsecase(repo domains.UserRepository) UserUsecase {
	return UserUsecase{
		Repo: repo,
	}
}

func (a AuthUsecase) Login(user domains.User) (domains.User, error) {
	user, err := a.Repo.Login(user.Email, user.Password)
	if err != nil {
		return user, err
	}
	if user.Role == "admin" {
		user.Token = a.TokenAuth.GenerateToken(user.ID, user.Email, true, false)
	} else {
		user.Token = a.TokenAuth.GenerateToken(user.ID, user.Email, false, false)
	}
	return user, nil
}

func (a AuthUsecase) Register(user domains.CreateUser) (domains.User, error) {
	isEmailExist := a.Repo.CheckAccountEmail(user.Email)
	if isEmailExist {
		return domains.User{}, exceptions.ErrUserAlreadyExists
	}

	createdUser, err := a.Repo.Create(user.Fullname, user.Email, user.Password, user.Address, user.PhoneNumber, 1)
	if err != nil {
		return createdUser, err
	}
	return createdUser, nil
}

func (u *UserUsecase) UpdateUserProfile(user domains.UpdateUser, id int64) (domains.UpdateUser, error) {
	_, err := u.Repo.FetchUserByID(id)
	if err != nil {
		return domains.UpdateUser{}, exceptions.ErrUserNotFound
	}

	userReturn, err := u.Repo.UpdateUserProfile(id, user.Fullname, user.Address, user.PhoneNumber, user.Photo)

	if err != nil {
		return domains.UpdateUser{}, err
	}

	return userReturn, nil
}

func (u *UserUsecase) FetchUserByID(id int64) (domains.User, error) {
	user, err := u.Repo.FetchUserByID(id)
	if err != nil {
		return domains.User{}, nil
	}

	return user, nil
}
