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

	stmt, err := db.Prepare("INSERT INTO t_news_v2_detail(news_id, lang, title, banner, banner_checksum, content, content_checksum) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	stmt2, err := db.Prepare("INSERT INTO t_news_v2_detail(news_id, lang, title, banner, banner_checksum, content, content_checksum) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	stmt3, err := db.Prepare("INSERT INTO t_news_v2(name, release_date, type) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	// news_id := r.Form.Get("news_id")
	// lang := r.Form.Get("lang")
	titleEN := r.Form.Get("titleEN")
	titleIN := r.Form.Get("titleIN")
	news_type := r.Form.Get("type")
	EN := "en"
	IN := "in"
	banner, banner_checksum, err := CheckorUpload(r, "banner")
	if err != nil {
		panic(err)
	}
	contentEN, content_checksumEN, err := UploadFile(r, "contentEN", "Test", EN)
	if err != nil {
		panic(err)
	}
	contentIN, content_checksumIN, err := UploadFile(r, "contentIN", "Test", IN)
	if err != nil {
		panic(err)
	}
	contentEN = EN + "/" + contentEN
	contentIN = IN + "/" + contentIN

	release_date := r.Form.Get("release_date")

	_, err = stmt3.Exec(titleEN, release_date, news_type)
	if err != nil {
		panic(err)
	}

	queryID, err := db.Query("SELECT MAX(id) as news_id FROM lokapala_accountdb.t_news_v2")
	if err != nil {
		panic(err.Error())
	}

	var newsId model.News_detail

	for queryID.Next() {

		err := queryID.Scan(&newsId.News_id)
		if err != nil {
			panic(err.Error())
		}
	}

	news_id := newsId.News_id

	_, err = stmt.Exec(news_id, EN, titleEN, banner, banner_checksum, contentEN, content_checksumEN)
	if err != nil {
		panic(err)
	}

	_, err = stmt2.Exec(news_id, IN, titleIN, banner, banner_checksum, contentIN, content_checksumIN)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode("Success")
}

func GetNewsDetails(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var details []model.News_detail

	result, err := db.Query("SELECT * FROM t_news_v2_detail")
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

	var details []model.News_detail
	result, err := db.Query("SELECT * FROM t_news_v2_detail WHERE news_id = ?", id)
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

func UpdateNewsBanner(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err.Error())
	}

	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("UPDATE t_news_v2_detail SET banner = ?, banner_checksum = ? where news_id = ?")
	if err != nil {
		panic(err.Error())
	}

	banner, banner_checksum, err := CheckorUpload(r, "banner")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(banner, banner_checksum, id)
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

	stmt, err := db.Prepare("UPDATE t_news_v2_detail SET content = ?, content_checksum = ? where news_id = ? AND lang = ?")
	if err != nil {
		panic(err.Error())
	}

	content, content_checksum, err := UploadFile(r, "content", "Test", lang)
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

func UpdateNewsTitle(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err.Error())
	}

	id := r.URL.Query().Get("id")
	lang := r.URL.Query().Get("lang")

	stmt, err := db.Prepare("UPDATE t_news_v2_detail SET title = ? WHERE news_id = ? AND lang = ?")
	if err != nil {
		panic(err.Error())
	}

	title := r.Form.Get("title")

	_, err = stmt.Exec(title, id, lang)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteNewsDetail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM t_news_v2_detail WHERE news_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	stmt2, err := db.Prepare("DELETE FROM t_news_v2 WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt2.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
