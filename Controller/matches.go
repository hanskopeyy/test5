package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func GetMatches(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	var matches []model.Matches

	query, err := db.Prepare("call lokapala_admindb.p_room_list_get(?,?,0)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var match model.Matches
		err := result.Scan(&match.Room_id, &match.Room_name, &match.Match_id, &match.Game_mode, &match.Server_ip, &match.Server_port, &match.Start_time, &match.Can_timeout)
		if err != nil {
			panic(err.Error())
		}

		matches = append(matches, match)

	}

	json.NewEncoder(w).Encode(matches)

}

func GetMatch(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	room_id := r.URL.Query().Get("room_id")

	var match model.Matches

	query, err := db.Prepare("call lokapala_admindb.p_room_list_get(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(count, offset, room_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&match.Room_id, &match.Room_name, &match.Match_id, &match.Game_mode, &match.Server_ip, &match.Server_port, &match.Start_time, &match.Can_timeout)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(match)

}

func CancelMatch(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	room_id := r.URL.Query().Get("room_id")

	var match int

	query, err := db.Prepare("call lokapala_admindb.p_room_timeout(0,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(room_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&match)
		if err != nil {
			panic(err.Error())
		}

	}

	if match == 1 {
		json.NewEncoder(w).Encode("Success")
		json.NewEncoder(w).Encode(match)
	} else {
		json.NewEncoder(w).Encode("error")
		json.NewEncoder(w).Encode(match)
	}

}
