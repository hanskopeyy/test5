package gacha

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_gacha(start_date, end_date, random_value) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")
	random_value := r.Form.Get("random_value")

	_, err = stmt.Exec(start_date, end_date, random_value)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var gachas []model.Gacha

	result, err := db.Query("SELECT * from lokapala_accountdb.t_gacha")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var gacha model.Gacha
		err := result.Scan(&gacha.Gacha_id, &gacha.Start_date, &gacha.End_date, &gacha.Random_value)
		if err != nil {
			panic(err.Error())
		}

		gachas = append(gachas, gacha)

	}

	json.NewEncoder(w).Encode(gachas)

}

func GetGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var gacha model.Gacha
	result, err := db.Query("SELECT * from lokapala_accountdb.t_gacha where gacha_id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&gacha.Gacha_id, &gacha.Start_date, &gacha.End_date, &gacha.Random_value)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(gacha)

}

func UpdateGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_gacha SET start_date = ?, end_date = ?, random_value = ? where gacha_id = ?")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")
	random_value := r.Form.Get("random_value")

	_, err = stmt.Exec(start_date, end_date, random_value, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteGacha(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_gacha WHERE gacha_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
