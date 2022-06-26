package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/presenter"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type Borrowing struct {
	TotalCost int64   `json:"total_cost" form:"total_cost"`
	CartIDs   []int64 `json:"cart_ids" form:"cart_ids"`
}

// func (r Borrowing) ToBorrowingDomain() domains.Borrowing {
// 	return domains.Borrowing{
// 		TotalCost: r.TotalCost,
// 	}
// }

type BorrowingController struct {
	borrowingUsecase domains.BorrowingUsecase
}

func NewBorrowingController(borrowingUsecase domains.BorrowingUsecase) BorrowingController {
	return BorrowingController{
		borrowingUsecase: borrowingUsecase,
	}
}

// func (bc BorrowingController) GetBorrowingByID(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
// 		return
// 	}

// 	cart, err := bc.borrowingUsecase.GetBorrowingByID(id)
// 	if err != nil {
// 		if err == exceptions.ErrBadRequest {
// 				presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
// 				return
// 			} else if err == exceptions.ErrUnauthorized {
// 				presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
// 				return
// 			} else if err == exceptions.ErrBorrowingNotFound {
// 				presenter.ErrorResponse(c, http.StatusNotFound, exceptions.ErrBorrowingNotFound)
// 				return
// 			}
// 			presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
// 			return
// 		}
// 	resFromDomain := presenter.BorrowingFromDomain(cart)
// 	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
// }

// func (bc BorrowingController) ShowBorrowingByUserID(c *gin.Context) {
// 	userID := int64(c.Keys["id"].(float64))
// 	res, err := bc.borrowingUsecase.ShowBorrowingByUserID(userID)
// 	if err != nil {
// 		if err == exceptions.ErrBadRequest {
// 			presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
// 			return
// 		} else if err == exceptions.ErrUnauthorized {
// 			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
// 			return
// 		}
// 		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
// 		return
// 	}
// 	resFromDomain := presenter.BorrowingListFromDomain(res)

// 	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
// }

func (bc BorrowingController) InsertToBorrowing(c *gin.Context) {
	req := Borrowing{}
	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	userID := int64(c.Keys["id"].(float64))
	fmt.Printf("%v", req)
	res, err := bc.borrowingUsecase.InsertToBorrowing(userID, req.CartIDs, req.TotalCost)
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
	resFromDomain := presenter.BorrowingWithBookFromDomain(res)

	presenter.SuccessResponse(c, http.StatusCreated, resFromDomain)
}
