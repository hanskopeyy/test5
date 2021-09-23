package ksatriya

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddKsatriyaSkinFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_ksatriya_skin_fragment(ksatriya_skin_id, amount_needed, sell_currency_id, sell_value) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	ksatriya_skin_id := r.Form.Get("ksatriya_skin_id")
	amount_needed := r.Form.Get("amount_needed")
	sell_currency_id := r.Form.Get("sell_currency_id")
	sell_value := r.Form.Get("sell_value")

	_, err = stmt.Exec(ksatriya_skin_id, amount_needed, sell_currency_id, sell_value)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllKsatriyaSkinFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var ksatriya_skin_fragments []model.Ksatriya_skin_fragment

	result, err := db.Query("SELECT * from lokapala_accountdb.t_ksatriya_skin")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var ksatriya_skin_fragment model.Ksatriya_skin_fragment
		err := result.Scan(&ksatriya_skin_fragment.Ksatriya_skin_id, &ksatriya_skin_fragment.Amount_needed, &ksatriya_skin_fragment.Sell_currency_id, &ksatriya_skin_fragment.Sell_value)
		if err != nil {
			panic(err.Error())
		}

		ksatriya_skin_fragments = append(ksatriya_skin_fragments, ksatriya_skin_fragment)

	}

	json.NewEncoder(w).Encode(ksatriya_skin_fragments)

}

func GetKsatriyaSkinFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	ksatriya_skin_id := r.URL.Query().Get("ksatriya_skin_id")

	var ksatriya_skin_fragment model.Ksatriya_skin_fragment
	result, err := db.Query("SELECT * from lokapala_accountdb.t_ksatriya_skin where ksatriya_skin_id = ? ", ksatriya_skin_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&ksatriya_skin_fragment.Ksatriya_skin_id, &ksatriya_skin_fragment.Amount_needed, &ksatriya_skin_fragment.Sell_currency_id, &ksatriya_skin_fragment.Sell_value)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(ksatriya_skin_fragment)

}

func UpdateKsatriyaSkinFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	ksatriya_skin_id := r.URL.Query().Get("ksatriya_skin_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_ksatriya_skin_fragment SET amount_needed = ?, sell_currency_id = ?, sell_value = ? where ksatriya_skin_id = ?")
	if err != nil {
		panic(err.Error())
	}

	amount_needed := r.Form.Get("amount_needed")
	sell_currency_id := r.Form.Get("sell_currency_id")
	sell_value := r.Form.Get("sell_value")

	_, err = stmt.Exec(amount_needed, sell_currency_id, sell_value, ksatriya_skin_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteKsatriyaSkinFragment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	ksatriya_skin_id := r.URL.Query().Get("ksatriya_skin_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_ksatriya_skin_fragment WHERE ksatriya_skin_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(ksatriya_skin_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
