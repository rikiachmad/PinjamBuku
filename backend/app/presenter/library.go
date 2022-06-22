package presenter

import "github.com/rg-km/final-project-engineering-16/backend/domains"

type Library struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phone_number"`
	Photo         string `json:"picture_profile"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankName      string `json:"bank_name"`
}

type UpdateLibrary struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"picture_profile"`
}

func FetchLibraryDefault(l domains.Library) Library {
	return Library{
		ID:            l.ID,
		Name:          l.Name,
		Email:         l.Email,
		Address:       l.Address,
		PhoneNumber:   l.PhoneNumber,
		Photo:         l.Photo,
		AccountNumber: l.AccountNumber,
		AccountName:   l.AccountName,
		BankName:      l.BankName,
	}
}

func FetchUpdateLibrary(l domains.UpdateLibrary) UpdateLibrary {
	return UpdateLibrary{
		ID:          l.ID,
		Name:        l.Name,
		Address:     l.Address,
		PhoneNumber: l.PhoneNumber,
		Photo:       l.Photo,
	}
}
