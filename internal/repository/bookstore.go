package repository

import (
	"database/sql"

	"github.com/doanandreas/kinder-library/internal/data"

	"github.com/lib/pq"
)

type BookStore interface {
	Insert(book *data.Book) error
	Update(book *data.Book) error
	Delete(id int64) error
	List(filters data.Filters) ([]data.Book, data.Pagination, error)
}

type PGBookStore struct {
	DB *sql.DB
}

func (pg *PGBookStore) Insert(book *data.Book) error {
	query := `
		INSERT INTO books (title, author, pages, description, rating, genres)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at`

	args := []any{book.Title, book.Author, book.Pages, book.Description, book.Rating, pq.Array(book.Genres)}

	return pg.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
}

func (pg *PGBookStore) Update(book *data.Book) error {
	if book.ID < 1 {
		return sql.ErrNoRows
	}

	query := `
		UPDATE books
		SET title = $1, author = $2, pages = $3, description = $4, rating = $5, genres = $6, updated_at = NOW()
		WHERE id = $7
		RETURNING id, created_at, updated_at`

	args := []any{
		book.Title,
		book.Author,
		book.Pages,
		book.Description,
		book.Rating,
		pq.Array(book.Genres),
		book.ID,
	}

	return pg.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
}

func (pg *PGBookStore) Delete(id int64) error {
	if id < 1 {
		return sql.ErrNoRows
	}

	query := `DELETE FROM books WHERE id = $1`
	res, err := pg.DB.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (pg *PGBookStore) List(filters data.Filters) ([]data.Book, data.Pagination, error) {
	books := []data.Book{
		{
			ID:          1,
			Title:       "The Pragmatic Programmer",
			Author:      "Andrew Hunt and David Thomas",
			Pages:       352,
			Description: "A classic book offering practical advice for programmers on how to think and work effectively.",
			Rating:      4.7,
			Genres:      []string{"Programming", "Software Engineering"},
		},
		{
			ID:          2,
			Title:       "Clean Code",
			Author:      "Robert C. Martin",
			Pages:       464,
			Description: "Teaches best practices for writing clean, maintainable, and efficient code.",
			Rating:      4.6,
			Genres:      []string{"Programming", "Software Engineering"},
		},
		{
			ID:          3,
			Title:       "Go Programming Language",
			Author:      "Alan A. A. Donovan and Brian W. Kernighan",
			Pages:       380,
			Description: "A thorough introduction to the Go programming language by one of its creators.",
			Rating:      4.5,
			Genres:      []string{"Programming", "Go"},
		},
		{
			ID:          4,
			Title:       "Design Patterns: Elements of Reusable Object-Oriented Software",
			Author:      "Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides",
			Pages:       395,
			Description: "The foundational book on design patterns in object-oriented programming.",
			Rating:      4.4,
			Genres:      []string{"Programming", "Design Patterns"},
		},
		{
			ID:          5,
			Title:       "Refactoring: Improving the Design of Existing Code",
			Author:      "Martin Fowler",
			Pages:       448,
			Description: "Explains how to restructure existing code for improved readability and performance.",
			Rating:      4.6,
			Genres:      []string{"Programming", "Refactoring"},
		},
	}

	pagination := data.Pagination{
		CurrentPage:  filters.Page,
		PageSize:     filters.PageSize,
		FirstPage:    1,
		LastPage:     10,
		TotalRecords: 61,
	}

	return books, pagination, nil
}
