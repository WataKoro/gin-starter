package repository

import (
	"context"
    "strings"
	"gin-starter/entity"
	"github.com/pkg/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
    "fmt"
)

// BookRepository is an struct for Book repository
type BookRepository struct {
	gormDB *gorm.DB
}

// BookRepositoryUseCase is an interface for Book repository use case
type BookRepositoryUseCase interface {
    // CreateBook creates a new Book
    CreateBook(ctx context.Context, book *entity.Book) error
    // GetBooks returns a list of books
    GetBooks(ctx context.Context, page int, pageSize int, sortBy string, sortOrder string, search string, author string, genre string) ([]*entity.Book, error)
    // GetBookByID returns a book by its ID
    GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error)
    DeleteBookByID(ctx context.Context, id uuid.UUID) error
    FindBookByTitle(ctx context.Context, title string) (int, error)
    UpdateBook(ctx context.Context, book *entity.Book, bookID uuid.UUID) error
}

// NewBookRepository creates a new Book repository
func NewBookRepository(
	db *gorm.DB,
) *BookRepository {
	return &BookRepository{
		gormDB: db,
	}
}

// CreateBook creates a new Book
func (repo *BookRepository) CreateBook(ctx context.Context, book *entity.Book) error {
    if err := repo.gormDB.
        WithContext(ctx).
        Create(book).
        Error; err != nil {
        return errors.Wrap(err, "[BookRepository-CreateBook]")
    }
    return nil
}

// GetBooks returns a list of books
func (repo *BookRepository) GetBooks(ctx context.Context, page int, pageSize int, sortBy string, sortOrder string, search string, author string, genre string) ([]*entity.Book, error) {
    offset := (page - 1) * pageSize
    models := make([]*entity.Book, 0)

    conditions := []interface{}{"%" + search + "%", "%" + search + "%", "%" + search + "%"}

    if err := repo.gormDB.
        WithContext(ctx).
        Model(&entity.Book{}).
        Offset(offset).
        Limit(pageSize).
        Where("title LIKE ? OR genre LIKE ? OR author LIKE ?", conditions...).
        Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).
        Find(&models).
        Error; err != nil {
        return nil, errors.Wrap(err, "[BookRepository-GetBooks]")
    }
    return models, nil
}

// GetBookByID returns a book by its ID
func (repo *BookRepository) GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
    book := new(entity.Book)
    if err := repo.gormDB.
        WithContext(ctx).
        Model(&entity.Book{}).
        Where("id = ?", id).
        First(book).
        Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // Book not found
        }
        return nil, errors.Wrap(err, "[BookRepository-GetBookByID]")
    }
    return book, nil
}

// DeleteBookByID deletes a book by its ID
func (repo *BookRepository) DeleteBookByID(ctx context.Context, id uuid.UUID) error {
	if err := repo.gormDB.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.Book{}).
		Error; err != nil {
		return errors.Wrap(err, "[BookRepository-DeleteBookByID]")
	}
	return nil
}

func (repo *BookRepository) FindBookByTitle(ctx context.Context, title string) (int, error) {
	models := make([]*entity.Book, 0)
    if err := repo.gormDB.
		WithContext(ctx).
        Model(&entity.Book{}).
		Where("REPLACE(lower(title), ' ', '') = ?", strings.ToLower(title)).
		Find(&models).
		Error; err != nil {
		return 0, errors.Wrap(err, "[BookRepository-FindBookByTitle]")
	}
	return len(models), nil
}


func (repo *BookRepository) UpdateBook(ctx context.Context, book *entity.Book, bookID uuid.UUID) error {
    if err := repo.gormDB.
        WithContext(ctx).
        Model(&entity.Book{}).
        Where("id = ?", bookID).
        Updates(book).
        Error; err != nil {
        return errors.Wrap(err, "[BookRepository-UpdateBook]")
    }
    return nil
}