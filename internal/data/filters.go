package data

import (
	"math"
	"strconv"

	"github.com/doanandreas/kinder-library/internal/validator"
)

type Pagination struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func CalculatePaginationData(totalRecords, page, pageSize int) Pagination {
	if totalRecords == 0 {
		return Pagination{}
	}

	return Pagination{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

type FiltersRequest struct {
	Page     string
	PageSize string
}

func (fr *FiltersRequest) Validate(v *validator.Validator) {
	_, err := strconv.Atoi(fr.Page)
	if err != nil {
		v.AddError("page", "must be an integer")
	}

	_, err = strconv.Atoi(fr.PageSize)
	if err != nil {
		v.AddError("page_size", "must be an integer")
	}
}

type Filters struct {
	Page     int
	PageSize int
}

func ParseFilters(fr *FiltersRequest) Filters {
	page, _ := strconv.Atoi(fr.Page)
	pageSize, _ := strconv.Atoi(fr.PageSize)

	res := Filters{
		Page:     page,
		PageSize: pageSize,
	}

	return res
}

func (f Filters) Validate(v *validator.Validator) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be less than or equal to 10_000_000")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be less than or equal to 100")
}

func (f Filters) Limit() int {
	return f.PageSize
}

func (f Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}
