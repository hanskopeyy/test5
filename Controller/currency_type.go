package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func AddCurrencyType(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_currency_type(name) VALUES (?)")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")

	_, err = stmt.Exec(name)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllCurrencyTypes(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var currencies []model.Currency

	result, err := db.Query("SELECT * from lokapala_accountdb.t_currency_type")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var currency model.Currency
		err := result.Scan(&currency.Currency_id, &currency.Name)
		if err != nil {
			panic(err.Error())
		}

		currencies = append(currencies, currency)

	}

	json.NewEncoder(w).Encode(currencies)

}

func GetCurrencyType(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var currency model.Currency
	result, err := db.Query("SELECT * from lokapala_accountdb.t_currency_type where currency_id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&currency.Currency_id, &currency.Name)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(currency)

}

func UpdateCurrencyType(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_currency_type SET name = ? where currency_id = ?")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")

	_, err = stmt.Exec(name, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteCurrencyType(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_currency_type WHERE currency_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
