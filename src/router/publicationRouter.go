package router

import (
	"api/src/controllers"
	conf "api/src/router/dto"
	"net/http"
)

var (
	PublicationRouter = []conf.ConfigRouter{
		{
			URI:             "/publication",
			Method:          http.MethodPost,
			Func:            controllers.Create,
			AuthRequirement: true,
		},
	}
)