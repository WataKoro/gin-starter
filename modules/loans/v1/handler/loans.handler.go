package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "gin-starter/common/errors"
    "gin-starter/common/interfaces"
    "gin-starter/modules/loans/v1/service"
    "gin-starter/resource"
    "gin-starter/response"
    // "github.com/google/uuid"
    // "gin-starter/utils"
)

type LoanRequestHandler struct {
    loanCreator service.LoanCreatorUseCase // Replace with the actual loan request service type
    cloudStorage interfaces.CloudStorageUseCase
}

func NewLoanRequestHandler(
    loanCreator service.LoanCreatorUseCase,
    cloudStorage interfaces.CloudStorageUseCase,
) *LoanRequestHandler {
    return &LoanRequestHandler{
        loanCreator: loanCreator,
        cloudStorage: cloudStorage,
    }
}

// CreateLoanRequest is a handler for creating a loan request for a book
func (lrh *LoanRequestHandler) CreateLoanRequest(c *gin.Context) {
    var request resource.CreateLoanRequest
    if err := c.ShouldBind(&request); err != nil {
        c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
        c.Abort()
        return
    }

    loan, err := lrh.loanCreator.CreateLoanRequest(
        c,
        request.BookID,
        request.UserID,
    )

    if err != nil {
        parseError := errors.ParseError(err)
        c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
        c.Abort()
        return
    }

    c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.CreateLoanResponse(loan)))
}