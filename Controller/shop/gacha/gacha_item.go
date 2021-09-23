package gacha

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddGachaItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_gacha_item(item_type, item_id, amount) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_type, item_id, amount)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllGachaItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var gacha_items []model.Gacha_item

	result, err := db.Query("SELECT * from lokapala_accountdb.t_gacha_item")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var gacha_item model.Gacha_item
		err := result.Scan(&gacha_item.Gacha_item_id, &gacha_item.Item_type, &gacha_item.Item_id, &gacha_item.Amount)
		if err != nil {
			panic(err.Error())
		}

		gacha_items = append(gacha_items, gacha_item)

	}

	json.NewEncoder(w).Encode(gacha_items)

}

func GetGachaItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var gacha_item model.Gacha_item
	result, err := db.Query("SELECT * from lokapala_accountdb.t_gacha_item where gacha_item_id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&gacha_item.Gacha_item_id, &gacha_item.Item_type, &gacha_item.Item_id, &gacha_item.Amount)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(gacha_item)

}

func UpdateGachaItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_gacha_item SET item_type = ?, item_id = ?, amount = ? where gacha_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_type, item_id, amount, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteGachaItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_gacha_item WHERE gacha_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
