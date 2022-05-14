package web

type PaymentCreateRequest struct {
	OrdeId   int    `json:"orde_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}
