package icon

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddiconAvatar(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_icon_avatar(avatar_id, description, release_date) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	avatar_id := r.Form.Get("avatar_id")
	description := r.Form.Get("description")
	release_date := r.Form.Get("release_date")

	_, err = stmt.Exec(avatar_id, description, release_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GeticonAvatars(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var avatars []model.Icon_avatar
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_icon_avatar")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var avatar model.Icon_avatar
		err := result.Scan(&avatar.Avatar_id, &avatar.Description, &avatar.Release_date, &avatar.Free)
		if err != nil {
			panic(err.Error())
		}

		avatars = append(avatars, avatar)
	}

	json.NewEncoder(w).Encode(avatars)
}

func GeticonAvatar(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("avatar_id")

	var avatar model.Icon_avatar
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_icon_avatar WHERE avatar_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&avatar.Avatar_id, &avatar.Description, &avatar.Release_date, &avatar.Free)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(avatar)
}

func UpdateiconAvatar(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("avatar_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_icon_avatar SET description = ? WHERE avatar_id = ?")
	if err != nil {
		panic(err.Error())
	}

	description := r.Form.Get("description")
	// release_date := r.Form.Get("release_date")

	_, err = stmt.Exec(description, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteiconAvatar(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("avatar_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_icon_avatar WHERE avatar_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
