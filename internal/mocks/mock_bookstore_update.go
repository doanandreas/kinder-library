package mocks

import (
	"github.com/doanandreas/kinder-library/internal/data"
	"github.com/doanandreas/kinder-library/internal/repository"
)

func UpdateBookMock(existId int64, existTitle string) *MockBookStore {
	return &MockBookStore{
		UpdateFunc: func(book *data.Book) error {
			if book.ID < 1 {
				return repository.ErrRecordNotFound
			}

			if book.ID != existId {
				return repository.ErrRecordNotFound
			}

			if book.Title == existTitle {
				return repository.ErrDuplicateTitle
			}

			return nil
		},
	}
}
