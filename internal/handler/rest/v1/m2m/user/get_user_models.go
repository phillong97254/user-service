package user

import (
	"net/http"
	"strings"

	"github.com/phil/go-lib/apperr"
	"github.com/phil/user-service/internal/service/profile"
)

type GetUserRequest struct {
	ID string
}

func (r *GetUserRequest) Bind(req *http.Request) { r.ID = strings.TrimSpace(req.PathValue("id")) }

func (r GetUserRequest) Validate() error {
	if r.ID == "" {
		return apperr.BadRequest("id is required")
	}
	return nil
}

type GetUserResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

func toGetUserResponse(out profile.GetProfileOutput) GetUserResponse {
	return GetUserResponse{ID: out.ID, Name: out.Name, Email: out.Email, Department: out.Department}
}
