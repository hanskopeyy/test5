package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func AddEnergy(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_energy(energy_id, description, target) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	energy_id := r.Form.Get("energy_id")
	description := r.Form.Get("description")
	target := r.Form.Get("target")

	_, err = stmt.Exec(energy_id, description, target)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetEnergies(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var energies []model.Energy
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_energy")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var energy model.Energy
		err := result.Scan(&energy.Energy_id, &energy.Description, &energy.Target)
		if err != nil {
			panic(err.Error())
		}

		energies = append(energies, energy)
	}

	json.NewEncoder(w).Encode(energies)
}

func GetEnergy(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("energy_id")

	var energy model.Energy
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_energy WHERE energy_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&energy.Energy_id, &energy.Description, &energy.Target)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(energy)
}

func UpdateEnergy(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("energy_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_energy SET description = ?, target = ? WHERE energy_id = ?")
	if err != nil {
		panic(err.Error())
	}

	description := r.Form.Get("description")
	target := r.Form.Get("target")

	_, err = stmt.Exec(description, target, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteEnergy(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("energy_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_energy WHERE energy_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
