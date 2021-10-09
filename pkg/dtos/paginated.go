package dtos

type PaginatedData struct {
	Page       int64       `json:"page,omitempty"`
	PerPage    int64       `json:"per_page,omitempty"`
	Total      int64       `json:"total,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
