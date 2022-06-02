package web

type RoleResponse struct {
	Id     int    `json:"id,omitempty"`
	IdRole string `json:"id_role,omitempty"`
	Name   string `json:"name,omitempty"`
}
