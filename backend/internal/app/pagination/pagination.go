package pagination

import (
	"gorm.io/gorm"

	"private-llm-backend/pkg/paginator"
)

type Pagination interface {
	// Page returns current page
	Page() (int, error)
	// Nums returns the total number of records
	Nums() (int64, error)
	// HasPages returns true if there is more than one page
	HasPages() (bool, error)
	// HasNext returns true if current page is not the last page
	HasNext() (bool, error)
	// PrevPage returns previous page number or paginator.ErrNoPrevPage if current page is first page
	PrevPage() (int, error)
	// NextPage returns next page number or paginator.ErrNoNextPage if current page is last page
	NextPage() (int, error)
	// HasPrev returns true if current page is not the first page
	HasPrev() (bool, error)
	// PageNums returns the total number of pages
	PageNums() (int, error)
}

type Params interface {
	GetPaginator(query *gorm.DB) paginator.Paginator
}
