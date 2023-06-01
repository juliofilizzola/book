package dto

import "net/http"

type ConfigRouter struct {
	URI             string
	Method          string
	Func            func(w http.ResponseWriter, r *http.Request)
	AuthRequirement bool
}
