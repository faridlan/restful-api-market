package service

import "context"

type SeederService interface {
	Delete(ctx context.Context)
}
