package admin

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := controller.OpenGMAdmin()
	defer db.Close()
	var users []model.User_details

	result, err := db.Query("SELECT A.id, A.name, A.email, B.id, GROUP_CONCAT(B.role_name) as roles_name from users A left join users_roles on users_roles.user_id = A.id left join roles B on B.id = users_roles.role_id GROUP BY A.id, B.id ORDER BY A.id ASC")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var user model.User_details
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Role_id, &user.Role_name)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)

	}

	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := controller.OpenGMAdmin()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var user model.User_details
	result, err := db.Query("SELECT A.id, A.name, A.email, B.id, GROUP_CONCAT(B.role_name) as roles_name from users A left join users_roles on users_roles.user_id = A.id left join roles B on B.id = users_roles.role_id WHERE A.id = ? GROUP BY A.id, B.id ORDER BY A.id ASC  ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Role_id, &user.Role_name)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(user)

}

func GetCurrentUserLogin(w http.ResponseWriter, r *http.Request) {
	db := controller.OpenGMAdmin()
	defer db.Close()
	id := r.Context().Value("user_id").(string)

	var user model.User_details
	result, err := db.Query("SELECT A.id, A.name, A.email, B.id, GROUP_CONCAT(B.role_name) as roles_name from users A left join users_roles on users_roles.user_id = A.id left join roles B on B.id = users_roles.role_id WHERE A.id = ? GROUP BY A.id, B.id ORDER BY A.id ASC", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Role_id, &user.Role_name)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := controller.OpenGMAdmin()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}

	name_new := r.Form.Get("name")
	email_new := r.Form.Get("email")

	_, err = stmt.Exec(name_new, email_new, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	db := controller.OpenGMAdmin()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE users SET password = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}

	password_new := r.Form.Get("password")

	var pwd_new = []byte(password_new)

	pwdhash_new, err := bcrypt.GenerateFromPassword(pwd_new, bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(pwdhash_new, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := controller.OpenGMAdmin()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt1, err := db.Prepare("DELETE FROM users_roles WHERE user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt1.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
