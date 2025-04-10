package data

import (
	"github.com/doanandreas/kinder-library/internal/validator"
	"time"
)

type BookListResponse struct {
	Metadata Metadata `json:"metadata"`
	Books    []Book   `json:"books"`
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

}

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`

	Description string   `json:"description,omitempty"`
	Rating      float64  `json:"rating,omitempty"`
	Genres      []string `json:"genres,omitempty"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}
