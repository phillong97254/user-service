package profile

import (
	"context"

	"github.com/phil/user-service/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (domain.Profile, error)
}
