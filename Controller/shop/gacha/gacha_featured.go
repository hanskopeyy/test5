package gacha

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddFeaturedGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_gacha_feature(gacha_id, gacha_item_id, priority) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	gacha_id := r.Form.Get("gacha_id")
	gacha_item_id := r.Form.Get("gacha_item_id")
	priority := r.Form.Get("priority")

	_, err = stmt.Exec(gacha_id, gacha_item_id, priority)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllFeaturedGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var gacha_features []model.Gacha_feature

	result, err := db.Query("SELECT * from lokapala_accountdb.t_gacha_feature")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var gacha_feature model.Gacha_feature
		err := result.Scan(&gacha_feature.Gacha_feature_id, &gacha_feature.Gacha_id, &gacha_feature.Gacha_item_id, &gacha_feature.Priority)
		if err != nil {
			panic(err.Error())
		}

		gacha_features = append(gacha_features, gacha_feature)

	}

	json.NewEncoder(w).Encode(gacha_features)

}

func GetFeaturedGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var gacha_feature model.Gacha_feature
	result, err := db.Query("SELECT * from lokapala_accountdb.t_gacha_feature where gacha_feature_id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&gacha_feature.Gacha_feature_id, &gacha_feature.Gacha_id, &gacha_feature.Gacha_item_id, &gacha_feature.Priority)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(gacha_feature)

}

func UpdateFeaturedGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_gacha_feature SET gacha_id = ?, gacha_item_id = ?, priority = ? where gacha_feature_id = ?")
	if err != nil {
		panic(err.Error())
	}

	gacha_id := r.Form.Get("gacha_id")
	gacha_item_id := r.Form.Get("gacha_item_id")
	priority := r.Form.Get("priority")

	_, err = stmt.Exec(gacha_id, gacha_item_id, priority, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteFeaturedGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_gacha_feature WHERE gacha_feature_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
