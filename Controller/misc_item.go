package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func GetMiscItems(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var misc_items []model.Misc_items
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_misc_item")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var misc_item model.Misc_items
		err := result.Scan(&misc_item.Misc_id, &misc_item.Misc_name, &misc_item.Amount)
		if err != nil {
			panic(err.Error())
		}
		misc_items = append(misc_items, misc_item)
	}

	json.NewEncoder(w).Encode(misc_items)
}

func GetMiscItem(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("misc_id")

	var misc_item model.Misc_items
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_misc_item WHERE misc_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&misc_item.Misc_id, &misc_item.Misc_name, &misc_item.Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(misc_item)
}
