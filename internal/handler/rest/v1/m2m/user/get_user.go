package user

import (
	"net/http"

	"github.com/phillong97254/go-lib/apperr"
	auth "github.com/phillong97254/go-lib/authkit"
	"github.com/phillong97254/user-service/internal/service/profile"
)

func (i *impl) GetUser(w http.ResponseWriter, r *http.Request) {
	var req GetUserRequest
	req.Bind(r)
	if err := req.Validate(); err != nil {
		apperr.Write(w, err)
		return
	}

	c, _ := auth.ClaimsFrom(r.Context())
	i.logger.Info("get user", "id", req.ID, "by", c.Subject, "client", c.ClientID)

	out, err := i.profileSvc.GetProfile(r.Context(), profile.GetProfileInput{ID: req.ID})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toGetUserResponse(out))
}
