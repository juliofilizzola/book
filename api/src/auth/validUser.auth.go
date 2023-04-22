package auth

import "net/http"

func ValidUser(r *http.Request, userId uint64) bool {
	userIdToken, err := GetUserId(r)
	if err != nil {
		return false
	}

	if userId == userIdToken {
		return false
	}

	return true
}
