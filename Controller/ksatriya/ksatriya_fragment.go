package ksatriya

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddKsatriyaFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_ksatriya_fragment(ksatriya_id, amount_needed, sell_currency_id, sell_value) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	ksatriya_id := r.Form.Get("ksatriya_id")
	amount_needed := r.Form.Get("amount_needed")
	sell_currency_id := r.Form.Get("sell_currency_id")
	sell_value := r.Form.Get("sell_value")

	_, err = stmt.Exec(ksatriya_id, amount_needed, sell_currency_id, sell_value)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetKsatriyaFragments(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var fragments []model.Ksatriya_fragment
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_ksatriya_fragment")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var fragment model.Ksatriya_fragment
		err := result.Scan(&fragment.Ksatriya_id, &fragment.Amount_needed, &fragment.Sell_currency_id, &fragment.Sell_value)
		if err != nil {
			panic(err.Error())
		}

		fragments = append(fragments, fragment)
	}

	json.NewEncoder(w).Encode(fragments)
}

func GetKsatriyaFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_id")

	var fragment model.Ksatriya_fragment
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_ksatriya_fragment WHERE ksatriya_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&fragment.Ksatriya_id, &fragment.Amount_needed, &fragment.Sell_currency_id, &fragment.Sell_value)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(fragment)
}

func UpdateKsatriyaFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_ksatriya_fragment SET amount_needed = ?, sell_currency_id = ?, sell_value = ? WHERE ksatriya_id = ?")
	if err != nil {
		panic(err.Error())
	}

	amount_needed := r.Form.Get("amount_needed")
	sell_currency_id := r.Form.Get("sell_currency_id")
	sell_value := r.Form.Get("sell_value")

	_, err = stmt.Exec(amount_needed, sell_currency_id, sell_value, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteKsatriyaFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_ksatriya_fragment WHERE ksatriya_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
