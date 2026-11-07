package v1

import "github.com/phillong97254/user-service/internal/handler/rest/v1/m2m"

type Handler interface {
	GetM2MHandler() m2m.Handler
}
