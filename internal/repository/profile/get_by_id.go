package profile

import (
	"context"

	"github.com/phil/user-service/internal/domain"
)

func (i *impl) GetByID(_ context.Context, id string) (domain.Profile, error) {
	p, ok := i.profiles[id]
	if !ok {
		return domain.Profile{}, ErrNotFound
	}
	return p, nil
}
