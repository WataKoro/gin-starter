package repository

import (
    "context"
    "gin-starter/entity" // Import your loan entity
    "github.com/pkg/errors"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "log"
    //"strings"
)

type LoanRepository struct {
    gormDB *gorm.DB
}

// LoanRepositoryUseCase is an interface for Loan repository use case
type LoanRepositoryUseCase interface {
    CreateLoanRequest(ctx context.Context, loan *entity.Loan) error
    GetLoanRequests(ctx context.Context, page int, pageSize int, sortOrder string, status int) ([]*entity.Loan, error)
    GetLoanRequestByID(ctx context.Context, id uuid.UUID) (*entity.Loan, error)
    UpdateLoan(ctx context.Context, loan *entity.Loan, loanID uuid.UUID) error
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
func (repo *LoanRepository) GetLoanRequests(ctx context.Context, page int, pageSize int, sortOrder string, status int) ([]*entity.Loan, error) {
    offset := (page - 1) * pageSize
    models := make([]*entity.Loan, 0)

    query := repo.gormDB.
        WithContext(ctx).
        Model(&entity.Loan{}).
        Select("loanrequest.*, u.name as name, b.title as title").
        Joins("INNER JOIN main.users u ON loanrequest.user_id = u.id").
        Joins("INNER JOIN main.book b ON loanrequest.book_id = b.id").
        Offset(offset).
        Limit(pageSize)

    if status != 0 {
        query = query.Where("loanrequest.status = ?", status)
    }

    if err := query.Order("loanrequest.requested_at ASC").
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
        Select("loanrequest.*, u.name as name, u.email as email, b.title as title, b.author as author, b.genre as genre, b.desc as desc").
        Joins("INNER JOIN main.users u ON loanrequest.user_id = u.id").
        Joins("INNER JOIN main.book b ON loanrequest.book_id = b.id").
        Where("loanrequest.id = ?", id).
        First(loan).
        Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, errors.Wrap(err, "[LoanRepository-GetLoanRequestByID]")
    }
    return loan, nil
}


// UpdateLoan updates a loan
func (repo *LoanRepository) UpdateLoan(ctx context.Context, loan *entity.Loan, loanID uuid.UUID) error {
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