package service

import (
	// "bytes"
	"context"
	// "fmt"
	// "gin-starter/common/constant"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/book/v1/repository"
	// "gin-starter/utils"
	"github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
	// "html/template"
	"log"
	"strings"
)

// BookUpdater is a struct that contains the dependencies of BookUpdater
type BookUpdater struct {
    cfg      config.Config
    bookRepo repository.BookRepositoryUseCase
}

type BookUpdaterUseCase interface {
    UpdateBook(ctx context.Context, book *entity.Book, bookID uuid.UUID) error
}

// NewBookUpdater is a constructor for the Book updater
func NewBookUpdater(cfg config.Config, bookRepo repository.BookRepositoryUseCase) *BookUpdater {
    return &BookUpdater{
        cfg:      cfg,
        bookRepo: bookRepo,
    }
}

// UpdateBook updates a book.
func (uu *BookUpdater) UpdateBook(ctx context.Context, book *entity.Book, bookID uuid.UUID) error {

	title := strings.TrimSpace(book.Title)
	trimmed := strings.TrimSpace(strings.ReplaceAll(title, " ", ""))
    
    log.Print(trimmed)

	if trimmed == "" {
        return errors.ErrEmptyData.Error()
    }
    
    count, err := uu.bookRepo.FindBookByTitle(ctx, trimmed)
    
    if err != nil {
        return errors.ErrInternalServerError.Error()
    }
    
    if count > 0 {
        return errors.ErrDuplicateEntry.Error()
    }

	if err := uu.bookRepo.UpdateBook(ctx, book, bookID); err != nil {
        return errors.ErrInternalServerError.Error()
    }

    log.Println("Update book jalan")

    return nil
}