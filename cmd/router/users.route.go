package router

import (
	UserControllers "api/cmd/controllers"
	conf "api/cmd/router/dto"
	"net/http"
)

var (
	User = []conf.ConfigRouter{
		// create user
		{
			URI:             "/users",
			Method:          http.MethodPost,
			Func:            UserControllers.CreateUser,
			AuthRequirement: false,
		},
		// get all user
		{
			URI:             "/users",
			Method:          http.MethodGet,
			Func:            UserControllers.GetUsers,
			AuthRequirement: false,
		},
		// get unique user
		{
			URI:             "/users/{id}",
			Method:          http.MethodGet,
			Func:            UserControllers.GetUser,
			AuthRequirement: false,
		},

		{
			URI:             "/users/{id}",
			Method:          http.MethodPut,
			Func:            UserControllers.UpdateUser,
			AuthRequirement: false,
		},
		{
			URI:             "/users/{id}",
			Method:          http.MethodDelete,
			Func:            UserControllers.DeleteUser,
			AuthRequirement: true,
		},
	}
)
