package dto

type PaginationQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (p *PaginationQuery) Normalize() (page, limit, offset int) {
	page = p.Page
	if page < 1 {
		page = 1
	}
	limit = p.Limit
	if limit < 1 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}
	offset = (page - 1) * limit
	return page, limit, offset
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

func NewPaginationMeta(page, limit int, total int64) PaginationMeta {
	totalPages := 0
	if limit > 0 && total > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}
	return PaginationMeta{
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: totalPages,
	}
}