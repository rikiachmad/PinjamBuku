package usecases

import (
	"fmt"

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

func (c CartUsecase) GetCartByID(id int64) (domains.Cart, error) {
	cart, err := c.Repo.FetchCartByID(id)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (c CartUsecase) ShowCartByUserID(userID int64) ([]domains.Cart, error) {
	if userID == 0 {
		return nil, exceptions.ErrUnauthorized
	}
	carts, err := c.Repo.FetchCartByUserID(userID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", carts)
	return carts, nil
}

func (c CartUsecase) InsertToCart(userID, bookID int64) (domains.Cart, error) {
	if userID == 0 || bookID == 0 {
		return domains.Cart{}, exceptions.ErrBadRequest
	}

	cart, err := c.Repo.CheckCartByUserIDAndBookID(userID, bookID)
	if err != nil {
		if err == exceptions.ErrCartNotFound {
			cart, err = c.Repo.InsertToCart(userID, bookID)
			if err != nil {
				return domains.Cart{}, err
			}
			return cart, nil
		}
		return domains.Cart{}, err
	}
	return cart, exceptions.ErrCartAlreadyExists
}

func (c CartUsecase) DeleteCartByID(id int64) error {
	if id == 0 {
		return exceptions.ErrBadRequest
	}

	cart, err := c.Repo.FetchCartByID(id)
	if err != nil {
		return err
	}
	if cart == (domains.Cart{}) || cart.ID == 0 {
		return exceptions.ErrCartNotFound
	}

	err = c.Repo.DeleteCartByID(id)
	if err != nil {
		return err
	}
	return nil
}