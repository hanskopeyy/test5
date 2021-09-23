package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func AddRune(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_rune(rune_id, name, description, rune_color) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	rune_id := r.Form.Get("rune_id")
	name := r.Form.Get("name")
	description := r.Form.Get("description")
	rune_color := r.Form.Get("rune_color")

	_, err = stmt.Exec(rune_id, name, description, rune_color)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetRunes(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var runes []model.Rune
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_rune")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var runee model.Rune
		err := result.Scan(&runee.Rune_id, &runee.Name, &runee.Description, &runee.Rune_color)
		if err != nil {
			panic(err.Error())
		}
		runes = append(runes, runee)
	}

	json.NewEncoder(w).Encode(runes)
}

func GetRune(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("rune_id")

	var rune model.Rune
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_rune WHERE rune_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&rune.Rune_id, &rune.Name, &rune.Description, &rune.Rune_color)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(rune)
}

func UpdateRune(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("rune_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_rune SET name = ?, description = ?, rune_color = ? WHERE rune_id = ?")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")
	description := r.Form.Get("description")
	rune_color := r.Form.Get("rune_color")

	_, err = stmt.Exec(name, description, rune_color, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteRune(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("rune_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_rune WHERE rune_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
