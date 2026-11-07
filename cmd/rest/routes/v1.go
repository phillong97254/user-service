package routes

import (
	"net/http"

	v1 "github.com/phil/user-service/internal/handler/rest/v1"
	auth "github.com/phil/user-service/pkg/authkit"
)

func V1Routes(mux *http.ServeMux, verifier *auth.Verifier, h v1.Handler) {
	v1M2MRoutes(mux, verifier, h)
}
