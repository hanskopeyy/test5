package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteReports(r *mux.Router) *mux.Router {

	route_match := r.PathPrefix("/report").Subrouter()
	route_match.Use(middleware.Middleware, middleware.CheckRoleReport)

	route_match.HandleFunc("/getReports", controller.GetAllPlayerReports).Methods("GET")
	route_match.HandleFunc("/getReportsbyRoom", controller.GetAllPlayerReportsByRoom).Methods("GET")
	route_match.HandleFunc("/getReportsbyUser", controller.GetAllPlayerReportsByUser).Methods("GET")
	route_match.HandleFunc("/getReport", controller.GetPlayerReport).Methods("GET")

	return r
}
