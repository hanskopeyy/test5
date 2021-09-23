package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func GetAllUsersRoles(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	var users_roles []model.Users_roles

	result, err := db.Query("select users.id, users.name, roles.id, GROUP_CONCAT(roles.role_name) as roles_name from users left join users_roles on users_roles.user_id = users.id left join roles on roles.id = users_roles.role_id GROUP BY users.id ORDER BY users.id ASC")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var users_role model.Users_roles
		err := result.Scan(&users_role.User_id, &users_role.User_name, &users_role.Role_id, &users_role.Role_name)
		if err != nil {
			panic(err.Error())
		}

		users_roles = append(users_roles, users_role)

	}

	json.NewEncoder(w).Encode(users_roles)

}

func GetUserRole(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var users_role model.Users_roles
	result, err := db.Query("select users.id, users.name, roles.id, GROUP_CONCAT(roles.role_name) as roles_name from users left join users_roles on users_roles.user_id = users.id left join roles on roles.id = users_roles.role_id WHERE users.id = ? GROUP BY users.id ORDER BY users.id ASC", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&users_role.User_id, &users_role.User_name, &users_role.Role_id, &users_role.Role_name)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(users_role)

}

func AddNewUserToRole(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")
	role_id := r.URL.Query().Get("role_id")

	stmt, err := db.Prepare("INSERT INTO users_roles(user_id, role_id) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(user_id, role_id)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode("Success")

}

func RemoveUserFromRole(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	user_id := r.URL.Query().Get("user_id")
	role_id := r.URL.Query().Get("role_id")

	stmt, err := db.Prepare("DELETE FROM users_roles WHERE user_id = ? AND role_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(role_id, user_id)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode("Success")
}
