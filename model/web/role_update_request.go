package web

type RoleUpdateRequest struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
