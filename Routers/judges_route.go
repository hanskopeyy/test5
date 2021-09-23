package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteJudges(r *mux.Router) *mux.Router {

	route_judge := r.PathPrefix("/judge").Subrouter()
	route_judge.Use(middleware.Middleware, middleware.CheckRoleJudge)

	route_judge.HandleFunc("/addJudge", controller.RegisterJudge).Methods("POST")
	route_judge.HandleFunc("/getAllJudges", controller.GetAllJudge).Methods("GET")
	route_judge.HandleFunc("/getJudge", controller.GetJudge).Methods("GET")
	route_judge.HandleFunc("/updateJudgeUsername", controller.UpdateJudgeName).Methods("PUT")
	route_judge.HandleFunc("/updateJudgePassword", controller.UpdateJudgePassword).Methods("PUT")
	route_judge.HandleFunc("/deleteJudge", controller.DeleteJudge).Methods("DELETE")

	return r
}
