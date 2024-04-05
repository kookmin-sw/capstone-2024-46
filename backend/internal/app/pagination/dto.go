package pagination

import (
	"errors"
	"strconv"

	"gorm.io/gorm"

	"private-llm-backend/pkg/errorutil"
	"private-llm-backend/pkg/paginator"
	"private-llm-backend/pkg/pointerutil"
	"private-llm-backend/pkg/typeconv"
)

const defaultLimit = 10

type pagePaginationParams struct {
	Limit int
	Page  int
}

func (p *pagePaginationParams) GetPaginator(query *gorm.DB) paginator.Paginator {
	pg := paginator.New(paginator.NewGORMAdapter(query), p.Limit)
	pg.SetPage(p.Page)
	return pg
}

func NewPaginationParams(pageToken *string, pageSize *int32) Params {
	return &pagePaginationParams{
		Limit: withDefaultLimit(pageSize),
		Page:  withDefaultPage(pageToken),
	}
}

func withDefaultLimit(size *int32) int {
	if size == nil || *size <= 0 {
		return defaultLimit
	}
	return int(*size)
}

func withDefaultPage(pageToken *string) int {
	if pageToken == nil {
		return 1
	}
	page, _ := strconv.Atoi(*pageToken)
	if page <= 0 {
		return 1
	}
	return page
}

type PageToken struct {
	NextPageToken     *string
	PreviousPageToken *string
	TotalItemCount    string
}

func GetPageToken(p Pagination) (*PageToken, error) {
	totalCount, err := p.Nums()
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to get total count"))
	}
	var nextPageToken, previousPageToken *string
	nextPage, err := p.NextPage()
	if err == nil {
		nextPageToken = pointerutil.String(strconv.Itoa(nextPage))
	} else if !errors.Is(err, paginator.ErrNoNextPage) {
		return nil, errorutil.WithDetail(err, errors.New("failed to get next page"))
	}
	previousPage, err := p.PrevPage()
	if err == nil {
		previousPageToken = pointerutil.String(strconv.Itoa(previousPage))
	} else if !errors.Is(err, paginator.ErrNoPrevPage) {
		return nil, errorutil.WithDetail(err, errors.New("failed to get previous page"))
	}
	return &PageToken{
		NextPageToken:     nextPageToken,
		PreviousPageToken: previousPageToken,
		TotalItemCount:    typeconv.I64toA(totalCount),
	}, nil
}
