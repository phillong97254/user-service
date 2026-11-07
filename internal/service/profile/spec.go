package profile

import "context"

type Service interface {
	GetProfile(ctx context.Context, input GetProfileInput) (GetProfileOutput, error)
}
