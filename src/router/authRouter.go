package router

import (
	authController "api/src/controllers"
	conf "api/src/router/dto"
	"net/http"
)

var (
	AuthsRouter = []conf.ConfigRouter{
		{
			URI:             "/auth/{userId}",
			Method:          http.MethodPost,
			Func:            authController.UpdatePassword,
			AuthRequirement: true,
		},
	}
)
