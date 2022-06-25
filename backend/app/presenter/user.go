package presenter

import "github.com/rg-km/final-project-engineering-16/backend/domains"

type UpdateUser struct {
	ID          int64  `json:"id"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"picture_profile"`
}

type User struct {
	ID          int64  `json:"id"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"picture_profile"`
}

func FetchUpdateUser(u domains.UpdateUser) UpdateUser {
	return UpdateUser{
		ID:          u.ID,
		Fullname:    u.Fullname,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		Photo:       u.Photo,
	}
}

func FetchUserDetail(u domains.User) User {
	return User{
		ID:          u.ID,
		Fullname:    u.Fullname,
		Address:     u.Address,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Photo:       u.Photo,
	}
}
