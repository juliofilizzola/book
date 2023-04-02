package configRouter

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type ConfigRouter struct {
	URI             string
	Method          string
	Func            func(w http.ResponseWriter, r *http.Request)
	AuthRequirement bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRoute
	routes = append(routes, loginRoute...)

	for _, route := range routes {
		if route.AuthRequirement {
			r.HandleFunc(route.URI, middlewares.Authentication(route.Func)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Func).Methods(route.Method)
		}
	}

	return r
}
