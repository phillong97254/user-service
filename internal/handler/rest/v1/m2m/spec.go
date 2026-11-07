package m2m

import "github.com/phillong97254/user-service/internal/handler/rest/v1/m2m/user"

// Handler cho các endpoint service-to-service (token client_credentials).
type Handler interface {
	GetUserHandler() user.Handler
}
