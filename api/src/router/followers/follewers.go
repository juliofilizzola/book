package followers

import (
	FollowersController "api/src/controllers/followers"
	"api/src/router/configRouter"
	"net/http"
)

var FollowersRouter = []configRouter.ConfigRouter{
	{
		URI:             "/followers/{userId}",
		Method:          http.MethodPost,
		Func:            FollowersController.FollowerUser,
		AuthRequirement: true,
	},
}
