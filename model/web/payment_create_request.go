package web

import "github.com/faridlan/restful-api-market/model"

type PaymentCreateRequest struct {
	OrdeId   int               `json:"orde_id,omitempty"`
	ImageUrl *model.NullString `json:"image_url,omitempty"`
}
