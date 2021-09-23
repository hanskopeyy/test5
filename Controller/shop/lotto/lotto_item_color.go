package lotto

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddlottoColor(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_lotto_item_color(color_name,weight) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	color_name := r.Form.Get("color_name")
	weight := r.Form.Get("weight")

	_, err = stmt.Exec(color_name, weight)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetlottoColors(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var l_colors []model.Lotto_item_color
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_item_color")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var l_color model.Lotto_item_color
		err := result.Scan(&l_color.Color_id, &l_color.Color_name, &l_color.Weight)
		if err != nil {
			panic(err.Error())
		}

		l_colors = append(l_colors, l_color)
	}
	json.NewEncoder(w).Encode(l_colors)
}

func GetlottoColor(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var l_color model.Lotto_item_color
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_item_color WHERE color_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&l_color.Color_id, &l_color.Color_name, &l_color.Weight)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(l_color)
}

func UpdatelottoColor(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_lotto_item_color SET color_name = ?, weight = ? WHERE color_id = ?")
	if err != nil {
		panic(err.Error())
	}

	color_name := r.Form.Get("color_name")
	weight := r.Form.Get("weight")

	_, err = stmt.Exec(color_name, weight, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeletelottoColor(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_lotto_item_color WHERE color_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
