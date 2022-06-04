package web

type RoleUpdateRequest struct {
	Id     int    `json:"id,omitempty"`
	IdRole string `json:"id_role,omitempty"`
	Name   string `json:"role_name,omitempty"`
}
