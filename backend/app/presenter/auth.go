package presenter

import (
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type Login struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Token    string `json:"token,omitempty"`
}

type LoginLibrary struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token,omitempty"`
}

func LoginFromDomain(u domains.User) Login {
	return Login{
		ID:       u.ID,
		Email:    u.Email,
		Fullname: u.Fullname,
		Token:    u.Token,
	}
}

func LoginLibraryFromDomain(l domains.Library) LoginLibrary {
	return LoginLibrary{
		ID:    l.ID,
		Email: l.Email,
		Name:  l.Name,
		Token: l.Token,
	}
}

type CreateUser struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Role        string `json:"role,omitempty"`
	Photo       string `json:"photo,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func CreateUserFromDomain(u domains.User) CreateUser {
	return CreateUser{
		ID:          u.ID,
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
