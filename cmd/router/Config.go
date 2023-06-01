package router

import (
	"api/cmd/middlewares"
	"github.com/gorilla/mux"
)

func Config(r *mux.Router) *mux.Router {
	routes := User
	routes = append(routes, LoginsRoute...)
	routes = append(routes, FollowerRouter...)
	routes = append(routes, AuthsRouter...)
	routes = append(routes, PublicationRouter...)

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
