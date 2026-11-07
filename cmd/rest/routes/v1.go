package routes

import (
	"net/http"

	auth "github.com/phillong97254/go-lib/authkit"
	v1 "github.com/phillong97254/user-service/internal/handler/rest/v1"
)

func V1Routes(mux *http.ServeMux, verifier *auth.Verifier, h v1.Handler) {
	v1M2MRoutes(mux, verifier, h)
}
