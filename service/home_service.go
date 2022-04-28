package service

import (
	"context"

	"github.com/faridlan/restful-api-market/model/web"
)

type HomeService interface {
	Product(ctx context.Context) []web.ProductResponse
}
