package newsinfo

import (
	"encoding/json"
	"net/http"

	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddNewsDetail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO t_news_detail(news_id, lang, title, banner, banner_checksum, content, content_checksum) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	news_id := r.Form.Get("news_id")
	lang := r.Form.Get("lang")
	title := r.Form.Get("title")
	banner, banner_checksum, err := CheckorUpload(r, "banner")
	if err != nil {
		panic(err)
	}
	content, content_checksum, err := UploadFile(r, "content", "storage", lang)
	if err != nil {
		panic(err)
	}
	content = lang + "/" + content

	_, err = stmt.Exec(news_id, lang, title, banner, banner_checksum, content, content_checksum)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode("Success")
}

func GetNewsDetails(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var details []model.News_detail

	result, err := db.Query("SELECT * FROM t_news_detail")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var detail model.News_detail
		err := result.Scan(&detail.News_id, &detail.Lang, &detail.Title, &detail.Banner, &detail.Banner_checksum, &detail.Content, &detail.Content_checksum)
		if err != nil {
			panic(err.Error())
		}

		details = append(details, detail)
	}

	json.NewEncoder(w).Encode(details)
}

func GetNewsDetail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")
	lang := r.URL.Query().Get("lang")

	var detail model.News_detail
	result, err := db.Query("SELECT * FROM t_news_detail WHERE news_id = ? AND lang = ?", id, lang)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&detail.News_id, &detail.Lang, &detail.Title, &detail.Banner, &detail.Banner_checksum, &detail.Content, &detail.Content_checksum)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(detail)
}

func UpdateNewsBanner(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err.Error())
	}

	id := r.URL.Query().Get("id")
	lang := r.URL.Query().Get("lang")

	stmt, err := db.Prepare("UPDATE t_news_detail SET banner = ?, banner_checksum = ? where news_id = ? AND lang = ?")
	if err != nil {
		panic(err.Error())
	}

	banner, banner_checksum, err := CheckorUpload(r, "banner")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(banner, banner_checksum, id, lang)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdateNewsContent(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err.Error())
	}

	id := r.URL.Query().Get("id")
	lang := r.URL.Query().Get("lang")

	stmt, err := db.Prepare("UPDATE t_news_detail SET content = ?, content_checksum = ? where news_id = ? AND lang = ?")
	if err != nil {
		panic(err.Error())
	}

	content, content_checksum, err := UploadFile(r, "content", "storage", lang)
	if err != nil {
		panic(err)
	}
	content = lang + "/" + content

	_, err = stmt.Exec(content, content_checksum, id, lang)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdateNewsDetail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err.Error())
	}

	id := r.URL.Query().Get("id")
	lang := r.URL.Query().Get("lang")

	stmt, err := db.Prepare("UPDATE t_news_detail SET news_id = ?, lang = ?, title = ? WHERE news_id = ? AND lang = ?")
	if err != nil {
		panic(err.Error())
	}

	news_id := r.Form.Get("news_id")
	language := r.Form.Get("language")
	title := r.Form.Get("title")

	_, err = stmt.Exec(news_id, language, title, id, lang)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteNewsDetail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")
	lang := r.URL.Query().Get("lang")

	stmt, err := db.Prepare("DELETE FROM t_news_detail WHERE news_id = ? AND lang = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id, lang)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
