package profile

import (
	"log/slog"

	"github.com/phil/user-service/internal/repository"
)

type impl struct {
	logger *slog.Logger
	repo   repository.Registry
}

func New(logger *slog.Logger, repo repository.Registry) Service {
	return &impl{logger: logger.With("service", "profile"), repo: repo}
}
