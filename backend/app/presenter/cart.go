package presenter

import (
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type InsertCart struct {
	ID     int64 `json:"id"`
	BookID int64 `json:"bookId"`
	UserID int64 `json:"userId"`
}

func InsertCartFromDomain(c domains.Cart) InsertCart {
	return InsertCart{
		ID:     c.ID,
		BookID: c.BookID,
		UserID: c.UserID,
	}
}

type Cart struct {
	ID        int64      `json:"id"`
	BookID    int64      `json:"bookId,omitempty"`
	UserID    int64      `json:"userId,omitempty"`
	User      CreateUser `json:"user"`
	Book      Book       `json:"book"`
	CreatedAt string     `json:"createdAt,omitempty"`
	DeletedAt string     `json:"deletedAt,omitempty"`
}

func CartFromDomain(u domains.Cart) Cart {
	return Cart{
		ID:        u.ID,
		UserID:    u.UserID,
		BookID:    u.BookID,
		User:      CreateUserFromDomain(u.User),
		Book:      FetchBookDefault(u.Book),
		CreatedAt: u.CreatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func CartListFromDomain(u []domains.Cart) []Cart {
	var carts []Cart
	for _, v := range u {
		carts = append(carts, CartFromDomain(v))
	}
	return carts
}
