package publication

import (
	loginController "api/src/controllers/login"
	conf "api/src/router/dto"
	"net/http"
)

var (
	Router = []conf.ConfigRouter{
		{
			URI:             "/publication",
			Method:          http.MethodPost,
			Func:            loginController.Login,
			AuthRequirement: true,
		},
	}
)
