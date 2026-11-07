package profile

import (
	"context"

	"github.com/phillong97254/user-service/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (domain.Profile, error)
}
