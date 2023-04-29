package auth

import (
	"fmt"
	"net/http"
)

func ValidUser(r *http.Request, userId uint64) bool {
	userIdToken, err := GetUserId(r)
	if err != nil {
		fmt.Printf("oi")
		return false
	}

	if userId == userIdToken {
		return false
	}

	return true
}
