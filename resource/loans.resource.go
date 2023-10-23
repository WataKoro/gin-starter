package resource

import (
    "gin-starter/entity"
    // "github.com/google/uuid"
)

type CreateLoanRequest struct {
    BookID      string `form:"bookid" json:"bookid" binding:"required"`
    UserID      string `form:"userid" json:"userid" binding:"required"`
}

type LoanResponse struct {
    ID          string `json:"id"`
    BookID      string `json:"bookid"`
    UserID      string `json:"userid"`
    RequestedAt string `json:"requestedat"`
    Approved    bool   `json:"approved"`
    Returned    bool   `json:"returned"`
    DueDate     string `json:"duedate"`
}

type UpdateLoanRequest struct {
    Approved  bool `json:"approved"`
    Returned  bool `json:"returned"`
}


type LoanListResponse struct {
    List  []*LoanResponse `json:"list"`
    Total int64           `json:"total"`
}

func CreateLoanResponse(loan *entity.Loan) *LoanResponse {
    return &LoanResponse{
        ID:          loan.ID.String(),
        BookID:      loan.BookID,
        UserID:      loan.UserID,
        RequestedAt: loan.RequestedAt.Format(timeFormat),
        Approved:    loan.Approved,
        Returned:    loan.Returned,
        DueDate:     loan.DueDate.Format(timeFormat),
    }
}

func UpdateLoanResponse(loan *entity.Loan) *LoanResponse {
    return &LoanResponse{
        ID:          loan.ID.String(),
        BookID:      loan.BookID,
        UserID:      loan.UserID,
        RequestedAt: loan.RequestedAt.Format(timeFormat),
        Approved:    loan.Approved,
        Returned:    loan.Returned,
        DueDate:     loan.DueDate.Format(timeFormat),
    }
}
