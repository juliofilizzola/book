package router

import (
	loginController "api/src/controllers"
	conf "api/src/router/dto"
	"net/http"
)

var (
	LoginsRoute = []conf.ConfigRouter{
		{
			URI:             "/login",
			Method:          http.MethodPost,
			Func:            loginController.Login,
			AuthRequirement: false,
		},
	}
)
