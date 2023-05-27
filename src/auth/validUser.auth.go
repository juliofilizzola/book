package auth

import (
	"net/http"
)

func ValidUser(r *http.Request, userId string) bool {
	userIdToken, err := GetUserId(r)
	if err != nil {
		return false
	}

	if userId == userIdToken {
		return false
	}

	return true
}
