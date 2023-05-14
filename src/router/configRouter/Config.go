package configRouter

import (
	"api/src/middlewares"
	"api/src/router/auth"
	"api/src/router/followers"
	"api/src/router/login"
	"api/src/router/publication"
	"api/src/router/users"
	"github.com/gorilla/mux"
)

func Config(r *mux.Router) *mux.Router {
	routes := users.UserRoute
	routes = append(routes, login.LoginsRoute...)
	routes = append(routes, followers.FollowerRouter...)
	routes = append(routes, auth.AuthsRouter...)
	routes = append(routes, publication.Router...)

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
