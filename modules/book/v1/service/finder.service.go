package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/book/v1/repository" // Import the book repository package
	"github.com/google/uuid"
)

// BookFinder is a service for books
type BookFinder struct {
	cfg          config.Config
	bookRepo     repository.BookRepositoryUseCase // Import the book repository interface
}

// BookFinderUseCase is a use case for books
type BookFinderUseCase interface {
	GetBooks(ctx context.Context) ([]*entity.Book, error)
	GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error)
}

// NewBookFinder creates a new BookFinder
func NewBookFinder(cfg config.Config, bookRepo repository.BookRepositoryUseCase) *BookFinder {
	return &BookFinder{
		cfg:      cfg,
		bookRepo: bookRepo,
	}
}

// GetBooks gets all books
func (bf *BookFinder) GetBooks(ctx context.Context) ([]*entity.Book, error) {
    books, err := bf.bookRepo.GetBooks(ctx)

    if err != nil {
        return nil, err
    }

    return books, nil
}

// GetBookByID gets a book by ID
func (bf *BookFinder) GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
	book, err := bf.bookRepo.GetBookByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	if book == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return book, nil
}
