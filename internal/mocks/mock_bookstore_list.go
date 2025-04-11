package mocks

import (
	"time"

	"github.com/doanandreas/kinder-library/internal/data"
)

func ListBookMock() *MockBookStore {
	return &MockBookStore{
		ListFunc: func(filters data.Filters) ([]data.Book, data.Pagination, error) {
			mockBooks := []data.Book{
				{
					ID:          1,
					Title:       "Mock Title 1",
					Author:      "Author 1",
					Pages:       100,
					Description: "This is a mock book.",
					Rating:      4.51,
					Genres:      []string{"fiction", "mock"},
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					ID:          2,
					Title:       "Mock Title 2",
					Author:      "Author 2",
					Pages:       200,
					Description: "This is another mock book.",
					Rating:      4.24,
					Genres:      []string{"science", "mock"},
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			}

			totalRecords := len(mockBooks)
			pagination := data.CalculatePaginationData(totalRecords, filters.Page, filters.PageSize)

			return mockBooks, pagination, nil
		},
	}
}
