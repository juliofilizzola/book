package router

import (
	"api/cmd/controllers"
	conf "api/cmd/router/dto"
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
		{
			URI:             "/publication",
			Method:          http.MethodGet,
			Func:            controllers.GetMyPublications,
			AuthRequirement: true,
		},

		{
			URI:             "/publication/all",
			Method:          http.MethodGet,
			Func:            controllers.GetAllPublication,
			AuthRequirement: true,
		},

		{
			URI:             "/publication/find/{id}",
			Method:          http.MethodGet,
			Func:            controllers.GetPublication,
			AuthRequirement: true,
		},
		{
			URI:             "/publication/up/{id}",
			Method:          http.MethodPatch,
			Func:            controllers.UpdatePublication,
			AuthRequirement: true,
		},

		{
			URI:             "/publication/del/{id}",
			Method:          http.MethodDelete,
			Func:            controllers.DeletedPublication,
			AuthRequirement: true,
		},

		{
			URI:             "/publication/like/{id}",
			Method:          http.MethodPatch,
			Func:            controllers.LikePublication,
			AuthRequirement: true,
		},

		{
			URI:             "/publication/dislike/{id}",
			Method:          http.MethodPatch,
			Func:            controllers.DislikePublication,
			AuthRequirement: true,
		},
	}
)
