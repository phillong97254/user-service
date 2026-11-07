package routes

import (
	"net/http"

	auth "github.com/phillong97254/go-lib/authkit"
	v1 "github.com/phillong97254/user-service/internal/handler/rest/v1"
)

// Backoffice gọi các endpoint này bằng token của CHÍNH backoffice (client_credentials), cần scope users:read.
func v1M2MRoutes(mux *http.ServeMux, verifier *auth.Verifier, h v1.Handler) {
	authn := auth.Authenticate(verifier)
	usersRead := auth.Policy{Scopes: []string{"users:read"}}

	userHandler := h.GetM2MHandler().GetUserHandler()
	mux.Handle("GET /v1/users/{id}", authn(auth.Require(usersRead, userHandler.GetUser)))
}
