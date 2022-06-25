package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/presenter"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type UserID struct {
	ID int64 `json:"id" form:"id"`
}

func (u *UserID) ToGetByIdUserDomain() domains.User {
	return domains.User{
		ID: u.ID,
	}
}

type UpdateUser struct {
	Fullname    string `json:"fullname" form:"fullname"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Photo       string `json:"picture_profile" form:"picture_profile"`
}

func (u *UpdateUser) UpdateUserToDomain() domains.UpdateUser {
	return domains.UpdateUser{
		Fullname:    u.Fullname,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		Photo:       u.Photo,
	}
}

type UserController struct {
	userUsecase domains.UserUsecase
}

func NewUserController(userUsercase domains.UserUsecase) UserController {
	return UserController{
		userUsecase: userUsercase,
	}
}

func (u *UserController) UpdateUserProfile(c *gin.Context) {
	paramsId := c.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}

	req := UpdateUser{}

	if err = c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	domain := req.UpdateUserToDomain()

	res, err := u.userUsecase.UpdateUserProfile(domain, int64(id))

	if err != nil {
		presenter.ErrorResponse(c, http.StatusNotFound, exceptions.ErrUserNotFound)
		return
	}

	responseFromDomain := presenter.FetchUpdateUser(res)
	presenter.SuccessResponse(c, http.StatusOK, responseFromDomain)
}

func (u *UserController) GetUserDetail(c *gin.Context) {
	paramsId := c.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}

	res, err := u.userUsecase.FetchUserByID(int64(id))

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	responseFromDomain := presenter.FetchUserDetail(res)
	presenter.SuccessResponse(c, http.StatusOK, responseFromDomain)

}
