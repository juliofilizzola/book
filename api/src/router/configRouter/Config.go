package configRouter

import (
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
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
