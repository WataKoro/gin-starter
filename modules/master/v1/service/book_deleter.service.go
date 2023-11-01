package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/modules/master/v1/repository"
	"github.com/google/uuid"
)

// BookDeleter is a service for book deletion
type BookDeleter struct {
	cfg      config.Config
	bookRepo repository.BookRepositoryUseCase
}

// BookDeleterUseCase is a use case for book deletion
type BookDeleterUseCase interface {
	DeleteBookByID(ctx context.Context, id uuid.UUID) error
}

// NewBookDeleter creates a new BookDeleter
func NewBookDeleter(
	cfg config.Config,
	bookRepo repository.BookRepositoryUseCase,
) *BookDeleter {
	return &BookDeleter{
		cfg:      cfg,
		bookRepo: bookRepo,
	}
}

// DeleteBookByID deletes a book by its ID
func (bd *BookDeleter) DeleteBookByID(ctx context.Context, id uuid.UUID) error {
	if err := bd.bookRepo.DeleteBookByID(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
