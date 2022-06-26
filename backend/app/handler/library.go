package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/presenter"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type LibraryId struct {
	ID int64 `json:"id" form:"id"`
}

func (l *LibraryId) ToGetByIdLibraryDomain() domains.Library {
	return domains.Library{
		ID: l.ID,
	}
}

type UpdateLibrary struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Photo       string `json:"picture_profile" form:"picture_profile"`
}

func (u *UpdateLibrary) UpdateLibraryToDomain() domains.UpdateLibrary {
	return domains.UpdateLibrary{
		Name:        u.Name,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		Photo:       u.Photo,
	}
}

type LibraryController struct {
	libraryUsecase domains.LibraryUsecase
}

func NewLibraryController(libraryUsecase domains.LibraryUsecase) LibraryController {
	return LibraryController{
		libraryUsecase: libraryUsecase,
	}
}

func (l *LibraryController) GetAllLibrary(c *gin.Context) {
	res, err := l.libraryUsecase.GetAllLibrary()
	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrInternalServerError)
		return
	}

	response := make([]presenter.Library, len(res))

	for i, library := range res {
		response[i] = presenter.FetchLibraryDefault(library)
	}

	presenter.SuccessResponse(c, http.StatusOK, response)
}

func (l *LibraryController) GetLibraryByID(c *gin.Context) {
	paramsId := c.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}

	req := LibraryId{
		ID: int64(id),
	}

	if err = c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	domain := req.ToGetByIdLibraryDomain()

	res, err := l.libraryUsecase.GetLibraryByID(domain.ID)

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	responseFromDomain := presenter.FetchLibraryDefault(res)
	presenter.SuccessResponse(c, http.StatusOK, responseFromDomain)
}

func (l *LibraryController) UpdateLibraryProfileByID(c *gin.Context) {
	paramsId := c.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}

	req := UpdateLibrary{}

	if err = c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	domain := req.UpdateLibraryToDomain()

	res, err := l.libraryUsecase.UpdateLibraryProfileByID(domain, int64(id))

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	responseFromDomain := presenter.FetchUpdateLibrary(res)
	presenter.SuccessResponse(c, http.StatusOK, responseFromDomain)
}

func (l *LibraryController) GetAllBookById(c *gin.Context) {
	paramsId := c.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}

	req := LibraryId{
		ID: int64(id),
	}

	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	res, err := l.libraryUsecase.GetAllBookById(int64(id))

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	response := make([]presenter.Book, len(res))

	for i, book := range res {
		response[i] = presenter.FetchBookDefault(book)
	}

	presenter.SuccessResponse(c, http.StatusOK, response)

}

func (l LibraryController) InsertBook(c *gin.Context) {
	paramsId := c.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	}

	// req := LibraryId{
	// 	ID: int64(id),
	// }

	// if err := c.Bind(&req); err != nil {
	// 	presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
	// 	return
	// }

	reqBook := Books{}

	if err := c.ShouldBindJSON(&reqBook); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	dom := reqBook.ToInsertBookDomain()

	err = l.libraryUsecase.CreateBook(dom, int64(id))

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	presenter.SuccessResponse(c, http.StatusCreated, "success create book")
}

func (l LibraryController) UpdateBook(c *gin.Context) {
	paramsId := c.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	req := Books{}

	if err := c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	domain := req.ToInsertBookDomain()

	err = l.libraryUsecase.UpdateBook(domain, int64(id))

	if err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	presenter.SuccessResponse(c, http.StatusOK, "success update book")
}
