package icon

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddiconFrame(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_icon_frame(frame_id, description, release_date) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	frame_id := r.Form.Get("frame_id")
	description := r.Form.Get("description")
	release_date := r.Form.Get("release_date")

	_, err = stmt.Exec(frame_id, description, release_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GeticonFrames(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var frames []model.Icon_frame
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_icon_frame")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var frame model.Icon_frame
		err := result.Scan(&frame.Frame_id, &frame.Description, &frame.Release_date)
		if err != nil {
			panic(err.Error())
		}

		frames = append(frames, frame)
	}

	json.NewEncoder(w).Encode(frames)
}

func GeticonFrame(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("frame_id")

	var frame model.Icon_frame
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_icon_frame WHERE frame_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&frame.Frame_id, &frame.Description, &frame.Release_date)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(frame)
}

func UpdateiconFrame(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("frame_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_icon_frame SET description = ? WHERE frame_id = ?")
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

func DeleteiconFrame(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("frame_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_icon_frame WHERE frame_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
