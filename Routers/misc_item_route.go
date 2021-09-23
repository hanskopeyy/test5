package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteMiscItem(r *mux.Router) *mux.Router {

	route_chest := r.PathPrefix("/miscItem").Subrouter()
	route_chest.Use(middleware.Middleware)

	route_chest.HandleFunc("/getMiscItems", controller.GetMiscItems).Methods("GET")
	route_chest.HandleFunc("/getMiscItem", controller.GetMiscItem).Methods("GET")

	return r
}
