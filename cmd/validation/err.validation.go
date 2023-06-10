package validation

import (
	"api/cmd/response"
	"log"
	"net/http"
)

func Err(w http.ResponseWriter, status int, err error) {
	if err != nil {
		response.Err(w, status, err)
		log.Fatal(err)
	}
}
