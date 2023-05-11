package auth

import (
	authController "api/src/controllers/auth"
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
