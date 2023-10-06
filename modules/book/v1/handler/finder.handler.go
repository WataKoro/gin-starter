package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/book/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/gin-gonic/gin"
	"net/http"

	"log"
)

// BookFinderHandler is a handler for book finder
type BookFinderHandler struct {
	bookFinder service.BookFinderUseCase
}

// NewBookFinderHandler is a constructor for BookFinderHandler
func NewBookFinderHandler(
	bookFinder service.BookFinderUseCase,
) *BookFinderHandler {
	return &BookFinderHandler{
		bookFinder: bookFinder,
	}
}

// GetBooks is a handler for getting all books
func (mf *BookFinderHandler) GetBooks(c *gin.Context) {
	log.Println("masuk getbooks")
	books, err := mf.bookFinder.GetBooks(c.Request.Context())
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	
	res := make([]*resource.BookResponse, 0)

	for _, book := range books {
		res = append(res, resource.NewBookResponse(book))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.BookListResponse{
		List:  res,
		Total: int64(len(res)),
	}))
}