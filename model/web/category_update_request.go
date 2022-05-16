package web

type CategoryUpdateRequest struct {
	Id           int    `json:"id,omitempty" validate:"required"`
	CategoryName string `json:"category_name,omitempty" validate:"required"`
}
