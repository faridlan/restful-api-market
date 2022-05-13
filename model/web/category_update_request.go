package web

type CategoryUpdateRequest struct {
	Id           int    `json:"id,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
