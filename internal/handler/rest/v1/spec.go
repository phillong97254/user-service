package v1

import "github.com/phil/user-service/internal/handler/rest/v1/m2m"

type Handler interface {
	GetM2MHandler() m2m.Handler
}
