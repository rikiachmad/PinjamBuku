package presenter

import (
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type Login struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Token    string `json:"token,omitempty"`
}

type LoginLibrary struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token,omitempty"`
}

func LoginFromDomain(u domains.User) Login {
	return Login{
		Email:    u.Email,
		Fullname: u.Fullname,
		Token:    u.Token,
	}
}

func LoginLibraryFromDomain(l domains.Library) LoginLibrary {
	return LoginLibrary{
		Email: l.Email,
		Name:  l.Name,
		Token: l.Token,
	}
}

type CreateUser struct {
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Photo       string `json:"photo"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func CreateUserFromDomain(u domains.User) CreateUser {
	return CreateUser{
		Email:       u.Email,
		Fullname:    u.Fullname,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		Role:        u.Role,
		Photo:       u.Photo,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
