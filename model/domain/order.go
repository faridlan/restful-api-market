package domain

import "time"

type Order struct {
	Id        int       `json:"id,omitempty"`
	User      User      `json:"user,omitempty"`
	Address   Address   `json:"address,omitempty"`
	Total     int       `json:"total,omitempty"`
	OrderDate time.Time `json:"order_date,omitempty"`
}
