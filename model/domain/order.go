package domain

import (
	"time"

	"github.com/faridlan/restful-api-market/model"
)

type Order struct {
	Id        int               `json:"id,omitempty"`
	User      User              `json:"user,omitempty"`
	Address   Address           `json:"address,omitempty"`
	Total     int               `json:"total,omitempty"`
	OrderDate time.Time         `json:"order_date,omitempty"`
	Status    StatusOrder       `json:"status_id,omitempty"`
	Payment   *model.NullString `json:"payment"`
}
