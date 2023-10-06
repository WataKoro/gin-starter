package service

import (
    "context"
    "strings"
    "log"
    "gin-starter/common/errors"
    "gin-starter/config"
    "gin-starter/entity"
    "gin-starter/modules/book/v1/repository" // Import the book repository package
    "github.com/google/uuid"
)

// BookCreator is a struct that contains all the dependencies for the Book creator
type BookCreator struct {
    cfg         config.Config
    bookRepo    repository.BookRepositoryUseCase // Import the book repository interface
}

// BookCreatorUseCase is a use case for the Book creator
type BookCreatorUseCase interface {
    // CreateBook creates a new book
    CreateBook(ctx context.Context, title, author, genre, description string) (*entity.Book, error)
}

// NewBookCreator is a constructor for the Book creator
func NewBookCreator(
    cfg config.Config,
    bookRepo repository.BookRepositoryUseCase, // Import the book repository interface
) *BookCreator {
    return &BookCreator{
        cfg:      cfg,
        bookRepo: bookRepo,
    }
}

// CreateBook creates a new book
func (bc *BookCreator) CreateBook(ctx context.Context, title, author, genre, description string) (*entity.Book, error) {
    title = strings.TrimSpace(title)
	trimmed := strings.TrimSpace(strings.ReplaceAll(title, " ", ""))
    
    log.Print(trimmed)
    
    if trimmed == "" {
        return nil, errors.ErrEmptyData.Error()
    }
    
    count, err := bc.bookRepo.FindBookByTitle(ctx, trimmed)
    
    if err != nil {
        return nil, errors.ErrInternalServerError.Error()
    }
    
    if count > 0 {
        return nil, errors.ErrDuplicateEntry.Error()
    }

    

    log.Println("jalan tanpa masalah")

    book := entity.NewBook(
        uuid.New(),
        title,
        author,
        genre,
        description,
    )

    if err := bc.bookRepo.CreateBook(ctx, book); err != nil {
        return nil, errors.ErrInternalServerError.Error()
    }

    return book, nil
}
