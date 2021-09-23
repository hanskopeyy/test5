package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
	"github.com/sen329/test5/Routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db := controller.Open()

	// defer db.Close()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	router := mux.NewRouter()

	router.HandleFunc("/api/login", controller.Login).Methods("POST")

	// ---- normal route ---- //
	route := router.PathPrefix("/api").Subrouter()
	route.Use(middleware.Middleware)
	route.HandleFunc("/test", controller.Test).Methods("GET")

	route.Handle("/", Routers.RouteIcon(route))
	route.Handle("/", Routers.RouteKsatriya(route))

	//	rune
	route.Handle("/", Routers.RouteRune(route))

	//	premium
	route.Handle("/", Routers.RoutePremium(route))

	//	energy
	route.Handle("/", Routers.RouteEnergy(route))

	//currency type
	route.Handle("/", Routers.RouteCurrency(route))

	//misc items
	route.Handle("/", Routers.RouteMiscItem(route))

	//box
	route.Handle("/", Routers.RouteBox(route))

	//chest
	route.Handle("/", Routers.RouteChest(route))

	// ---- Mail Subroute ---- //
	route.Handle("/", Routers.RouteMail(route))

	// ---- News Subroute ---- //
	route.Handle("/", Routers.RouteNews(route))

	// ---- Lotto Subroute ---- //
	route.Handle("/", Routers.RouteLotto(route))

	// ---- Gacha Subroute ---- //
	route.Handle("/", Routers.RouteGacha(route))

	// ---- Shop Subroute ---- //
	route.Handle("/", Routers.RouteLotus(route))
	route.Handle("/", Routers.RouteShop(route))

	// ---- Role Subroute ---- //
	route.Handle("/", Routers.RouteRole(route))

	route.Handle("/", Routers.RoutePlayers(route))

	route.Handle("/", Routers.RouteKsaRot(route))

	route.Handle("/", Routers.RouteUser(route))

	route.Handle("/", Routers.RouteMatches(route))

	route.Handle("/", Routers.RouteReports(route))

	route.Handle("/", Routers.RouteBlacklists(route))

	route.Handle("/", Routers.RouteVouchers(route))

	route.Handle("/", Routers.RouteWarning(route))

	route.Handle("/", Routers.RouteJudges(route))

	route.Handle("/", Routers.RoutePlayerStats(route))

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
