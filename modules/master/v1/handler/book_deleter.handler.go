package handler

import (
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/modules/master/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"log"
)

// BookDeleterHandler is a handler for book deleter
type BookDeleterHandler struct {
	bookDeleter  service.BookDeleterUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewBookDeleterHandler is a constructor for BookDeleterHandler
func NewBookDeleterHandler(
	bookDeleter service.BookDeleterUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *BookDeleterHandler {
	return &BookDeleterHandler{
		bookDeleter:  bookDeleter,
		cloudStorage: cloudStorage,
	}
}

// DeleteBook is a handler for deleting a book
func (bd *BookDeleterHandler) DeleteBook(c *gin.Context) {
	var request resource.DeleteBookRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	reqID, err := uuid.Parse(request.ID)

	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}

	log.Println("jalan delete")
	
	if err := bd.bookDeleter.DeleteBookByID(c, reqID); err != nil { // Updated method name
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", nil))
}
