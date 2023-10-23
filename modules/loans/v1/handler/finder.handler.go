package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/loans/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/gin-gonic/gin"
	"net/http"

	"log"
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

// GetBooks is a handler for getting all books
func (mf *LoanFinderHandler) GetLoanRequests(c *gin.Context) {
	log.Println("masuk getloans")
	loans, err := mf.loanFinder.GetLoanRequests(c.Request.Context())
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