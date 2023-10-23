package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/loans/v1/repository"
	"github.com/google/uuid"
)

// BookFinder is a service for books
type LoanFinder struct {
	cfg          config.Config
	loanRepo     repository.LoanRepositoryUseCase
}

// BookFinderUseCase is a use case for books
type LoanFinderUseCase interface {
	GetLoanRequests(ctx context.Context) ([]*entity.Loan, error)
	GetLoanRequestByID(ctx context.Context, id uuid.UUID) (*entity.Loan, error)
}

// NewBookFinder creates a new BookFinder
func NewLoanFinder(cfg config.Config, loanRepo repository.LoanRepositoryUseCase) *LoanFinder {
	return &LoanFinder{
		cfg:      cfg,
		loanRepo: loanRepo,
	}
}

// GetBooks gets all books
func (lf *LoanFinder) GetLoanRequests(ctx context.Context) ([]*entity.Loan, error) {
    loans, err := lf.loanRepo.GetLoanRequests(ctx)

    if err != nil {
        return nil, err
    }

    return loans, nil
}

// GetBookByID gets a book by ID
func (lf *LoanFinder) GetLoanRequestByID(ctx context.Context, id uuid.UUID) (*entity.Loan, error) {
	loan, err := lf.loanRepo.GetLoanRequestByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	if loan == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return loan, nil
}
