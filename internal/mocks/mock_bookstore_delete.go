package mocks

import "github.com/doanandreas/kinder-library/internal/repository"

func DeleteBookMock(existId int64) *MockBookStore {
	return &MockBookStore{
		DeleteFunc: func(id int64) error {
			if id < 1 {
				return repository.ErrRecordNotFound
			}

			if id != existId {
				return repository.ErrRecordNotFound
			}

			return nil
		},
	}
}
