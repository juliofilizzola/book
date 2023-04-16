package configRouter

import (
	"api/src/middlewares"
	"api/src/router/followers"
	"api/src/router/login"
	"api/src/router/users"
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
	routes := users.UserRoute
	routes = append(routes, login.LoginRoute...)
	routes = append(routes, followers.FollowersRouter...)

	for _, route := range routes {
		if route.AuthRequirement {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authentication(route.Func)),
			).Methods(route.Method)

		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
