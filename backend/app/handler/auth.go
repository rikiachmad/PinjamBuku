package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/presenter"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/commons/helpers"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type Auth struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (a Auth) ToDomain() domains.User {
	return domains.User{
		Email:    a.Email,
		Password: a.Password,
	}
}

type AuthController struct {
	loginUsecase domains.UserUsecase
}

// NewAuthController creates new user controller
func NewAuthController(loginUsecase domains.UserUsecase) AuthController {
	return AuthController{
		loginUsecase: loginUsecase,
	}
}

func (a AuthController) SignIn(c *gin.Context) {
	req := Auth{}
	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}
	domain := req.ToDomain()
	res, err := a.loginUsecase.Login(domain)
	resFromDomain := presenter.AuthFromDomain(res)
	if err != nil {
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			presenter.ErrorResponse(c, http.StatusConflict, exceptions.ErrInvalidCredentials)
			return
		}
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}
	cookie := helpers.CreateCookie(resFromDomain.Token)
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}
