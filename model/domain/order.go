package domain

import (
	"time"

	"github.com/faridlan/restful-api-market/model"
)

type Order struct {
	Id        int
	User      User
	Address   Address
	Total     int
	OrderDate time.Time
	Status    StatusOrder
	Payment   *model.NullString
	// Payment string `json:"payment"`
}
