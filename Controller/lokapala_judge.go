package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func RegisterJudge(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_judgedb.t_judge(username, password) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	_, err = stmt.Exec(username, password)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllJudge(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var judges []model.Judge_acc

	result, err := db.Query("SELECT * from lokapala_judgedb.t_judge")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var judge model.Judge_acc
		err := result.Scan(&judge.User_id, &judge.Username, &judge.Password)
		if err != nil {
			panic(err.Error())
		}

		judges = append(judges, judge)

	}

	json.NewEncoder(w).Encode(judges)

}

func GetJudge(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var judge model.Judge_acc
	result, err := db.Query("SELECT * from lokapala_judgedb.t_judge where user_id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&judge.User_id, &judge.Username, &judge.Password)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(judge)

}

func UpdateJudgeName(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_judgedb.t_judge SET username = ? where user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	username := r.Form.Get("username")

	_, err = stmt.Exec(username, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdateJudgePassword(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_judgedb.t_judge SET password = ? where user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	password := r.Form.Get("password")

	_, err = stmt.Exec(password, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteJudge(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_judgedb.t_judge WHERE user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
