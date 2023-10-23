package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/book/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
    "strings"

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

    page := 1
    pageSize := 10
    sortOrder := "asc"
    search := ""
    author := c.DefaultQuery("author", "")
    genre := c.DefaultQuery("genre", "")
    
    if pageParam := c.DefaultQuery("page", ""); pageParam != "" {
        page, _ = strconv.Atoi(pageParam)
    }

    if pageSizeParam := c.DefaultQuery("pageSize", ""); pageSizeParam != "" {
        pageSize, _ = strconv.Atoi(pageSizeParam)
    }

    if sortOrderParam := c.DefaultQuery("sortOrder", ""); sortOrderParam != "" {
        sortOrder = strings.ToLower(sortOrderParam)
    }

    if searchParam := c.DefaultQuery("search", ""); searchParam != "" {
        search = searchParam
    }

    books, err := mf.bookFinder.GetBooks(c.Request.Context(), page, pageSize, sortOrder, search, author, genre)
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
        Page: page,
        Total: int64(len(res)),
    }))
}
