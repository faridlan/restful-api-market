package web

import "github.com/faridlan/restful-api-market/model"

type PaymentCreateRequest struct {
	OrdeId   int               `json:"id_order,omitempty"`
	ImageUrl *model.NullString `json:"image_url,omitempty"`
}
