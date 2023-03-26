package configRouter

import (
	loginController "api/src/controllers/login"
	"net/http"
)

var loginRoute = []ConfigRouter{
	{
		URI:             "/login",
		Method:          http.MethodPost,
		Func:            loginController.Login,
		AuthRequirement: false,
	},
}
