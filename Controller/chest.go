package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func AddChest(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_chest(duration) VALUES (?)")
	if err != nil {
		panic(err.Error())
	}

	duration := r.Form.Get("duration")

	_, err = stmt.Exec(duration)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllChest(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var chests []model.Chest

	result, err := db.Query("SELECT * from lokapala_accountdb.t_chest")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var chest model.Chest
		err := result.Scan(&chest.Duration)
		if err != nil {
			panic(err.Error())
		}

		chests = append(chests, chest)

	}

	json.NewEncoder(w).Encode(chests)

}

func GetChest(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	duration := r.URL.Query().Get("duration")

	var chest model.Chest
	result, err := db.Query("SELECT * from lokapala_accountdb.t_chest where duration = ? ", duration)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&chest.Duration)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(chest)

}

func UpdateChest(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	duration := r.URL.Query().Get("duration")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_chest SET duration = ? where duration = ?")
	if err != nil {
		panic(err.Error())
	}

	duration_new := r.Form.Get("duration")

	_, err = stmt.Exec(duration_new, duration)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteChest(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	duration := r.URL.Query().Get("duration")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_chest WHERE duration = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(duration)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
