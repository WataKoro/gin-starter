package repository

import (
    "context"
    "gin-starter/entity" // Import your loan entity
    "github.com/pkg/errors"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "log"
)

type LoanRepository struct {
    gormDB *gorm.DB
}

// LoanRepositoryUseCase is an interface for Loan repository use case
type LoanRepositoryUseCase interface {
    CreateLoanRequest(ctx context.Context, loan *entity.Loan) error
    GetLoanRequests(ctx context.Context) ([]*entity.Loan, error)
    GetLoanRequestByID(ctx context.Context, id uuid.UUID) (*entity.Loan, error)
    ApproveLoan(ctx context.Context, loan *entity.Loan, loanID uuid.UUID) error
}

func NewLoanRepository(
    db *gorm.DB,
) *LoanRepository {
    return &LoanRepository{
        gormDB: db,
    }
}

// CreateLoanRequest creates a new loan request
func (repo *LoanRepository) CreateLoanRequest(ctx context.Context, loan *entity.Loan) error {
    if err := repo.gormDB.
        WithContext(ctx).
        Create(loan).
        Error; err != nil {
        return errors.Wrap(err, "[LoanRepository-CreateLoanRequest]")
    }
    return nil
}

// GetLoanRequests returns a list of loan requests
func (repo *LoanRepository) GetLoanRequests(ctx context.Context) ([]*entity.Loan, error) {
    models := make([]*entity.Loan, 0)
    if err := repo.gormDB.
        WithContext(ctx).
        Model(&entity.Loan{}).
        Find(&models).
        Error; err != nil {
        return nil, errors.Wrap(err, "[LoanRepository-GetLoanRequests]")
    }
    return models, nil
}

// GetLoanRequestByID returns a loan request by its ID
func (repo *LoanRepository) GetLoanRequestByID(ctx context.Context, id uuid.UUID) (*entity.Loan, error) {
    loan := new(entity.Loan)
    if err := repo.gormDB.
        WithContext(ctx).
        Model(&entity.Loan{}).
        Where("id = ?", id).
        First(loan).
        Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // Loan request not found
        }
        return nil, errors.Wrap(err, "[LoanRepository-GetLoanRequestByID]")
    }
    return loan, nil
}

func (repo *LoanRepository) ApproveLoan(ctx context.Context, loan *entity.Loan, loanID uuid.UUID) error {
    if err := repo.gormDB.
        WithContext(ctx).
        Model(&entity.Loan{}).
        Where("id = ?", loanID).
        Updates(loan).
        Error; err != nil {
        return errors.Wrap(err, "[LoanRepository-UpdateLoan]")
    }
    log.Println("jalan")
    return nil
}