package router

import "github.com/gorilla/mux"

func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}
