package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func AddBox(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_box(box_name, rand_value) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	box_name := r.Form.Get("box_name")
	rand_value := r.Form.Get("rand_value")

	_, err = stmt.Exec(box_name, rand_value)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllBox(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var boxes []model.Box

	result, err := db.Query("SELECT * from lokapala_accountdb.t_box")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var box model.Box
		err := result.Scan(&box.Box_id, &box.Box_name, &box.Rand_value)
		if err != nil {
			panic(err.Error())
		}

		boxes = append(boxes, box)

	}

	json.NewEncoder(w).Encode(boxes)

}

func GetBox(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var box model.Box
	result, err := db.Query("SELECT * from lokapala_accountdb.t_box where box_id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&box.Box_id, &box.Box_name, &box.Rand_value)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(box)

}

func UpdateBox(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_box SET box_name = ?, rand_value = ? where box_id = ?")
	if err != nil {
		panic(err.Error())
	}

	box_name := r.Form.Get("box_name")
	rand_value := r.Form.Get("rand_value")

	_, err = stmt.Exec(box_name, rand_value, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteBox(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_box WHERE box_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func AddBoxLoot(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_box_loot_table(box_id, item_id, item_type, amount, chance, min, max) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	box_id := r.Form.Get("box_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	chance := r.Form.Get("chance")
	min := r.Form.Get("min")
	max := r.Form.Get("max")

	_, err = stmt.Exec(box_id, item_id, item_type, amount, chance, min, max)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllBoxLoot(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var loot_boxes []model.Box_loot_table

	result, err := db.Query("SELECT * from lokapala_accountdb.t_box_loot_table")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var loot_box model.Box_loot_table
		err := result.Scan(&loot_box.Uid, &loot_box.Box_id, &loot_box.Item_id, &loot_box.Item_type, &loot_box.Amount, &loot_box.Chance, &loot_box.Min, &loot_box.Max)
		if err != nil {
			panic(err.Error())
		}

		loot_boxes = append(loot_boxes, loot_box)

	}

	json.NewEncoder(w).Encode(loot_boxes)

}

func GetBoxLoot(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	uid := r.URL.Query().Get("uid")

	var loot_box model.Box_loot_table
	result, err := db.Query("SELECT * from lokapala_accountdb.t_box_loot_table where uid = ? ", uid)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&loot_box.Uid, &loot_box.Box_id, &loot_box.Item_id, &loot_box.Item_type, &loot_box.Amount, &loot_box.Chance, &loot_box.Min, &loot_box.Max)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(loot_box)

}

func UpdateBoxLoot(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	uid := r.URL.Query().Get("uid")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_box_loot_table SET box_id = ?, item_id = ?, item_type = ?, amount = ?, chance = ?, min = ?, max =? where uid = ?")
	if err != nil {
		panic(err.Error())
	}

	box_id := r.Form.Get("box_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	chance := r.Form.Get("chance")
	min := r.Form.Get("min")
	max := r.Form.Get("max")

	_, err = stmt.Exec(box_id, item_id, item_type, amount, chance, min, max, uid)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteBoxLoot(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	uid := r.URL.Query().Get("uid")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_box_loot_table WHERE uid = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(uid)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
