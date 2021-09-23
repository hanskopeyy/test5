package ksatriya

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddnewKsatriyaRotation(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_ksatriya_rotation(ksatriya_id,start_date, end_date) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	ksatriya_id := r.Form.Get("ksatriya_id")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(ksatriya_id, start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllKsatriyasRotation(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var ksatriyas_rotation []model.Ksatriya_rotation
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_ksatriya_rotation")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var ksatriya_rotation model.Ksatriya_rotation
		err := result.Scan(&ksatriya_rotation.Ksatriya_rotation_id, &ksatriya_rotation.Ksatriya_id, &ksatriya_rotation.Start_date, &ksatriya_rotation.End_date)
		if err != nil {
			panic(err.Error())
		}

		ksatriyas_rotation = append(ksatriyas_rotation, ksatriya_rotation)
	}

	json.NewEncoder(w).Encode(ksatriyas_rotation)
}

func GetKsatriyaRotation(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_rotation_id")

	var ksatriya_rotation model.Ksatriya_rotation
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_ksatriya_rotation WHERE ksatriya_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&ksatriya_rotation.Ksatriya_rotation_id, &ksatriya_rotation.Ksatriya_id, &ksatriya_rotation.Start_date, &ksatriya_rotation.End_date)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(ksatriya_rotation)
}

func UpdateKsatriyaRotation(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_rotation_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_ksatriya_rotation SET ksatriya_id = ?, start_date = ?, end_date = ? WHERE ksatriya_rotation_id = ?")
	if err != nil {
		panic(err.Error())
	}

	ksatriya_id := r.Form.Get("ksatriya_id")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(ksatriya_id, start_date, end_date, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteKsatriyaRotation(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("ksatriya_rotation_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_ksatriya_rotation WHERE ksatriya_rotation_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
