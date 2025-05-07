package models

type PaginatedResponse struct {
	Rows       interface{} `json:"rows"`
	Page       float64     `json:"page"`
	Limit      float64     `json:"limit"`
	TotalRows  float64     `json:"totalRows"`
	TotalPages float64     `json:"totalPages"`
}
