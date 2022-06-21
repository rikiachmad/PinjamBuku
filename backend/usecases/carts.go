package usecases

import (
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type CartUsecase struct {
	Repo domains.CartRepository
}

func NewCartUsecase(repo domains.CartRepository) CartUsecase {
	return CartUsecase{
		Repo: repo,
	}
}

func (c CartUsecase) ShowCartByUserID(userID int64) ([]domains.Cart, error) {
	carts, err := c.Repo.FetchCartByUserID(userID)
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (c CartUsecase) InsertToCart(userID, bookID int64) (domains.Cart, error) {
	if userID == 0 || bookID == 0 {
		return domains.Cart{}, exceptions.ErrBadRequest
	}
	cart, err := c.Repo.InsertToCart(userID, bookID)
	if err != nil {
		return domains.Cart{}, err
	}
	return cart, nil
}