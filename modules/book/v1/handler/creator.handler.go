package handler

import (
    "gin-starter/common/errors"
    "gin-starter/common/interfaces"
    "gin-starter/modules/book/v1/service" // Import the book service package
    "gin-starter/resource"
    "gin-starter/response"
    "github.com/gin-gonic/gin"
    "net/http"

    "log"
)

// BookCreatorHandler is a handler for book creation and management
type BookCreatorHandler struct {
    bookCreator  service.BookCreatorUseCase // Import the book service interface
    cloudStorage interfaces.CloudStorageUseCase
}

// NewBookCreatorHandler is a constructor for BookCreatorHandler
func NewBookCreatorHandler(
    bookCreator service.BookCreatorUseCase, // Import the book service interface
    cloudStorage interfaces.CloudStorageUseCase,
) *BookCreatorHandler {
    return &BookCreatorHandler{
        bookCreator: bookCreator,
        cloudStorage: cloudStorage,
    }
}

// CreateBook is a handler for creating a book
func (bc *BookCreatorHandler) CreateBook(c *gin.Context) {
    var request resource.CreateBookRequest
    if err := c.ShouldBind(&request); err != nil {
        c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
        c.Abort()
        return
    }

    log.Println("jalan create book")

    // Create the book using the bookCreator service method
    book, err := bc.bookCreator.CreateBook(
        c,
        request.Title,
        request.Author,
        request.Genre,
        request.Desc,
    )

    if err != nil {
        parseError := errors.ParseError(err)
        c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
        c.Abort()
        return
    }

    c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewBookResponse(book)))
}
