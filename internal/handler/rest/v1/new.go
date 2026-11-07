package v1

import (
	"log/slog"

	"github.com/phillong97254/user-service/internal/handler/rest/v1/m2m"
	"github.com/phillong97254/user-service/internal/service/profile"
)

type impl struct {
	m2m m2m.Handler
}

func New(logger *slog.Logger, profileSvc profile.Service) Handler {
	return &impl{m2m: m2m.New(logger, profileSvc)}
}

func (i *impl) GetM2MHandler() m2m.Handler { return i.m2m }
