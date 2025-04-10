package data

import (
	"github.com/doanandreas/kinder-library/internal/validator"
	"time"
)

type BookListResponse struct {
	Pagination Pagination `json:"pagination"`
	Books      []Book     `json:"books"`
}

type BookResponse struct {
	Book Book `json:"book"`
}

type BookRequest struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Pages       int      `json:"pages"`
	Description string   `json:"description"`
	Rating      float64  `json:"rating"`
	Genres      []string `json:"genres"`
}

func (b *BookRequest) Validate(v *validator.Validator) {
	v.Check(b.Title != "", "title", "must be provided")
	v.Check(b.Author != "", "author", "must be provided")
	v.Check(b.Pages != 0, "pages", "must be provided")
}

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Pages       int       `json:"pages"`
	Description string    `json:"description,omitempty"`
	Rating      float64   `json:"rating,omitempty"`
	Genres      []string  `json:"genres,omitempty"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
