package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RoutePlayerStats(r *mux.Router) *mux.Router {

	route_player := r.PathPrefix("/stats").Subrouter()
	route_player.Use(middleware.Middleware, middleware.CheckRolePlayerStats)

	route_player.HandleFunc("/getDailyUser", controller.GetDailyUserCount).Methods("GET")
	route_player.HandleFunc("/getDailyUserUnique", controller.GetDailyUserCountUnique).Methods("GET")
	route_player.HandleFunc("/getConcurrentUser", controller.GetConcurrentUserCount).Methods("GET")

	return r
}
