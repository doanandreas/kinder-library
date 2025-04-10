package mocks

import "github.com/doanandreas/kinder-library/internal/data"

type MockBookStore struct {
	InsertFunc func(book *data.Book) error
	UpdateFunc func(book *data.Book) error
	DeleteFunc func(id int64) error
	ListFunc   func(filters data.Filters) ([]data.Book, data.Pagination, error)
}

func (m *MockBookStore) Insert(book *data.Book) error {
	return m.InsertFunc(book)
}

func (m *MockBookStore) Update(book *data.Book) error {
	return m.UpdateFunc(book)
}

func (m *MockBookStore) Delete(id int64) error {
	return m.DeleteFunc(id)
}

func (m *MockBookStore) List(filters data.Filters) ([]data.Book, data.Pagination, error) {
	return m.ListFunc(filters)
}
