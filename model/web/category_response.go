package web

type CategoryResponse struct {
	Id           int    `json:"id,omitempty"`
	IdCategory   string `json:"id_category,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
