package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func AddPremium(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_premium(item_id, duration) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	duration := r.Form.Get("duration")

	_, err = stmt.Exec(item_id, duration)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetPremiums(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var premiums []model.Premium
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_premium")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var premium model.Premium
		err := result.Scan(&premium.Item_id, &premium.Duration)
		if err != nil {
			panic(err.Error())
		}

		premiums = append(premiums, premium)
	}

	json.NewEncoder(w).Encode(premiums)
}

func GetPremium(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("item_id")

	var premium model.Premium
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_premium WHERE item_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&premium.Item_id, &premium.Duration)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(premium)
}

func UpdatePremium(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("item_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_premium SET duration = ? WHERE item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	duration := r.Form.Get("duration")

	_, err = stmt.Exec(duration, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeletePremium(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("item_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_premium WHERE item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
