package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteWarning(r *mux.Router) *mux.Router {

	route_premium := r.PathPrefix("/warning").Subrouter()
	route_premium.Use(middleware.Middleware)

	route_premium.HandleFunc("/ksaRot", controller.GetWarningKsaRotation).Methods("GET")
	route_premium.HandleFunc("/gacha", controller.GetWarningGacha).Methods("GET")
	route_premium.HandleFunc("/lotto", controller.GetWarningLotto).Methods("GET")
	route_premium.HandleFunc("/lotus", controller.GetWarningLotus).Methods("GET")
	route_premium.HandleFunc("/season", controller.GetWarningSeason).Methods("GET")

	return r
}
