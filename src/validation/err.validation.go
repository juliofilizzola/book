package validation

import (
	"api/src/response"
	"net/http"
)

func Err(w http.ResponseWriter, status int, err error) {
	if err != nil {
		response.Err(w, status, err)
		return
	}
}
