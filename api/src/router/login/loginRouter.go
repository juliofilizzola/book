package login

import (
	loginController "api/src/controllers/login"
	"api/src/router/configRouter"
	"net/http"
)

var LoginRoute = []configRouter.ConfigRouter{
	{
		URI:             "/login",
		Method:          http.MethodPost,
		Func:            loginController.Login,
		AuthRequirement: false,
	},
}
