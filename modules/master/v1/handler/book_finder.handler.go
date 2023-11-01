package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/master/v1/service"
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
    sortBy := "title"
    sortOrder := "asc"
    search := ""
    author := ""
    genre := ""

    if pageParam := c.DefaultQuery("page", ""); pageParam != "" {
        page, _ = strconv.Atoi(pageParam)
    }

    if pageSizeParam := c.DefaultQuery("pageSize", ""); pageSizeParam != "" {
        pageSize, _ = strconv.Atoi(pageSizeParam)
    }

    if sortByParam := c.DefaultQuery("sortBy", ""); sortByParam != "" {
        sortBy = strings.ToLower(sortByParam)
    }

    if sortOrderParam := c.DefaultQuery("sortOrder", ""); sortOrderParam != "" {
        sortOrder = strings.ToLower(sortOrderParam)
    }

    if searchParam := c.DefaultQuery("search", ""); searchParam != "" {
        search = searchParam
    }

    books, err := mf.bookFinder.GetBooks(c.Request.Context(), 1, 1000, sortBy, sortOrder, "", "", "")

    res := make([]*resource.BookResponse, 0)

    for _, book := range books {
        res = append(res, resource.NewBookResponse(book))
    }

    totalBooks := int64(len(res))

    books, err = mf.bookFinder.GetBooks(c.Request.Context(), page, pageSize, sortBy, sortOrder, search, author, genre)
    if err != nil {
        c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
        c.Abort()
        return
    }

    res = res[:0]

    for _, book := range books {
        res = append(res, resource.NewBookResponse(book))
    }

    log.Println("len :",  totalBooks)
    totalPages := totalBooks / int64(pageSize)
    if totalBooks % int64(pageSize) != 0 {
        totalPages++
    }

    c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.BookListResponse{
        List:  res,
        Page: page,
        PerPage: int64(len(res)),
        TotalBooks: totalBooks,
        TotalPages: totalPages,
    }))
}