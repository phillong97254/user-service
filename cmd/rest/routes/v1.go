package routes

import (
	"net/http"

	auth "github.com/phil/go-lib/authkit"
	v1 "github.com/phil/user-service/internal/handler/rest/v1"
)

func V1Routes(mux *http.ServeMux, verifier *auth.Verifier, h v1.Handler) {
	v1M2MRoutes(mux, verifier, h)
}
