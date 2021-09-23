package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteMatches(r *mux.Router) *mux.Router {

	route_match := r.PathPrefix("/match").Subrouter()
	route_match.Use(middleware.Middleware, middleware.CheckRoleMatches)

	route_match.HandleFunc("/getMatches", controller.GetMatches).Methods("GET")
	route_match.HandleFunc("/getmatch", controller.GetMatch).Methods("GET")
	route_match.HandleFunc("/cancelmatch", controller.CancelMatch).Methods("GET")

	return r
}
