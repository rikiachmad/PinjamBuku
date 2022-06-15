package usecases

import (
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type UserUsecase struct {
	Repo domains.UserRepository
}

func NewUserUsecase(repo domains.UserRepository) domains.UserUsecase {
	return &UserUsecase{
		Repo:           repo,
	}
}

func (u *UserUsecase) Login(user domains.User) (domains.User, error) {
	return u.Repo.Login(user.Email, user.Password)
}