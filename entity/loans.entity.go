package entity

import (
    "time"

    "github.com/google/uuid"
)

const (
    loanRequestTableName = "main.loanrequest"
)

type Loan struct {
    ID          uuid.UUID    `json:"id"`
    BookID      string       `json:"bookid"`
    UserID      string       `json:"userid"`
    RequestedAt time.Time    `json:"requestedat"`
    Status      int          `json:"status"`
    DueDate     time.Time    `json:"duedate"`
    Title       string       `json:"title"`
    Author      string       `json:"author"`
    Genre       string       `json:"genre"`
    Desc        string       `json:"desc"`
    Name        string       `json:"name"`
    Email       string       `json:"email"`
}

type LoanDetail struct {
    ID          uuid.UUID    `json:"id"`
    BookID      string       `json:"bookid"`
    UserID      string       `json:"userid"`
    RequestedAt time.Time    `json:"requestedat"`
    Status      int          `json:"status"`
    DueDate     time.Time    `json:"duedate"`
    Title       string       `json:"title"`
    Author      string       `json:"author"`
    Genre       string       `json:"genre"`
    Desc        string       `json:"desc"`
    Name        string       `json:"name"`
    Email       string       `json:"email"`
}

// TableName specifies table name
func (model *Loan) TableName() string {
    return loanRequestTableName
}

func NewLoan(
    id uuid.UUID,
    bookID string,
    userID string,
    // requestedAt time.Time,
    // dueDate time.Time,
) *Loan {
    return &Loan{
        ID:          id,
        BookID:      bookID,
        Title:       "",
        UserID:      userID,
        Name:        "",
        RequestedAt: time.Now(),
        Status:      0,
        DueDate:     time.Now(),
    }
}

// MapUpdateFrom mapping from model
func (model *Loan) MapUpdateFrom(from *Loan) *map[string]interface{} {
    if from == nil {
        return &map[string]interface{}{
            "bookid":      model.BookID,
            "userid":      model.UserID,
            "requestedat": model.RequestedAt,
            "status":      model.Status,
            "duedate":     model.DueDate,
        }
    }

    mapped := make(map[string]interface{})

    if model.BookID != from.BookID {
        mapped["bookid"] = from.BookID
    }

    if model.UserID != from.UserID {
        mapped["userid"] = from.UserID
    }

    if model.Status != from.Status {
        mapped["status"] = from.Status
    }

    mapped["requestedat"] = time.Now()
    mapped["duedate"] = time.Now()
    return &mapped
}
