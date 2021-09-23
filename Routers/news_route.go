package Routers

import (
	"github.com/gorilla/mux"
	newsinfo "github.com/sen329/test5/Controller/news_info"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteNews(r *mux.Router) *mux.Router {

	route_news := r.PathPrefix("/news").Subrouter()
	route_news.Use(middleware.Middleware)

	route_news.HandleFunc("/add", newsinfo.AddNews).Methods("POST")
	route_news.HandleFunc("/getAll", newsinfo.GetAllNews).Methods("GET")
	route_news.HandleFunc("/get", newsinfo.GetNews).Methods("GET")
	route_news.HandleFunc("/update", newsinfo.UpdateNews).Methods("PUT")
	route_news.HandleFunc("/delete", newsinfo.DeleteNews).Methods("DELETE")

	route_news.HandleFunc("/addDetail", newsinfo.AddNewsDetail).Methods("POST")
	route_news.HandleFunc("/getDetails", newsinfo.GetNewsDetails).Methods("GET")
	route_news.HandleFunc("/getDetail", newsinfo.GetNewsDetail).Methods("GET")
	route_news.HandleFunc("/updateDetailTitle", newsinfo.UpdateNewsTitle).Methods("PUT")
	route_news.HandleFunc("/updateDetailBanner", newsinfo.UpdateNewsBanner).Methods("PUT")
	route_news.HandleFunc("/updateDetailContent", newsinfo.UpdateNewsContent).Methods("PUT")
	route_news.HandleFunc("/deleteDetail", newsinfo.DeleteNewsDetail).Methods("DELETE")

	route_news.HandleFunc("/addImage", newsinfo.AddImage).Methods("POST")
	route_news.HandleFunc("/getImages", newsinfo.GetImages).Methods("GET")
	route_news.HandleFunc("/getImage", newsinfo.GetImage).Methods("GET")
	route_news.HandleFunc("/getFavImages", newsinfo.GetyourFavImages).Methods("GET")
	route_news.HandleFunc("/updateImages", newsinfo.UpdateImage).Methods("PUT")
	route_news.HandleFunc("/deleteImage", newsinfo.DeleteImage).Methods("DELETE")

	route_news.HandleFunc("/getTypes", newsinfo.GetNewsTypes).Methods("GET")
	route_news.HandleFunc("/getType", newsinfo.GetNewsType).Methods("GET")

	return r
}
