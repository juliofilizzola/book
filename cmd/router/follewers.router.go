package router

import (
	FollowersController "api/cmd/controllers"
	conf "api/cmd/router/dto"
	"net/http"
)

var (
	FollowerRouter = []conf.ConfigRouter{
		{
			URI:             "/follow/{userId}",
			Method:          http.MethodPost,
			Func:            FollowersController.FollowerUser,
			AuthRequirement: true,
		},
		{
			URI:             "/unfollow/{userId}",
			Method:          http.MethodPost,
			Func:            FollowersController.Unfollow,
			AuthRequirement: true,
		},
		{
			URI:             "/follow/{userId}/allFollow",
			Method:          http.MethodGet,
			Func:            FollowersController.GetFollow,
			AuthRequirement: true,
		},
		{
			URI:             "/follow/{followId}/allFollowers",
			Method:          http.MethodGet,
			Func:            FollowersController.GetFollowers,
			AuthRequirement: true,
		},
	}
)
