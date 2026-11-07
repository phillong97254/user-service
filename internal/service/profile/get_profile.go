package profile

import (
	"context"
	"errors"

	"github.com/phil/go-lib/apperr"
	profileRepo "github.com/phil/user-service/internal/repository/profile"
)

func (i *impl) GetProfile(ctx context.Context, input GetProfileInput) (GetProfileOutput, error) {
	p, err := i.repo.Profile().GetByID(ctx, input.ID)
	if errors.Is(err, profileRepo.ErrNotFound) {
		return GetProfileOutput{}, apperr.NotFound("user not found")
	}
	if err != nil {
		return GetProfileOutput{}, apperr.Internal(err)
	}
	return GetProfileOutput{ID: p.ID, Name: p.Name, Email: p.Email, Department: p.Department}, nil
}
