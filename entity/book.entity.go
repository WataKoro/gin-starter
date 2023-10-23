package entity

import (
    "time"

    "github.com/google/uuid"
)

const (
    bookTableName = "main.book"
)

type Book struct {
    ID          uuid.UUID    `json:"id"`
    Title       string       `json:"title"`
    Author      string       `json:"author"`
    Genre       string       `json:"genre"`
    Desc        string       `json:"desc"`
    CreatedAt   time.Time    `json:"created_at"`
    UpdatedAt   time.Time    `json:"updated_at"`
}

// TableName specifies table name
func (model *Book) TableName() string {
    return bookTableName
}

func NewBook(
    id uuid.UUID,
    title string,
    author string,
    genre string,
    desc string,
) *Book {
    return &Book{
        ID:          id,
        Title:       title,
        Author:      author,
        Genre:       genre,
        Desc:        desc,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
}

// MapUpdateFrom mapping from model
func (model *Book) MapUpdateFrom(from *Book) *map[string]interface{} {
    if from == nil {
        return &map[string]interface{}{
            "title":       model.Title,
            "author":      model.Author,
            "genre":       model.Genre,
            "desc":        model.Desc,
            "updated_at":  model.UpdatedAt,
        }
    }

    mapped := make(map[string]interface{})

    if model.Title != from.Title {
        mapped["title"] = from.Title
    }

    if model.Author != from.Author {
        mapped["author"] = from.Author
    }

    if model.Genre != from.Genre {
        mapped["genre"] = from.Genre
    }

    if model.Desc != from.Desc {
        mapped["desc"] = from.Desc
    }

    mapped["updated_at"] = time.Now()
    return &mapped
}