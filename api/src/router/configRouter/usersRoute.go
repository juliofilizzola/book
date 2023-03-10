package configRouter

import (
	UserControllers "api/src/controllers"
	"net/http"
)

var userRoute = []ConfigRouter{
	// create user
	{
		URI:             "/users",
		Method:          http.MethodPost,
		Func:            UserControllers.CreateUser,
		AuthRequeriment: false,
	},
	// get all user
	{
		URI:             "/users",
		Method:          http.MethodGet,
		Func:            UserControllers.GetUsers,
		AuthRequeriment: false,
	},
	// get unique user
	{
		URI:             "/users/{id}",
		Method:          http.MethodGet,
		Func:            UserControllers.GetUser,
		AuthRequeriment: false,
	},

	{
		URI:             "/users/{id}",
		Method:          http.MethodPatch,
		Func:            UserControllers.UpdateUser,
		AuthRequeriment: false,
	},
	{
		URI:             "/users/{id}",
		Method:          http.MethodDelete,
		Func:            UserControllers.DeleteUser,
		AuthRequeriment: false,
	},
}
