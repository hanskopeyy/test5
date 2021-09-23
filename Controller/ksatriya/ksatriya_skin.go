package ksatriya

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddKsatriyaSkin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_ksatriya_skin(ksatriya_skin_id, ksatriya_id, release_date) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	ksatriya_skin_id := r.Form.Get("ksatriya_skin_id")
	ksatriya_id := r.Form.Get("ksatriya_id")
	release_date := r.Form.Get("release_date")

	_, err = stmt.Exec(ksatriya_skin_id, ksatriya_id, release_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllKsatriyaSkin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var ksatriya_skins []model.Ksatriya_skin

	result, err := db.Query("SELECT B.ksatriya_skin_id, B.ksatriya_id, A.ksatriya_name, B.release_date FROM lokapala_accountdb.t_ksatriya A LEFT JOIN lokapala_accountdb.t_ksatriya_skin B ON A.ksatriya_id = B.ksatriya_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var ksatriya_skin model.Ksatriya_skin
		err := result.Scan(&ksatriya_skin.Ksatriya_skin_id, &ksatriya_skin.Ksatriya_id, &ksatriya_skin.Ksatriya_name, &ksatriya_skin.Release_date)
		if err != nil {
			panic(err.Error())
		}

		ksatriya_skins = append(ksatriya_skins, ksatriya_skin)

	}

	json.NewEncoder(w).Encode(ksatriya_skins)

}

func GetKsatriyaSkin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	ksatriya_skin_id := r.URL.Query().Get("ksatriya_skin_id")

	var ksatriya_skin model.Ksatriya_skin
	result, err := db.Query("SELECT B.ksatriya_skin_id, B.ksatriya_id, A.ksatriya_name, B.release_date FROM lokapala_accountdb.t_ksatriya A LEFT JOIN lokapala_accountdb.t_ksatriya_skin B ON A.ksatriya_id = B.ksatriya_id where ksatriya_skin_id = ? ", ksatriya_skin_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&ksatriya_skin.Ksatriya_skin_id, &ksatriya_skin.Ksatriya_id, &ksatriya_skin.Ksatriya_name, &ksatriya_skin.Release_date)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(ksatriya_skin)

}

func UpdateKsatriyaSkin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	ksatriya_skin_id := r.URL.Query().Get("ksatriya_skin_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_ksatriya_skin SET ksatriya_id = ? where ksatriya_skin_id = ?")
	if err != nil {
		panic(err.Error())
	}

	ksatriya_id := r.Form.Get("ksatriya_id")

	_, err = stmt.Exec(ksatriya_id, ksatriya_skin_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteKsatriyaSkin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	ksatriya_skin_id := r.URL.Query().Get("ksatriya_skin_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_ksatriya_skin WHERE ksatriya_skin_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(ksatriya_skin_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
