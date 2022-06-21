package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/presenter"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type Cart struct {
	BookID int64 `json:"book_id" form:"book_id"`
	UserID int64 `json:"user_id" form:"user_id"`
}

func (r Cart) ToCartDomain() domains.Cart {
	return domains.Cart{
		BookID: r.BookID,
		UserID: r.UserID,
	}
}

type CartController struct {
	cartUsecase domains.CartUsecase
}

func NewCartController(cartUsecase domains.CartUsecase) CartController {
	return CartController{
		cartUsecase: cartUsecase,
	}
}

func (cc CartController) ShowCartByUserID(c *gin.Context) {
	req := Cart{}
	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}
	
	fmt.Printf("userID: %d\n", req)
	domain := req.ToCartDomain()
	fmt.Printf("userID: %d\n", domain.UserID)
	res, err := cc.cartUsecase.ShowCartByUserID(domain.UserID)
	if err != nil {
		if err == exceptions.ErrBadRequest {
			presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
			return
		} else if err == exceptions.ErrUnauthorized {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			return
		}
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}
	resFromDomain := presenter.CartListFromDomain(res)

	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (cc CartController) InsertToCart(c *gin.Context) {
	req := Cart{}
	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}
	
	domain := req.ToCartDomain()
	res, err := cc.cartUsecase.InsertToCart(domain.UserID, domain.BookID)
	if err != nil {
		if err == exceptions.ErrBadRequest {
			presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
			return
		}
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}
	resFromDomain := presenter.InsertCartFromDomain(res)

	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}