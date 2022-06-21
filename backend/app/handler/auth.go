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

func (a Auth) ToUserDomain() domains.User {
	return domains.User{
		Email:    a.Email,
		Password: a.Password,
	}
}

func (a Auth) ToLibraryDomain() domains.Library {
	return domains.Library{
		Email:    a.Email,
		Password: a.Password,
	}
}

type RegisterUser struct {
	Fullname    string `json:"fullname" form:"fullname"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Role        int    `json:"role" form:"role"`
	Photo       string `json:"photo" form:"photo"`
}

func (r RegisterUser) ToCreateUserDomain() domains.CreateUser {
	return domains.CreateUser{
		Fullname:    r.Fullname,
		Email:       r.Email,
		Password:    r.Password,
		Address:     r.Address,
		PhoneNumber: r.PhoneNumber,
		Role:        r.Role,
		Photo:       r.Photo,
	}
}

type AuthController struct {
	authUsecase domains.AuthUsecase
}

type LibraryAuthController struct {
	authUsecase domains.LibraryAuthUsecase
}

func NewAuthController(authUsecase domains.AuthUsecase) AuthController {
	return AuthController{
		authUsecase: authUsecase,
	}
}

func NewLibraryAuthController(authUsecase domains.LibraryAuthUsecase) LibraryAuthController {
	return LibraryAuthController{
		authUsecase: authUsecase,
	}
}

func (a AuthController) SignIn(c *gin.Context) {
	req := Auth{}
	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	domain := req.ToUserDomain()
	res, err := a.authUsecase.Login(domain)
	resFromDomain := presenter.LoginFromDomain(res)
	if err != nil {
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrInvalidCredentials)
			return
		}
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}
	cookie := helpers.CreateCookie(resFromDomain.Token)
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (l LibraryAuthController) SingInForLibrary(c *gin.Context) {
	req := Auth{}
	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	domain := req.ToLibraryDomain()
	res, err := l.authUsecase.Login(domain)
	resFromDomain := presenter.LoginLibraryFromDomain(res)
	if err != nil {
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrInvalidCredentials)
			return
		}
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}
	cookie := helpers.CreateCookie(resFromDomain.Token)
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}

func (a AuthController) Register(c *gin.Context) {
	req := RegisterUser{}
	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}
	domain := req.ToCreateUserDomain()
	res, err := a.authUsecase.Register(domain)
	if err != nil {
		if errors.Is(err, exceptions.ErrUserAlreadyExists) {
			presenter.ErrorResponse(c, http.StatusConflict, exceptions.ErrUserAlreadyExists)
			return
		}
		presenter.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	resFromDomain := presenter.CreateUserFromDomain(res)

	presenter.SuccessResponse(c, http.StatusOK, resFromDomain)
}
