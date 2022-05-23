package web

import "github.com/faridlan/restful-api-market/model"

type PaymentResponse struct {
	Id       int               `json:"id,omitempty"`
	ImageUrl *model.NullString `json:"image_url"`
}
