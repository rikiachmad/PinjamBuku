package presenter

import (
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type Login struct {
	Email          string     `json:"email"`
	Fullname       string     `json:"fullname"`
	Token          string     `json:"token,omitempty"`
}

func LoginFromDomain(u domains.User) Login {
	return Login{
		Email:          u.Email,
		Fullname:       u.Fullname,
		Token:          u.Token,
	}
}

type CreateUser struct {
	Email          string     `json:"email"`
	Fullname       string     `json:"fullname"`
	Address		string     `json:"address"`
	PhoneNumber	string     `json:"phoneNumber"`
	Role		string		`json:"role"`
	Photo		string     `json:"photo"`
	CreatedAt	string     `json:"createdAt"`
	UpdatedAt	string     `json:"updatedAt"`
}

func CreateUserFromDomain(u domains.User) CreateUser {
	return CreateUser{
		Email:          u.Email,
		Fullname:       u.Fullname,
		Address:		u.Address,
		PhoneNumber:	u.PhoneNumber,
		Role:			u.Role,
		Photo:			u.Photo,
		CreatedAt:		u.CreatedAt,
		UpdatedAt:		u.UpdatedAt,
	}
}
