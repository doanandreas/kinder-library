package repository

import (
	"database/sql"
	"errors"
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

	err := pg.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return ErrDuplicateTitle
		}
		return err
	}

	return nil
}

func (pg *PGBookStore) Update(book *data.Book) error {
	if book.ID < 1 {
		return ErrRecordNotFound
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

	err := pg.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrRecordNotFound
		}

		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return ErrDuplicateTitle
		}
		return err
	}

	return nil
}

func (pg *PGBookStore) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
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
		return ErrRecordNotFound
	}

	return nil
}

func (pg *PGBookStore) List(filters data.Filters) ([]data.Book, data.Pagination, error) {
	query := `
		SELECT count(*) OVER(), id, title, author, pages, description, rating, genres, created_at, updated_at
		FROM books
		ORDER BY title ASC
		LIMIT $1 OFFSET $2`

	args := []any{filters.Limit(), filters.Offset()}

	rows, err := pg.DB.Query(query, args...)
	if err != nil {
		return nil, data.Pagination{}, err
	}
	defer rows.Close()

	totalRecords := 0
	books := []data.Book{}

	for rows.Next() {
		var book data.Book

		err := rows.Scan(
			&totalRecords,
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Pages,
			&book.Description,
			&book.Rating,
			pq.Array(&book.Genres),
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return nil, data.Pagination{}, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, data.Pagination{}, err
	}

	if filters.Page > 1 && totalRecords == 0 {
		return nil, data.Pagination{}, ErrPageOutOfBounds
	}

	pagination := data.CalculatePaginationData(totalRecords, filters.Page, filters.PageSize)

	return books, pagination, nil
}
