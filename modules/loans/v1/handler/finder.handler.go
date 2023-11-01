package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/loans/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"log"
	"github.com/google/uuid"
)

// BookFinderHandler is a handler for book finder
type LoanFinderHandler struct {
	loanFinder service.LoanFinderUseCase
}

// NewBookFinderHandler is a constructor for BookFinderHandler
func NewLoanFinderHandler(
	loanFinder service.LoanFinderUseCase,
) *LoanFinderHandler {
	return &LoanFinderHandler{
		loanFinder: loanFinder,
	}
}

// GetLoanRequests is a handler for getting all loans
func (mf *LoanFinderHandler) GetLoanRequests(c *gin.Context) {
	log.Println("masuk getloans")

	page := 1
    pageSize := 10
	status := 0
	sortOrder := "asc"

	if pageParam := c.DefaultQuery("page", ""); pageParam != "" {
        page, _ = strconv.Atoi(pageParam)
    }

    if pageSizeParam := c.DefaultQuery("pageSize", ""); pageSizeParam != "" {
        pageSize, _ = strconv.Atoi(pageSizeParam)
    }

	if sortOrderParam := c.DefaultQuery("sortOrder", ""); sortOrderParam != "" {
        sortOrder = strings.ToLower(sortOrderParam)
    }

	if statusParam := c.DefaultQuery("status", ""); statusParam != "" {
        status, _ = strconv.Atoi(statusParam)
    }

	loans, err := mf.loanFinder.GetLoanRequests(c.Request.Context(), page, pageSize, sortOrder, status)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.LoanResponse, 0)

	for _, loan := range loans {
		res = append(res, resource.CreateLoanResponse(loan))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.LoanListResponse{
		List:  res,
		Total: int64(len(res)),
	}))
}

func (lf *LoanFinderHandler) GetLoanRequestByID(c *gin.Context) {
    var request resource.GetLoanByIDRequest

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

    loanDetail, err := lf.loanFinder.GetLoanRequestByID(c, reqID)

    if err != nil {
        parseError := errors.ParseError(err)
        c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
        c.Abort()
        return
    }

    c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewLoanDetail(loanDetail)))
}
