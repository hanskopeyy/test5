package lotto

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddlottoItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_lotto_item(item_type,item_id,amount,color_id,default_amount,item_name) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	color_id := r.Form.Get("color_id")
	default_amount := r.Form.Get("default_amount")
	item_name := r.Form.Get("item_name")

	_, err = stmt.Exec(item_type, item_id, amount, color_id, default_amount, item_name)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetlottoItems(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var l_items []model.Lotto_item
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_item")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var l_item model.Lotto_item
		err := result.Scan(&l_item.Lotto_item_id, &l_item.Item_type, &l_item.Item_id, &l_item.Amount, &l_item.Color_id, &l_item.Default_amount, &l_item.Item_name)
		if err != nil {
			panic(err.Error())
		}

		l_items = append(l_items, l_item)
	}

	json.NewEncoder(w).Encode(l_items)
}

func GetlottoItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var l_item model.Lotto_item
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_item WHERE lotto_item_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&l_item.Lotto_item_id, &l_item.Item_type, &l_item.Item_id, &l_item.Amount, &l_item.Color_id, &l_item.Default_amount, &l_item.Item_name)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(l_item)

}

func UpdatelottoItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_lotto_item SET item_type = ?, item_id = ?, amount = ?, color_id = ?, default_amount = ?, item_name = ? WHERE lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	color_id := r.Form.Get("color_id")
	default_amount := r.Form.Get("default_amount")
	item_name := r.Form.Get("item_name")

	_, err = stmt.Exec(item_type, item_id, amount, color_id, default_amount, item_name, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeletelottoItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_lotto_item WHERE lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
