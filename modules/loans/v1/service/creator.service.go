package service

import (
    "context"
    "gin-starter/common/errors"
    "gin-starter/config"
    "gin-starter/entity"
    "gin-starter/modules/loans/v1/repository" // Import the loan repository package
    "github.com/google/uuid"
)

// LoanCreator is a struct that contains all the dependencies for the Loan creator
type LoanCreator struct {
    cfg         config.Config
    loanRepo    repository.LoanRepositoryUseCase // Import the loan repository interface
}

// LoanCreatorUseCase is a use case for the Loan creator
type LoanCreatorUseCase interface {
    CreateLoanRequest(ctx context.Context, bookID, userID string) (*entity.Loan, error)
}

// NewLoanCreator is a constructor for the Loan creator
func NewLoanCreator(
    cfg config.Config,
    loanRepo repository.LoanRepositoryUseCase, // Import the loan repository interface
) *LoanCreator {
    return &LoanCreator{
        cfg:      cfg,
        loanRepo: loanRepo,
    }
}

// CreateLoanRequest creates a new loan request
func (lc *LoanCreator) CreateLoanRequest(ctx context.Context, bookID, userID string) (*entity.Loan, error) {

    loan := entity.NewLoan(
        uuid.New(),
        bookID,
        userID,
        // time.Now(),
        // time.Now(),
    )

    if err := lc.loanRepo.CreateLoanRequest(ctx, loan); err!= nil {
        return nil, errors.ErrInternalServerError.Error()
    }

    return loan, nil
}