package user

import "net/http"

type Handler interface {
	// GET /v1/users/{id}
	GetUser(w http.ResponseWriter, r *http.Request)
}
