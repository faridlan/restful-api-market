package web

type CategoryResponse struct {
	Id           int    `json:"id,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
