package dto

import "backend/pkg/pagination"

type PaginationDTO struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	SortBy  string `json:"sortBy"`
	OrderBy string `json:"orderBy"`
}

func (p *PaginationDTO) Normalize() {
	p.Page, p.Limit = pagination.Normalize(p.Page, p.Limit, 1, 10, 100)

	if p.SortBy == "" {
		p.SortBy = "created_at"
	}

	if p.OrderBy == "" {
		p.OrderBy = "desc"
	}
}

func (p *PaginationDTO) Offset() int {
	return pagination.Offset(p.Page, p.Limit)
}
