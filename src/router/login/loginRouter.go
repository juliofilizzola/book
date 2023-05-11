package login

import (
	loginController "api/src/controllers/login"
	conf "api/src/router/dto"
	"net/http"
)

var (
	LoginsRoute = []conf.ConfigRouter{
		{
			URI:             "/auth/update-password/{userId}",
			Method:          http.MethodPost,
			Func:            loginController.Login,
			AuthRequirement: true,
		},
	}
)
