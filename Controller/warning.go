package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetWarningKsaRotation(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	result, err := db.Query("SELECT IF((SELECT NOW()) > date_add((SELECT MAX(end_date) FROM lokapala_accountdb.t_ksatriya_rotation), INTERVAL -2 WEEK), true, false) as flag")
	if err != nil {
		panic(err.Error())
	}

	var flag string

	for result.Next() {

		err := result.Scan(&flag)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(flag)
}

func GetWarningGacha(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	result, err := db.Query("SELECT IF((SELECT NOW()) > date_add((SELECT MAX(end_date) FROM lokapala_accountdb.t_gacha), INTERVAL -2 WEEK), true, false) as flag")
	if err != nil {
		panic(err.Error())
	}

	var flag string

	for result.Next() {

		err := result.Scan(&flag)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(flag)
}

func GetWarningLotto(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	result, err := db.Query("SELECT IF((SELECT NOW()) > date_add((SELECT MAX(end_date) FROM lokapala_accountdb.t_lotto), INTERVAL -2 WEEK), true, false) as flag")
	if err != nil {
		panic(err.Error())
	}

	var flag string

	for result.Next() {

		err := result.Scan(&flag)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(flag)
}

func GetWarningLotus(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	result, err := db.Query("SELECT IF((SELECT NOW()) > date_add((SELECT MAX(end_date) FROM lokapala_accountdb.t_shop_lotus_period), INTERVAL -2 WEEK), true, false) as flag")
	if err != nil {
		panic(err.Error())
	}

	var flag string

	for result.Next() {

		err := result.Scan(&flag)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(flag)
}

func GetWarningSeason(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	result, err := db.Query("SELECT IF((SELECT NOW()) > date_add((SELECT MAX(end_date) FROM lokapala_accountdb.t_season), INTERVAL -2 WEEK), true, false) as flag")
	if err != nil {
		panic(err.Error())
	}

	var flag string

	for result.Next() {

		err := result.Scan(&flag)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(flag)
}
