package router

import (
	authController "api/cmd/controllers"
	conf "api/cmd/router/dto"
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
