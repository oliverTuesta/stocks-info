package domain

type PaginationRequest struct {
	Page     int    `json:"page" form:"page" binding:"min=1"`
	PageSize int    `json:"page_size" form:"page_size" binding:"min=1,max=100"`
	Search   string `json:"search" form:"search"`
}

type PaginationResponse struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type PaginatedResult[T any] struct {
	Data       []T                `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

func (p *PaginationResponse) CalculateTotalPages() {
	if p.PageSize > 0 {
		p.TotalPages = int((p.Total + int64(p.PageSize) - 1) / int64(p.PageSize))
	}
}
