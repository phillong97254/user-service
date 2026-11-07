package user

import (
	"log/slog"

	"github.com/phil/user-service/internal/service/profile"
)

type impl struct {
	logger     *slog.Logger
	profileSvc profile.Service
}

func New(logger *slog.Logger, profileSvc profile.Service) Handler {
	return &impl{logger: logger.With("handler", "user"), profileSvc: profileSvc}
}
