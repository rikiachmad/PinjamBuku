package handler

import (
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

func (cc CartController) InsertToCart(c *gin.Context) {
	req := Cart{}
	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}
	
	domain := req.ToCartDomain()
	res, err := cc.cartUsecase.InsertToCart(domain.UserID, domain.BookID)
	resFromDomain := presenter.InsertCartFromDomain(res)
	if err != nil {
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}

	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}