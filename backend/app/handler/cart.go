package handler

import (
	"net/http"
	"strconv"

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

func (cc CartController) GetCartByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	cart, err := cc.cartUsecase.GetCartByID(id)
	if err != nil {
		if err == exceptions.ErrBadRequest {
				presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
				return
			} else if err == exceptions.ErrUnauthorized {
				presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
				return
			} else if err == exceptions.ErrCartNotFound {
				presenter.ErrorResponse(c, http.StatusNotFound, exceptions.ErrCartNotFound)
				return
			}
			presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
			return
		}
	resFromDomain := presenter.CartFromDomain(cart)
	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (cc CartController) ShowCartByUserID(c *gin.Context) {
	userID := int64(c.Keys["id"].(float64))
	res, err := cc.cartUsecase.ShowCartByUserID(userID)
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
	
	userID := int64(c.Keys["id"].(float64))
	domain := req.ToCartDomain()
	res, err := cc.cartUsecase.InsertToCart(userID, domain.BookID)
	if err != nil {
		if err == exceptions.ErrBadRequest {
			presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
			return
		} else if err == exceptions.ErrUnauthorized {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			return
		} else if err == exceptions.ErrCartAlreadyExists {
			presenter.ErrorResponse(c, http.StatusConflict, exceptions.ErrCartAlreadyExists)
			return
		}
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}
	resFromDomain := presenter.InsertCartFromDomain(res)

	presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (cc CartController) DeleteCartByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}
	err = cc.cartUsecase.DeleteCartByID(intID)
	if err != nil {
		if err == exceptions.ErrBadRequest {
			presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
			return
		} else if err == exceptions.ErrUnauthorized {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			return
		} else if err == exceptions.ErrCartNotFound {
			presenter.ErrorResponse(c, http.StatusNotFound, exceptions.ErrCartNotFound)
			return
		}
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}
	presenter.SuccessResponse(c, http.StatusOK, nil)
}
