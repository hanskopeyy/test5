package newsinfo

import (
	"encoding/json"
	"net/http"

	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func GetNewsTypes(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	var newsTypes []model.News_type

	result, err := db.Query("SELECT * FROM t_news_type")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var newsType model.News_type
		err := result.Scan(&newsType.Id, &newsType.Name)
		if err != nil {
			panic(err.Error())
		}

		newsTypes = append(newsTypes, newsType)
	}
	json.NewEncoder(w).Encode(newsTypes)
}

func GetNewsType(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var newsType model.News_type
	result, err := db.Query("SELECT * FROM t_news_type WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&newsType.Id, &newsType.Name)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(newsType)
}
