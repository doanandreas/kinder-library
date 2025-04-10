package repository

import "errors"

type Models struct {
	Books BookStore
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateTitle = errors.New("duplicate title")
)
