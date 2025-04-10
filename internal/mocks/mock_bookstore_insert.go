package mocks

import (
	"github.com/doanandreas/kinder-library/internal/data"
	"github.com/doanandreas/kinder-library/internal/repository"
)

func InsertBookMock(existTitle string) *MockBookStore {
	return &MockBookStore{
		InsertFunc: func(book *data.Book) error {
			if book.Title == existTitle {
				return repository.ErrDuplicateTitle
			}
			return nil
		},
	}
}
