package presenter

import (
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type Auth struct {
	Email          string     `json:"email"`
	Fullname       string     `json:"fullname"`
	Token          string     `json:"token,omitempty"`
}

func AuthFromDomain(u domains.User) Auth {
	return Auth{
		Email:          u.Email,
		Fullname:       u.Fullname,
		Token:          u.Token,
	}
}