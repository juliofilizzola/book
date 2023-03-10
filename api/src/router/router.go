package router

import (
	"api/src/router/configRouter"
	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return configRouter.Config(r)
}
