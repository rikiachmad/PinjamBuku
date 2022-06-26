package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/presenter"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
	"github.com/rg-km/final-project-engineering-16/backend/domains"
)

type idBook struct {
	Id int64 `json:"id" form:"id"`
}

type InsertBook struct {
	KatalogId   string `json:"katalogId" form:"katalogId"`
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	Description string `json:"description" form:"description"`
	Cover       string `json:"cover" form:"cover"`
	PageNumber  int64  `json:"pageNumber" form:"pageNumber"`
	Stock       int64  `json:"stock" form:"stock"`
	Deposit     int64  `json:"deposit" form:"deposit"`
	CategoryId  int64  `json:"categoryId" form:"categoryId"`
	LibraryId   int64  `json:"libraryId" form:"libraryId"`
	IsPublish   bool   `json:"isPublish" form:"isPublish"`
}

type Books struct {
	Id          int64  `json:"id" form:"id"`
	KatalogId   string `json:"katalogId" form:"katalogId"`
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	Description string `json:"description" form:"description"`
	Cover       string `json:"cover" form:"cover"`
	PageNumber  int64  `json:"pageNumber" form:"pageNumber"`
	Stock       int64  `json:"stock" form:"stock"`
	Deposit     int64  `json:"deposit" form:"deposit"`
	CategoryId  int64  `json:"categoryId" form:"categoryId"`
	LibraryId   int64  `json:"libraryId" form:"libraryId"`
	IsPublish   bool   `json:"isPublish" form:"isPublish"`
}

func (i *idBook) ToGetByIdBookDomain() domains.Book {
	return domains.Book{
		ID: i.Id,
	}
}

func (b *Books) ToInsertBookDomain() domains.Book {
	return domains.Book{
		ID:          b.Id,
		KatalogId:   b.KatalogId,
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
		Cover:       b.Cover,
		PageNumber:  b.PageNumber,
		Stock:       b.Stock,
		Deposit:     b.Deposit,
		CategoryId:  b.CategoryId,
		LibraryId:   b.LibraryId,
		IsPublish:   b.IsPublish,
	}
}

type BookController struct {
	bookUsecase domains.BookUsecase
}

func NewBookController(b domains.BookUsecase) BookController {
	return BookController{
		bookUsecase: b,
	}
}

func (b BookController) GetBookById(c *gin.Context) {
	paramsId := c.Param("id")

	id, err := strconv.Atoi(paramsId)

	req := idBook{
		Id: int64(id),
	}

	if err = c.Bind(&req); err != nil {
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
		return
	}

	domain := req.ToGetByIdBookDomain()

	res, err := b.bookUsecase.FetchById(domain)

	if err != nil {
		if errors.Is(err, exceptions.ErrBadRequest) {
			log.Printf("error handler-book FetchById %s", err)
			presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
			return
		}
	}

	responseFromDomain := presenter.FetchBookDefault(res)
	presenter.SuccessResponse(c, http.StatusOK, responseFromDomain)
}

func (b BookController) GetAllBook(c *gin.Context) {
	res, err := b.bookUsecase.FetchAll()
	if err != nil {
		log.Printf("error handler-book GetAllBook %v", err)
		presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrInternalServerError)
		return
	}

	response := make([]presenter.Book, len(res))

	for i, book := range res {
		response[i] = presenter.FetchBookDefault(book)
	}

	presenter.SuccessResponse(c, http.StatusOK, response)
}

func (b BookController) GetSearchBook(c *gin.Context) {
	key := c.Query("q")
	words := c.Query("w")
	sort := c.Query("sort")

	if len(key) >= 1 && len(words) >= 1 {

		switch key {
		case "title":
			res, err := b.bookUsecase.FetchSearchBook(key, words)

			if err != nil {
				log.Printf("error handler-book GetSearchBook %v", err)
				presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrInternalServerError)
				return
			}

			response := make([]presenter.Book, len(res))

			for i, book := range res {
				response[i] = presenter.FetchBookDefault(book)
			}

			presenter.SuccessResponse(c, http.StatusOK, response)
			return

		}

	} else if len(sort) >= 1 {
		res, err := b.bookUsecase.FetchSort(sort)

		if err != nil {
			log.Printf("error handler-book FetchSort %v", err)
			presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
			return
		}

		response := make([]presenter.Book, len(res))

		for i, book := range res {
			response[i] = presenter.FetchBookDefault(book)
		}

		presenter.SuccessResponse(c, http.StatusOK, response)
		return
	}

	presenter.ErrorResponse(c, http.StatusBadRequest, exceptions.ErrBadRequest)
}
