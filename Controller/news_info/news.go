package newsinfo

import (
	"encoding/json"
	"net/http"

	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddNews(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO t_news(name, release_date, type) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")
	release_date := r.Form.Get("release_date")
	news_type := r.Form.Get("type")

	_, err = stmt.Exec(name, release_date, news_type)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllNews(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var allNews []model.News

	result, err := db.Query("SELECT * FROM t_news")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var news model.News
		err := result.Scan(&news.Id, &news.Name, &news.Release_date, &news.Type)
		if err != nil {
			panic(err.Error())
		}
		allNews = append(allNews, news)
	}

	json.NewEncoder(w).Encode(allNews)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var news model.News
	result, err := db.Query("SELECT * from t_news where id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&news.Id, &news.Name, &news.Release_date, &news.Type)
		if err != nil {
			panic(err.Error())
		}

	}
	json.NewEncoder(w).Encode(news)
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_news SET name = ?, release_date = ?, type = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")
	release_date := r.Form.Get("release_date")
	news_type := r.Form.Get("type")

	_, err = stmt.Exec(name, release_date, news_type, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM t_news WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
