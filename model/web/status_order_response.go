package web

type StatusOrderResponse struct {
	Id         int    `json:"id,omitempty"`
	IdStatus   string `json:"id_status,omitempty"`
	StatusName string `json:"status_name,omitempty"`
}
