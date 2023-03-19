package configRouter

import "net/http"

var loginRoute = []ConfigRouter{
	{
		URI:    "/login",
		Method: http.MethodPost,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequirement: true,
	},
}
