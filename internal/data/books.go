package data

import (
	"math"
	"time"

	"github.com/doanandreas/kinder-library/internal/validator"
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
	v.Check(len(b.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(b.Author != "", "author", "must be provided")
	v.Check(len(b.Author) <= 200, "author", "must not be more than 200 bytes long")

	v.Check(b.Pages != 0, "pages", "must be provided")
	v.Check(b.Pages > 0, "pages", "must be a positive integer")

	v.Check(len(b.Description) <= 1000, "description", "must not be more than 1000 bytes long")

	v.Check(b.Rating >= 0 && b.Rating <= 5, "rating", "must be between 0.00 and 5.00")
	v.Check(hasTwoDecimalPlaces(b.Rating), "rating", "must be two decimal places at max")

	v.Check(len(b.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(b.Genres), "genres", "must not contain duplicate values")
	v.Check(validator.ContainsEmptyString(b.Genres) == false, "genres", "must not contain empty values")
}

func hasTwoDecimalPlaces(f float64) bool {
	rounded := math.Round(f*100) / 100
	return math.Abs(f-rounded) < 1e-6
}

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Pages       int       `json:"pages"`
	Description string    `json:"description,omitempty"`
	Rating      float64   `json:"rating,omitempty"`
	Genres      []string  `json:"genres,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
