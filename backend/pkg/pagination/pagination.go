package pagination

import "math"

type Meta struct {
	Page        int   `json:"page"`
	Limit       int   `json:"limit"`
	TotalData   int64 `json:"totalData"`
	TotalPage   int   `json:"totalPage"`
	HasNext     bool  `json:"hasNext"`
	HasPrevious bool  `json:"hasPrevious"`
}

func Normalize(page int, limit int, defaultPage int, defaultLimit int, maxLimit int) (int, int) {
	if page <= 0 {
		page = defaultPage
	}

	if limit <= 0 {
		limit = defaultLimit
	}

	if maxLimit > 0 && limit > maxLimit {
		limit = maxLimit
	}

	return page, limit
}

func Offset(page int, limit int) int {
	return (page - 1) * limit
}

func BuildMeta(page int, limit int, totalData int64) Meta {
	totalPage := 0
	if limit > 0 && totalData > 0 {
		totalPage = int(math.Ceil(float64(totalData) / float64(limit)))
	}

	return Meta{
		Page:        page,
		Limit:       limit,
		TotalData:   totalData,
		TotalPage:   totalPage,
		HasNext:     page < totalPage,
		HasPrevious: page > 1,
	}
}
