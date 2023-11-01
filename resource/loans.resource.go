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
    Title       string `json:"title"`
    UserID      string `json:"userid"`
    Name        string `json:"name"`
    RequestedAt string `json:"requestedat"`
    Status      int    `json:"status"`
    DueDate     string `json:"duedate"`
}

type LoanDetailResponse struct {
    ID          string `json:"id"`
    BookID      string `json:"bookid"`
    UserID      string `json:"userid"`
    RequestedAt string `json:"requestedat"`
    Status      int    `json:"status"`
    DueDate     string `json:"duedate"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    Genre       string `json:"genre"`
    Desc        string `json:"desc"`
    Name        string `json:"name"`
    Email       string `json:"email"`
}

type UpdateLoanRequest struct {
    Status  int `json:"status"`
}


type LoanListResponse struct {
    List  []*LoanResponse `json:"list"`
    Total int64           `json:"total"`
}

type GetLoanByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

func CreateLoanResponse(loan *entity.Loan) *LoanResponse {
    return &LoanResponse{
        ID:          loan.ID.String(),
        BookID:      loan.BookID,
        Title:       loan.Title,
        UserID:      loan.UserID,
        Name:        loan.Name,
        RequestedAt: loan.RequestedAt.Format(timeFormat),
        Status:      loan.Status,
        DueDate:     loan.DueDate.Format(timeFormat),
    }
}

func UpdateLoanResponse(loan *entity.Loan) *LoanResponse {
    return &LoanResponse{
        ID:          loan.ID.String(),
        BookID:      loan.BookID,
        Title:       loan.Title,
        UserID:      loan.UserID,
        Name:        loan.Name,
        RequestedAt: loan.RequestedAt.Format(timeFormat),
        Status:      loan.Status,
        DueDate:     loan.DueDate.Format(timeFormat),
    }
}

func NewLoanDetail(loan *entity.Loan) *LoanDetailResponse {
    return &LoanDetailResponse{
        ID:          loan.ID.String(),
        BookID:      loan.BookID,
        UserID:      loan.UserID,
        RequestedAt: loan.RequestedAt.Format(timeFormat),
        Status:      loan.Status,
        DueDate:     loan.DueDate.Format(timeFormat),
        Title:       loan.Title,
        Author:      loan.Author,
        Genre:       loan.Genre,
        Desc:        loan.Desc,
        Name:        loan.Name,
        Email:       loan.Email,
    }
}
