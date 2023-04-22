package followers

import (
	FollowersController "api/src/controllers/followers"
	conf "api/src/router/dto"

	"net/http"
)

var FollowersRouter = []conf.ConfigRouter{
	{
		URI:             "/followers/{userId}",
		Method:          http.MethodPost,
		Func:            FollowersController.FollowerUser,
		AuthRequirement: true,
	},
}
