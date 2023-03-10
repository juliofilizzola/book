package router

import "net/http"

type ConfigRoute struct {
	URI             string
	Method          string
	Func            func(w http.ResponseWriter, r *http.Request)
	AuthRequeriment bool
}
