package profile

import (
	"log/slog"

	"github.com/phillong97254/user-service/internal/repository"
)

type impl struct {
	logger *slog.Logger
	repo   repository.Registry
}

func New(logger *slog.Logger, repo repository.Registry) Service {
	return &impl{logger: logger.With("service", "profile"), repo: repo}
}
