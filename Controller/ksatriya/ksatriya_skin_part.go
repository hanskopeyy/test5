package ksatriya

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddKsatriyaSkinPart(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_ksatriya_skin_part(skin_part_id, release_date) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	skin_part_id := r.Form.Get("skin_part_id")
	release_date := r.Form.Get("release_date")

	_, err = stmt.Exec(skin_part_id, release_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetKsatriyaSkinParts(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var parts []model.Ksatriya_skin_part
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_ksatriya_skin_part")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var part model.Ksatriya_skin_part
		err := result.Scan(&part.Skin_part_id, &part.Release_date)
		if err != nil {
			panic(err.Error())
		}

		parts = append(parts, part)
	}

	json.NewEncoder(w).Encode(parts)
}

func GetKsatriyaSkinPart(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var part model.Ksatriya_skin_part
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_ksatriya_skin_part WHERE skin_part_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&part.Skin_part_id, &part.Release_date)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(part)
}

/* func UpdateKsatriyaSkinPart(w http.ResponseWriter, r *http.Request) {
     db := controller.Open()
    defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_ksatriya_skin_part SET release_date = ? WHERE skin_part_id = ?")
	if err != nil {
		panic(err.Error())
	}

	release_date := r.Form.Get("release_date")

	_, err = stmt.Exec(release_date, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
} */

func DeleteKsatriyaSkinPart(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_ksatriya_skin_part WHERE skin_part_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
