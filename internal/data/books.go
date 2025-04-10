package data

import "time"

type BookListResponse struct {
	Metadata Metadata `json:"metadata"`
	Books    []Book   `json:"books"`
}

type BookResponse struct {
	Book Book `json:"book"`
}

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int32  `json:"pages"`

	Description string   `json:"description,omitempty"`
	Rating      float64  `json:"rating,omitempty"`
	Genres      []string `json:"genres,omitempty"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Metadata struct {
	CurrentPage  int32 `json:"current_page,omitempty"`
	PageSize     int32 `json:"PageSize,omitempty"`
	FirstPage    int32 `json:"first_page,omitempty"`
	LastPage     int32 `json:"last_page,omitempty"`
	TotalRecords int32 `json:"total_records,omitempty"`
}
