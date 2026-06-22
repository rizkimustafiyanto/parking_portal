package dto

import "backend/pkg/pagination"

type PaginationDTO struct {
	Page    int    `form:"page" json:"page"`
	Limit   int    `form:"limit" json:"limit"`
	SortBy  string `form:"sortBy" json:"sortBy"`
	OrderBy string `form:"orderBy" json:"orderBy"`
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
