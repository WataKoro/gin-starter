package service

import (
	// "bytes"
	"context"
	// "fmt"
	// "gin-starter/common/constant"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/loans/v1/repository"
	// "gin-starter/utils"
	"github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
	// "html/template"
	"log"
	// "strings"
)

// LoanUpdater is a struct that contains the dependencies of LoanUpdater
type LoanUpdater struct {
    cfg      config.Config
    loanRepo repository.LoanRepositoryUseCase
}

type LoanUpdaterUseCase interface {
    UpdateLoan(ctx context.Context, loan *entity.Loan, loanID uuid.UUID) error
}

// NewLoanUpdater is a constructor for the Loan updater
func NewLoanUpdater(cfg config.Config, loanRepo repository.LoanRepositoryUseCase) *LoanUpdater {
    return &LoanUpdater{
        cfg:      cfg,
        loanRepo: loanRepo,
    }
}

// UpdateLoan updates a loan.
func (uu *LoanUpdater) UpdateLoan(ctx context.Context, loan *entity.Loan, loanID uuid.UUID) error {   
	if err := uu.loanRepo.UpdateLoan(ctx, loan, loanID); err != nil {
        return errors.ErrInternalServerError.Error()
    }
    log.Println("Update loan jalan")
    return nil
}