package web

type StatusOrderUpdate struct {
	Id            int    `json:"id,omitempty"`
	IdStatusOrder string `json:"id_status_order,omitempty"`
	StatusName    string `json:"status_name,omitempty"`
}
