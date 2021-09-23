package lotto

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddlottoFeature(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_lotto_feature(lotto_id,lotto_item_id,priority) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	lotto_id := r.Form.Get("lotto_id")
	lotto_item_id := r.Form.Get("lotto_item_id")
	priority := r.Form.Get("priority")

	_, err = stmt.Exec(lotto_id, lotto_item_id, priority)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetlottoFeatures(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var l_features []model.Lotto_feature
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_feature")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var l_feature model.Lotto_feature
		err := result.Scan(&l_feature.Lotto_feature_id, &l_feature.Lotto_id, &l_feature.Lotto_item_id, &l_feature.Priority)
		if err != nil {
			panic(err.Error())
		}

		l_features = append(l_features, l_feature)
	}
	json.NewEncoder(w).Encode(l_features)
}

func GetlottoFeature(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var l_feature model.Lotto_feature
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_feature WHERE lotto_feature_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&l_feature.Lotto_feature_id, &l_feature.Lotto_id, &l_feature.Lotto_item_id, &l_feature.Priority)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(l_feature)
}

func GetlottoFeatureByLottoId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var l_features []model.Lotto_feature
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_feature WHERE lotto_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var l_feature model.Lotto_feature
		err := result.Scan(&l_feature.Lotto_feature_id, &l_feature.Lotto_id, &l_feature.Lotto_item_id, &l_feature.Priority)
		if err != nil {
			panic(err.Error())
		}
		l_features = append(l_features, l_feature)
	}
	json.NewEncoder(w).Encode(l_features)
}

func UpdatelottoFeature(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_lotto_feature SET lotto_id = ?, lotto_item_id = ?, priority = ? WHERE lotto_feature_id = ?")
	if err != nil {
		panic(err.Error())
	}

	lotto_id := r.Form.Get("lotto_id")
	lotto_item_id := r.Form.Get("lotto_item_id")
	priority := r.Form.Get("priority")

	_, err = stmt.Exec(lotto_id, lotto_item_id, priority, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeletelottoFeature(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_lotto_feature WHERE lotto_feature_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
