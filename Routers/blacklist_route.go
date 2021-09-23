package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteBlacklists(r *mux.Router) *mux.Router {

	route_match := r.PathPrefix("/report").Subrouter()
	route_match.Use(middleware.Middleware, middleware.CheckRoleBlacklist)

	route_match.HandleFunc("/blacklistPlayer", controller.BlacklistPlayer).Methods("POST")
	route_match.HandleFunc("/getBlacklists", controller.GetAllBlacklists).Methods("GET")
	route_match.HandleFunc("/getBlacklist", controller.GetBlacklist).Methods("GET")
	route_match.HandleFunc("/unblacklistPlayer", controller.UnblacklistPlayer).Methods("DELETE")

	return r
}
