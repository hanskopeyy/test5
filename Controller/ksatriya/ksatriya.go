package ksatriya

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddnewKsatriya(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_ksatriya(ksatriya_id, role, release_date, ksatriya_name) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	ksatriya_id := r.Form.Get("ksatriya_id")
	role := r.Form.Get("role")
	release_date := r.Form.Get("release_date")
	ksatriya_name := r.Form.Get("ksatriya_name")

	_, err = stmt.Exec(ksatriya_id, role, release_date, ksatriya_name)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetKsatriyas(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var ksatriyas []model.Ksatriya
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_ksatriya")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var ksatriya model.Ksatriya
		err := result.Scan(&ksatriya.Ksatriya_id, &ksatriya.Role, &ksatriya.Release_date, &ksatriya.Ksatriya_name)
		if err != nil {
			panic(err.Error())
		}

		ksatriyas = append(ksatriyas, ksatriya)
	}

	json.NewEncoder(w).Encode(ksatriyas)
}

func GetKsatriya(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_id")

	var ksatriya model.Ksatriya
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_ksatriya WHERE ksatriya_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&ksatriya.Ksatriya_id, &ksatriya.Role, &ksatriya.Release_date, &ksatriya.Ksatriya_name)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(ksatriya)
}

func UpdateKsatriya(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_ksatriya SET role = ?, release_date = ?, ksatriya_name = ? WHERE ksatriya_id = ?")
	if err != nil {
		panic(err.Error())
	}

	role := r.Form.Get("role")
	release_date := r.Form.Get("release_date")
	ksatriya_name := r.Form.Get("ksatriya_name")

	_, err = stmt.Exec(role, release_date, ksatriya_name, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteKsatriya(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_ksatriya WHERE ksatriya_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
