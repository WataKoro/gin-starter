package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "gin-starter/common/errors"
    "gin-starter/response"
    "gin-starter/modules/book/v1/service"
    "gin-starter/resource"
)

// BookUpdaterHandler is a handler for book updater
type BookUpdaterHandler struct {
    bookUpdater  service.BookUpdaterUseCase
    bookFinder   service.BookFinderUseCase
}

// NewBookUpdaterHandler is a constructor for BookUpdaterHandler
func NewBookUpdaterHandler(
    bookUpdater service.BookUpdaterUseCase,
    bookFinder service.BookFinderUseCase,
) *BookUpdaterHandler {
    return &BookUpdaterHandler{
        bookUpdater: bookUpdater,
        bookFinder:  bookFinder,
    }
}

// UpdateBook is a handler for updating a book
func (bu *BookUpdaterHandler) UpdateBook(c *gin.Context) {
    var request resource.UpdateBookRequest

    if err := c.ShouldBind(&request); err != nil {
        c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
        c.Abort()
        return
    }

    bookIDStr := c.Param("id")
    bookID, err := uuid.Parse(bookIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, "Invalid book ID"))
        c.Abort()
        return
    }

    book, err := bu.bookFinder.GetBookByID(c, bookID)
    if err != nil {
        parseError := errors.ParseError(err)
        c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
        c.Abort()
        return
    }

    if book == nil {
        c.JSON(http.StatusNotFound, response.ErrorAPIResponse(http.StatusNotFound, "Book not found"))
        c.Abort()
        return
    }

    // Update the book entity based on the request
    book.Title = request.Title
    book.Author = request.Author
    book.Genre = request.Genre
    book.Desc = request.Desc

    if err := bu.bookUpdater.UpdateBook(c, book, bookID); err != nil {
        parseError := errors.ParseError(err)
        c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
        c.Abort()
        return
    }

    c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Book updated successfully", nil))
}
