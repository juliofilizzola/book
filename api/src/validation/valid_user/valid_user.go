package validation

import (
	"api/src/auth"
	"api/src/response"
	"net/http"
)

func ValidUser(r *http.Request, w http.ResponseWriter, ID uint64) error {
	userIdToken, err := auth.GetIdToken(r)

	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
	}

	if userIdToken != ID {
		response.Err(w, http.StatusForbidden, err)
	}

	return nil
}
