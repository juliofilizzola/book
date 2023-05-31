package router

import (
	loginController "api/cmd/controllers"
	conf "api/cmd/router/dto"
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
