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
    Approved    bool         `json:"approved"`
    Returned    bool         `json:"returned"`
    DueDate     time.Time    `json:"duedate"`
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
        UserID:      userID,
        RequestedAt: time.Now(),
        Approved:    false,
        Returned:    false,
        DueDate:     time.Now(),
    }
}

// MapUpdateFrom mapping from model
func (model *Loan) MapUpdateFrom(from *Loan) *map[string]interface{} {
    if from == nil {
        return &map[string]interface{}{
            "bookid":      model.BookID,
            "userid":      model.UserID,
            "requested_at": model.RequestedAt,
            "approved":     model.Approved,
            "returned":     model.Returned,
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

    if model.Approved != from.Approved {
        mapped["approved"] = from.Approved
    }

    if model.Returned != from.Returned {
        mapped["returned"] = from.Returned
    }
    mapped["requestedat"] = time.Now()
    mapped["duedate"] = time.Now()
    return &mapped
}
