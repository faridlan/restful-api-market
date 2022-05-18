package web

type StatusOrderUpdate struct {
	Id         int    `json:"id,omitempty"`
	StatusName string `json:"status_name,omitempty"`
}
