package login

import (
	loginController "api/src/controllers/login"
	conf "api/src/router/dto"
	"net/http"
)

var LoginRoute = []conf.ConfigRouter{
	{
		URI:             "/login",
		Method:          http.MethodPost,
		Func:            loginController.Login,
		AuthRequirement: false,
	},
}
