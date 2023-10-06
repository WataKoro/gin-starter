package resource

import (
	"gin-starter/entity"
)

type CreateBookRequest struct {
	Title       string                `form:"title" json:"title"`
	Genre	    string                `form:"genre" json:"genre"`
	Author      string                `form:"author" json:"author"`
	Desc 		string                `form:"desc" json:"desc"`
}

type BookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre		string `json:"genre"`
	Author      string `json:"author"`
	Desc		string `json:"desc"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateBookRequest struct {
    Title  string `json:"title" binding:"required"`
    Author string `json:"author" binding:"required"`
    Genre  string `json:"genre" binding:"required"`
    Desc   string `json:"desc" binding:"required"`
}

type BookListResponse struct {
	List  []*BookResponse `json:"list"`
	Total int64       `json:"total"`
}

type DeleteBookRequest struct {
	ID string `uri:"id" binding:"required"`
}

func NewBookResponse(book *entity.Book) *BookResponse {
	return &BookResponse{
		ID:          book.ID.String(),
		Title:       book.Title,
		Genre:		 book.Genre,
		Author:      book.Author,
		Desc: 		 book.Desc,
		CreatedAt:   book.CreatedAt.Format(timeFormat),
		UpdatedAt:   book.UpdatedAt.Format(timeFormat),
	}
}
