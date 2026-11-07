package m2m

import (
	"log/slog"

	"github.com/phillong97254/user-service/internal/handler/rest/v1/m2m/user"
	"github.com/phillong97254/user-service/internal/service/profile"
)

type impl struct {
	user user.Handler
}

func New(logger *slog.Logger, profileSvc profile.Service) Handler {
	return &impl{user: user.New(logger, profileSvc)}
}

func (i *impl) GetUserHandler() user.Handler { return i.user }
