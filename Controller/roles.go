package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func AddRoles(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO roles(role_name, description) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	role_name := r.Form.Get("role_name")
	description := r.Form.Get("description")

	_, err = stmt.Exec(role_name, description)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	var roles []model.Roles

	result, err := db.Query("SELECT id, role_name, description from roles")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var role model.Roles
		err := result.Scan(&role.Id, &role.Role_name, &role.Description)
		if err != nil {
			panic(err.Error())
		}

		roles = append(roles, role)

	}

	json.NewEncoder(w).Encode(roles)

}

func GetRole(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var role model.Roles
	result, err := db.Query("SELECT id, role_name, description from roles where id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&role.Id, &role.Role_name, &role.Description)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(role)

}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE roles SET role_name = ?, description = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}

	role_name_new := r.Form.Get("role_name_new")
	description_new := r.Form.Get("description_new")

	_, err = stmt.Exec(role_name_new, description_new, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM roles WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
