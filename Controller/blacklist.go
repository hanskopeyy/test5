package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func BlacklistPlayer(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_blacklist(blacklist_user_id, target_user_id) VALUES (0,?)")
	if err != nil {
		panic(err.Error())
	}

	target_user_id := r.Form.Get("target_user_id")

	_, err = stmt.Exec(target_user_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllBlacklists(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var blacklists []model.Blacklist
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_blacklist")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var blacklist model.Blacklist
		err := result.Scan(&blacklist.Blacklist_user_id, &blacklist.Target_user_id, &blacklist.Blacklist_date)
		if err != nil {
			panic(err.Error())
		}

		blacklists = append(blacklists, blacklist)
	}

	json.NewEncoder(w).Encode(blacklists)
}

func GetBlacklist(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	user_id := r.URL.Query().Get("user_id")
	target_id := r.URL.Query().Get("target_id")

	var blacklist model.Blacklist
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_blacklist WHERE blacklist_user_id = ? AND target_user_id = ?", user_id, target_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&blacklist.Blacklist_user_id, &blacklist.Target_user_id, &blacklist.Blacklist_date)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(blacklist)
}

func UnblacklistPlayer(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	user_id := r.URL.Query().Get("user_id")
	target_id := r.URL.Query().Get("target_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_blacklist WHERE blacklist_user_id = ? AND target_user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(user_id, target_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
