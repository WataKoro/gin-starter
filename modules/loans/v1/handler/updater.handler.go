package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "gin-starter/common/errors"
    "gin-starter/response"
    "gin-starter/modules/loans/v1/service"
    "gin-starter/resource"
    "log"
)

// LoanUpdaterHandler is a handler for loan updater
type LoanUpdaterHandler struct {
    loanUpdater  service.LoanUpdaterUseCase
    loanFinder   service.LoanFinderUseCase
}

// NewLoanUpdaterHandler is a constructor for LoanUpdaterHandler
func NewLoanUpdaterHandler(
    loanUpdater service.LoanUpdaterUseCase,
    loanFinder  service.LoanFinderUseCase,
) *LoanUpdaterHandler {
    return &LoanUpdaterHandler{
        loanUpdater: loanUpdater,
        loanFinder:  loanFinder,
    }
}

// ApproveLoan is a handler for approving a loan
func (lrh *LoanUpdaterHandler) ApproveLoan(c *gin.Context) {
    var request resource.UpdateLoanRequest

    if err := c.ShouldBind(&request); err != nil {
        c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
        c.Abort()
        return
    }

    loanIDStr := c.Param("id")
    loanID, err := uuid.Parse(loanIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, "Invalid loan ID"))
        c.Abort()
        return
    }

    loan, err := lrh.loanFinder.GetLoanRequestByID(c, loanID)
    if err != nil {
        parseError := errors.ParseError(err)
        c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
        c.Abort()
        return
    }

    if loan == nil {
        c.JSON(http.StatusNotFound, response.ErrorAPIResponse(http.StatusNotFound, "Loan not found"))
        c.Abort()
        return
    }

    log.Print(request.Approved)

    loan.Approved = request.Approved
    loan.Returned = request.Returned

    if err := lrh.loanUpdater.ApproveLoan(c, loan, loanID); err != nil {
        parseError := errors.ParseError(err)
        c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
        c.Abort()
        return
    }

    c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Loan updated successfully", resource.UpdateLoanResponse(loan)))
}